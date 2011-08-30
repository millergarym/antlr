lexer grammar t006lexer;
options {language = Go;}
@header {package main;}

FOO: 'f' ('o' | 'a')*;
