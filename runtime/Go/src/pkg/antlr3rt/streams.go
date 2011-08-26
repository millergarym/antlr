package antlr3rt

import (
    "fmt"
)

const (
    EOF = -1
    INVALID_TOKEN_TYPE = 0
)

var (
    EOF_TOKEN Token = &commonToken{ttype: EOF,}
    SKIP_TOKEN Token = &commonToken{ttype: INVALID_TOKEN_TYPE,}
    _ Token = &commonToken{}
)

type Token interface {
    getType() int
    setTokenIndex(int)
}

type CharStream interface {
    LA(int) int 
    Consume()
    index() int
    getCharIndex() int
    getCharPositionInLine() int
    getLine() int
    substring(start int, stop int) string
}

type TokenStream interface {
    GetTokens() []Token
    nextToken() Token
}

type TokenSource interface {
    NextToken() Token
}

type commonToken struct {
    input *CharStream
    ttype int
    index int
    start, stop int
    charPos, line int
    channel uint8
}
func (s *commonToken) String() string {
    return fmt.Sprintf("<typ:%v,idx:%v,'%v',pos[%v:%v],chan:%v>", s.ttype, s.index, (*s.input).substring(s.start,s.stop), s.line, s.charPos,s.channel)
}
func (self *commonToken) getType() (ret int) {
    defer un(trace("getType"))
    return self.ttype
}
func (self *commonToken) setTokenIndex(i int) {
    defer un(trace("setTokenIndex"))
    self.index = i
}

type stringStream struct {
    data string
    // How many characters are actually in the buffer 
    n int
    // 0..n-1 index into string of next char 
    p int
    // line number 1..n within the input
    line int
    // The index of the character relative to the beginning of the line 0..n-1 
    charPositionInLine int
/*    
    // tracks how deep mark() calls are nested
    markDepth int
    // A list of CharStreamState objects that tracks the stream state
    //  values line, charPositionInLine, and p that can change as you
    //  move through the input stream.  Indexed from 1..markDepth.
    //  A null is kept @ index 0.  Create upon first call to mark().
    markers []int
    // Track the last mark() call result value for use in rewind().
    lastMarker int
    // What is name or source of this char stream?
    name string
*/
}

var _ CharStream = &stringStream{}

type commonTokenStream struct {
    p int
    tokens []Token
    tokenSource TokenSource    
}

func (self *stringStream) substring(start int, stop int) (ret string) {
    return self.data[start:stop]
}

func (self *commonTokenStream) nextToken() (ret Token) {
    defer un(trace("nextToken"))
    panic("TODO commonTokenStream.nextToken")
    return
}

func NewStringStream(data string) (ret *stringStream ) {
    defer un(trace("new string stream"))
    return &stringStream{
        data: data,
        line: 1,
        n: len(data),
    }
}

func NewCommonTokenStream(lexer Lexer) (ret *commonTokenStream) {
    defer un(trace("new common token stream"))
    return &commonTokenStream {
        p: -1,
        tokenSource: lexer,
    }
}

func (s *stringStream) LA(i int) (ret int) {
    defer un(trace("LA"))
    if i == 0 {
        return 0; // undefined
    }
    if i < 0 {
        i++ // e.g., translate LA(-1) to use offset i=0; then data[p+0-1]
        if ( (s.p+i-1) < 0 ) {
            return EOF // invalid; no char before first char
        }
    }
    if (s.p+i-1) >= s.n {
        return EOF
    }
    return int(s.data[s.p+i-1])
}
func (s *stringStream) Consume() {
    if s.p < s.n {
        s.charPositionInLine++
        if s.data[s.p] == '\n' {
            s.line++
            s.charPositionInLine=0
        }
        s.p++
    }
}
func (s *stringStream) index() (ret int) {
    defer un(trace("stringStream.index"))
    return s.p
}
func (s *stringStream) getCharIndex() (ret int) {
    defer un(trace("stringStream.get char index"))
    return s.p
}
func (s *stringStream) getCharPositionInLine() (ret int) {
    defer un(trace("stringStream.get char pos in line"))
    return s.charPositionInLine
}
func (s *stringStream) getLine() (ret int) {
    defer un(trace("stringStream.get line"))
    return s.line
}

func (s *commonTokenStream) GetTokens() (ret []Token) {
    defer un(trace("common token stream.get tokens"))    
    if s.p == -1  {
        s.fillBuffer()
    }
    return s.tokens
}

func  (self *commonTokenStream) fillBuffer() {
    defer un(trace("common token stream.fill buffer"))    
    index := 0
    t := self.tokenSource.NextToken()
    for t != nil && t.getType() != EOF {
        discard := false
        // is there a channel override for token type?
//        if ( channelOverrideMap!=null ) {
//            Integer channelI = (Integer)
//                channelOverrideMap.get(new Integer(t.getType()));
//            if ( channelI!=null ) {
//                t.setChannel(channelI.intValue());
//            }
//        }
//        if ( discardSet!=null &&
//             discardSet.contains(new Integer(t.getType())) )
//        {
//            discard = true;
//        }
//        else if ( discardOffChannelTokens && t.getChannel()!=this.channel ) {
//            discard = true;
//        }
        if !discard {
            t.setTokenIndex(index)
            self.tokens = append( self.tokens, t)
            index++
        }
        t = self.tokenSource.NextToken()
    }
    // leave p pointing at first token on channel
    self.p = 0
    self.p = self.skipOffTokenChannels(self.p)    
}

func (self *commonTokenStream) skipOffTokenChannels(i int) (ret int) {
    defer un(trace("common token stream.skip off token channels"))    
    return 0
}
