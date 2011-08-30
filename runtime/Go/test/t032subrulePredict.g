grammar t032subrulePredict;
options {language = Go;}
@header {package main;}

a: 'BEGIN' b WS+ 'END';
b: ( WS+ 'A' )+;
WS: ' ';
