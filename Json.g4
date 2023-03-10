grammar Json;

json: object
    | array
    ;

object: '{' pair (',' pair)* '}'
    | '{' '}'
    ;
pair: STRING ':' value;

array: '[' value(',' value)* ']'
    | '[' ']'
    ;

value:
    STRING
    | NUMBER
    | object
    | array
    | 'true'
    | 'false'
    | 'null'
    ;

STRING: '"' (ESC | ~["\\])* '"';
NUMBER: '-'? INT '.' INT EXP?
    | '-'? INT EXP
    | '-'? INT
    ;

fragment ESC : '\\' (["\\/bfnrt]| UNICODE);
fragment UNICODE: 'u' HEX HEX HEX ;
fragment HEX : [0-9a-fA-F];
fragment INT: '0' | [1-9][0-9]*;
fragment EXP: [Ee] [+\-]? INT;

WS : [ \t\n\r]+ -> skip;