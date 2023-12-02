grammar calcit2;

GREATER: '>';
LESS: '<';
PLUS: '+';
MINUS: '-';
MULTIPLY: '*';
DIVIDE: '/';
MODULO: '%';
NOT: '!';
LPAREN: '(';
RPAREN: ')';
COLON: ':';
EQUALS: '=';
POW: '**';
COLON_EQUALS: ':=';
SEMICOLON: ';';
BIT_AND: '&';
BIT_OR: '|';
BIT_XOR: '^';
BIT_CLEAR: '&^';
LSHIFT: '<<';
RSHIFT: '>>';
LOGICAL_AND: '&&';
LOGICAL_OR: '||';
EQUAL: '==';
NOT_EQUAL: '!=';
LESS_OR_EQUAL: '<=';
GREATER_OR_EQUAL: '>=';
// Chars are expressed as integers
HEX_TYPE: 'hex';
BIN_TYPE: 'bin';
OCT_TYPE: 'oct';
F_TYPE: 'float';
I_TYPE: 'int';
B_TYPE: 'bool';
S_TYPE: 'string';
P_CAST: 'pcast'; //special case to pointer cast float into integer, does nothing if anything else is casted
FUNCTION_NAME: 'f' [a-zA-Z0-9_]*;
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
HEX_LITERAL: '0x' [0-9a-fA-F]+;
BIN_LITERAL: '0b' [01]+;
OCT_LITERAL: '0o' [0-7]+;
INTEGER_LITERAL: [0-9]+;
FLOAT_LITERAL: [0-9]+ '.' [0-9]+;
CHAR_LITERAL: '\'' . '\'';
WS: [ \t\r\n]+ -> skip; // Skip whitespaces

//WIP implement grammar flow later

// Parser rules
program: statement*;
statement: function_assignment | expression;
// math functions like f(x) = x + 1
function_assignment: FUNCTION_NAME LPAREN IDENTIFIER RPAREN EQUALS expression;
// expression like x + 1, any operation resulting in a bool belongs to the bool_expression rule
expression: integer_expression | float_expression | bool_expression;
// Main integer and float expression rules
number
    : integer_expression
    | float_expression;

any
    : number
    | bool_expression;

integer_expression
    : LPAREN right=integer_expression RPAREN
    | integer_cast_expression
    | integer_pcast_expression
    | op=MINUS right=integer_expression
    | op=NOT right=integer_expression
    | left=integer_expression op=POW right=integer_expression
    | left=integer_expression op=(MULTIPLY | DIVIDE | MODULO) right=integer_expression
    | left=integer_expression op=(PLUS | MINUS) right=integer_expression
    | left=integer_expression op=(BIT_AND | BIT_OR | BIT_XOR | BIT_CLEAR | LSHIFT | RSHIFT) right=integer_expression
    | (INTEGER_LITERAL | HEX_LITERAL | BIN_LITERAL | OCT_LITERAL | IDENTIFIER);

integer_cast_expression: op=(HEX_TYPE | BIN_TYPE | OCT_TYPE | I_TYPE | F_TYPE) LPAREN right=number RPAREN;

integer_pcast_expression: op=P_CAST LPAREN right=float_expression RPAREN;

float_expression
    : LPAREN right=float_expression RPAREN
    | float_cast_expression
    | float_pcast_expression
    | op=MINUS right=float_expression
    | left=float_expression op=POW right=float_expression
    | left=float_expression op=(MULTIPLY | DIVIDE) right=float_expression
    | left=float_expression op=(PLUS | MINUS) right=float_expression
    | (FLOAT_LITERAL | IDENTIFIER);

float_cast_expression: op=F_TYPE LPAREN right=number RPAREN;

float_pcast_expression: op=P_CAST LPAREN right=integer_expression RPAREN;

bool_expression
    : LPAREN right=bool_expression RPAREN
    | bool_cast_expression
    | bool_number_operation
    | left=bool_expression op=(EQUAL | NOT_EQUAL) right=bool_expression
    | left=bool_expression op=(LOGICAL_AND | LOGICAL_OR) right=bool_expression
    // Cases for floats and intergers
    | op=NOT right=bool_expression
    | ('true' | 'false');

bool_cast_expression: op=B_TYPE LPAREN right=any RPAREN;

bool_number_operation: left=number op=(EQUAL | NOT_EQUAL | LESS | GREATER | LESS_OR_EQUAL | GREATER_OR_EQUAL) right=number;