// BabyDuck.g4
grammar BabyDuck;

//Tokens

// Reserved Keywords
PROGRAM: 'program';
MAIN: 'main';
END: 'end';
VAR: 'var';
PRINT: 'print';
IF: 'if';
ELSE: 'else';
WHILE: 'while';
DO: 'do';
VOID: 'void';
INT: 'int';
FLOAT: 'float';
STRING: 'string';

// Arithmetic Operators
OP_MULTIPLY: '*';
OP_DIVIDE: '/';
OP_ADD: '+';
OP_SUBTRACT: '-';

// Relational Operators
GREATER: '>';
LESS: '<';
NOT_EQUAL: '!=';

// Assignment Operator
ASSIGN: '=';

// Delimiters
LPAREN: '(';
RPAREN: ')';
LBRACKET: '[';
RBRACKET: ']';
LBRACE: '{';
RBRACE: '}';

// Punctuation
COLON: ':';
COMMA: ',';
SEMICOLON: ';';

// Constants
CONST_INT: [0-9]+;
CONST_FLOAT: [0-9]*[.][0-9]+;
CONST_STRING:'"' ~('"')* '"';

WHITESPACE: [ \r\n\t]+ -> skip;

// Identifier
ID: [a-zA-Z_][a-zA-Z0-9_]*;

// Rules
program: programHeader (varsSection)? (functionsSection)* programBody;
programHeader: PROGRAM ID SEMICOLON;
programBody: MAIN codeBlock END;

identifier: ID;
parenthesizedExpression: LPAREN expression RPAREN;

// Type Definitions
type: INT | FLOAT;

// Variable Declarations (VARS)
idList: ID (COMMA ID)*;
varDecl: idList COLON type SEMICOLON;
varsSection: VAR varDecl+;

// Constants (CTE)
constant: CONST_INT | CONST_FLOAT;

// Function Declarations (FUNCS)
parameter: identifier COLON type;
functionBody: LBRACKET (varsSection)? codeBlock RBRACKET SEMICOLON;
functionDeclaration: VOID identifier LPAREN (parameter (COMMA parameter)*)? RPAREN functionBody;
functionsSection: functionDeclaration;

// Code Block Structure (BODY)
codeBlock: LBRACE (statement)* RBRACE;

// Statement Types (STATEMENT)
statement: assignment | conditional | loop | functionCall | printStatement;

// Assignment Statement (ASSIGN)
assignment: identifier ASSIGN expression SEMICOLON;

// Conditional Statement (CONDITION)
ifBlock: IF parenthesizedExpression codeBlock;
elseBlock: ELSE codeBlock;
conditional: ifBlock (elseBlock)? SEMICOLON;

// Loop Statement (CYCLE)
loop: WHILE parenthesizedExpression DO codeBlock SEMICOLON;

// Function Call (F_CALL)
argumentList: expression (COMMA expression)*;
functionCall: identifier LPAREN (argumentList)? RPAREN SEMICOLON;

// Print Statement (PRINT)
printable: expression | CONST_STRING;
printStatement: PRINT LPAREN printable (COMMA printable)*  RPAREN SEMICOLON;

// Expression (EXPRESSION)
expression: arithmeticExpression (relationalExpression)?;
relationalOperator: GREATER | LESS | NOT_EQUAL;
relationalExpression: relationalOperator arithmeticExpression;

// Arithmetic Expressions (EXP)
additiveOperator: OP_ADD | OP_SUBTRACT;
arithmeticExpression: term (additiveOperator term)*;

// Term Expressions (TERM)
multiplicativeOperator: OP_MULTIPLY | OP_DIVIDE;
term: factor (multiplicativeOperator factor)*;

// Factor Expressions (FACTOR)
factor: parenthesizedExpression | valueWithOptionalSign;
valueWithOptionalSign: (additiveOperator)? value;
value: identifier | constant;
