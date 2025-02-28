// Code generated from Challenge.g4 by ANTLR 4.13.2. DO NOT EDIT.

package challenge

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ChallengeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ChallengeLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func challengelexerLexerInit() {
	staticData := &ChallengeLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'\\t'", "' '", "'!'", "'\"'", "'#'", "'$'", "'%'", "'&'", "'''",
		"'('", "')'", "'*'", "'+'", "','", "'-'", "'.'", "'/'", "", "':'", "';'",
		"'<'", "'='", "'>'", "'?'", "'@'", "", "'['", "'\\'", "']'", "'^'",
		"'_'", "'`'", "'{'", "'|'", "'}'", "'~'",
	}
	staticData.SymbolicNames = []string{
		"", "HTAB", "SP", "EXCLAMATION_MARK", "DQUOTE", "HASH", "DOLLAR", "PERCENT",
		"AMPERSAND", "SQUOTE", "OPEN_PARENS", "CLOSE_PARENS", "ASTERISK", "PLUS",
		"COMMA", "MINUS", "PERIOD", "SLASH", "DIGIT", "COLON", "SEMICOLON",
		"LESS_THAN", "EQUALS", "GREATER_THAN", "QUESTION", "AT", "ALPHA", "OPEN_BRACKET",
		"BACKSLASH", "CLOSE_BRACKET", "CARET", "UNDERSCORE", "GRAVE", "OPEN_BRACE",
		"PIPE", "CLOSE_BRACE", "TILDE", "EXTENDED_ASCII",
	}
	staticData.RuleNames = []string{
		"HTAB", "SP", "EXCLAMATION_MARK", "DQUOTE", "HASH", "DOLLAR", "PERCENT",
		"AMPERSAND", "SQUOTE", "OPEN_PARENS", "CLOSE_PARENS", "ASTERISK", "PLUS",
		"COMMA", "MINUS", "PERIOD", "SLASH", "DIGIT", "COLON", "SEMICOLON",
		"LESS_THAN", "EQUALS", "GREATER_THAN", "QUESTION", "AT", "ALPHA", "OPEN_BRACKET",
		"BACKSLASH", "CLOSE_BRACKET", "CARET", "UNDERSCORE", "GRAVE", "OPEN_BRACE",
		"PIPE", "CLOSE_BRACE", "TILDE", "EXTENDED_ASCII",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 37, 149, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1,
		32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 0, 0,
		37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 1,
		0, 2, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 148, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1, 0, 0,
		0, 7, 81, 1, 0, 0, 0, 9, 83, 1, 0, 0, 0, 11, 85, 1, 0, 0, 0, 13, 87, 1,
		0, 0, 0, 15, 89, 1, 0, 0, 0, 17, 91, 1, 0, 0, 0, 19, 93, 1, 0, 0, 0, 21,
		95, 1, 0, 0, 0, 23, 97, 1, 0, 0, 0, 25, 99, 1, 0, 0, 0, 27, 101, 1, 0,
		0, 0, 29, 103, 1, 0, 0, 0, 31, 105, 1, 0, 0, 0, 33, 107, 1, 0, 0, 0, 35,
		109, 1, 0, 0, 0, 37, 111, 1, 0, 0, 0, 39, 113, 1, 0, 0, 0, 41, 115, 1,
		0, 0, 0, 43, 117, 1, 0, 0, 0, 45, 119, 1, 0, 0, 0, 47, 121, 1, 0, 0, 0,
		49, 123, 1, 0, 0, 0, 51, 125, 1, 0, 0, 0, 53, 127, 1, 0, 0, 0, 55, 129,
		1, 0, 0, 0, 57, 131, 1, 0, 0, 0, 59, 133, 1, 0, 0, 0, 61, 135, 1, 0, 0,
		0, 63, 137, 1, 0, 0, 0, 65, 139, 1, 0, 0, 0, 67, 141, 1, 0, 0, 0, 69, 143,
		1, 0, 0, 0, 71, 145, 1, 0, 0, 0, 73, 147, 1, 0, 0, 0, 75, 76, 5, 9, 0,
		0, 76, 2, 1, 0, 0, 0, 77, 78, 5, 32, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5,
		33, 0, 0, 80, 6, 1, 0, 0, 0, 81, 82, 5, 34, 0, 0, 82, 8, 1, 0, 0, 0, 83,
		84, 5, 35, 0, 0, 84, 10, 1, 0, 0, 0, 85, 86, 5, 36, 0, 0, 86, 12, 1, 0,
		0, 0, 87, 88, 5, 37, 0, 0, 88, 14, 1, 0, 0, 0, 89, 90, 5, 38, 0, 0, 90,
		16, 1, 0, 0, 0, 91, 92, 5, 39, 0, 0, 92, 18, 1, 0, 0, 0, 93, 94, 5, 40,
		0, 0, 94, 20, 1, 0, 0, 0, 95, 96, 5, 41, 0, 0, 96, 22, 1, 0, 0, 0, 97,
		98, 5, 42, 0, 0, 98, 24, 1, 0, 0, 0, 99, 100, 5, 43, 0, 0, 100, 26, 1,
		0, 0, 0, 101, 102, 5, 44, 0, 0, 102, 28, 1, 0, 0, 0, 103, 104, 5, 45, 0,
		0, 104, 30, 1, 0, 0, 0, 105, 106, 5, 46, 0, 0, 106, 32, 1, 0, 0, 0, 107,
		108, 5, 47, 0, 0, 108, 34, 1, 0, 0, 0, 109, 110, 7, 0, 0, 0, 110, 36, 1,
		0, 0, 0, 111, 112, 5, 58, 0, 0, 112, 38, 1, 0, 0, 0, 113, 114, 5, 59, 0,
		0, 114, 40, 1, 0, 0, 0, 115, 116, 5, 60, 0, 0, 116, 42, 1, 0, 0, 0, 117,
		118, 5, 61, 0, 0, 118, 44, 1, 0, 0, 0, 119, 120, 5, 62, 0, 0, 120, 46,
		1, 0, 0, 0, 121, 122, 5, 63, 0, 0, 122, 48, 1, 0, 0, 0, 123, 124, 5, 64,
		0, 0, 124, 50, 1, 0, 0, 0, 125, 126, 7, 1, 0, 0, 126, 52, 1, 0, 0, 0, 127,
		128, 5, 91, 0, 0, 128, 54, 1, 0, 0, 0, 129, 130, 5, 92, 0, 0, 130, 56,
		1, 0, 0, 0, 131, 132, 5, 93, 0, 0, 132, 58, 1, 0, 0, 0, 133, 134, 5, 94,
		0, 0, 134, 60, 1, 0, 0, 0, 135, 136, 5, 95, 0, 0, 136, 62, 1, 0, 0, 0,
		137, 138, 5, 96, 0, 0, 138, 64, 1, 0, 0, 0, 139, 140, 5, 123, 0, 0, 140,
		66, 1, 0, 0, 0, 141, 142, 5, 124, 0, 0, 142, 68, 1, 0, 0, 0, 143, 144,
		5, 125, 0, 0, 144, 70, 1, 0, 0, 0, 145, 146, 5, 126, 0, 0, 146, 72, 1,
		0, 0, 0, 147, 148, 2, 128, 255, 0, 148, 74, 1, 0, 0, 0, 1, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ChallengeLexerInit initializes any static state used to implement ChallengeLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewChallengeLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ChallengeLexerInit() {
	staticData := &ChallengeLexerLexerStaticData
	staticData.once.Do(challengelexerLexerInit)
}

// NewChallengeLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewChallengeLexer(input antlr.CharStream) *ChallengeLexer {
	ChallengeLexerInit()
	l := new(ChallengeLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ChallengeLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Challenge.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ChallengeLexer tokens.
const (
	ChallengeLexerHTAB             = 1
	ChallengeLexerSP               = 2
	ChallengeLexerEXCLAMATION_MARK = 3
	ChallengeLexerDQUOTE           = 4
	ChallengeLexerHASH             = 5
	ChallengeLexerDOLLAR           = 6
	ChallengeLexerPERCENT          = 7
	ChallengeLexerAMPERSAND        = 8
	ChallengeLexerSQUOTE           = 9
	ChallengeLexerOPEN_PARENS      = 10
	ChallengeLexerCLOSE_PARENS     = 11
	ChallengeLexerASTERISK         = 12
	ChallengeLexerPLUS             = 13
	ChallengeLexerCOMMA            = 14
	ChallengeLexerMINUS            = 15
	ChallengeLexerPERIOD           = 16
	ChallengeLexerSLASH            = 17
	ChallengeLexerDIGIT            = 18
	ChallengeLexerCOLON            = 19
	ChallengeLexerSEMICOLON        = 20
	ChallengeLexerLESS_THAN        = 21
	ChallengeLexerEQUALS           = 22
	ChallengeLexerGREATER_THAN     = 23
	ChallengeLexerQUESTION         = 24
	ChallengeLexerAT               = 25
	ChallengeLexerALPHA            = 26
	ChallengeLexerOPEN_BRACKET     = 27
	ChallengeLexerBACKSLASH        = 28
	ChallengeLexerCLOSE_BRACKET    = 29
	ChallengeLexerCARET            = 30
	ChallengeLexerUNDERSCORE       = 31
	ChallengeLexerGRAVE            = 32
	ChallengeLexerOPEN_BRACE       = 33
	ChallengeLexerPIPE             = 34
	ChallengeLexerCLOSE_BRACE      = 35
	ChallengeLexerTILDE            = 36
	ChallengeLexerEXTENDED_ASCII   = 37
)
