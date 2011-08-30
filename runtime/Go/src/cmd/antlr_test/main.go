package main

import (
   "antlr3rt"
   "fmt"
)

func main() {
    T_test()
    t001lexer_test()
    t002lexer_test()
}

func T_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("A00\"ab0\"A01")
    var lexer antlr3rt.Lexer 
    lexer = NewT( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")
}

func t001lexer_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("000")
    var lexer antlr3rt.Lexer 
    lexer = Newt001lexer( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t002lexer_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("0110101")
    var lexer antlr3rt.Lexer 
    lexer = Newt002lexer( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}