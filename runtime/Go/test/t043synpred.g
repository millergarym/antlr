grammar t043synpred;
options {language = Go;}
@header {package main;}

a: ((s+ P)=> s+ b)? E;
b: P 'foo';

s: S;


S: ' ';
P: '+';
E: '>';
