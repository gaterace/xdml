grammar DML;

/*
Copyright 2017-2019 Demian Harvill
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

tokens 
{
}

@header 
{
// package org.gaterace.dml;
}

@members 
{

}

DML : 'dml';
DMLFIELDS : 'dmlfields';
DMLFIELD : 'dmlfield';
DMLINPUTS  : 'dmlinputs';
DMLOUTPUTS  : 'dmloutputs';
DMLQUALID : 'dmlqualid';
ENUM :  'enum' ;
PACKAGE : 'package' ;
IMPORT : 'import';
FIELDSET : 'fieldset' ;
CLASS : 'class';
TABLE : 'table';
ASSOCIATION : 'association';
SERVICE : 'service';
MEDIATES : 'mediates';
METHOD : 'method';
RETURNS : 'returns';
DOUBLE : 'double';
FLOAT : 'float';
INT32 : 'int32';
INT64 : 'int64';
UINT32 : 'uint32';
UINT64 : 'uint64';
BOOL : 'bool';
STRING : 'string';
BYTES : 'bytes';
CHARARRAY : 'chararray';
BYTEARRAY : 'bytearray';
DATETIME : 'datetime';
GUID : 'guid';
DECIMAL : 'decimal';
AUTO : 'auto';
ID_KEYWORD : 'id';
IDGEN : 'idgen';
REQUIRED : 'required';
OPTIONAL : 'optional';
REPEATED : 'repeated';
VIRTUAL : 'virtual';
MAP : 'map';
LIST : 'list';
REFERENCES : 'references';
INDEX : 'index';
PRIMARY : 'primary';
UNIQUE  : 'unique';
NONUNIQUE : 'nonunique';
PATTERN : 'pattern';
INSERT : 'insert';
UPDATE : 'update';
SELECT : 'select';
DELETE : 'delete';
EXTENDS : 'extends';



INTEGER_NUM		: ('0'..'9')+ ;

SEMI	: ';' ;
COMMA   : ',' ;
COLON	: ':' ;

DOT	: '.' ;
PLUS : '+' ;
MINUS : '-' ;
RPAREN	: ')' ;
LPAREN	: '(' ;
RBRACK	: ']' ;
LBRACK	: '[' ;
LBRACE  : '{' ;
RBRACE  : '}' ;
EQUALS  : '=' ;

WHITE_SPACE	: ( ' '|'\r'|'\t'|'\n' ) -> skip ;

fragment HEX_DIGIT_FRAGMENT: ( 'a'..'f' | 'A'..'F' | '0'..'9' ) ;
HEX_DIGIT:
	(  '0x'     (HEX_DIGIT_FRAGMENT)+  )
	|
	(  'X' '\'' (HEX_DIGIT_FRAGMENT)+ '\''  ) 
;

REAL_NUMBER:
	(  INTEGER_NUM DOT INTEGER_NUM | INTEGER_NUM DOT | DOT INTEGER_NUM | INTEGER_NUM  )
	(  ('E'|'e') ( PLUS | MINUS )? INTEGER_NUM  )? 
;

TEXT_STRING
    : (  '\''  ~('\'')* '\''  )
    | (  '\"'  ~('\"')* '\"'  )
;

ID
    : ( 'A'..'Z' | 'a'..'z' ) ( 'A'..'Z' | 'a'..'z' | '_' | '0'..'9' )*
;

COMMENT
	: '/*' .*? '*/' -> skip;
	

KEEP_COMMENT
    : '///' ~('\r' | '\n')* ('\r' | '\n')+
    ;

LINE_COMMENT
	: '//' ~('/') ~('\r' | '\n')* ('\r' | '\n')+ -> skip;
	

model
    : modelParts // {System.out.println($modelParts.tree.toStringTree());}
    ;

modelParts
	: packageElement importElements* sourceElements EOF 
	;

	
sourceElements
	: sourceElement (sourceElement)*
	;

sourceElement
    :  enumElement
    |  fieldsetElement
    |  classElement
    |  associationElement
    |  serviceElement
    ;

enumElement
    : KEEP_COMMENT* ENUM ID  LBRACE enumValues RBRACE 
    ;

enumValues
    : enumValue+ 
    ;

integerConst : (PLUS | MINUS)? INTEGER_NUM;

enumValue
    : KEEP_COMMENT* ID (EQUALS integerConst)? SEMI 
    ;

fieldsetElement
    : KEEP_COMMENT* FIELDSET ID  (extendsQualifier)? LBRACE fieldsetValues RBRACE 
    ;

extendsQualifier
    : EXTENDS ID 
    ;

fieldsetValues
    : fieldsetValue+  
    ;

fieldsetValue
    : KEEP_COMMENT* fieldType ID (EQUALS integerConst)? SEMI 
    ;

qualifiedId
    : ID (DOT ID)* 
    ;


fieldType
    : DOUBLE
    | FLOAT
    | INT32
    | INT64
    | UINT32
    | UINT64
    | BOOL
    | STRING ( LPAREN integerConst RPAREN )?  
    | BYTES  ( LPAREN integerConst RPAREN )?  
    | CHARARRAY ( LPAREN integerConst RPAREN )?  
    | BYTEARRAY ( LPAREN integerConst RPAREN )?  
    | DATETIME
    | GUID
    | DECIMAL
    | qualifiedId
    ;


indexType
    : PRIMARY
    | UNIQUE
    | NONUNIQUE
    ;

indexElement
    : INDEX indexType LPAREN ID (  COMMA ID )* RPAREN SEMI 
    ;


classElement
    : KEEP_COMMENT* CLASS ID classOptions LBRACE classFields indexElement* RBRACE 
    ;

classOptions
    : classOption*
    ;


classOption
    : fieldsetOption
    | tableOption
    ;

fieldsetOption
    : FIELDSET LPAREN ID RPAREN 
    ;

tableOption
    : TABLE LPAREN ID RPAREN 
    ;

mediatesOption
    : MEDIATES LPAREN ID RPAREN 
    ;

classFields
    : classField+ 
    ;

referencesModifier
    : REFERENCES ID (LPAREN ID RPAREN)? 
    ;

classField
    : fieldModifier ID referencesModifier?  SEMI 
    ;

parameterModifier
    : REQUIRED
    | OPTIONAL
    | REPEATED
    ;


fieldModifier
    : AUTO
    | ID_KEYWORD
    | IDGEN
    | VIRTUAL
    | MAP
    | LIST
    | parameterModifier
    ;

associationElement
    : KEEP_COMMENT* ASSOCIATION ID associationOptions LBRACE classFields RBRACE 
    ;

associationOptions
    : associationOption*
    ;

associationOption
    : fieldsetOption
    | tableOption
    ;

serviceElement
    : KEEP_COMMENT* SERVICE ID fieldsetOption? LBRACE methods RBRACE 
    ;


patternType
    : INSERT
    | UPDATE
    | SELECT
    | DELETE
    ;


methodPattern
    : PATTERN LPAREN patternType qualifiedId RPAREN 
    ;

methods
    : (method)+
    ;

method
    : KEEP_COMMENT* METHOD ID methodPattern? LBRACE inputList RBRACE methodReturn? 
    ;

inputList
    : fieldList 
    ;

fieldList
    : fieldInstance+  
    ;

fieldInstance
    : parameterModifier ID SEMI 
    ;

methodReturn
    : RETURNS LBRACE fieldList RBRACE 
    ;
        
packageElement
    : PACKAGE qualifiedId SEMI 
    ;


importElements
    : IMPORT TEXT_STRING SEMI 
    ;
