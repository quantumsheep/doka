// Code generated from Dokafile.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Dokafile

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDokafileVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDokafileVisitor) VisitTranslation(ctx *TranslationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitInstructions(ctx *InstructionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitInstruction(ctx *InstructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitNothing(ctx *NothingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitComment(ctx *CommentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitNative_instruction(ctx *Native_instructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitIf_instruction(ctx *If_instructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitElif_instruction(ctx *Elif_instructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitElse_instruction(ctx *Else_instructionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitIf_expression(ctx *If_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDokafileVisitor) VisitInclude_instruction(ctx *Include_instructionContext) interface{} {
	return v.VisitChildren(ctx)
}
