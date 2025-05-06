// Code generated from ./BabyDuck.g4 by ANTLR 4.13.2. DO NOT EDIT.

package generated // BabyDuck

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type BabyDuckParser struct {
	*antlr.BaseParser
}

var BabyDuckParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func babyduckParserInit() {
	staticData := &BabyDuckParserStaticData
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
		"program", "programHeader", "programBody", "identifier", "parenthesizedExpression",
		"type", "idList", "varDecl", "varsSection", "constant", "parameter",
		"functionBody", "functionDeclaration", "functionsSection", "codeBlock",
		"statement", "assignment", "ifBlock", "elseBlock", "conditional", "loop",
		"argumentList", "functionCall", "printable", "printStatement", "expression",
		"relationalOperator", "relationalExpression", "additiveOperator", "arithmeticExpression",
		"multiplicativeOperator", "term", "factor", "valueWithOptionalSign",
		"value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 35, 267, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 1, 0, 3, 0, 73, 8,
		0, 1, 0, 5, 0, 76, 8, 0, 10, 0, 12, 0, 79, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 5, 6, 102, 8, 6, 10, 6, 12, 6, 105, 9, 6,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 4, 8, 114, 8, 8, 11, 8, 12, 8,
		115, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 3, 11, 126,
		8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		12, 5, 12, 138, 8, 12, 10, 12, 12, 12, 141, 9, 12, 3, 12, 143, 8, 12, 1,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 152, 8, 14, 10, 14,
		12, 14, 155, 9, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3,
		15, 164, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 3, 19, 180, 8, 19, 1, 19, 1,
		19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 5, 21,
		193, 8, 21, 10, 21, 12, 21, 196, 9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 201,
		8, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 208, 8, 23, 1, 24, 1,
		24, 1, 24, 1, 24, 1, 24, 5, 24, 215, 8, 24, 10, 24, 12, 24, 218, 9, 24,
		1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 3, 25, 225, 8, 25, 1, 26, 1, 26, 1,
		27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 238,
		8, 29, 10, 29, 12, 29, 241, 9, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 5, 31, 249, 8, 31, 10, 31, 12, 31, 252, 9, 31, 1, 32, 1, 32, 3, 32,
		256, 8, 32, 1, 33, 3, 33, 259, 8, 33, 1, 33, 1, 33, 1, 34, 1, 34, 3, 34,
		265, 8, 34, 1, 34, 0, 0, 35, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 0, 5, 1, 0, 11, 12, 1, 0, 31, 32, 1, 0, 18, 20, 1,
		0, 16, 17, 1, 0, 14, 15, 254, 0, 70, 1, 0, 0, 0, 2, 82, 1, 0, 0, 0, 4,
		86, 1, 0, 0, 0, 6, 90, 1, 0, 0, 0, 8, 92, 1, 0, 0, 0, 10, 96, 1, 0, 0,
		0, 12, 98, 1, 0, 0, 0, 14, 106, 1, 0, 0, 0, 16, 111, 1, 0, 0, 0, 18, 117,
		1, 0, 0, 0, 20, 119, 1, 0, 0, 0, 22, 123, 1, 0, 0, 0, 24, 131, 1, 0, 0,
		0, 26, 147, 1, 0, 0, 0, 28, 149, 1, 0, 0, 0, 30, 163, 1, 0, 0, 0, 32, 165,
		1, 0, 0, 0, 34, 170, 1, 0, 0, 0, 36, 174, 1, 0, 0, 0, 38, 177, 1, 0, 0,
		0, 40, 183, 1, 0, 0, 0, 42, 189, 1, 0, 0, 0, 44, 197, 1, 0, 0, 0, 46, 207,
		1, 0, 0, 0, 48, 209, 1, 0, 0, 0, 50, 222, 1, 0, 0, 0, 52, 226, 1, 0, 0,
		0, 54, 228, 1, 0, 0, 0, 56, 231, 1, 0, 0, 0, 58, 233, 1, 0, 0, 0, 60, 242,
		1, 0, 0, 0, 62, 244, 1, 0, 0, 0, 64, 255, 1, 0, 0, 0, 66, 258, 1, 0, 0,
		0, 68, 264, 1, 0, 0, 0, 70, 72, 3, 2, 1, 0, 71, 73, 3, 16, 8, 0, 72, 71,
		1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 77, 1, 0, 0, 0, 74, 76, 3, 26, 13,
		0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78,
		1, 0, 0, 0, 78, 80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 3, 4, 2, 0,
		81, 1, 1, 0, 0, 0, 82, 83, 5, 1, 0, 0, 83, 84, 5, 35, 0, 0, 84, 85, 5,
		30, 0, 0, 85, 3, 1, 0, 0, 0, 86, 87, 5, 2, 0, 0, 87, 88, 3, 28, 14, 0,
		88, 89, 5, 3, 0, 0, 89, 5, 1, 0, 0, 0, 90, 91, 5, 35, 0, 0, 91, 7, 1, 0,
		0, 0, 92, 93, 5, 22, 0, 0, 93, 94, 3, 50, 25, 0, 94, 95, 5, 23, 0, 0, 95,
		9, 1, 0, 0, 0, 96, 97, 7, 0, 0, 0, 97, 11, 1, 0, 0, 0, 98, 103, 5, 35,
		0, 0, 99, 100, 5, 29, 0, 0, 100, 102, 5, 35, 0, 0, 101, 99, 1, 0, 0, 0,
		102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104,
		13, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 107, 3, 12, 6, 0, 107, 108,
		5, 28, 0, 0, 108, 109, 3, 10, 5, 0, 109, 110, 5, 30, 0, 0, 110, 15, 1,
		0, 0, 0, 111, 113, 5, 4, 0, 0, 112, 114, 3, 14, 7, 0, 113, 112, 1, 0, 0,
		0, 114, 115, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 115, 116, 1, 0, 0, 0, 116,
		17, 1, 0, 0, 0, 117, 118, 7, 1, 0, 0, 118, 19, 1, 0, 0, 0, 119, 120, 3,
		6, 3, 0, 120, 121, 5, 28, 0, 0, 121, 122, 3, 10, 5, 0, 122, 21, 1, 0, 0,
		0, 123, 125, 5, 24, 0, 0, 124, 126, 3, 16, 8, 0, 125, 124, 1, 0, 0, 0,
		125, 126, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 3, 28, 14, 0, 128,
		129, 5, 25, 0, 0, 129, 130, 5, 30, 0, 0, 130, 23, 1, 0, 0, 0, 131, 132,
		5, 10, 0, 0, 132, 133, 3, 6, 3, 0, 133, 142, 5, 22, 0, 0, 134, 139, 3,
		20, 10, 0, 135, 136, 5, 29, 0, 0, 136, 138, 3, 20, 10, 0, 137, 135, 1,
		0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0, 0, 0, 139, 140, 1, 0, 0,
		0, 140, 143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 142, 134, 1, 0, 0, 0, 142,
		143, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 5, 23, 0, 0, 145, 146,
		3, 22, 11, 0, 146, 25, 1, 0, 0, 0, 147, 148, 3, 24, 12, 0, 148, 27, 1,
		0, 0, 0, 149, 153, 5, 26, 0, 0, 150, 152, 3, 30, 15, 0, 151, 150, 1, 0,
		0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0,
		154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157, 5, 27, 0, 0, 157,
		29, 1, 0, 0, 0, 158, 164, 3, 32, 16, 0, 159, 164, 3, 38, 19, 0, 160, 164,
		3, 40, 20, 0, 161, 164, 3, 44, 22, 0, 162, 164, 3, 48, 24, 0, 163, 158,
		1, 0, 0, 0, 163, 159, 1, 0, 0, 0, 163, 160, 1, 0, 0, 0, 163, 161, 1, 0,
		0, 0, 163, 162, 1, 0, 0, 0, 164, 31, 1, 0, 0, 0, 165, 166, 3, 6, 3, 0,
		166, 167, 5, 21, 0, 0, 167, 168, 3, 50, 25, 0, 168, 169, 5, 30, 0, 0, 169,
		33, 1, 0, 0, 0, 170, 171, 5, 6, 0, 0, 171, 172, 3, 8, 4, 0, 172, 173, 3,
		28, 14, 0, 173, 35, 1, 0, 0, 0, 174, 175, 5, 7, 0, 0, 175, 176, 3, 28,
		14, 0, 176, 37, 1, 0, 0, 0, 177, 179, 3, 34, 17, 0, 178, 180, 3, 36, 18,
		0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 181, 1, 0, 0, 0, 181,
		182, 5, 30, 0, 0, 182, 39, 1, 0, 0, 0, 183, 184, 5, 8, 0, 0, 184, 185,
		3, 8, 4, 0, 185, 186, 5, 9, 0, 0, 186, 187, 3, 28, 14, 0, 187, 188, 5,
		30, 0, 0, 188, 41, 1, 0, 0, 0, 189, 194, 3, 50, 25, 0, 190, 191, 5, 29,
		0, 0, 191, 193, 3, 50, 25, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1, 0, 0,
		0, 194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 43, 1, 0, 0, 0, 196,
		194, 1, 0, 0, 0, 197, 198, 3, 6, 3, 0, 198, 200, 5, 22, 0, 0, 199, 201,
		3, 42, 21, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1,
		0, 0, 0, 202, 203, 5, 23, 0, 0, 203, 204, 5, 30, 0, 0, 204, 45, 1, 0, 0,
		0, 205, 208, 3, 50, 25, 0, 206, 208, 5, 33, 0, 0, 207, 205, 1, 0, 0, 0,
		207, 206, 1, 0, 0, 0, 208, 47, 1, 0, 0, 0, 209, 210, 5, 5, 0, 0, 210, 211,
		5, 22, 0, 0, 211, 216, 3, 46, 23, 0, 212, 213, 5, 29, 0, 0, 213, 215, 3,
		46, 23, 0, 214, 212, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0,
		0, 0, 216, 217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0,
		219, 220, 5, 23, 0, 0, 220, 221, 5, 30, 0, 0, 221, 49, 1, 0, 0, 0, 222,
		224, 3, 58, 29, 0, 223, 225, 3, 54, 27, 0, 224, 223, 1, 0, 0, 0, 224, 225,
		1, 0, 0, 0, 225, 51, 1, 0, 0, 0, 226, 227, 7, 2, 0, 0, 227, 53, 1, 0, 0,
		0, 228, 229, 3, 52, 26, 0, 229, 230, 3, 58, 29, 0, 230, 55, 1, 0, 0, 0,
		231, 232, 7, 3, 0, 0, 232, 57, 1, 0, 0, 0, 233, 239, 3, 62, 31, 0, 234,
		235, 3, 56, 28, 0, 235, 236, 3, 62, 31, 0, 236, 238, 1, 0, 0, 0, 237, 234,
		1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0, 239, 240, 1, 0,
		0, 0, 240, 59, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0, 242, 243, 7, 4, 0, 0,
		243, 61, 1, 0, 0, 0, 244, 250, 3, 64, 32, 0, 245, 246, 3, 60, 30, 0, 246,
		247, 3, 64, 32, 0, 247, 249, 1, 0, 0, 0, 248, 245, 1, 0, 0, 0, 249, 252,
		1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 63, 1, 0,
		0, 0, 252, 250, 1, 0, 0, 0, 253, 256, 3, 8, 4, 0, 254, 256, 3, 66, 33,
		0, 255, 253, 1, 0, 0, 0, 255, 254, 1, 0, 0, 0, 256, 65, 1, 0, 0, 0, 257,
		259, 3, 56, 28, 0, 258, 257, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260,
		1, 0, 0, 0, 260, 261, 3, 68, 34, 0, 261, 67, 1, 0, 0, 0, 262, 265, 3, 6,
		3, 0, 263, 265, 3, 18, 9, 0, 264, 262, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0,
		265, 69, 1, 0, 0, 0, 20, 72, 77, 103, 115, 125, 139, 142, 153, 163, 179,
		194, 200, 207, 216, 224, 239, 250, 255, 258, 264,
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

// BabyDuckParserInit initializes any static state used to implement BabyDuckParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewBabyDuckParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func BabyDuckParserInit() {
	staticData := &BabyDuckParserStaticData
	staticData.once.Do(babyduckParserInit)
}

// NewBabyDuckParser produces a new parser instance for the optional input antlr.TokenStream.
func NewBabyDuckParser(input antlr.TokenStream) *BabyDuckParser {
	BabyDuckParserInit()
	this := new(BabyDuckParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &BabyDuckParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "BabyDuck.g4"

	return this
}

// BabyDuckParser tokens.
const (
	BabyDuckParserEOF          = antlr.TokenEOF
	BabyDuckParserPROGRAM      = 1
	BabyDuckParserMAIN         = 2
	BabyDuckParserEND          = 3
	BabyDuckParserVAR          = 4
	BabyDuckParserPRINT        = 5
	BabyDuckParserIF           = 6
	BabyDuckParserELSE         = 7
	BabyDuckParserWHILE        = 8
	BabyDuckParserDO           = 9
	BabyDuckParserVOID         = 10
	BabyDuckParserINT          = 11
	BabyDuckParserFLOAT        = 12
	BabyDuckParserSTRING       = 13
	BabyDuckParserOP_MULTIPLY  = 14
	BabyDuckParserOP_DIVIDE    = 15
	BabyDuckParserOP_ADD       = 16
	BabyDuckParserOP_SUBTRACT  = 17
	BabyDuckParserGREATER      = 18
	BabyDuckParserLESS         = 19
	BabyDuckParserNOT_EQUAL    = 20
	BabyDuckParserASSIGN       = 21
	BabyDuckParserLPAREN       = 22
	BabyDuckParserRPAREN       = 23
	BabyDuckParserLBRACKET     = 24
	BabyDuckParserRBRACKET     = 25
	BabyDuckParserLBRACE       = 26
	BabyDuckParserRBRACE       = 27
	BabyDuckParserCOLON        = 28
	BabyDuckParserCOMMA        = 29
	BabyDuckParserSEMICOLON    = 30
	BabyDuckParserCONST_INT    = 31
	BabyDuckParserCONST_FLOAT  = 32
	BabyDuckParserCONST_STRING = 33
	BabyDuckParserWHITESPACE   = 34
	BabyDuckParserID           = 35
)

// BabyDuckParser rules.
const (
	BabyDuckParserRULE_program                 = 0
	BabyDuckParserRULE_programHeader           = 1
	BabyDuckParserRULE_programBody             = 2
	BabyDuckParserRULE_identifier              = 3
	BabyDuckParserRULE_parenthesizedExpression = 4
	BabyDuckParserRULE_type                    = 5
	BabyDuckParserRULE_idList                  = 6
	BabyDuckParserRULE_varDecl                 = 7
	BabyDuckParserRULE_varsSection             = 8
	BabyDuckParserRULE_constant                = 9
	BabyDuckParserRULE_parameter               = 10
	BabyDuckParserRULE_functionBody            = 11
	BabyDuckParserRULE_functionDeclaration     = 12
	BabyDuckParserRULE_functionsSection        = 13
	BabyDuckParserRULE_codeBlock               = 14
	BabyDuckParserRULE_statement               = 15
	BabyDuckParserRULE_assignment              = 16
	BabyDuckParserRULE_ifBlock                 = 17
	BabyDuckParserRULE_elseBlock               = 18
	BabyDuckParserRULE_conditional             = 19
	BabyDuckParserRULE_loop                    = 20
	BabyDuckParserRULE_argumentList            = 21
	BabyDuckParserRULE_functionCall            = 22
	BabyDuckParserRULE_printable               = 23
	BabyDuckParserRULE_printStatement          = 24
	BabyDuckParserRULE_expression              = 25
	BabyDuckParserRULE_relationalOperator      = 26
	BabyDuckParserRULE_relationalExpression    = 27
	BabyDuckParserRULE_additiveOperator        = 28
	BabyDuckParserRULE_arithmeticExpression    = 29
	BabyDuckParserRULE_multiplicativeOperator  = 30
	BabyDuckParserRULE_term                    = 31
	BabyDuckParserRULE_factor                  = 32
	BabyDuckParserRULE_valueWithOptionalSign   = 33
	BabyDuckParserRULE_value                   = 34
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ProgramHeader() IProgramHeaderContext
	ProgramBody() IProgramBodyContext
	VarsSection() IVarsSectionContext
	AllFunctionsSection() []IFunctionsSectionContext
	FunctionsSection(i int) IFunctionsSectionContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) ProgramHeader() IProgramHeaderContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramHeaderContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramHeaderContext)
}

func (s *ProgramContext) ProgramBody() IProgramBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IProgramBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IProgramBodyContext)
}

func (s *ProgramContext) VarsSection() IVarsSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsSectionContext)
}

func (s *ProgramContext) AllFunctionsSection() []IFunctionsSectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionsSectionContext); ok {
			len++
		}
	}

	tst := make([]IFunctionsSectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionsSectionContext); ok {
			tst[i] = t.(IFunctionsSectionContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) FunctionsSection(i int) IFunctionsSectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsSectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsSectionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, BabyDuckParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.ProgramHeader()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserVAR {
		{
			p.SetState(71)
			p.VarsSection()
		}

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserVOID {
		{
			p.SetState(74)
			p.FunctionsSection()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.ProgramBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgramHeaderContext is an interface to support dynamic dispatch.
type IProgramHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PROGRAM() antlr.TerminalNode
	ID() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode

	// IsProgramHeaderContext differentiates from other interfaces.
	IsProgramHeaderContext()
}

type ProgramHeaderContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramHeaderContext() *ProgramHeaderContext {
	var p = new(ProgramHeaderContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_programHeader
	return p
}

func InitEmptyProgramHeaderContext(p *ProgramHeaderContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_programHeader
}

func (*ProgramHeaderContext) IsProgramHeaderContext() {}

func NewProgramHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramHeaderContext {
	var p = new(ProgramHeaderContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_programHeader

	return p
}

func (s *ProgramHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramHeaderContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserPROGRAM, 0)
}

func (s *ProgramHeaderContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *ProgramHeaderContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *ProgramHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterProgramHeader(s)
	}
}

func (s *ProgramHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitProgramHeader(s)
	}
}

func (s *ProgramHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitProgramHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) ProgramHeader() (localctx IProgramHeaderContext) {
	localctx = NewProgramHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, BabyDuckParserRULE_programHeader)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(BabyDuckParserPROGRAM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IProgramBodyContext is an interface to support dynamic dispatch.
type IProgramBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAIN() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	END() antlr.TerminalNode

	// IsProgramBodyContext differentiates from other interfaces.
	IsProgramBodyContext()
}

type ProgramBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramBodyContext() *ProgramBodyContext {
	var p = new(ProgramBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_programBody
	return p
}

func InitEmptyProgramBodyContext(p *ProgramBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_programBody
}

func (*ProgramBodyContext) IsProgramBodyContext() {}

func NewProgramBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramBodyContext {
	var p = new(ProgramBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_programBody

	return p
}

func (s *ProgramBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramBodyContext) MAIN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserMAIN, 0)
}

func (s *ProgramBodyContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *ProgramBodyContext) END() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserEND, 0)
}

func (s *ProgramBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterProgramBody(s)
	}
}

func (s *ProgramBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitProgramBody(s)
	}
}

func (s *ProgramBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitProgramBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) ProgramBody() (localctx IProgramBodyContext) {
	localctx = NewProgramBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, BabyDuckParserRULE_programBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(BabyDuckParserMAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.CodeBlock()
	}
	{
		p.SetState(88)
		p.Match(BabyDuckParserEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, BabyDuckParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParenthesizedExpressionContext is an interface to support dynamic dispatch.
type IParenthesizedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode

	// IsParenthesizedExpressionContext differentiates from other interfaces.
	IsParenthesizedExpressionContext()
}

type ParenthesizedExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenthesizedExpressionContext() *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_parenthesizedExpression
	return p
}

func InitEmptyParenthesizedExpressionContext(p *ParenthesizedExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_parenthesizedExpression
}

func (*ParenthesizedExpressionContext) IsParenthesizedExpressionContext() {}

func NewParenthesizedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_parenthesizedExpression

	return p
}

func (s *ParenthesizedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenthesizedExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *ParenthesizedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesizedExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenthesizedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitParenthesizedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) ParenthesizedExpression() (localctx IParenthesizedExpressionContext) {
	localctx = NewParenthesizedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, BabyDuckParserRULE_parenthesizedExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Expression()
	}
	{
		p.SetState(94)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) INT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserINT, 0)
}

func (s *TypeContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserFLOAT, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, BabyDuckParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserINT || _la == BabyDuckParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdListContext is an interface to support dynamic dispatch.
type IIdListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsIdListContext differentiates from other interfaces.
	IsIdListContext()
}

type IdListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdListContext() *IdListContext {
	var p = new(IdListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_idList
	return p
}

func InitEmptyIdListContext(p *IdListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_idList
}

func (*IdListContext) IsIdListContext() {}

func NewIdListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdListContext {
	var p = new(IdListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_idList

	return p
}

func (s *IdListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdListContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserID)
}

func (s *IdListContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserID, i)
}

func (s *IdListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *IdListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *IdListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterIdList(s)
	}
}

func (s *IdListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitIdList(s)
	}
}

func (s *IdListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitIdList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) IdList() (localctx IIdListContext) {
	localctx = NewIdListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, BabyDuckParserRULE_idList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(BabyDuckParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserCOMMA {
		{
			p.SetState(99)
			p.Match(BabyDuckParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(BabyDuckParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdList() IIdListContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext
	SEMICOLON() antlr.TerminalNode

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) IdList() IIdListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdListContext)
}

func (s *VarDeclContext) COLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOLON, 0)
}

func (s *VarDeclContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *VarDeclContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (s *VarDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitVarDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, BabyDuckParserRULE_varDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.IdList()
	}
	{
		p.SetState(107)
		p.Match(BabyDuckParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(108)
		p.Type_()
	}
	{
		p.SetState(109)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarsSectionContext is an interface to support dynamic dispatch.
type IVarsSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext

	// IsVarsSectionContext differentiates from other interfaces.
	IsVarsSectionContext()
}

type VarsSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarsSectionContext() *VarsSectionContext {
	var p = new(VarsSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_varsSection
	return p
}

func InitEmptyVarsSectionContext(p *VarsSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_varsSection
}

func (*VarsSectionContext) IsVarsSectionContext() {}

func NewVarsSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsSectionContext {
	var p = new(VarsSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_varsSection

	return p
}

func (s *VarsSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsSectionContext) VAR() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserVAR, 0)
}

func (s *VarsSectionContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarsSectionContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *VarsSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterVarsSection(s)
	}
}

func (s *VarsSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitVarsSection(s)
	}
}

func (s *VarsSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitVarsSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) VarsSection() (localctx IVarsSectionContext) {
	localctx = NewVarsSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, BabyDuckParserRULE_varsSection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(BabyDuckParserVAR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == BabyDuckParserID {
		{
			p.SetState(112)
			p.VarDecl()
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST_INT() antlr.TerminalNode
	CONST_FLOAT() antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CONST_INT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCONST_INT, 0)
}

func (s *ConstantContext) CONST_FLOAT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCONST_FLOAT, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, BabyDuckParserRULE_constant)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserCONST_INT || _la == BabyDuckParserCONST_FLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	COLON() antlr.TerminalNode
	Type_() ITypeContext

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ParameterContext) COLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOLON, 0)
}

func (s *ParameterContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, BabyDuckParserRULE_parameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Identifier()
	}
	{
		p.SetState(120)
		p.Match(BabyDuckParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		p.Type_()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	RBRACKET() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	VarsSection() IVarsSectionContext

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionBody
	return p
}

func InitEmptyFunctionBodyContext(p *FunctionBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionBody
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLBRACKET, 0)
}

func (s *FunctionBodyContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *FunctionBodyContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRBRACKET, 0)
}

func (s *FunctionBodyContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *FunctionBodyContext) VarsSection() IVarsSectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarsSectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarsSectionContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (s *FunctionBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFunctionBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, BabyDuckParserRULE_functionBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(BabyDuckParserLBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserVAR {
		{
			p.SetState(124)
			p.VarsSection()
		}

	}
	{
		p.SetState(127)
		p.CodeBlock()
	}
	{
		p.SetState(128)
		p.Match(BabyDuckParserRBRACKET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VOID() antlr.TerminalNode
	Identifier() IIdentifierContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	FunctionBody() IFunctionBodyContext
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) VOID() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserVOID, 0)
}

func (s *FunctionDeclarationContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *FunctionDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *FunctionDeclarationContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *FunctionDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *FunctionDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, BabyDuckParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(BabyDuckParserVOID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Identifier()
	}
	{
		p.SetState(133)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserID {
		{
			p.SetState(134)
			p.Parameter()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == BabyDuckParserCOMMA {
			{
				p.SetState(135)
				p.Match(BabyDuckParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(136)
				p.Parameter()
			}

			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(144)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.FunctionBody()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionsSectionContext is an interface to support dynamic dispatch.
type IFunctionsSectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionDeclaration() IFunctionDeclarationContext

	// IsFunctionsSectionContext differentiates from other interfaces.
	IsFunctionsSectionContext()
}

type FunctionsSectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsSectionContext() *FunctionsSectionContext {
	var p = new(FunctionsSectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionsSection
	return p
}

func InitEmptyFunctionsSectionContext(p *FunctionsSectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionsSection
}

func (*FunctionsSectionContext) IsFunctionsSectionContext() {}

func NewFunctionsSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsSectionContext {
	var p = new(FunctionsSectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_functionsSection

	return p
}

func (s *FunctionsSectionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsSectionContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *FunctionsSectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsSectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsSectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFunctionsSection(s)
	}
}

func (s *FunctionsSectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFunctionsSection(s)
	}
}

func (s *FunctionsSectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFunctionsSection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) FunctionsSection() (localctx IFunctionsSectionContext) {
	localctx = NewFunctionsSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, BabyDuckParserRULE_functionsSection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.FunctionDeclaration()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICodeBlockContext is an interface to support dynamic dispatch.
type ICodeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsCodeBlockContext differentiates from other interfaces.
	IsCodeBlockContext()
}

type CodeBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeBlockContext() *CodeBlockContext {
	var p = new(CodeBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_codeBlock
	return p
}

func InitEmptyCodeBlockContext(p *CodeBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_codeBlock
}

func (*CodeBlockContext) IsCodeBlockContext() {}

func NewCodeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContext {
	var p = new(CodeBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_codeBlock

	return p
}

func (s *CodeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLBRACE, 0)
}

func (s *CodeBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRBRACE, 0)
}

func (s *CodeBlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *CodeBlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CodeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterCodeBlock(s)
	}
}

func (s *CodeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitCodeBlock(s)
	}
}

func (s *CodeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitCodeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) CodeBlock() (localctx ICodeBlockContext) {
	localctx = NewCodeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, BabyDuckParserRULE_codeBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(BabyDuckParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34359738720) != 0 {
		{
			p.SetState(150)
			p.Statement()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(BabyDuckParserRBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Conditional() IConditionalContext
	Loop() ILoopContext
	FunctionCall() IFunctionCallContext
	PrintStatement() IPrintStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) Conditional() IConditionalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *StatementContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) PrintStatement() IPrintStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, BabyDuckParserRULE_statement)
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(158)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(159)
			p.Conditional()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Loop()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(161)
			p.FunctionCall()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(162)
			p.PrintStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserASSIGN, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, BabyDuckParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Identifier()
	}
	{
		p.SetState(166)
		p.Match(BabyDuckParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Expression()
	}
	{
		p.SetState(168)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	ParenthesizedExpression() IParenthesizedExpressionContext
	CodeBlock() ICodeBlockContext

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_ifBlock
	return p
}

func InitEmptyIfBlockContext(p *IfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_ifBlock
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) IF() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserIF, 0)
}

func (s *IfBlockContext) ParenthesizedExpression() IParenthesizedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesizedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesizedExpressionContext)
}

func (s *IfBlockContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (s *IfBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitIfBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, BabyDuckParserRULE_ifBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(BabyDuckParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.ParenthesizedExpression()
	}
	{
		p.SetState(172)
		p.CodeBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	CodeBlock() ICodeBlockContext

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserELSE, 0)
}

func (s *ElseBlockContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (s *ElseBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitElseBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, BabyDuckParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(BabyDuckParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.CodeBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfBlock() IIfBlockContext
	SEMICOLON() antlr.TerminalNode
	ElseBlock() IElseBlockContext

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_conditional
	return p
}

func InitEmptyConditionalContext(p *ConditionalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_conditional
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *ConditionalContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *ConditionalContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (s *ConditionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitConditional(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Conditional() (localctx IConditionalContext) {
	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, BabyDuckParserRULE_conditional)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(177)
		p.IfBlock()
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserELSE {
		{
			p.SetState(178)
			p.ElseBlock()
		}

	}
	{
		p.SetState(181)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHILE() antlr.TerminalNode
	ParenthesizedExpression() IParenthesizedExpressionContext
	DO() antlr.TerminalNode
	CodeBlock() ICodeBlockContext
	SEMICOLON() antlr.TerminalNode

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) WHILE() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserWHILE, 0)
}

func (s *LoopContext) ParenthesizedExpression() IParenthesizedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesizedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesizedExpressionContext)
}

func (s *LoopContext) DO() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserDO, 0)
}

func (s *LoopContext) CodeBlock() ICodeBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICodeBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *LoopContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, BabyDuckParserRULE_loop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(BabyDuckParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.ParenthesizedExpression()
	}
	{
		p.SetState(185)
		p.Match(BabyDuckParserDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.CodeBlock()
	}
	{
		p.SetState(187)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, BabyDuckParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Expression()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserCOMMA {
		{
			p.SetState(190)
			p.Match(BabyDuckParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Expression()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *FunctionCallContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, BabyDuckParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Identifier()
	}
	{
		p.SetState(198)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&40806580224) != 0 {
		{
			p.SetState(199)
			p.ArgumentList()
		}

	}
	{
		p.SetState(202)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintableContext is an interface to support dynamic dispatch.
type IPrintableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	CONST_STRING() antlr.TerminalNode

	// IsPrintableContext differentiates from other interfaces.
	IsPrintableContext()
}

type PrintableContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintableContext() *PrintableContext {
	var p = new(PrintableContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printable
	return p
}

func InitEmptyPrintableContext(p *PrintableContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printable
}

func (*PrintableContext) IsPrintableContext() {}

func NewPrintableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintableContext {
	var p = new(PrintableContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_printable

	return p
}

func (s *PrintableContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintableContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintableContext) CONST_STRING() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCONST_STRING, 0)
}

func (s *PrintableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterPrintable(s)
	}
}

func (s *PrintableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitPrintable(s)
	}
}

func (s *PrintableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitPrintable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Printable() (localctx IPrintableContext) {
	localctx = NewPrintableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, BabyDuckParserRULE_printable)
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserOP_ADD, BabyDuckParserOP_SUBTRACT, BabyDuckParserLPAREN, BabyDuckParserCONST_INT, BabyDuckParserCONST_FLOAT, BabyDuckParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Expression()
		}

	case BabyDuckParserCONST_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(206)
			p.Match(BabyDuckParserCONST_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintStatementContext is an interface to support dynamic dispatch.
type IPrintStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRINT() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllPrintable() []IPrintableContext
	Printable(i int) IPrintableContext
	RPAREN() antlr.TerminalNode
	SEMICOLON() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPrintStatementContext differentiates from other interfaces.
	IsPrintStatementContext()
}

type PrintStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintStatementContext() *PrintStatementContext {
	var p = new(PrintStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printStatement
	return p
}

func InitEmptyPrintStatementContext(p *PrintStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_printStatement
}

func (*PrintStatementContext) IsPrintStatementContext() {}

func NewPrintStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintStatementContext {
	var p = new(PrintStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_printStatement

	return p
}

func (s *PrintStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintStatementContext) PRINT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserPRINT, 0)
}

func (s *PrintStatementContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLPAREN, 0)
}

func (s *PrintStatementContext) AllPrintable() []IPrintableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrintableContext); ok {
			len++
		}
	}

	tst := make([]IPrintableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrintableContext); ok {
			tst[i] = t.(IPrintableContext)
			i++
		}
	}

	return tst
}

func (s *PrintStatementContext) Printable(i int) IPrintableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintableContext)
}

func (s *PrintStatementContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserRPAREN, 0)
}

func (s *PrintStatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserSEMICOLON, 0)
}

func (s *PrintStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(BabyDuckParserCOMMA)
}

func (s *PrintStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(BabyDuckParserCOMMA, i)
}

func (s *PrintStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterPrintStatement(s)
	}
}

func (s *PrintStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitPrintStatement(s)
	}
}

func (s *PrintStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitPrintStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) PrintStatement() (localctx IPrintStatementContext) {
	localctx = NewPrintStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, BabyDuckParserRULE_printStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		p.Match(BabyDuckParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(BabyDuckParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Printable()
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserCOMMA {
		{
			p.SetState(212)
			p.Match(BabyDuckParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Printable()
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(219)
		p.Match(BabyDuckParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(BabyDuckParserSEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ArithmeticExpression() IArithmeticExpressionContext
	RelationalExpression() IRelationalExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ArithmeticExpression() IArithmeticExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticExpressionContext)
}

func (s *ExpressionContext) RelationalExpression() IRelationalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, BabyDuckParserRULE_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.ArithmeticExpression()
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1835008) != 0 {
		{
			p.SetState(223)
			p.RelationalExpression()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalOperatorContext is an interface to support dynamic dispatch.
type IRelationalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GREATER() antlr.TerminalNode
	LESS() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode

	// IsRelationalOperatorContext differentiates from other interfaces.
	IsRelationalOperatorContext()
}

type RelationalOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalOperatorContext() *RelationalOperatorContext {
	var p = new(RelationalOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relationalOperator
	return p
}

func InitEmptyRelationalOperatorContext(p *RelationalOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relationalOperator
}

func (*RelationalOperatorContext) IsRelationalOperatorContext() {}

func NewRelationalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalOperatorContext {
	var p = new(RelationalOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_relationalOperator

	return p
}

func (s *RelationalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalOperatorContext) GREATER() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserGREATER, 0)
}

func (s *RelationalOperatorContext) LESS() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserLESS, 0)
}

func (s *RelationalOperatorContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserNOT_EQUAL, 0)
}

func (s *RelationalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterRelationalOperator(s)
	}
}

func (s *RelationalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitRelationalOperator(s)
	}
}

func (s *RelationalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitRelationalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) RelationalOperator() (localctx IRelationalOperatorContext) {
	localctx = NewRelationalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, BabyDuckParserRULE_relationalOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1835008) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationalOperator() IRelationalOperatorContext
	ArithmeticExpression() IArithmeticExpressionContext

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) RelationalOperator() IRelationalOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalOperatorContext)
}

func (s *RelationalExpressionContext) ArithmeticExpression() IArithmeticExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArithmeticExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArithmeticExpressionContext)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, BabyDuckParserRULE_relationalExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.RelationalOperator()
	}
	{
		p.SetState(229)
		p.ArithmeticExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveOperatorContext is an interface to support dynamic dispatch.
type IAdditiveOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_ADD() antlr.TerminalNode
	OP_SUBTRACT() antlr.TerminalNode

	// IsAdditiveOperatorContext differentiates from other interfaces.
	IsAdditiveOperatorContext()
}

type AdditiveOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveOperatorContext() *AdditiveOperatorContext {
	var p = new(AdditiveOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_additiveOperator
	return p
}

func InitEmptyAdditiveOperatorContext(p *AdditiveOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_additiveOperator
}

func (*AdditiveOperatorContext) IsAdditiveOperatorContext() {}

func NewAdditiveOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveOperatorContext {
	var p = new(AdditiveOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_additiveOperator

	return p
}

func (s *AdditiveOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveOperatorContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserOP_ADD, 0)
}

func (s *AdditiveOperatorContext) OP_SUBTRACT() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserOP_SUBTRACT, 0)
}

func (s *AdditiveOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterAdditiveOperator(s)
	}
}

func (s *AdditiveOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitAdditiveOperator(s)
	}
}

func (s *AdditiveOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitAdditiveOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) AdditiveOperator() (localctx IAdditiveOperatorContext) {
	localctx = NewAdditiveOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, BabyDuckParserRULE_additiveOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserOP_ADD || _la == BabyDuckParserOP_SUBTRACT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArithmeticExpressionContext is an interface to support dynamic dispatch.
type IArithmeticExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTerm() []ITermContext
	Term(i int) ITermContext
	AllAdditiveOperator() []IAdditiveOperatorContext
	AdditiveOperator(i int) IAdditiveOperatorContext

	// IsArithmeticExpressionContext differentiates from other interfaces.
	IsArithmeticExpressionContext()
}

type ArithmeticExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArithmeticExpressionContext() *ArithmeticExpressionContext {
	var p = new(ArithmeticExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_arithmeticExpression
	return p
}

func InitEmptyArithmeticExpressionContext(p *ArithmeticExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_arithmeticExpression
}

func (*ArithmeticExpressionContext) IsArithmeticExpressionContext() {}

func NewArithmeticExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArithmeticExpressionContext {
	var p = new(ArithmeticExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_arithmeticExpression

	return p
}

func (s *ArithmeticExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArithmeticExpressionContext) AllTerm() []ITermContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITermContext); ok {
			len++
		}
	}

	tst := make([]ITermContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITermContext); ok {
			tst[i] = t.(ITermContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticExpressionContext) Term(i int) ITermContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITermContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ArithmeticExpressionContext) AllAdditiveOperator() []IAdditiveOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAdditiveOperatorContext); ok {
			len++
		}
	}

	tst := make([]IAdditiveOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAdditiveOperatorContext); ok {
			tst[i] = t.(IAdditiveOperatorContext)
			i++
		}
	}

	return tst
}

func (s *ArithmeticExpressionContext) AdditiveOperator(i int) IAdditiveOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveOperatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveOperatorContext)
}

func (s *ArithmeticExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArithmeticExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArithmeticExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterArithmeticExpression(s)
	}
}

func (s *ArithmeticExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitArithmeticExpression(s)
	}
}

func (s *ArithmeticExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitArithmeticExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) ArithmeticExpression() (localctx IArithmeticExpressionContext) {
	localctx = NewArithmeticExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, BabyDuckParserRULE_arithmeticExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Term()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserOP_ADD || _la == BabyDuckParserOP_SUBTRACT {
		{
			p.SetState(234)
			p.AdditiveOperator()
		}
		{
			p.SetState(235)
			p.Term()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeOperatorContext is an interface to support dynamic dispatch.
type IMultiplicativeOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OP_MULTIPLY() antlr.TerminalNode
	OP_DIVIDE() antlr.TerminalNode

	// IsMultiplicativeOperatorContext differentiates from other interfaces.
	IsMultiplicativeOperatorContext()
}

type MultiplicativeOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeOperatorContext() *MultiplicativeOperatorContext {
	var p = new(MultiplicativeOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_multiplicativeOperator
	return p
}

func InitEmptyMultiplicativeOperatorContext(p *MultiplicativeOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_multiplicativeOperator
}

func (*MultiplicativeOperatorContext) IsMultiplicativeOperatorContext() {}

func NewMultiplicativeOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeOperatorContext {
	var p = new(MultiplicativeOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_multiplicativeOperator

	return p
}

func (s *MultiplicativeOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeOperatorContext) OP_MULTIPLY() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserOP_MULTIPLY, 0)
}

func (s *MultiplicativeOperatorContext) OP_DIVIDE() antlr.TerminalNode {
	return s.GetToken(BabyDuckParserOP_DIVIDE, 0)
}

func (s *MultiplicativeOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterMultiplicativeOperator(s)
	}
}

func (s *MultiplicativeOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitMultiplicativeOperator(s)
	}
}

func (s *MultiplicativeOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitMultiplicativeOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) MultiplicativeOperator() (localctx IMultiplicativeOperatorContext) {
	localctx = NewMultiplicativeOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, BabyDuckParserRULE_multiplicativeOperator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		_la = p.GetTokenStream().LA(1)

		if !(_la == BabyDuckParserOP_MULTIPLY || _la == BabyDuckParserOP_DIVIDE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFactor() []IFactorContext
	Factor(i int) IFactorContext
	AllMultiplicativeOperator() []IMultiplicativeOperatorContext
	MultiplicativeOperator(i int) IMultiplicativeOperatorContext

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_term
	return p
}

func InitEmptyTermContext(p *TermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_term
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFactorContext); ok {
			len++
		}
	}

	tst := make([]IFactorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFactorContext); ok {
			tst[i] = t.(IFactorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFactorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllMultiplicativeOperator() []IMultiplicativeOperatorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeOperatorContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeOperatorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeOperatorContext); ok {
			tst[i] = t.(IMultiplicativeOperatorContext)
			i++
		}
	}

	return tst
}

func (s *TermContext) MultiplicativeOperator(i int) IMultiplicativeOperatorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeOperatorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeOperatorContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (s *TermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, BabyDuckParserRULE_term)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Factor()
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == BabyDuckParserOP_MULTIPLY || _la == BabyDuckParserOP_DIVIDE {
		{
			p.SetState(245)
			p.MultiplicativeOperator()
		}
		{
			p.SetState(246)
			p.Factor()
		}

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ParenthesizedExpression() IParenthesizedExpressionContext
	ValueWithOptionalSign() IValueWithOptionalSignContext

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factor
	return p
}

func InitEmptyFactorContext(p *FactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_factor
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) ParenthesizedExpression() IParenthesizedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParenthesizedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParenthesizedExpressionContext)
}

func (s *FactorContext) ValueWithOptionalSign() IValueWithOptionalSignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueWithOptionalSignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueWithOptionalSignContext)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (s *FactorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitFactor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, BabyDuckParserRULE_factor)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.ParenthesizedExpression()
		}

	case BabyDuckParserOP_ADD, BabyDuckParserOP_SUBTRACT, BabyDuckParserCONST_INT, BabyDuckParserCONST_FLOAT, BabyDuckParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.ValueWithOptionalSign()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueWithOptionalSignContext is an interface to support dynamic dispatch.
type IValueWithOptionalSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext
	AdditiveOperator() IAdditiveOperatorContext

	// IsValueWithOptionalSignContext differentiates from other interfaces.
	IsValueWithOptionalSignContext()
}

type ValueWithOptionalSignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueWithOptionalSignContext() *ValueWithOptionalSignContext {
	var p = new(ValueWithOptionalSignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_valueWithOptionalSign
	return p
}

func InitEmptyValueWithOptionalSignContext(p *ValueWithOptionalSignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_valueWithOptionalSign
}

func (*ValueWithOptionalSignContext) IsValueWithOptionalSignContext() {}

func NewValueWithOptionalSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueWithOptionalSignContext {
	var p = new(ValueWithOptionalSignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_valueWithOptionalSign

	return p
}

func (s *ValueWithOptionalSignContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueWithOptionalSignContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueWithOptionalSignContext) AdditiveOperator() IAdditiveOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveOperatorContext)
}

func (s *ValueWithOptionalSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueWithOptionalSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueWithOptionalSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterValueWithOptionalSign(s)
	}
}

func (s *ValueWithOptionalSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitValueWithOptionalSign(s)
	}
}

func (s *ValueWithOptionalSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitValueWithOptionalSign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) ValueWithOptionalSign() (localctx IValueWithOptionalSignContext) {
	localctx = NewValueWithOptionalSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, BabyDuckParserRULE_valueWithOptionalSign)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == BabyDuckParserOP_ADD || _la == BabyDuckParserOP_SUBTRACT {
		{
			p.SetState(257)
			p.AdditiveOperator()
		}

	}
	{
		p.SetState(260)
		p.Value()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Constant() IConstantContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = BabyDuckParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = BabyDuckParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(BabyDuckListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case BabyDuckVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *BabyDuckParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, BabyDuckParserRULE_value)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case BabyDuckParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(262)
			p.Identifier()
		}

	case BabyDuckParserCONST_INT, BabyDuckParserCONST_FLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(263)
			p.Constant()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
