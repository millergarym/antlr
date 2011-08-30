lexer grammar t029synpredgate;
options {language = Go;}
@header {package main;}

FOO
    : ('ab')=> A
    | ('ac')=> B
    ;

fragment
A: 'a';

fragment
B: 'a';

