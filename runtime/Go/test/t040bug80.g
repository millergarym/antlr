lexer grammar t040bug80; 
options {language = Go;}
@header {package main;}
 
ID_LIKE
    : 'defined' 
    | {False}? Identifier 
    | Identifier 
    ; 
 
fragment 
Identifier: 'a'..'z'+ ; // with just 'a', output compiles 
