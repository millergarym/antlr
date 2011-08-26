lexer grammar T;
options { language = Go; }
@header {
package main;
}

ZEROZERO: '00';
ZERO: '0';
ONE: '1';
WORD: ('a'..'z'|'A'..'Z')+;
QUOTE:  '"';
STRING: QUOTE ~(QUOTE)* QUOTE;
