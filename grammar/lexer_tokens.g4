lexer grammar lexer_tokens;

IMPORT: 'import';
USE: 'use';
PACKAGE: 'package';
TYPE: 'type';
CONST: 'const';
SERVICE: 'service';
GATEWAY: 'gateway';
BACKEND: 'backend';

// Types 
INT: 'int';
LONG: 'int64';
SHORT: 'int16';
BYTE: 'byte';
BOOL: 'bool';
STRING: 'string';
DATETIME: 'time';
UNKNOWN: 'struct';


// API Methods
GET: 'get';
POST: 'post';
PUT: 'put';
PATCH: 'patch';
DELETE: 'delete';

// Constants 
DATA: '$data';

// Symbols
EQ: '=';
COLON: ':';
DOUBLE_QUOTE: '"';
GT: '>';
LT: '<';
LP: '(';
RP: ')';
LCB: '{';
RCB: '}';
SEMI: ';';
AT: '@';

// Patterns
IDENT: [A-Za-z_]([A-Za-z0-9_] | ('.')+)*;
NUMBER: [0-9]+('.'[0-9]+)*;
NEWLINE: [\r\n]+ -> skip;
WS: [ \t\r\n]+ -> skip;