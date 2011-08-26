package antlr3rt

import (
    "fmt"
    "os"
//    "reflect"
)

func trace(s string) string {
    fmt.Println( "->", s )
    return s
}
func un(s string) {
    fmt.Println( "<-", s )
}

const (
    DEFAULT_TOKEN_CHANNEL uint8 = 1
    DEFAULT_CHANNEL = 0
)

//type ItemType int16
type RecognitionException interface {
    GetUnexpectedType() int
}
type recognitionException struct {
}
var _ RecognitionException = &recognitionException{}
func (s *recognitionException) GetUnexpectedType() (ret int) {
    return
}

type MismatchedTokenException struct {
    *recognitionException
    expected int
}
type MismatchedRangeException struct {
    *recognitionException
    a,b int
}
type NoViableAltException struct {}
type EarlyExitException struct {}

type State struct {
    mType int
    channel uint8
    backtracking int
    failed bool
    token Token
    tokenStartCharIndex int
    tokenStartCharPositionInLine int
    tokenStartLine int
    text *string 
}
func (self *State) SetType(i int) {
    self.mType = i
}
func (self *State) SetChannel(i uint8) {
    self.channel = i
}

type Lexer interface {
    Test()
    State() *State
    Input() *CharStream
    MatchChar( int )    
    MatchString( string )
    MatchRange(int,int)
    Recover()// func()
    ReportError( interface{} )
    NextToken() Token
    MTokens()
    Emit() Token
    SetClient(*Lexer)
}

type LexerImpl struct {
    state State
    input CharStream
    client *Lexer
}
var _ Lexer = &LexerImpl{}

func NewLexerImpl(input CharStream) (ret *LexerImpl) {
    defer un(trace("NewLexerImpl"))
    return &LexerImpl{
        state: *new(State),
        input: input,
    }
}


func (self *LexerImpl) SetClient(c *Lexer) {
   defer un(trace("lexer impl.set client"))    
   self.client = c
}

func (self *LexerImpl) MTokens() {
    defer un(trace("lexer impl.mtokens"))    
    panic( "Implment MToken() *func() in client" )
    return
}

func (self *LexerImpl) ReportError( iface interface{} ) {
    defer un(trace("lexer impl.report error"))
    fmt.Fprintln(os.Stderr, "Lexer error")    
}

func (self *LexerImpl) Test() {
    fmt.Println("Antlr3rt")
}
func (self *LexerImpl) Recover() {
    defer un(trace("lexer impl.Recover"))    
//    defer func () {
//        if x := recover(); x != nil {
//            fmt.Println("run time panic: %v", x)
//            self.input.consume()
//        }
//    }()
}
func (self *LexerImpl) State() (ret *State) {
    defer un(trace("lexer impl.state"))    
    return &self.state
}
func (self *LexerImpl) Input() (ret *CharStream) {
    defer un(trace("lexer impl.state"))    
    return &self.input
}
 
func (self *LexerImpl) MatchChar(c int) {
    defer un(trace("lexer impl.match char"))    
    if self.input.LA(1) != c {
        if self.state.backtracking > 0 {
            self.state.failed = true;
            return
        }
        panic( &MismatchedTokenException{expected:c} )
    }
    self.input.Consume()
    self.state.failed = false
}

func (self *LexerImpl) MatchRange(a,b int) {
    if self.input.LA(1) < a || self.input.LA(1) > b {
        if self.state.backtracking > 0 {
            self.state.failed = true;
            return;
        }
        panic( &MismatchedRangeException{a:a,b:b} )
    }
    self.input.Consume()
    self.state.failed = false    
}

func (self *LexerImpl) MatchString(s string) {
    defer un(trace("lexer impl.match string"))    
    i := 0
    for i < len(s) {
        if self.input.LA(1) != int(s[i]) { // TODO(garym) type convert?
            if self.state.backtracking > 0 {
                self.state.failed = true
                return
            }
            panic( &MismatchedTokenException{expected:int(s[i])} ) // TODO(garym) type convert?
//            recover(mte);
        }
        i++
        self.input.Consume()
        self.state.failed = false
    }
}

func (self *LexerImpl) Emit() (ret Token) {
    defer un(trace("lexer impl.emit"))    
    t := &commonToken{
           input: &self.input, 
           ttype: self.state.mType, 
           channel: self.state.channel, 
           start: self.state.tokenStartCharIndex,
           stop: self.input.getCharIndex(), 
//           index: // set in fillBuffer
           charPos: self.input.getCharPositionInLine(),
           line: self.input.getLine(),
    }
//    t.setLine(state.tokenStartLine);
//    t.setText(state.text);
//    t.setCharPositionInLine(state.tokenStartCharPositionInLine);
    self.state.token = t
    return t
}

func (self *LexerImpl) NextToken() (ret Token) {
    defer un(trace("LexerImpl NextToken"))
    defer func () {
        if x := recover(); x != nil {
           if _, ok := x.(*RecognitionException); ok {
               self.ReportError( x )
               self.Recover() //?
           } else if _, ok := x.(*NoViableAltException); ok {
               self.ReportError( x )
               self.Recover()
           } else {
               panic(x)
           }
        }
    }()
    for (true) {
        self.state.token = nil
        self.state.channel = DEFAULT_CHANNEL
        self.state.tokenStartCharIndex = self.input.index()
        self.state.tokenStartCharPositionInLine = self.input.getCharPositionInLine()
        self.state.tokenStartLine = self.input.getLine()
        self.state.text = nil
        if ( self.input.LA(1) == EOF ) {
            return EOF_TOKEN
        }
        (*self.client).MTokens()
        if ( self.state.token == nil ) {
            self.Emit()
        } else if  self.state.token == SKIP_TOKEN {
            continue
        }
        return self.state.token
    }
    panic( "can't get here, return is for loop" )
    return nil
}





