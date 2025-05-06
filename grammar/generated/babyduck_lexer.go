// Code generated from ./BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type BabyDuckLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var BabyDuckLexerLexerStaticData struct {
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

func babyducklexerLexerInit() {
	staticData := &BabyDuckLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'program'", "'main'", "'end'", "'var'", "'print'", "'if'", "'else'",
		"'while'", "'do'", "'void'", "'int'", "'float'", "'string'", "'*'",
		"'/'", "'+'", "'-'", "'>'", "'<'", "'!='", "'='", "'('", "')'", "'['",
		"']'", "'{'", "'}'", "':'", "','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "PROGRAM", "MAIN", "END", "VAR", "PRINT", "IF", "ELSE", "WHILE",
		"DO", "VOID", "INT", "FLOAT", "STRING", "OP_MULTIPLY", "OP_DIVIDE",
		"OP_ADD", "OP_SUBTRACT", "GREATER", "LESS", "NOT_EQUAL", "ASSIGN", "LPAREN",
		"RPAREN", "LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "COLON", "COMMA",
		"SEMICOLON", "CONST_INT", "CONST_FLOAT", "CONST_STRING", "WHITESPACE",
		"ID",
	}
	staticData.RuleNames = []string{
		"PROGRAM", "MAIN", "END", "VAR", "PRINT", "IF", "ELSE", "WHILE", "DO",
		"VOID", "INT", "FLOAT", "STRING", "OP_MULTIPLY", "OP_DIVIDE", "OP_ADD",
		"OP_SUBTRACT", "GREATER", "LESS", "NOT_EQUAL", "ASSIGN", "LPAREN", "RPAREN",
		"LBRACKET", "RBRACKET", "LBRACE", "RBRACE", "COLON", "COMMA", "SEMICOLON",
		"CONST_INT", "CONST_FLOAT", "CONST_STRING", "WHITESPACE", "ID",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 35, 212, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 4, 30, 174,
		8, 30, 11, 30, 12, 30, 175, 1, 31, 5, 31, 179, 8, 31, 10, 31, 12, 31, 182,
		9, 31, 1, 31, 1, 31, 4, 31, 186, 8, 31, 11, 31, 12, 31, 187, 1, 32, 1,
		32, 5, 32, 192, 8, 32, 10, 32, 12, 32, 195, 9, 32, 1, 32, 1, 32, 1, 33,
		4, 33, 200, 8, 33, 11, 33, 12, 33, 201, 1, 33, 1, 33, 1, 34, 1, 34, 5,
		34, 208, 8, 34, 10, 34, 12, 34, 211, 9, 34, 0, 0, 35, 1, 1, 3, 2, 5, 3,
		7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13,
		27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
		45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31,
		63, 32, 65, 33, 67, 34, 69, 35, 1, 0, 6, 1, 0, 48, 57, 1, 0, 46, 46, 1,
		0, 34, 34, 3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 65, 90, 95, 95, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 217, 0, 1, 1, 0, 0, 0, 0, 3, 1,
		0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1,
		0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19,
		1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0,
		27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0,
		0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0,
		0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0,
		0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1,
		0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65,
		1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 1, 71, 1, 0, 0, 0, 3,
		79, 1, 0, 0, 0, 5, 84, 1, 0, 0, 0, 7, 88, 1, 0, 0, 0, 9, 92, 1, 0, 0, 0,
		11, 98, 1, 0, 0, 0, 13, 101, 1, 0, 0, 0, 15, 106, 1, 0, 0, 0, 17, 112,
		1, 0, 0, 0, 19, 115, 1, 0, 0, 0, 21, 120, 1, 0, 0, 0, 23, 124, 1, 0, 0,
		0, 25, 130, 1, 0, 0, 0, 27, 137, 1, 0, 0, 0, 29, 139, 1, 0, 0, 0, 31, 141,
		1, 0, 0, 0, 33, 143, 1, 0, 0, 0, 35, 145, 1, 0, 0, 0, 37, 147, 1, 0, 0,
		0, 39, 149, 1, 0, 0, 0, 41, 152, 1, 0, 0, 0, 43, 154, 1, 0, 0, 0, 45, 156,
		1, 0, 0, 0, 47, 158, 1, 0, 0, 0, 49, 160, 1, 0, 0, 0, 51, 162, 1, 0, 0,
		0, 53, 164, 1, 0, 0, 0, 55, 166, 1, 0, 0, 0, 57, 168, 1, 0, 0, 0, 59, 170,
		1, 0, 0, 0, 61, 173, 1, 0, 0, 0, 63, 180, 1, 0, 0, 0, 65, 189, 1, 0, 0,
		0, 67, 199, 1, 0, 0, 0, 69, 205, 1, 0, 0, 0, 71, 72, 5, 112, 0, 0, 72,
		73, 5, 114, 0, 0, 73, 74, 5, 111, 0, 0, 74, 75, 5, 103, 0, 0, 75, 76, 5,
		114, 0, 0, 76, 77, 5, 97, 0, 0, 77, 78, 5, 109, 0, 0, 78, 2, 1, 0, 0, 0,
		79, 80, 5, 109, 0, 0, 80, 81, 5, 97, 0, 0, 81, 82, 5, 105, 0, 0, 82, 83,
		5, 110, 0, 0, 83, 4, 1, 0, 0, 0, 84, 85, 5, 101, 0, 0, 85, 86, 5, 110,
		0, 0, 86, 87, 5, 100, 0, 0, 87, 6, 1, 0, 0, 0, 88, 89, 5, 118, 0, 0, 89,
		90, 5, 97, 0, 0, 90, 91, 5, 114, 0, 0, 91, 8, 1, 0, 0, 0, 92, 93, 5, 112,
		0, 0, 93, 94, 5, 114, 0, 0, 94, 95, 5, 105, 0, 0, 95, 96, 5, 110, 0, 0,
		96, 97, 5, 116, 0, 0, 97, 10, 1, 0, 0, 0, 98, 99, 5, 105, 0, 0, 99, 100,
		5, 102, 0, 0, 100, 12, 1, 0, 0, 0, 101, 102, 5, 101, 0, 0, 102, 103, 5,
		108, 0, 0, 103, 104, 5, 115, 0, 0, 104, 105, 5, 101, 0, 0, 105, 14, 1,
		0, 0, 0, 106, 107, 5, 119, 0, 0, 107, 108, 5, 104, 0, 0, 108, 109, 5, 105,
		0, 0, 109, 110, 5, 108, 0, 0, 110, 111, 5, 101, 0, 0, 111, 16, 1, 0, 0,
		0, 112, 113, 5, 100, 0, 0, 113, 114, 5, 111, 0, 0, 114, 18, 1, 0, 0, 0,
		115, 116, 5, 118, 0, 0, 116, 117, 5, 111, 0, 0, 117, 118, 5, 105, 0, 0,
		118, 119, 5, 100, 0, 0, 119, 20, 1, 0, 0, 0, 120, 121, 5, 105, 0, 0, 121,
		122, 5, 110, 0, 0, 122, 123, 5, 116, 0, 0, 123, 22, 1, 0, 0, 0, 124, 125,
		5, 102, 0, 0, 125, 126, 5, 108, 0, 0, 126, 127, 5, 111, 0, 0, 127, 128,
		5, 97, 0, 0, 128, 129, 5, 116, 0, 0, 129, 24, 1, 0, 0, 0, 130, 131, 5,
		115, 0, 0, 131, 132, 5, 116, 0, 0, 132, 133, 5, 114, 0, 0, 133, 134, 5,
		105, 0, 0, 134, 135, 5, 110, 0, 0, 135, 136, 5, 103, 0, 0, 136, 26, 1,
		0, 0, 0, 137, 138, 5, 42, 0, 0, 138, 28, 1, 0, 0, 0, 139, 140, 5, 47, 0,
		0, 140, 30, 1, 0, 0, 0, 141, 142, 5, 43, 0, 0, 142, 32, 1, 0, 0, 0, 143,
		144, 5, 45, 0, 0, 144, 34, 1, 0, 0, 0, 145, 146, 5, 62, 0, 0, 146, 36,
		1, 0, 0, 0, 147, 148, 5, 60, 0, 0, 148, 38, 1, 0, 0, 0, 149, 150, 5, 33,
		0, 0, 150, 151, 5, 61, 0, 0, 151, 40, 1, 0, 0, 0, 152, 153, 5, 61, 0, 0,
		153, 42, 1, 0, 0, 0, 154, 155, 5, 40, 0, 0, 155, 44, 1, 0, 0, 0, 156, 157,
		5, 41, 0, 0, 157, 46, 1, 0, 0, 0, 158, 159, 5, 91, 0, 0, 159, 48, 1, 0,
		0, 0, 160, 161, 5, 93, 0, 0, 161, 50, 1, 0, 0, 0, 162, 163, 5, 123, 0,
		0, 163, 52, 1, 0, 0, 0, 164, 165, 5, 125, 0, 0, 165, 54, 1, 0, 0, 0, 166,
		167, 5, 58, 0, 0, 167, 56, 1, 0, 0, 0, 168, 169, 5, 44, 0, 0, 169, 58,
		1, 0, 0, 0, 170, 171, 5, 59, 0, 0, 171, 60, 1, 0, 0, 0, 172, 174, 7, 0,
		0, 0, 173, 172, 1, 0, 0, 0, 174, 175, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0,
		175, 176, 1, 0, 0, 0, 176, 62, 1, 0, 0, 0, 177, 179, 7, 0, 0, 0, 178, 177,
		1, 0, 0, 0, 179, 182, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0, 180, 181, 1, 0,
		0, 0, 181, 183, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 183, 185, 7, 1, 0, 0,
		184, 186, 7, 0, 0, 0, 185, 184, 1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187,
		185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 64, 1, 0, 0, 0, 189, 193, 5,
		34, 0, 0, 190, 192, 8, 2, 0, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0,
		0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 196, 1, 0, 0, 0, 195,
		193, 1, 0, 0, 0, 196, 197, 5, 34, 0, 0, 197, 66, 1, 0, 0, 0, 198, 200,
		7, 3, 0, 0, 199, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 199, 1, 0,
		0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 204, 6, 33, 0, 0,
		204, 68, 1, 0, 0, 0, 205, 209, 7, 4, 0, 0, 206, 208, 7, 5, 0, 0, 207, 206,
		1, 0, 0, 0, 208, 211, 1, 0, 0, 0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0,
		0, 0, 210, 70, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 7, 0, 175, 180, 187, 193,
		201, 209, 1, 6, 0, 0,
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

// BabyDuckLexerInit initializes any static state used to implement BabyDuckLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewBabyDuckLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func BabyDuckLexerInit() {
	staticData := &BabyDuckLexerLexerStaticData
	staticData.once.Do(babyducklexerLexerInit)
}

// NewBabyDuckLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewBabyDuckLexer(input antlr.CharStream) *BabyDuckLexer {
	BabyDuckLexerInit()
	l := new(BabyDuckLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &BabyDuckLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "BabyDuck.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// BabyDuckLexer tokens.
const (
	BabyDuckLexerPROGRAM      = 1
	BabyDuckLexerMAIN         = 2
	BabyDuckLexerEND          = 3
	BabyDuckLexerVAR          = 4
	BabyDuckLexerPRINT        = 5
	BabyDuckLexerIF           = 6
	BabyDuckLexerELSE         = 7
	BabyDuckLexerWHILE        = 8
	BabyDuckLexerDO           = 9
	BabyDuckLexerVOID         = 10
	BabyDuckLexerINT          = 11
	BabyDuckLexerFLOAT        = 12
	BabyDuckLexerSTRING       = 13
	BabyDuckLexerOP_MULTIPLY  = 14
	BabyDuckLexerOP_DIVIDE    = 15
	BabyDuckLexerOP_ADD       = 16
	BabyDuckLexerOP_SUBTRACT  = 17
	BabyDuckLexerGREATER      = 18
	BabyDuckLexerLESS         = 19
	BabyDuckLexerNOT_EQUAL    = 20
	BabyDuckLexerASSIGN       = 21
	BabyDuckLexerLPAREN       = 22
	BabyDuckLexerRPAREN       = 23
	BabyDuckLexerLBRACKET     = 24
	BabyDuckLexerRBRACKET     = 25
	BabyDuckLexerLBRACE       = 26
	BabyDuckLexerRBRACE       = 27
	BabyDuckLexerCOLON        = 28
	BabyDuckLexerCOMMA        = 29
	BabyDuckLexerSEMICOLON    = 30
	BabyDuckLexerCONST_INT    = 31
	BabyDuckLexerCONST_FLOAT  = 32
	BabyDuckLexerCONST_STRING = 33
	BabyDuckLexerWHITESPACE   = 34
	BabyDuckLexerID           = 35
)
