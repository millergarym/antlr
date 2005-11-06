/*
 [The "BSD licence"]
 Copyright (c) 2005 Terence Parr
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions
 are met:
 1. Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.
 2. Redistributions in binary form must reproduce the above copyright
    notice, this list of conditions and the following disclaimer in the
    documentation and/or other materials provided with the distribution.
 3. The name of the author may not be used to endorse or promote products
    derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE AUTHOR ``AS IS'' AND ANY EXPRESS OR
 IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
 OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
 IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
 INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
 NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
 THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package org.antlr.tool;

import org.antlr.analysis.*;
import org.antlr.stringtemplate.StringTemplate;
import org.antlr.stringtemplate.StringTemplateGroup;
import org.antlr.stringtemplate.language.AngleBracketTemplateLexer;
import org.antlr.tool.Grammar;
import org.antlr.tool.ANTLRParser;
import org.antlr.misc.Utils;
import org.antlr.Tool;

import java.util.*;
import java.io.FileWriter;
import java.io.IOException;
import java.io.File;

/** The DOT (part of graphviz) generation aspect.
 *  TODO: move templates out of test!
 */
public class DOTGenerator {
	public static final boolean STRIP_NONREDUCED_STATES = false;

    /** Library of output templates; use <attrname> format */
    public static StringTemplateGroup stlib =
            new StringTemplateGroup("toollib", AngleBracketTemplateLexer.class);

    /** To prevent infinite recursion when walking state machines, record
     *  which states we've visited.  Make a new set every time you start
     *  walking in case you reuse this object.
     */
    protected Set markedStates = null;

    protected Grammar grammar;

    /** This aspect is associated with a grammar */
	public DOTGenerator(Grammar grammar) {
		this.grammar = grammar;
	}

    /** Return a String containing a DOT description that, when displayed,
     *  will show the incoming state machine visually.  All nodes reachable
     *  from startState will be included.
     */
    public String getDOT(State startState) {
        // The output DOT graph for visualization
        StringTemplate dot = null;
		markedStates = new HashSet();
        if ( startState instanceof DFAState ) {
            dot = stlib.getInstanceOf("org/antlr/tool/templates/dot/dfa");
			dot.setAttribute("startState",
					new Integer(startState.stateNumber));
			dot.setAttribute("useBox",
							 new Boolean(Tool.internalOption_ShowNFConfigsInDFA));
			walkCreatingDOT(dot, startState);
        }
        else {
            dot = stlib.getInstanceOf("org/antlr/tool/templates/dot/nfa");
			dot.setAttribute("startState",
					new Integer(startState.stateNumber));
			walkRuleNFACreatingDOT(dot, startState);
        }
        return dot.toString();
    }

    /** Return a String containing a DOT description that, when displayed,
     *  will show the incoming state machine visually.  All nodes reachable
     *  from startState will be included.
    public String getRuleNFADOT(State startState) {
        // The output DOT graph for visualization
        StringTemplate dot = stlib.getInstanceOf("org/antlr/tool/templates/dot/nfa");

        markedStates = new HashSet();
        dot.setAttribute("startState",
                new Integer(startState.stateNumber));
        walkRuleNFACreatingDOT(dot, startState);
        return dot.toString();
    }
	 */

    /** Do a depth-first walk of the state machine graph and
     *  fill a DOT description template.  Keep filling the
     *  states and edges attributes.
     */
    protected void walkCreatingDOT(StringTemplate dot,
                                   State s)
    {
        if ( markedStates.contains(s) ) {
            return; // already visited this node
        }

        markedStates.add(s); // mark this node as completed.

        // first add this node
        StringTemplate st;
        if ( s.isAcceptState() ) {
            st = stlib.getInstanceOf("org/antlr/tool/templates/dot/stopstate");
        }
        else {
            st = stlib.getInstanceOf("org/antlr/tool/templates/dot/state");
        }
        st.setAttribute("name", getStateLabel(s));
        dot.setAttribute("states", st);

        // make a DOT edge for each transition
		for (int i = 0; i < s.getNumberOfTransitions(); i++) {
			Transition edge = (Transition) s.transition(i);
			//System.out.println("dfa "+s+" edge from s"+s.stateNumber+" ["+i+"] of "+s.getNumberOfTransitions());
			if ( STRIP_NONREDUCED_STATES ) {
				if ( edge.target instanceof DFAState &&
					((DFAState)edge.target).getAcceptStateReachable()!=DFA.REACHABLE_YES )
				{
 					continue; // don't generate nodes for terminal states
				}
			}
            st = stlib.getInstanceOf("org/antlr/tool/templates/dot/edge");
            st.setAttribute("label", getEdgeLabel(edge.label.toString(grammar)));
            st.setAttribute("src", getStateLabel(s));
            st.setAttribute("target", getStateLabel(edge.target));
            dot.setAttribute("edges", st);
            walkCreatingDOT(dot, edge.target); // keep walkin'
        }
    }

    /** Do a depth-first walk of the state machine graph and
     *  fill a DOT description template.  Keep filling the
     *  states and edges attributes.  We know this is an NFA
     *  for a rule so don't traverse edges to other rules and
     *  don't go past rule end state.
     */
    protected void walkRuleNFACreatingDOT(StringTemplate dot,
                                          State s)
    {
        if ( markedStates.contains(s) ) {
            return; // already visited this node
        }

        markedStates.add(s); // mark this node as completed.

        // first add this node
        StringTemplate st;
        if ( s.isAcceptState() ) {
            st = stlib.getInstanceOf("org/antlr/tool/templates/dot/stopstate");
        }
        else {
            st = stlib.getInstanceOf("org/antlr/tool/templates/dot/state");
        }
        st.setAttribute("name", getStateLabel(s));
        dot.setAttribute("states", st);

        if ( s.isAcceptState() )  {
            return; // don't go past end of rule node to the follow states
        }

        // special case: if decision point, then line up the alt start states
        // unless it's an end of block
        if ( ((NFAState)s).isDecisionState() &&
             ((NFAState)s).getAssociatedASTNode().getType()!=ANTLRParser.EOB )
        {
            st = stlib.getInstanceOf("org/antlr/tool/templates/dot/decision-rank");
            NFAState alt = (NFAState)s;
            while ( alt!=null ) {
                st.setAttribute("states", getStateLabel(alt));
                if ( alt.transition(1)!=null ) {
                    alt = (NFAState)alt.transition(1).target;
                }
                else {
                    alt=null;
                }
            }
            dot.setAttribute("decisionRanks", st);
        }

        // make a DOT edge for each transition
        for (int i = 0; i < s.getNumberOfTransitions(); i++) {
            Transition edge = (Transition) s.transition(i);
            if ( edge instanceof RuleClosureTransition ) {
                RuleClosureTransition rr = ((RuleClosureTransition)edge);
                // don't jump to other rules, but display edge to follow node
                st = stlib.getInstanceOf("org/antlr/tool/templates/dot/edge");
                st.setAttribute("label", "<"+grammar.getRuleName(rr.getRuleIndex())+">");
                st.setAttribute("src", getStateLabel(s));
                st.setAttribute("target", getStateLabel(rr.getFollowState()));
                dot.setAttribute("edges", st);
                walkRuleNFACreatingDOT(dot, rr.getFollowState());
                continue;
            }
            st = stlib.getInstanceOf("org/antlr/tool/templates/dot/edge");
            st.setAttribute("label", getEdgeLabel(edge.label.toString(grammar)));
            st.setAttribute("src", getStateLabel(s));
            st.setAttribute("target", getStateLabel(edge.target));
            dot.setAttribute("edges", st);
            walkRuleNFACreatingDOT(dot, edge.target); // keep walkin'
        }
    }

    /*
	public void writeDOTFilesForAllRuleNFAs() throws IOException {
        Collection rules = grammar.getRules();
        for (Iterator itr = rules.iterator(); itr.hasNext();) {
			Grammar.Rule r = (Grammar.Rule) itr.next();
            String ruleName = r.name;
            writeDOTFile(
                    ruleName,
                    getRuleNFADOT(grammar.getRuleStartState(ruleName)));
        }
    }
    */

    /*
	public void writeDOTFilesForAllDecisionDFAs() throws IOException {
        // for debugging, create a DOT file for each decision in
        // a directory named for the grammar.
        File grammarDir = new File(grammar.name+"_DFAs");
        grammarDir.mkdirs();
        List decisionList = grammar.getDecisionNFAStartStateList();
        if ( decisionList==null ) {
            return;
        }
        int i = 1;
        Iterator iter = decisionList.iterator();
        while (iter.hasNext()) {
            NFAState decisionState = (NFAState)iter.next();
            DFA dfa = decisionState.getDecisionASTNode().getLookaheadDFA();
            if ( dfa!=null ) {
                String dot = getDOT( dfa.startState );
                writeDOTFile(grammarDir+"/dec-"+i, dot);
            }
            i++;
        }
    }
    */

    /** Fix edge strings so they print out in DOT properly */
    protected String getEdgeLabel(String label) {
        label = Utils.replace(label,"\"", "\\\"");
        if ( label.equals(Label.EPSILON_STR) ) {
            label = "e";
        }
        return label;
    }

    protected String getStateLabel(State s) {
        if ( s==null ) {
            return "null";
        }
        String stateLabel = String.valueOf(s.stateNumber);
		if ( s instanceof DFAState ) {
            StringBuffer buf = new StringBuffer(250);
			buf.append('s');
			buf.append(s.stateNumber);
			if ( Tool.internalOption_ShowNFConfigsInDFA ) {
				buf.append("\\n");
				// separate alts
				Set alts = ((DFAState)s).getAltSet();
				List altList = new ArrayList();
				altList.addAll(alts);
				Collections.sort(altList);
				Set configurations = ((DFAState)s).getNFAConfigurations();
				for (int altIndex = 0; altIndex < altList.size(); altIndex++) {
					Integer altI = (Integer) altList.get(altIndex);
					int alt = altI.intValue();
					if ( altIndex>0 ) {
						buf.append("\\n");
					}
					buf.append("alt");
					buf.append(alt);
					buf.append(':');
					// get a list of configs for just this alt
					// it will help us print better later
					List configsInAlt = new ArrayList();
					for (Iterator it = configurations.iterator(); it.hasNext();) {
						NFAConfiguration c = (NFAConfiguration) it.next();
						if ( c.alt!=alt ) continue;
						configsInAlt.add(c);
					}
					int n = 0;
					for (int cIndex = 0; cIndex < configsInAlt.size(); cIndex++) {
						NFAConfiguration c =
							(NFAConfiguration)configsInAlt.get(cIndex);
						n++;
						buf.append(c.toString(false));
						if ( (cIndex+1)<configsInAlt.size() ) {
							buf.append(", ");
						}
						if ( n%5==0 && (configsInAlt.size()-cIndex)>3 ) {
							buf.append("\\n");
						}
					}
				}
			}
            stateLabel = buf.toString();
        }
		if ( (s instanceof NFAState) && ((NFAState)s).isDecisionState() ) {
			stateLabel = stateLabel+",d="+
					((NFAState)s).getDecisionNumber();
			if ( ((NFAState)s).endOfBlockStateNumber!=State.INVALID_STATE_NUMBER ) {
				stateLabel += ",eob="+((NFAState)s).endOfBlockStateNumber;
			}
		}
		else if ( (s instanceof NFAState) &&
			((NFAState)s).endOfBlockStateNumber!=State.INVALID_STATE_NUMBER)
		{
			stateLabel = stateLabel+",eob="+
					((NFAState)s).endOfBlockStateNumber;
		}
        else if ( s instanceof DFAState && ((DFAState)s).isAcceptState() ) {
            stateLabel = stateLabel+
                    "=>"+((DFAState)s).getUniquelyPredictedAlt();
        }
        return '"'+stateLabel+'"';
    }
}
