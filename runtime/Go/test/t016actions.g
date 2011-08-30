grammar t016actions;
options {language = Go;}
@header {package main;}

declaration returns [name]
    :   functionHeader ';'
        {$name = $functionHeader.name}
    ;

functionHeader returns [name]
    :   type ID
	{$name = $ID.text}
    ;

type
    :   'int'   
    |   'char'  
    |   'void'
    ;

ID  :   ('a'..'z'|'A'..'Z'|'_') ('a'..'z'|'A'..'Z'|'0'..'9'|'_')*
    ;

WS  :   (   ' '
        |   '\t'
        |   '\r'
        |   '\n'
        )+
        {$channel=HIDDEN}
    ;    
