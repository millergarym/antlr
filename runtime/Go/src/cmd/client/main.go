package main

import (
   "antlr3rt"
)

func main() {
    var lex antlr3rt.Lexer
    c := newC()
    lex = c
    lex.Test()  
}