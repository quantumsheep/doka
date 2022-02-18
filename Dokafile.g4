grammar Dokafile;

translation: instructions EOF;

instructions: ( Spacing? instruction '\n')*;

instruction:
	if_instruction
	| include_instruction
	| native_instruction
	| nothing
	| comment;

nothing:;

comment: Comment;

native_instruction: NativeInstructionCall;

if_instruction:
	If Spacing if_expression '\n' instructions (
		Spacing? elif_instruction
	)? Spacing? End;

elif_instruction:
	Elif Spacing if_expression '\n' instructions (
		Spacing? (elif_instruction | else_instruction)
	)?;

else_instruction: Else '\n' instructions;

if_expression:
	value Spacing? (OperatorEq | OperatorNe) Spacing? value;

value: StringLiteral | Variable;

include_instruction: Include Spacing StringLiteral;

If: '!IF';
Else: '!ELSE';
Elif: '!ELIF';
End: '!END';

OperatorEq: '==';
OperatorNe: '!=';

Variable: '$' Name;

Include: '!INCLUDE';

NativeInstructionCall:
	NativeInstruction Spacing (~ [\r\n] | '\\\n')*;

StringLiteral: ('"' StringChar* '"') | ('\'' CharChar+ '\'');

Comment: Spacing? '#' (~ [\r\n])*;

Spacing: [ \t]+;

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

fragment Name: [a-zA-Z_] [a-zA-Z0-9_]*;

fragment StringChar: ~ ["\r\n] | Escape;
fragment CharChar: ~ ['\r\n] | Escape;

fragment Escape: '\\\'' | '\\"' | '\\\\' | '\\n' | '\\r';
