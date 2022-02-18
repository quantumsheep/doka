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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 15, 104,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 3, 5, 3, 31, 10, 3, 3, 3, 3, 3, 3, 3, 7, 3,
	36, 10, 3, 12, 3, 14, 3, 39, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	46, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 60, 10, 8, 3, 8, 5, 8, 63, 10, 8, 3, 8, 5, 8, 66, 10,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 76, 10, 9, 3,
	9, 3, 9, 5, 9, 80, 10, 9, 5, 9, 82, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 5, 11, 90, 10, 11, 3, 11, 3, 11, 5, 11, 94, 10, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 2, 2, 14, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 4, 3, 2, 8, 9, 4, 2, 10, 10, 13,
	13, 2, 105, 2, 26, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 45, 3, 2, 2, 2, 8,
	47, 3, 2, 2, 2, 10, 49, 3, 2, 2, 2, 12, 51, 3, 2, 2, 2, 14, 53, 3, 2, 2,
	2, 16, 69, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2, 20, 87, 3, 2, 2, 2, 22, 97,
	3, 2, 2, 2, 24, 99, 3, 2, 2, 2, 26, 27, 5, 4, 3, 2, 27, 28, 7, 2, 2, 3,
	28, 3, 3, 2, 2, 2, 29, 31, 7, 15, 2, 2, 30, 29, 3, 2, 2, 2, 30, 31, 3,
	2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 33, 5, 6, 4, 2, 33, 34, 7, 3, 2, 2, 34,
	36, 3, 2, 2, 2, 35, 30, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2,
	2, 37, 38, 3, 2, 2, 2, 38, 5, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 46, 5,
	14, 8, 2, 41, 46, 5, 24, 13, 2, 42, 46, 5, 12, 7, 2, 43, 46, 5, 8, 5, 2,
	44, 46, 5, 10, 6, 2, 45, 40, 3, 2, 2, 2, 45, 41, 3, 2, 2, 2, 45, 42, 3,
	2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 7, 3, 2, 2, 2, 47,
	48, 3, 2, 2, 2, 48, 9, 3, 2, 2, 2, 49, 50, 7, 14, 2, 2, 50, 11, 3, 2, 2,
	2, 51, 52, 7, 12, 2, 2, 52, 13, 3, 2, 2, 2, 53, 54, 7, 4, 2, 2, 54, 55,
	7, 15, 2, 2, 55, 56, 5, 20, 11, 2, 56, 57, 7, 3, 2, 2, 57, 62, 5, 4, 3,
	2, 58, 60, 7, 15, 2, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 63, 5, 16, 9, 2, 62, 59, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2,
	63, 65, 3, 2, 2, 2, 64, 66, 7, 15, 2, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3,
	2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 68, 7, 7, 2, 2, 68, 15, 3, 2, 2, 2, 69,
	70, 7, 6, 2, 2, 70, 71, 7, 15, 2, 2, 71, 72, 5, 20, 11, 2, 72, 73, 7, 3,
	2, 2, 73, 81, 5, 4, 3, 2, 74, 76, 7, 15, 2, 2, 75, 74, 3, 2, 2, 2, 75,
	76, 3, 2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 80, 5, 16, 9, 2, 78, 80, 5, 18,
	10, 2, 79, 77, 3, 2, 2, 2, 79, 78, 3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81,
	75, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 17, 3, 2, 2, 2, 83, 84, 7, 5, 2,
	2, 84, 85, 7, 3, 2, 2, 85, 86, 5, 4, 3, 2, 86, 19, 3, 2, 2, 2, 87, 89,
	5, 22, 12, 2, 88, 90, 7, 15, 2, 2, 89, 88, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 90, 91, 3, 2, 2, 2, 91, 93, 9, 2, 2, 2, 92, 94, 7, 15, 2, 2, 93, 92,
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 96, 5, 22, 12,
	2, 96, 21, 3, 2, 2, 2, 97, 98, 9, 3, 2, 2, 98, 23, 3, 2, 2, 2, 99, 100,
	7, 11, 2, 2, 100, 101, 7, 15, 2, 2, 101, 102, 7, 13, 2, 2, 102, 25, 3,
	2, 2, 2, 13, 30, 37, 45, 59, 62, 65, 75, 79, 81, 89, 93,
}
var literalNames = []string{
	"", "'\n'", "'!IF'", "'!ELSE'", "'!ELIF'", "'!END'", "'=='", "'!='", "",
	"'!INCLUDE'",
}
var symbolicNames = []string{
	"", "", "If", "Else", "Elif", "End", "OperatorEq", "OperatorNe", "Variable",
	"Include", "NativeInstructionCall", "StringLiteral", "Comment", "Spacing",
}

var ruleNames = []string{
	"translation", "instructions", "instruction", "nothing", "comment", "native_instruction",
	"if_instruction", "elif_instruction", "else_instruction", "if_expression",
	"value", "include_instruction",
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
	DokafileParserIf                    = 2
	DokafileParserElse                  = 3
	DokafileParserElif                  = 4
	DokafileParserEnd                   = 5
	DokafileParserOperatorEq            = 6
	DokafileParserOperatorNe            = 7
	DokafileParserVariable              = 8
	DokafileParserInclude               = 9
	DokafileParserNativeInstructionCall = 10
	DokafileParserStringLiteral         = 11
	DokafileParserComment               = 12
	DokafileParserSpacing               = 13
)

// DokafileParser rules.
const (
	DokafileParserRULE_translation         = 0
	DokafileParserRULE_instructions        = 1
	DokafileParserRULE_instruction         = 2
	DokafileParserRULE_nothing             = 3
	DokafileParserRULE_comment             = 4
	DokafileParserRULE_native_instruction  = 5
	DokafileParserRULE_if_instruction      = 6
	DokafileParserRULE_elif_instruction    = 7
	DokafileParserRULE_else_instruction    = 8
	DokafileParserRULE_if_expression       = 9
	DokafileParserRULE_value               = 10
	DokafileParserRULE_include_instruction = 11
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
		p.SetState(24)
		p.Instructions()
	}
	{
		p.SetState(25)
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

func (s *InstructionsContext) AllSpacing() []antlr.TerminalNode {
	return s.GetTokens(DokafileParserSpacing)
}

func (s *InstructionsContext) Spacing(i int) antlr.TerminalNode {
	return s.GetToken(DokafileParserSpacing, i)
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(28)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == DokafileParserSpacing {
				{
					p.SetState(27)
					p.Match(DokafileParserSpacing)
				}

			}
			{
				p.SetState(30)
				p.Instruction()
			}
			{
				p.SetState(31)
				p.Match(DokafileParserT__0)
			}

		}
		p.SetState(37)
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

func (s *InstructionContext) Include_instruction() IInclude_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInclude_instructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInclude_instructionContext)
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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DokafileParserIf:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.If_instruction()
		}

	case DokafileParserInclude:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Include_instruction()
		}

	case DokafileParserNativeInstructionCall:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Native_instruction()
		}

	case DokafileParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Nothing()
		}

	case DokafileParserComment:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(42)
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
		p.SetState(47)
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
		p.SetState(49)
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

func (s *If_instructionContext) AllSpacing() []antlr.TerminalNode {
	return s.GetTokens(DokafileParserSpacing)
}

func (s *If_instructionContext) Spacing(i int) antlr.TerminalNode {
	return s.GetToken(DokafileParserSpacing, i)
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
		p.SetState(51)
		p.Match(DokafileParserIf)
	}
	{
		p.SetState(52)
		p.Match(DokafileParserSpacing)
	}
	{
		p.SetState(53)
		p.If_expression()
	}
	{
		p.SetState(54)
		p.Match(DokafileParserT__0)
	}
	{
		p.SetState(55)
		p.Instructions()
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DokafileParserSpacing {
			{
				p.SetState(56)
				p.Match(DokafileParserSpacing)
			}

		}
		{
			p.SetState(59)
			p.Elif_instruction()
		}

	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DokafileParserSpacing {
		{
			p.SetState(62)
			p.Match(DokafileParserSpacing)
		}

	}
	{
		p.SetState(65)
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

func (s *Elif_instructionContext) AllSpacing() []antlr.TerminalNode {
	return s.GetTokens(DokafileParserSpacing)
}

func (s *Elif_instructionContext) Spacing(i int) antlr.TerminalNode {
	return s.GetToken(DokafileParserSpacing, i)
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
		p.SetState(67)
		p.Match(DokafileParserElif)
	}
	{
		p.SetState(68)
		p.Match(DokafileParserSpacing)
	}
	{
		p.SetState(69)
		p.If_expression()
	}
	{
		p.SetState(70)
		p.Match(DokafileParserT__0)
	}
	{
		p.SetState(71)
		p.Instructions()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DokafileParserSpacing {
			{
				p.SetState(72)
				p.Match(DokafileParserSpacing)
			}

		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case DokafileParserElif:
			{
				p.SetState(75)
				p.Elif_instruction()
			}

		case DokafileParserElse:
			{
				p.SetState(76)
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
		p.SetState(81)
		p.Match(DokafileParserElse)
	}
	{
		p.SetState(82)
		p.Match(DokafileParserT__0)
	}
	{
		p.SetState(83)
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

func (s *If_expressionContext) AllSpacing() []antlr.TerminalNode {
	return s.GetTokens(DokafileParserSpacing)
}

func (s *If_expressionContext) Spacing(i int) antlr.TerminalNode {
	return s.GetToken(DokafileParserSpacing, i)
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
		p.SetState(85)
		p.Value()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DokafileParserSpacing {
		{
			p.SetState(86)
			p.Match(DokafileParserSpacing)
		}

	}
	{
		p.SetState(89)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DokafileParserOperatorEq || _la == DokafileParserOperatorNe) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DokafileParserSpacing {
		{
			p.SetState(90)
			p.Match(DokafileParserSpacing)
		}

	}
	{
		p.SetState(93)
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
		p.SetState(95)
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

// IInclude_instructionContext is an interface to support dynamic dispatch.
type IInclude_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInclude_instructionContext differentiates from other interfaces.
	IsInclude_instructionContext()
}

type Include_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclude_instructionContext() *Include_instructionContext {
	var p = new(Include_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DokafileParserRULE_include_instruction
	return p
}

func (*Include_instructionContext) IsInclude_instructionContext() {}

func NewInclude_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Include_instructionContext {
	var p = new(Include_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DokafileParserRULE_include_instruction

	return p
}

func (s *Include_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Include_instructionContext) Include() antlr.TerminalNode {
	return s.GetToken(DokafileParserInclude, 0)
}

func (s *Include_instructionContext) Spacing() antlr.TerminalNode {
	return s.GetToken(DokafileParserSpacing, 0)
}

func (s *Include_instructionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DokafileParserStringLiteral, 0)
}

func (s *Include_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Include_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Include_instructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case DokafileVisitor:
		return t.VisitInclude_instruction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *DokafileParser) Include_instruction() (localctx IInclude_instructionContext) {
	localctx = NewInclude_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DokafileParserRULE_include_instruction)

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
		p.SetState(97)
		p.Match(DokafileParserInclude)
	}
	{
		p.SetState(98)
		p.Match(DokafileParserSpacing)
	}
	{
		p.SetState(99)
		p.Match(DokafileParserStringLiteral)
	}

	return localctx
}
