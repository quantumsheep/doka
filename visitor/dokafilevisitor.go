package visitor

import (
	"fmt"
	"os"
	"runtime"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/quantumsheep/doka/parser"
)

type DokafileVisitor struct {
	*parser.BaseDokafileVisitor

	filepath  string
	variables map[string]string

	output string
}

func NewDokafileVisitor(filepath string, variables map[string]string) *DokafileVisitor {
	return &DokafileVisitor{
		filepath:  filepath,
		variables: variables,
	}
}

func (v *DokafileVisitor) ErrorAtToken(message string, token antlr.Token) error {
	return fmt.Errorf("%s:%d:%d: %s '%s'", v.filepath, token.GetLine(), token.GetColumn()+1, message, token.GetText())
}

func (v *DokafileVisitor) Visit(tree antlr.ParseTree) error {
	switch val := tree.(type) {
	case *parser.TranslationContext:
		return v.VisitTranslation(val)
	default:
		return fmt.Errorf("invalid context")
	}
}

func (v *DokafileVisitor) VisitTranslation(ctx *parser.TranslationContext) error {
	return v.VisitInstructions(ctx.Instructions().(*parser.InstructionsContext))
}

func (v *DokafileVisitor) VisitInstructions(ctx *parser.InstructionsContext) error {
	for _, instruction := range ctx.AllInstruction() {
		e := v.VisitInstruction(instruction.(*parser.InstructionContext))
		if e != nil {
			return e
		}
	}

	return nil
}

func (v *DokafileVisitor) VisitInstruction(ctx *parser.InstructionContext) error {
	if child := ctx.If_instruction(); child != nil {
		e := v.VisitIf_instruction(child.(*parser.If_instructionContext))
		if e != nil {
			return e
		}
	} else if child := ctx.Native_instruction(); child != nil {
		v.VisitNative_instruction(child.(*parser.Native_instructionContext))
	} else if child := ctx.Nothing(); child != nil {
		v.output += "\n"
	} else if child := ctx.Comment(); child != nil {
		v.VisitComment(child.(*parser.CommentContext))
	} else {
		return v.ErrorAtToken("unknown instruction", ctx.GetStart())
	}

	return nil
}

func (v *DokafileVisitor) VisitIf_instruction(ctx *parser.If_instructionContext) error {
	expression, e := v.VisitIf_expression(ctx.If_expression().(*parser.If_expressionContext))
	if e != nil {
		return e
	}

	if expression {
		return v.VisitInstructions(ctx.Instructions().(*parser.InstructionsContext))
	}

	return v.VisitElif_instruction(ctx.Elif_instruction().(*parser.Elif_instructionContext))
}

func (v *DokafileVisitor) VisitElif_instruction(ctx *parser.Elif_instructionContext) error {
	expression, e := v.VisitIf_expression(ctx.If_expression().(*parser.If_expressionContext))
	if e != nil {
		return e
	}

	if expression {
		return v.VisitInstructions(ctx.Instructions().(*parser.InstructionsContext))
	}

	if child := ctx.Elif_instruction(); child != nil {
		return v.VisitElif_instruction(child.(*parser.Elif_instructionContext))
	} else if child := ctx.Else_instruction(); child != nil {
		return v.VisitElse_instruction(child.(*parser.Else_instructionContext))
	}

	return v.ErrorAtToken("unknown context", ctx.GetStart())
}

func (v *DokafileVisitor) VisitElse_instruction(ctx *parser.Else_instructionContext) error {
	return v.VisitInstructions(ctx.Instructions().(*parser.InstructionsContext))
}

func (v *DokafileVisitor) VisitIf_expression(ctx *parser.If_expressionContext) (bool, error) {
	leftValue, e := v.VisitValue(ctx.Value(0).(*parser.ValueContext))
	if e != nil {
		return false, e
	}

	rightValue, e := v.VisitValue(ctx.Value(1).(*parser.ValueContext))
	if e != nil {
		return false, e
	}

	if ctx.OperatorEq() != nil {
		return leftValue == rightValue, nil
	} else if ctx.OperatorNe() != nil {
		return leftValue != rightValue, nil
	}

	return false, v.ErrorAtToken("operator not implemented", ctx.GetStart())
}

func (v *DokafileVisitor) VisitValue(ctx *parser.ValueContext) (string, error) {
	if child := ctx.StringLiteral(); child != nil {
		str := child.GetText()
		return str[1 : len(str)-1], nil
	} else if child := ctx.Variable(); child != nil {
		variableName := child.GetText()[1:]

		if value, ok := v.variables[variableName]; ok {
			return value, nil
		} else if value := os.Getenv(variableName); value != "" {
			return value, nil
		}

		return "", v.ErrorAtToken("unknown variable", child.GetSymbol())
	}

	return "", v.ErrorAtToken("unknown value type", ctx.GetStart())
}

func (v *DokafileVisitor) VisitNative_instruction(ctx *parser.Native_instructionContext) {
	v.output += ctx.GetText() + "\n"
}

func (v *DokafileVisitor) VisitComment(ctx *parser.CommentContext) {
	v.output += ctx.GetText() + "\n"
}

type DokaErrorListener struct {
	*antlr.DefaultErrorListener

	Filepath   string
	FullErrors bool
	Errors     []error
}

func NewDokaErrorListener(filepath string, fullErrors bool) *DokaErrorListener {
	return &DokaErrorListener{
		Filepath:   filepath,
		FullErrors: fullErrors,
		Errors:     make([]error, 0),
	}
}

func (listener *DokaErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	var err error = nil

	if listener.FullErrors {
		err = fmt.Errorf("%s:%d:%d: %s", listener.Filepath, line, column, msg)
	} else {
		switch token := offendingSymbol.(type) {
		case antlr.Token:
			err = fmt.Errorf("%s:%d:%d: syntax error at '%s'", listener.Filepath, line, column, token.GetText())
		default:
			err = fmt.Errorf("%s:%d:%d: %s", listener.Filepath, line, column, msg)
		}
	}

	listener.Errors = append(listener.Errors, err)
}

func ParseFile(filepath string, fullErrors bool, variables map[string]string) (string, []error) {
	input, e := antlr.NewFileStream(filepath)
	if e != nil {
		return "", []error{e}
	}

	errorListener := NewDokaErrorListener(filepath, fullErrors)

	lexer := parser.NewDokafileLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewDokafileParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	tree := parser.Translation()
	if tree == nil {
		return "", []error{fmt.Errorf("parse error")}
	}

	if len(errorListener.Errors) > 0 {
		return "", errorListener.Errors
	}

	visitor := NewDokafileVisitor(filepath, variables)

	// assign os if not already set
	if _, ok := variables["os"]; !ok {
		visitor.variables["os"] = runtime.GOOS
	}

	if _, ok := variables["arch"]; !ok {
		visitor.variables["arch"] = runtime.GOARCH
	}

	e = visitor.Visit(tree)
	if e != nil {
		return "", []error{e}
	}

	return visitor.output, []error{}
}
