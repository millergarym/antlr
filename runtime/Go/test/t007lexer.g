lexer grammar t007lexer;
options {language = Go;}
@header {package main;}

FOO: 'f' ('o' | 'a' 'b'+)*;
