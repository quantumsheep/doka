grammar Dokafile;

translation: instructions EOF;

instructions: ( Indent? instruction '\n')*;

instruction:
	if_instruction
	| native_instruction
	| nothing
	| comment;

nothing:;

comment: Comment;

native_instruction: NativeInstructionCall;

if_instruction:
	If ' '+ if_expression '\n' instructions Indent? End;

if_expression: value ' '* (OperatorEq | OperatorNe) ' '* value;

value: StringLiteral | Variable;

If: '!IF';
End: '!END';

OperatorEq: '==';
OperatorNe: '!=';

Variable: '$' Name;

NativeInstructionCall:
	NativeInstruction ' ' NativeInstructionCallChar*;

fragment NativeInstructionCallChar: ~ [\r\n] | '\\\n';

Indent: [ \t]+;

fragment NativeInstruction:
	'FROM'
	| 'WORKDIR'
	| 'COPY'
	| 'RUN'
	| 'CMD'
	| 'ENTRYPOINT';

StringLiteral: ('"' StringChar* '"') | '\'' CharChar+ '\'';

Comment: '#' (~ [\r\n])*;

fragment Name: [a-zA-Z_] [a-zA-Z0-9_]*;

fragment StringChar: ~ ["\r\n] | Escape;
fragment CharChar: ~ ['\r\n] | Escape;

fragment Escape: '\\\'' | '\\"' | '\\\\' | '\\n' | '\\r';
