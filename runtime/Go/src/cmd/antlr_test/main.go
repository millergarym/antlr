package main

import (
   "antlr3rt"
   "fmt"
)

func main() {
    T_test()
    t001lexer_test()
    t002lexer_test()
    t003lexer_test()
    t004lexer_test()
    t005lexer_test()
    t006lexer_test()
    t007lexer_test()
    t008lexer_test()
    t009lexer_test()
    t010lexer_test()
    t011lexer_test()
}

func T_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("A00\"ab0\"A01")
    var lexer antlr3rt.Lexer 
    lexer = TCreate( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")
}

func t001lexer_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("000")
    var lexer antlr3rt.Lexer 
    lexer = t001lexerCreate( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t002lexer_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("0110101")
    var lexer antlr3rt.Lexer 
    lexer = t002lexerCreate( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t003lexer_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("01fooze10fooze11fooze100")
    var lexer antlr3rt.Lexer 
    lexer = t003lexerCreate( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t004lexer_test() {
    var input antlr3rt.CharStream 
    input = antlr3rt.NewStringStream("ffofoo")
    var lexer antlr3rt.Lexer 
    lexer = t004lexerCreate( input ) 
    var tokenStream antlr3rt.TokenStream
    tokenStream = antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t005lexer_test() {
    input := antlr3rt.NewStringStream("fofofoo")
    lexer := t005lexerCreate( input ) 
    tokenStream := antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t006lexer_test() {
    input := antlr3rt.NewStringStream("fofafoa")
    lexer := t006lexerCreate( input )
    tokenStream := antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t007lexer_test() {
    input := antlr3rt.NewStringStream("fofoabfabfaboabb")
    lexer := t007lexerCreate( input )
    tokenStream := antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t008lexer_test() {
    input := antlr3rt.NewStringStream("ffa")
    lexer := t008lexerCreate( input )
    tokenStream := antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t009lexer_test() {
    input := antlr3rt.NewStringStream("1234567890")
    lexer := t009lexerCreate( input )
    tokenStream := antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t010lexer_test() {
    input := antlr3rt.NewStringStream("Hello World\nfinally got there\nw00t")
    lexer := t010lexerCreate( input )
    tokenStream := antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}

func t011lexer_test() {
    input := antlr3rt.NewStringStream("Hello World\nfinally got there\nw00t__")
    lexer := t011lexerCreate( input )
    tokenStream := antlr3rt.NewCommonTokenStream( lexer )
    fmt.Println( ":", tokenStream.GetTokens(), ":")    
}