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
	If ' '+ if_expression '\n' instructions (
		Indent? elif_instruction
	)? Indent? End;

elif_instruction:
	Elif ' '+ if_expression '\n' instructions (
		Indent? (elif_instruction | else_instruction)
	)?;

else_instruction: Else '\n' instructions;

if_expression: value ' '* (OperatorEq | OperatorNe) ' '* value;

value: StringLiteral | Variable;

If: '!IF';
Else: '!ELSE';
Elif: '!ELIF';
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
	| 'LABEL'
	| 'EXPOSE'
	| 'ENV'
	| 'ARG'
	| 'VOLUME'
	| 'USER'
	| 'ONBUILD'
	| 'STOPSIGNAL'
	| 'HEALTHCHECK'
	| 'SHELL'
	| 'ADD'
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
