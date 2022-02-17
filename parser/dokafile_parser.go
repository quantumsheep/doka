// Code generated from Dokafile.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Dokafile

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 111,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 3, 5, 3, 29, 10, 3, 3, 3, 3, 3, 3, 3, 7, 3, 34, 10, 3, 12,
	3, 14, 3, 37, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 43, 10, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 6, 8, 53, 10, 8, 13, 8, 14, 8, 54,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 61, 10, 8, 3, 8, 5, 8, 64, 10, 8, 3, 8, 5,
	8, 67, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 6, 9, 73, 10, 9, 13, 9, 14, 9, 74,
	3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 81, 10, 9, 3, 9, 3, 9, 5, 9, 85, 10, 9, 5,
	9, 87, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 95, 10,
	11, 12, 11, 14, 11, 98, 11, 11, 3, 11, 3, 11, 7, 11, 102, 10, 11, 12, 11,
	14, 11, 105, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 2, 2, 13, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 2, 4, 3, 2, 9, 10, 4, 2, 11, 11, 14,
	14, 2, 114, 2, 24, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8,
	44, 3, 2, 2, 2, 10, 46, 3, 2, 2, 2, 12, 48, 3, 2, 2, 2, 14, 50, 3, 2, 2,
	2, 16, 70, 3, 2, 2, 2, 18, 88, 3, 2, 2, 2, 20, 92, 3, 2, 2, 2, 22, 108,
	3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26, 7, 2, 2, 3, 26, 3, 3, 2, 2, 2,
	27, 29, 7, 13, 2, 2, 28, 27, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 30, 3,
	2, 2, 2, 30, 31, 5, 6, 4, 2, 31, 32, 7, 3, 2, 2, 32, 34, 3, 2, 2, 2, 33,
	28, 3, 2, 2, 2, 34, 37, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 36, 3, 2, 2,
	2, 36, 5, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 38, 43, 5, 14, 8, 2, 39, 43,
	5, 12, 7, 2, 40, 43, 5, 8, 5, 2, 41, 43, 5, 10, 6, 2, 42, 38, 3, 2, 2,
	2, 42, 39, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 7, 3,
	2, 2, 2, 44, 45, 3, 2, 2, 2, 45, 9, 3, 2, 2, 2, 46, 47, 7, 15, 2, 2, 47,
	11, 3, 2, 2, 2, 48, 49, 7, 12, 2, 2, 49, 13, 3, 2, 2, 2, 50, 52, 7, 5,
	2, 2, 51, 53, 7, 4, 2, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 54, 55, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 57, 5, 20, 11,
	2, 57, 58, 7, 3, 2, 2, 58, 63, 5, 4, 3, 2, 59, 61, 7, 13, 2, 2, 60, 59,
	3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64, 5, 16, 9, 2,
	63, 60, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2, 65, 67, 7,
	13, 2, 2, 66, 65, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68,
	69, 7, 8, 2, 2, 69, 15, 3, 2, 2, 2, 70, 72, 7, 7, 2, 2, 71, 73, 7, 4, 2,
	2, 72, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 72, 3, 2, 2, 2, 74, 75,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 5, 20, 11, 2, 77, 78, 7, 3, 2,
	2, 78, 86, 5, 4, 3, 2, 79, 81, 7, 13, 2, 2, 80, 79, 3, 2, 2, 2, 80, 81,
	3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 85, 5, 16, 9, 2, 83, 85, 5, 18, 10,
	2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 87, 3, 2, 2, 2, 86, 80,
	3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 17, 3, 2, 2, 2, 88, 89, 7, 6, 2, 2,
	89, 90, 7, 3, 2, 2, 90, 91, 5, 4, 3, 2, 91, 19, 3, 2, 2, 2, 92, 96, 5,
	22, 12, 2, 93, 95, 7, 4, 2, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2,
	96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98, 96, 3,
	2, 2, 2, 99, 103, 9, 2, 2, 2, 100, 102, 7, 4, 2, 2, 101, 100, 3, 2, 2,
	2, 102, 105, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104,
	106, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 107, 5, 22, 12, 2, 107, 21,
	3, 2, 2, 2, 108, 109, 9, 3, 2, 2, 109, 23, 3, 2, 2, 2, 15, 28, 35, 42,
	54, 60, 63, 66, 74, 80, 84, 86, 96, 103,
}
var literalNames = []string{
	"", "'\n'", "' '", "'!IF'", "'!ELSE'", "'!ELIF'", "'!END'", "'=='", "'!='",
}
var symbolicNames = []string{
	"", "", "", "If", "Else", "Elif", "End", "OperatorEq", "OperatorNe", "Variable",
	"NativeInstructionCall", "Indent", "StringLiteral", "Comment",
}

var ruleNames = []string{
	"translation", "instructions", "instruction", "nothing", "comment", "native_instruction",
	"if_instruction", "elif_instruction", "else_instruction", "if_expression",
	"value",
}

type DokafileParser struct {
	*antlr.BaseParser
}

// NewDokafileParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DokafileParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDokafileParser(input antlr.TokenStream) *DokafileParser {
	this := new(DokafileParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Dokafile.g4"

	return this
}

// DokafileParser tokens.
const (
	DokafileParserEOF                   = antlr.TokenEOF
	DokafileParserT__0                  = 1
	DokafileParserT__1                  = 2
	DokafileParserIf                    = 3
	DokafileParserElse                  = 4
	DokafileParserElif                  = 5
	DokafileParserEnd                   = 6
	DokafileParserOperatorEq            = 7
	DokafileParserOperatorNe            = 8
	DokafileParserVariable              = 9
	DokafileParserNativeInstructionCall = 10
	DokafileParserIndent                = 11
	DokafileParserStringLiteral         = 12
	DokafileParserComment               = 13
)

// DokafileParser rules.
const (
	DokafileParserRULE_translation        = 0
	DokafileParserRULE_instructions       = 1
	DokafileParserRULE_instruction        = 2
	DokafileParserRULE_nothing            = 3
	DokafileParserRULE_comment            = 4
	DokafileParserRULE_native_instruction = 5
	DokafileParserRULE_if_instruction     = 6
	DokafileParserRULE_elif_instruction   = 7
	DokafileParserRULE_else_instruction   = 8
	DokafileParserRULE_if_expression      = 9
	DokafileParserRULE_value              = 10
)

// ITranslationContext is an interface to support dynamic dispatch.
type ITranslationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTranslationContext differentiates from other interfaces.
	IsTranslationContext()
}

type TranslationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTranslationContext() *TranslationContext {
	var p = new(TranslationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_translation
	return p
}

func (*TranslationContext) IsTranslationContext() {}

func NewTranslationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TranslationContext {
	var p = new(TranslationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_translation

	return p
}

func (s *TranslationContext) GetParser() antlr.Parser { return s.parser }

func (s *TranslationContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *TranslationContext) EOF() antlr.TerminalNode {
	return s.GetToken(DokafileParserEOF, 0)
}

func (s *TranslationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TranslationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TranslationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitTranslation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Translation() (localctx ITranslationContext) {
	localctx = NewTranslationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DokafileParserRULE_translation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(22)
		p.Instructions()
	}
	{
		p.SetState(23)
		p.Match(DokafileParserEOF)
	}

	return localctx
}

// IInstructionsContext is an interface to support dynamic dispatch.
type IInstructionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionsContext differentiates from other interfaces.
	IsInstructionsContext()
}

type InstructionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionsContext() *InstructionsContext {
	var p = new(InstructionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_instructions
	return p
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_instructions

	return p
}

func (s *InstructionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *InstructionsContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstructionsContext) AllIndent() []antlr.TerminalNode {
	return s.GetTokens(DokafileParserIndent)
}

func (s *InstructionsContext) Indent(i int) antlr.TerminalNode {
	return s.GetToken(DokafileParserIndent, i)
}

func (s *InstructionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitInstructions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DokafileParserRULE_instructions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(26)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DokafileParserIndent {
				{
					p.SetState(25)
					p.Match(DokafileParserIndent)
				}

			}
			{
				p.SetState(28)
				p.Instruction()
			}
			{
				p.SetState(29)
				p.Match(DokafileParserT__0)
			}

		}
		p.SetState(35)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) If_instruction() IIf_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_instructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_instructionContext)
}

func (s *InstructionContext) Native_instruction() INative_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INative_instructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INative_instructionContext)
}

func (s *InstructionContext) Nothing() INothingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INothingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INothingContext)
}

func (s *InstructionContext) Comment() ICommentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommentContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DokafileParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DokafileParserIf:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.If_instruction()
		}

	case DokafileParserNativeInstructionCall:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Native_instruction()
		}

	case DokafileParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(38)
			p.Nothing()
		}

	case DokafileParserComment:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(39)
			p.Comment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INothingContext is an interface to support dynamic dispatch.
type INothingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNothingContext differentiates from other interfaces.
	IsNothingContext()
}

type NothingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNothingContext() *NothingContext {
	var p = new(NothingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_nothing
	return p
}

func (*NothingContext) IsNothingContext() {}

func NewNothingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NothingContext {
	var p = new(NothingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_nothing

	return p
}

func (s *NothingContext) GetParser() antlr.Parser { return s.parser }
func (s *NothingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NothingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NothingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitNothing(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Nothing() (localctx INothingContext) {
	localctx = NewNothingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DokafileParserRULE_nothing)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)

	return localctx
}

// ICommentContext is an interface to support dynamic dispatch.
type ICommentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommentContext differentiates from other interfaces.
	IsCommentContext()
}

type CommentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommentContext() *CommentContext {
	var p = new(CommentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_comment
	return p
}

func (*CommentContext) IsCommentContext() {}

func NewCommentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommentContext {
	var p = new(CommentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_comment

	return p
}

func (s *CommentContext) GetParser() antlr.Parser { return s.parser }

func (s *CommentContext) Comment() antlr.TerminalNode {
	return s.GetToken(DokafileParserComment, 0)
}

func (s *CommentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitComment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Comment() (localctx ICommentContext) {
	localctx = NewCommentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DokafileParserRULE_comment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(DokafileParserComment)
	}

	return localctx
}

// INative_instructionContext is an interface to support dynamic dispatch.
type INative_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNative_instructionContext differentiates from other interfaces.
	IsNative_instructionContext()
}

type Native_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNative_instructionContext() *Native_instructionContext {
	var p = new(Native_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_native_instruction
	return p
}

func (*Native_instructionContext) IsNative_instructionContext() {}

func NewNative_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Native_instructionContext {
	var p = new(Native_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_native_instruction

	return p
}

func (s *Native_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Native_instructionContext) NativeInstructionCall() antlr.TerminalNode {
	return s.GetToken(DokafileParserNativeInstructionCall, 0)
}

func (s *Native_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Native_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Native_instructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitNative_instruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Native_instruction() (localctx INative_instructionContext) {
	localctx = NewNative_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DokafileParserRULE_native_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(DokafileParserNativeInstructionCall)
	}

	return localctx
}

// IIf_instructionContext is an interface to support dynamic dispatch.
type IIf_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_instructionContext differentiates from other interfaces.
	IsIf_instructionContext()
}

type If_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_instructionContext() *If_instructionContext {
	var p = new(If_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_if_instruction
	return p
}

func (*If_instructionContext) IsIf_instructionContext() {}

func NewIf_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_instructionContext {
	var p = new(If_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_if_instruction

	return p
}

func (s *If_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *If_instructionContext) If() antlr.TerminalNode {
	return s.GetToken(DokafileParserIf, 0)
}

func (s *If_instructionContext) If_expression() IIf_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_expressionContext)
}

func (s *If_instructionContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *If_instructionContext) End() antlr.TerminalNode {
	return s.GetToken(DokafileParserEnd, 0)
}

func (s *If_instructionContext) Elif_instruction() IElif_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElif_instructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElif_instructionContext)
}

func (s *If_instructionContext) AllIndent() []antlr.TerminalNode {
	return s.GetTokens(DokafileParserIndent)
}

func (s *If_instructionContext) Indent(i int) antlr.TerminalNode {
	return s.GetToken(DokafileParserIndent, i)
}

func (s *If_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_instructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitIf_instruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) If_instruction() (localctx IIf_instructionContext) {
	localctx = NewIf_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DokafileParserRULE_if_instruction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(DokafileParserIf)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DokafileParserT__1 {
		{
			p.SetState(49)
			p.Match(DokafileParserT__1)
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.If_expression()
	}
	{
		p.SetState(55)
		p.Match(DokafileParserT__0)
	}
	{
		p.SetState(56)
		p.Instructions()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DokafileParserIndent {
			{
				p.SetState(57)
				p.Match(DokafileParserIndent)
			}

		}
		{
			p.SetState(60)
			p.Elif_instruction()
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DokafileParserIndent {
		{
			p.SetState(63)
			p.Match(DokafileParserIndent)
		}

	}
	{
		p.SetState(66)
		p.Match(DokafileParserEnd)
	}

	return localctx
}

// IElif_instructionContext is an interface to support dynamic dispatch.
type IElif_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElif_instructionContext differentiates from other interfaces.
	IsElif_instructionContext()
}

type Elif_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElif_instructionContext() *Elif_instructionContext {
	var p = new(Elif_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_elif_instruction
	return p
}

func (*Elif_instructionContext) IsElif_instructionContext() {}

func NewElif_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Elif_instructionContext {
	var p = new(Elif_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_elif_instruction

	return p
}

func (s *Elif_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Elif_instructionContext) Elif() antlr.TerminalNode {
	return s.GetToken(DokafileParserElif, 0)
}

func (s *Elif_instructionContext) If_expression() IIf_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_expressionContext)
}

func (s *Elif_instructionContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *Elif_instructionContext) Elif_instruction() IElif_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElif_instructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElif_instructionContext)
}

func (s *Elif_instructionContext) Else_instruction() IElse_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElse_instructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElse_instructionContext)
}

func (s *Elif_instructionContext) Indent() antlr.TerminalNode {
	return s.GetToken(DokafileParserIndent, 0)
}

func (s *Elif_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Elif_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Elif_instructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitElif_instruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Elif_instruction() (localctx IElif_instructionContext) {
	localctx = NewElif_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DokafileParserRULE_elif_instruction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(DokafileParserElif)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DokafileParserT__1 {
		{
			p.SetState(69)
			p.Match(DokafileParserT__1)
		}

		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.If_expression()
	}
	{
		p.SetState(75)
		p.Match(DokafileParserT__0)
	}
	{
		p.SetState(76)
		p.Instructions()
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DokafileParserIndent {
			{
				p.SetState(77)
				p.Match(DokafileParserIndent)
			}

		}
		p.SetState(82)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DokafileParserElif:
			{
				p.SetState(80)
				p.Elif_instruction()
			}

		case DokafileParserElse:
			{
				p.SetState(81)
				p.Else_instruction()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IElse_instructionContext is an interface to support dynamic dispatch.
type IElse_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElse_instructionContext differentiates from other interfaces.
	IsElse_instructionContext()
}

type Else_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElse_instructionContext() *Else_instructionContext {
	var p = new(Else_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_else_instruction
	return p
}

func (*Else_instructionContext) IsElse_instructionContext() {}

func NewElse_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Else_instructionContext {
	var p = new(Else_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_else_instruction

	return p
}

func (s *Else_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Else_instructionContext) Else() antlr.TerminalNode {
	return s.GetToken(DokafileParserElse, 0)
}

func (s *Else_instructionContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *Else_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Else_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Else_instructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitElse_instruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Else_instruction() (localctx IElse_instructionContext) {
	localctx = NewElse_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DokafileParserRULE_else_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(DokafileParserElse)
	}
	{
		p.SetState(87)
		p.Match(DokafileParserT__0)
	}
	{
		p.SetState(88)
		p.Instructions()
	}

	return localctx
}

// IIf_expressionContext is an interface to support dynamic dispatch.
type IIf_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_expressionContext differentiates from other interfaces.
	IsIf_expressionContext()
}

type If_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_expressionContext() *If_expressionContext {
	var p = new(If_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_if_expression
	return p
}

func (*If_expressionContext) IsIf_expressionContext() {}

func NewIf_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_expressionContext {
	var p = new(If_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_if_expression

	return p
}

func (s *If_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *If_expressionContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *If_expressionContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *If_expressionContext) OperatorEq() antlr.TerminalNode {
	return s.GetToken(DokafileParserOperatorEq, 0)
}

func (s *If_expressionContext) OperatorNe() antlr.TerminalNode {
	return s.GetToken(DokafileParserOperatorNe, 0)
}

func (s *If_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitIf_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) If_expression() (localctx IIf_expressionContext) {
	localctx = NewIf_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DokafileParserRULE_if_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Value()
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DokafileParserT__1 {
		{
			p.SetState(91)
			p.Match(DokafileParserT__1)
		}

		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DokafileParserOperatorEq || _la == DokafileParserOperatorNe) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DokafileParserT__1 {
		{
			p.SetState(98)
			p.Match(DokafileParserT__1)
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(104)
		p.Value()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DokafileParserStringLiteral, 0)
}

func (s *ValueContext) Variable() antlr.TerminalNode {
	return s.GetToken(DokafileParserVariable, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DokafileParserRULE_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DokafileParserVariable || _la == DokafileParserStringLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
