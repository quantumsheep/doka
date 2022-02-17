// Code generated from Dokafile.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 162,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 7, 9, 63, 10, 9, 12, 9, 14, 9, 66, 11, 9, 3, 10, 3, 10,
	3, 10, 5, 10, 71, 10, 10, 3, 11, 6, 11, 74, 10, 11, 13, 11, 14, 11, 75,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5,
	12, 109, 10, 12, 3, 13, 3, 13, 7, 13, 113, 10, 13, 12, 13, 14, 13, 116,
	11, 13, 3, 13, 3, 13, 3, 13, 6, 13, 121, 10, 13, 13, 13, 14, 13, 122, 3,
	13, 3, 13, 5, 13, 127, 10, 13, 3, 14, 3, 14, 7, 14, 131, 10, 14, 12, 14,
	14, 14, 134, 11, 14, 3, 15, 3, 15, 7, 15, 138, 10, 15, 12, 15, 14, 15,
	141, 11, 15, 3, 16, 3, 16, 5, 16, 145, 10, 16, 3, 17, 3, 17, 5, 17, 149,
	10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 5, 18, 161, 10, 18, 2, 2, 19, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 2, 21, 11, 23, 2, 25, 12, 27, 13, 29, 2, 31, 2, 33,
	2, 35, 2, 3, 2, 8, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 5, 2, 67,
	92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 12, 12,
	15, 15, 36, 36, 5, 2, 12, 12, 15, 15, 41, 41, 2, 174, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 3, 37, 3, 2, 2, 2, 5, 39, 3, 2,
	2, 2, 7, 41, 3, 2, 2, 2, 9, 45, 3, 2, 2, 2, 11, 50, 3, 2, 2, 2, 13, 53,
	3, 2, 2, 2, 15, 56, 3, 2, 2, 2, 17, 59, 3, 2, 2, 2, 19, 70, 3, 2, 2, 2,
	21, 73, 3, 2, 2, 2, 23, 108, 3, 2, 2, 2, 25, 126, 3, 2, 2, 2, 27, 128,
	3, 2, 2, 2, 29, 135, 3, 2, 2, 2, 31, 144, 3, 2, 2, 2, 33, 148, 3, 2, 2,
	2, 35, 160, 3, 2, 2, 2, 37, 38, 7, 12, 2, 2, 38, 4, 3, 2, 2, 2, 39, 40,
	7, 34, 2, 2, 40, 6, 3, 2, 2, 2, 41, 42, 7, 35, 2, 2, 42, 43, 7, 75, 2,
	2, 43, 44, 7, 72, 2, 2, 44, 8, 3, 2, 2, 2, 45, 46, 7, 35, 2, 2, 46, 47,
	7, 71, 2, 2, 47, 48, 7, 80, 2, 2, 48, 49, 7, 70, 2, 2, 49, 10, 3, 2, 2,
	2, 50, 51, 7, 63, 2, 2, 51, 52, 7, 63, 2, 2, 52, 12, 3, 2, 2, 2, 53, 54,
	7, 35, 2, 2, 54, 55, 7, 63, 2, 2, 55, 14, 3, 2, 2, 2, 56, 57, 7, 38, 2,
	2, 57, 58, 5, 29, 15, 2, 58, 16, 3, 2, 2, 2, 59, 60, 5, 23, 12, 2, 60,
	64, 7, 34, 2, 2, 61, 63, 5, 19, 10, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3, 2,
	2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 18, 3, 2, 2, 2, 66, 64,
	3, 2, 2, 2, 67, 71, 10, 2, 2, 2, 68, 69, 7, 94, 2, 2, 69, 71, 7, 12, 2,
	2, 70, 67, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 20, 3, 2, 2, 2, 72, 74,
	9, 3, 2, 2, 73, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2,
	75, 76, 3, 2, 2, 2, 76, 22, 3, 2, 2, 2, 77, 78, 7, 72, 2, 2, 78, 79, 7,
	84, 2, 2, 79, 80, 7, 81, 2, 2, 80, 109, 7, 79, 2, 2, 81, 82, 7, 89, 2,
	2, 82, 83, 7, 81, 2, 2, 83, 84, 7, 84, 2, 2, 84, 85, 7, 77, 2, 2, 85, 86,
	7, 70, 2, 2, 86, 87, 7, 75, 2, 2, 87, 109, 7, 84, 2, 2, 88, 89, 7, 69,
	2, 2, 89, 90, 7, 81, 2, 2, 90, 91, 7, 82, 2, 2, 91, 109, 7, 91, 2, 2, 92,
	93, 7, 84, 2, 2, 93, 94, 7, 87, 2, 2, 94, 109, 7, 80, 2, 2, 95, 96, 7,
	69, 2, 2, 96, 97, 7, 79, 2, 2, 97, 109, 7, 70, 2, 2, 98, 99, 7, 71, 2,
	2, 99, 100, 7, 80, 2, 2, 100, 101, 7, 86, 2, 2, 101, 102, 7, 84, 2, 2,
	102, 103, 7, 91, 2, 2, 103, 104, 7, 82, 2, 2, 104, 105, 7, 81, 2, 2, 105,
	106, 7, 75, 2, 2, 106, 107, 7, 80, 2, 2, 107, 109, 7, 86, 2, 2, 108, 77,
	3, 2, 2, 2, 108, 81, 3, 2, 2, 2, 108, 88, 3, 2, 2, 2, 108, 92, 3, 2, 2,
	2, 108, 95, 3, 2, 2, 2, 108, 98, 3, 2, 2, 2, 109, 24, 3, 2, 2, 2, 110,
	114, 7, 36, 2, 2, 111, 113, 5, 31, 16, 2, 112, 111, 3, 2, 2, 2, 113, 116,
	3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 117, 3, 2,
	2, 2, 116, 114, 3, 2, 2, 2, 117, 127, 7, 36, 2, 2, 118, 120, 7, 41, 2,
	2, 119, 121, 5, 33, 17, 2, 120, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2,
	122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124,
	125, 7, 41, 2, 2, 125, 127, 3, 2, 2, 2, 126, 110, 3, 2, 2, 2, 126, 118,
	3, 2, 2, 2, 127, 26, 3, 2, 2, 2, 128, 132, 7, 37, 2, 2, 129, 131, 10, 2,
	2, 2, 130, 129, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2,
	132, 133, 3, 2, 2, 2, 133, 28, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 139,
	9, 4, 2, 2, 136, 138, 9, 5, 2, 2, 137, 136, 3, 2, 2, 2, 138, 141, 3, 2,
	2, 2, 139, 137, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 30, 3, 2, 2, 2,
	141, 139, 3, 2, 2, 2, 142, 145, 10, 6, 2, 2, 143, 145, 5, 35, 18, 2, 144,
	142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 32, 3, 2, 2, 2, 146, 149, 10,
	7, 2, 2, 147, 149, 5, 35, 18, 2, 148, 146, 3, 2, 2, 2, 148, 147, 3, 2,
	2, 2, 149, 34, 3, 2, 2, 2, 150, 151, 7, 94, 2, 2, 151, 161, 7, 41, 2, 2,
	152, 153, 7, 94, 2, 2, 153, 161, 7, 36, 2, 2, 154, 155, 7, 94, 2, 2, 155,
	161, 7, 94, 2, 2, 156, 157, 7, 94, 2, 2, 157, 161, 7, 112, 2, 2, 158, 159,
	7, 94, 2, 2, 159, 161, 7, 116, 2, 2, 160, 150, 3, 2, 2, 2, 160, 152, 3,
	2, 2, 2, 160, 154, 3, 2, 2, 2, 160, 156, 3, 2, 2, 2, 160, 158, 3, 2, 2,
	2, 161, 36, 3, 2, 2, 2, 15, 2, 64, 70, 75, 108, 114, 122, 126, 132, 139,
	144, 148, 160, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'\n'", "' '", "'!IF'", "'!END'", "'=='", "'!='",
}

var lexerSymbolicNames = []string{
	"", "", "", "If", "End", "OperatorEq", "OperatorNe", "Variable", "NativeInstructionCall",
	"Indent", "StringLiteral", "Comment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "If", "End", "OperatorEq", "OperatorNe", "Variable", "NativeInstructionCall",
	"NativeInstructionCallChar", "Indent", "NativeInstruction", "StringLiteral",
	"Comment", "Name", "StringChar", "CharChar", "Escape",
}

type DokafileLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewDokafileLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *DokafileLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDokafileLexer(input antlr.CharStream) *DokafileLexer {
	l := new(DokafileLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Dokafile.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DokafileLexer tokens.
const (
	DokafileLexerT__0                  = 1
	DokafileLexerT__1                  = 2
	DokafileLexerIf                    = 3
	DokafileLexerEnd                   = 4
	DokafileLexerOperatorEq            = 5
	DokafileLexerOperatorNe            = 6
	DokafileLexerVariable              = 7
	DokafileLexerNativeInstructionCall = 8
	DokafileLexerIndent                = 9
	DokafileLexerStringLiteral         = 10
	DokafileLexerComment               = 11
)
