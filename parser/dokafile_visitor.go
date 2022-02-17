// Code generated from Dokafile.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Dokafile

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DokafileParser.
type DokafileVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DokafileParser#translation.
	VisitTranslation(ctx *TranslationContext) interface{}

	// Visit a parse tree produced by DokafileParser#instructions.
	VisitInstructions(ctx *InstructionsContext) interface{}

	// Visit a parse tree produced by DokafileParser#instruction.
	VisitInstruction(ctx *InstructionContext) interface{}

	// Visit a parse tree produced by DokafileParser#nothing.
	VisitNothing(ctx *NothingContext) interface{}

	// Visit a parse tree produced by DokafileParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Visit a parse tree produced by DokafileParser#native_instruction.
	VisitNative_instruction(ctx *Native_instructionContext) interface{}

	// Visit a parse tree produced by DokafileParser#if_instruction.
	VisitIf_instruction(ctx *If_instructionContext) interface{}

	// Visit a parse tree produced by DokafileParser#if_expression.
	VisitIf_expression(ctx *If_expressionContext) interface{}

	// Visit a parse tree produced by DokafileParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
