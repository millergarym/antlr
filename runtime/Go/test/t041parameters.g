grammar t041parameters;
options {language = Go;}
@header {package main;}

a[arg1, arg2] returns [l]
    : A+ EOF
        { 
            l = ($arg1, $arg2) 
            $arg1 = "gnarz"
        }
    ;

A: 'a'..'z';

WS: ' '+  { $channel = HIDDEN };
