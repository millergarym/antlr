lexer grammar t011lexer;
options {language = Go;}
@header {package main;

import . "fmt"
}

IDENTIFIER: 
        ('a'..'z'|'A'..'Z'|'_') 
        ('a'..'z'
        |'A'..'Z'
        |'0'..'9'
        |'_'
            { 
              Println( "Underscore" ) 
              Println( "foo" )
            }
        )*
    ;

WS: (' ' | '\n')+;
