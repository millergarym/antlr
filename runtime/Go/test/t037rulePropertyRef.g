grammar t037rulePropertyRef;
options {language = Go;}
@header {package main;}

a returns [bla]
@after {
    $bla = $start, $stop, $text
}
    : A+
    ;

A: 'a'..'z';

WS: ' '+  { $channel = HIDDEN };
