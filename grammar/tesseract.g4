grammar tesseract;
import lexer_tokens;

EscapedString : '"' ( '""' | ~["\r\n] )* '"';

program: packageStatement (importStatement | useStatement | typeStatement | constStatement | serviceStatement | backendStatement | gatewayStatement)*;
importStatement: IMPORT EscapedString SEMI;
useStatement: USE IDENT SEMI;
packageStatement: PACKAGE IDENT SEMI;
typeStatement: attribute* TYPE LP (attribute* IDENT LCB field (',' field)* RCB)* RP SEMI*;
constStatement: CONST LP (IDENT EQ (EscapedString | NUMBER))* RP; 
serviceStatement: attribute* SERVICE LP rpc* RP;
gatewayStatement: attribute* GATEWAY LP api* RP;
backendStatement: attribute* BACKEND LP (IDENT EQ GT call) RP;

type : IDENT | INT | LONG | SHORT | BYTE | BOOL | STRING | DATETIME | UNKNOWN;
list: '[]'type;
fieldType: (type | list);
field: (attribute)* IDENT fieldType;
assignment: IDENT EQ (EscapedString | NUMBER| | IDENT | call);
arg: (assignment | NUMBER | EscapedString | IDENT | DATA | call);
call:  IDENT LP arg* (',' arg)* RP;
attribute: AT call;
request: attribute* IDENT;
response: attribute* IDENT;
rpc: attribute* IDENT LP request RP response SEMI;
api: attribute* (GET | POST | PUT | PATCH | DELETE) IDENT LP request* RP response* SEMI;

