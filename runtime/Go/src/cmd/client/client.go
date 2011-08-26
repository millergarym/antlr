package main

import (
    . "antlr3rt"   
)

type client struct {
    Lexer
//    lexer Lexer
}

func newC() *client {
    c := &client{
        Lexer: new(LexerImpl),
    }
    return c 
}

func (self *client) Test2() {
    return
}

