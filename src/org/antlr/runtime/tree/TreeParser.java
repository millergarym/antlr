package org.antlr.runtime.tree;

import org.antlr.runtime.BaseRecognizer;

/** A parser for a stream of tree nodes.  "tree grammars" result in a subclass
 *  of this.  All the error reporting and recovery is shared with Parser via
 *  the BaseRecognizer superclass.
*/
public class TreeParser extends BaseRecognizer {
	protected TreeNodeStream input;

	public TreeParser(TreeNodeStream input) {
		setTreeNodeStream(input);
	}

	/** Set the input stream and reset the parser */
	public void setTreeNodeStream(TreeNodeStream input) {
		this.input = input;
		reset();
	}

	public TreeNodeStream getTreeNodeStream() {
		return input;
	}
}
