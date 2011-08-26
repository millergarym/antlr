// $ANTLR 3.4 antlr/test/T.g 2011-08-26 10:24:54

package main;


import "antlr3rt"
//import "fmt"

type T struct {
    antlr3rt.Lexer
}

const (
    EOF = -1
    ONE = 4
    QUOTE = 5
    STRING = 6
    WORD = 7
    ZERO = 8
    ZEROZERO = 9

)


// delegates
// delegators

func NewT(input antlr3rt.CharStream) *T {
    lex := antlr3rt.NewLexerImpl(input)
    var self antlr3rt.Lexer = &T {
      antlr3rt.Lexer: lex,
    }
    lex.SetClient( &self )
    return self.(*T);
}

func GetGrammarFileName() string { 
    return "antlr/test/T.g"
}
// $ANTLR start "ZEROZERO"
func (self *T) mZEROZERO() {
 
    
		_type := ZEROZERO
		_channel := antlr3rt.DEFAULT_TOKEN_CHANNEL
		// antlr/test/T.g:7:9: ( '00' )
		// antlr/test/T.g:7:11: '00'
		{
		self.MatchString("00"); 



		}

		self.State().SetType( _type )
		self.State().SetChannel( _channel )
}
// $ANTLR end "ZEROZERO"

// $ANTLR start "ZERO"
func (self *T) mZERO() {
 
    
		_type := ZERO
		_channel := antlr3rt.DEFAULT_TOKEN_CHANNEL
		// antlr/test/T.g:8:5: ( '0' )
		// antlr/test/T.g:8:7: '0'
		{
		self.MatchChar('0') 


		}

		self.State().SetType( _type )
		self.State().SetChannel( _channel )
}
// $ANTLR end "ZERO"

// $ANTLR start "ONE"
func (self *T) mONE() {
 
    
		_type := ONE
		_channel := antlr3rt.DEFAULT_TOKEN_CHANNEL
		// antlr/test/T.g:9:4: ( '1' )
		// antlr/test/T.g:9:6: '1'
		{
		self.MatchChar('1') 


		}

		self.State().SetType( _type )
		self.State().SetChannel( _channel )
}
// $ANTLR end "ONE"

// $ANTLR start "WORD"
func (self *T) mWORD() {
 
    
		_type := WORD
		_channel := antlr3rt.DEFAULT_TOKEN_CHANNEL
		// antlr/test/T.g:10:5: ( ( 'a' .. 'z' | 'A' .. 'Z' )+ )
		// antlr/test/T.g:10:7: ( 'a' .. 'z' | 'A' .. 'Z' )+
		{
		// antlr/test/T.g:10:7: ( 'a' .. 'z' | 'A' .. 'Z' )+
		cnt1 := 0
		loop1:
		for {
		    alt1 := 2
		    LA1_0 := (*self.Input()).LA(1)

		    if ( ((LA1_0 >= 'A' && LA1_0 <= 'Z')||(LA1_0 >= 'a' && LA1_0 <= 'z')) ) {
		        alt1=1
		    }


		    switch (alt1) {
			case 1 :
			    // antlr/test/T.g:
			    {
			    if ((*self.Input()).LA(1) >= 'A' && (*self.Input()).LA(1)<='Z')||((*self.Input()).LA(1) >= 'a' && (*self.Input()).LA(1)<='z') {
			        (*self.Input()).Consume()
			    } else {
			        panic( &antlr3rt.MismatchedTokenException{} )
			    }


			    }

		    
			default :
			    if cnt1 >= 1 {
			       break loop1;
			    }
		      panic( &antlr3rt.EarlyExitException{} )
		    }
		    cnt1++;
		}


		}

		self.State().SetType( _type )
		self.State().SetChannel( _channel )
}
// $ANTLR end "WORD"

// $ANTLR start "QUOTE"
func (self *T) mQUOTE() {
 
    
		_type := QUOTE
		_channel := antlr3rt.DEFAULT_TOKEN_CHANNEL
		// antlr/test/T.g:11:6: ( '\"' )
		// antlr/test/T.g:11:9: '\"'
		{
		self.MatchChar('"') 


		}

		self.State().SetType( _type )
		self.State().SetChannel( _channel )
}
// $ANTLR end "QUOTE"

// $ANTLR start "STRING"
func (self *T) mSTRING() {
 
    
		_type := STRING
		_channel := antlr3rt.DEFAULT_TOKEN_CHANNEL
		// antlr/test/T.g:12:7: ( QUOTE (~ ( QUOTE ) )* QUOTE )
		// antlr/test/T.g:12:9: QUOTE (~ ( QUOTE ) )* QUOTE
		{
		self.mQUOTE() 


		// antlr/test/T.g:12:15: (~ ( QUOTE ) )*
		loop2:
		for {
		    alt2 := 2
		    LA2_0 := (*self.Input()).LA(1)

		    if ( ((LA2_0 >= '\u0000' && LA2_0 <= '!')||(LA2_0 >= '#' && LA2_0 <= '\uFFFF')) ) {
		        alt2=1
		    }


		    switch (alt2) {
			case 1 :
			    // antlr/test/T.g:
			    {
			    if ((*self.Input()).LA(1) >= '\u0000' && (*self.Input()).LA(1)<='!')||((*self.Input()).LA(1) >= '#' && (*self.Input()).LA(1)<='\uFFFF') {
			        (*self.Input()).Consume()
			    } else {
			        panic( &antlr3rt.MismatchedTokenException{} )
			    }


			    }


			default :
			    break loop2;
		    }
		}


		self.mQUOTE() 


		}

		self.State().SetType( _type )
		self.State().SetChannel( _channel )
}
// $ANTLR end "STRING"

func (self *T) MTokens() {
    // antlr/test/T.g:1:8: ( ZEROZERO | ZERO | ONE | WORD | QUOTE | STRING )
    alt3 := 6
    switch ( (*self.Input()).LA(1) ) {
    case '0':
        {
        LA3_1 := (*self.Input()).LA(2)

        if ( (LA3_1=='0') ) {
            alt3=1
        } else {
            alt3=2

        }
        }
    case '1':
        {
        alt3=3
        }
    case 'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z','a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z':
        {
        alt3=4
        }
    case '"':
        {
        LA3_4 := (*self.Input()).LA(2)

        if ( ((LA3_4 >= '\u0000' && LA3_4 <= '\uFFFF')) ) {
            alt3=6
        } else {
            alt3=5

        }
        }
    default:
        panic( &antlr3rt.NoViableAltException{} )
    //   NoViableAltException nvae =
    //        new NoViableAltException("", 3, 0, input);

    }

    switch (alt3) {
        case 1 :
            // antlr/test/T.g:1:10: ZEROZERO
            {
            self.mZEROZERO() 


            }

        case 2 :
            // antlr/test/T.g:1:19: ZERO
            {
            self.mZERO() 


            }

        case 3 :
            // antlr/test/T.g:1:24: ONE
            {
            self.mONE() 


            }

        case 4 :
            // antlr/test/T.g:1:28: WORD
            {
            self.mWORD() 


            }

        case 5 :
            // antlr/test/T.g:1:33: QUOTE
            {
            self.mQUOTE() 


            }

        case 6 :
            // antlr/test/T.g:1:39: STRING
            {
            self.mSTRING() 


            }


    }

}


 
