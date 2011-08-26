package main

import (
   "antlr3rt"
   "fmt"
)

func main() {
    var input antlr3rt.CharStream
    input = antlr3rt.NewStringStream("00\"ab0\"A01")
    var lexer antlr3rt.Lexer 
    lexer = NewT( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")
    
}
