// Code generated from calcit2.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

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

type calcit2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Calcit2LexerLexerStaticData struct {
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

func calcit2lexerLexerInit() {
	staticData := &Calcit2LexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'true'", "'false'", "'>'", "'<'", "'+'", "'-'", "'*'", "'/'", "'%'",
		"'!'", "'('", "')'", "':'", "'='", "'**'", "':='", "';'", "'&'", "'|'",
		"'^'", "'&^'", "'<<'", "'>>'", "'&&'", "'||'", "'=='", "'!='", "'<='",
		"'>='", "'hex'", "'bin'", "'oct'", "'float'", "'int'", "'bool'", "'string'",
		"'pcast'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "GREATER", "LESS", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
		"MODULO", "NOT", "LPAREN", "RPAREN", "COLON", "EQUALS", "POW", "COLON_EQUALS",
		"SEMICOLON", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_CLEAR", "LSHIFT",
		"RSHIFT", "LOGICAL_AND", "LOGICAL_OR", "EQUAL", "NOT_EQUAL", "LESS_OR_EQUAL",
		"GREATER_OR_EQUAL", "HEX_TYPE", "BIN_TYPE", "OCT_TYPE", "F_TYPE", "I_TYPE",
		"B_TYPE", "S_TYPE", "P_CAST", "FUNCTION_NAME", "IDENTIFIER", "HEX_LITERAL",
		"BIN_LITERAL", "OCT_LITERAL", "INTEGER_LITERAL", "FLOAT_LITERAL", "CHAR_LITERAL",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "GREATER", "LESS", "PLUS", "MINUS", "MULTIPLY", "DIVIDE",
		"MODULO", "NOT", "LPAREN", "RPAREN", "COLON", "EQUALS", "POW", "COLON_EQUALS",
		"SEMICOLON", "BIT_AND", "BIT_OR", "BIT_XOR", "BIT_CLEAR", "LSHIFT",
		"RSHIFT", "LOGICAL_AND", "LOGICAL_OR", "EQUAL", "NOT_EQUAL", "LESS_OR_EQUAL",
		"GREATER_OR_EQUAL", "HEX_TYPE", "BIN_TYPE", "OCT_TYPE", "F_TYPE", "I_TYPE",
		"B_TYPE", "S_TYPE", "P_CAST", "FUNCTION_NAME", "IDENTIFIER", "HEX_LITERAL",
		"BIN_LITERAL", "OCT_LITERAL", "INTEGER_LITERAL", "FLOAT_LITERAL", "CHAR_LITERAL",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 46, 274, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1,
		14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 5, 37, 212, 8,
		37, 10, 37, 12, 37, 215, 9, 37, 1, 38, 1, 38, 5, 38, 219, 8, 38, 10, 38,
		12, 38, 222, 9, 38, 1, 39, 1, 39, 1, 39, 1, 39, 4, 39, 228, 8, 39, 11,
		39, 12, 39, 229, 1, 40, 1, 40, 1, 40, 1, 40, 4, 40, 236, 8, 40, 11, 40,
		12, 40, 237, 1, 41, 1, 41, 1, 41, 1, 41, 4, 41, 244, 8, 41, 11, 41, 12,
		41, 245, 1, 42, 4, 42, 249, 8, 42, 11, 42, 12, 42, 250, 1, 43, 4, 43, 254,
		8, 43, 11, 43, 12, 43, 255, 1, 43, 1, 43, 4, 43, 260, 8, 43, 11, 43, 12,
		43, 261, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 4, 45, 269, 8, 45, 11, 45,
		12, 45, 270, 1, 45, 1, 45, 0, 0, 46, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33,
		67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42,
		85, 43, 87, 44, 89, 45, 91, 46, 1, 0, 7, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 3, 0, 65, 90, 95, 95, 97, 122, 3, 0, 48, 57, 65, 70, 97, 102,
		1, 0, 48, 49, 1, 0, 48, 55, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32,
		282, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0,
		0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1,
		0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23,
		1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0,
		0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0,
		0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1,
		0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69,
		1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0,
		77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0,
		0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0,
		0, 1, 93, 1, 0, 0, 0, 3, 98, 1, 0, 0, 0, 5, 104, 1, 0, 0, 0, 7, 106, 1,
		0, 0, 0, 9, 108, 1, 0, 0, 0, 11, 110, 1, 0, 0, 0, 13, 112, 1, 0, 0, 0,
		15, 114, 1, 0, 0, 0, 17, 116, 1, 0, 0, 0, 19, 118, 1, 0, 0, 0, 21, 120,
		1, 0, 0, 0, 23, 122, 1, 0, 0, 0, 25, 124, 1, 0, 0, 0, 27, 126, 1, 0, 0,
		0, 29, 128, 1, 0, 0, 0, 31, 131, 1, 0, 0, 0, 33, 134, 1, 0, 0, 0, 35, 136,
		1, 0, 0, 0, 37, 138, 1, 0, 0, 0, 39, 140, 1, 0, 0, 0, 41, 142, 1, 0, 0,
		0, 43, 145, 1, 0, 0, 0, 45, 148, 1, 0, 0, 0, 47, 151, 1, 0, 0, 0, 49, 154,
		1, 0, 0, 0, 51, 157, 1, 0, 0, 0, 53, 160, 1, 0, 0, 0, 55, 163, 1, 0, 0,
		0, 57, 166, 1, 0, 0, 0, 59, 169, 1, 0, 0, 0, 61, 173, 1, 0, 0, 0, 63, 177,
		1, 0, 0, 0, 65, 181, 1, 0, 0, 0, 67, 187, 1, 0, 0, 0, 69, 191, 1, 0, 0,
		0, 71, 196, 1, 0, 0, 0, 73, 203, 1, 0, 0, 0, 75, 209, 1, 0, 0, 0, 77, 216,
		1, 0, 0, 0, 79, 223, 1, 0, 0, 0, 81, 231, 1, 0, 0, 0, 83, 239, 1, 0, 0,
		0, 85, 248, 1, 0, 0, 0, 87, 253, 1, 0, 0, 0, 89, 263, 1, 0, 0, 0, 91, 268,
		1, 0, 0, 0, 93, 94, 5, 116, 0, 0, 94, 95, 5, 114, 0, 0, 95, 96, 5, 117,
		0, 0, 96, 97, 5, 101, 0, 0, 97, 2, 1, 0, 0, 0, 98, 99, 5, 102, 0, 0, 99,
		100, 5, 97, 0, 0, 100, 101, 5, 108, 0, 0, 101, 102, 5, 115, 0, 0, 102,
		103, 5, 101, 0, 0, 103, 4, 1, 0, 0, 0, 104, 105, 5, 62, 0, 0, 105, 6, 1,
		0, 0, 0, 106, 107, 5, 60, 0, 0, 107, 8, 1, 0, 0, 0, 108, 109, 5, 43, 0,
		0, 109, 10, 1, 0, 0, 0, 110, 111, 5, 45, 0, 0, 111, 12, 1, 0, 0, 0, 112,
		113, 5, 42, 0, 0, 113, 14, 1, 0, 0, 0, 114, 115, 5, 47, 0, 0, 115, 16,
		1, 0, 0, 0, 116, 117, 5, 37, 0, 0, 117, 18, 1, 0, 0, 0, 118, 119, 5, 33,
		0, 0, 119, 20, 1, 0, 0, 0, 120, 121, 5, 40, 0, 0, 121, 22, 1, 0, 0, 0,
		122, 123, 5, 41, 0, 0, 123, 24, 1, 0, 0, 0, 124, 125, 5, 58, 0, 0, 125,
		26, 1, 0, 0, 0, 126, 127, 5, 61, 0, 0, 127, 28, 1, 0, 0, 0, 128, 129, 5,
		42, 0, 0, 129, 130, 5, 42, 0, 0, 130, 30, 1, 0, 0, 0, 131, 132, 5, 58,
		0, 0, 132, 133, 5, 61, 0, 0, 133, 32, 1, 0, 0, 0, 134, 135, 5, 59, 0, 0,
		135, 34, 1, 0, 0, 0, 136, 137, 5, 38, 0, 0, 137, 36, 1, 0, 0, 0, 138, 139,
		5, 124, 0, 0, 139, 38, 1, 0, 0, 0, 140, 141, 5, 94, 0, 0, 141, 40, 1, 0,
		0, 0, 142, 143, 5, 38, 0, 0, 143, 144, 5, 94, 0, 0, 144, 42, 1, 0, 0, 0,
		145, 146, 5, 60, 0, 0, 146, 147, 5, 60, 0, 0, 147, 44, 1, 0, 0, 0, 148,
		149, 5, 62, 0, 0, 149, 150, 5, 62, 0, 0, 150, 46, 1, 0, 0, 0, 151, 152,
		5, 38, 0, 0, 152, 153, 5, 38, 0, 0, 153, 48, 1, 0, 0, 0, 154, 155, 5, 124,
		0, 0, 155, 156, 5, 124, 0, 0, 156, 50, 1, 0, 0, 0, 157, 158, 5, 61, 0,
		0, 158, 159, 5, 61, 0, 0, 159, 52, 1, 0, 0, 0, 160, 161, 5, 33, 0, 0, 161,
		162, 5, 61, 0, 0, 162, 54, 1, 0, 0, 0, 163, 164, 5, 60, 0, 0, 164, 165,
		5, 61, 0, 0, 165, 56, 1, 0, 0, 0, 166, 167, 5, 62, 0, 0, 167, 168, 5, 61,
		0, 0, 168, 58, 1, 0, 0, 0, 169, 170, 5, 104, 0, 0, 170, 171, 5, 101, 0,
		0, 171, 172, 5, 120, 0, 0, 172, 60, 1, 0, 0, 0, 173, 174, 5, 98, 0, 0,
		174, 175, 5, 105, 0, 0, 175, 176, 5, 110, 0, 0, 176, 62, 1, 0, 0, 0, 177,
		178, 5, 111, 0, 0, 178, 179, 5, 99, 0, 0, 179, 180, 5, 116, 0, 0, 180,
		64, 1, 0, 0, 0, 181, 182, 5, 102, 0, 0, 182, 183, 5, 108, 0, 0, 183, 184,
		5, 111, 0, 0, 184, 185, 5, 97, 0, 0, 185, 186, 5, 116, 0, 0, 186, 66, 1,
		0, 0, 0, 187, 188, 5, 105, 0, 0, 188, 189, 5, 110, 0, 0, 189, 190, 5, 116,
		0, 0, 190, 68, 1, 0, 0, 0, 191, 192, 5, 98, 0, 0, 192, 193, 5, 111, 0,
		0, 193, 194, 5, 111, 0, 0, 194, 195, 5, 108, 0, 0, 195, 70, 1, 0, 0, 0,
		196, 197, 5, 115, 0, 0, 197, 198, 5, 116, 0, 0, 198, 199, 5, 114, 0, 0,
		199, 200, 5, 105, 0, 0, 200, 201, 5, 110, 0, 0, 201, 202, 5, 103, 0, 0,
		202, 72, 1, 0, 0, 0, 203, 204, 5, 112, 0, 0, 204, 205, 5, 99, 0, 0, 205,
		206, 5, 97, 0, 0, 206, 207, 5, 115, 0, 0, 207, 208, 5, 116, 0, 0, 208,
		74, 1, 0, 0, 0, 209, 213, 5, 102, 0, 0, 210, 212, 7, 0, 0, 0, 211, 210,
		1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0,
		0, 0, 214, 76, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 220, 7, 1, 0, 0,
		217, 219, 7, 0, 0, 0, 218, 217, 1, 0, 0, 0, 219, 222, 1, 0, 0, 0, 220,
		218, 1, 0, 0, 0, 220, 221, 1, 0, 0, 0, 221, 78, 1, 0, 0, 0, 222, 220, 1,
		0, 0, 0, 223, 224, 5, 48, 0, 0, 224, 225, 5, 120, 0, 0, 225, 227, 1, 0,
		0, 0, 226, 228, 7, 2, 0, 0, 227, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 227, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 80, 1, 0, 0, 0, 231, 232,
		5, 48, 0, 0, 232, 233, 5, 98, 0, 0, 233, 235, 1, 0, 0, 0, 234, 236, 7,
		3, 0, 0, 235, 234, 1, 0, 0, 0, 236, 237, 1, 0, 0, 0, 237, 235, 1, 0, 0,
		0, 237, 238, 1, 0, 0, 0, 238, 82, 1, 0, 0, 0, 239, 240, 5, 48, 0, 0, 240,
		241, 5, 111, 0, 0, 241, 243, 1, 0, 0, 0, 242, 244, 7, 4, 0, 0, 243, 242,
		1, 0, 0, 0, 244, 245, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0,
		0, 0, 246, 84, 1, 0, 0, 0, 247, 249, 7, 5, 0, 0, 248, 247, 1, 0, 0, 0,
		249, 250, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251,
		86, 1, 0, 0, 0, 252, 254, 7, 5, 0, 0, 253, 252, 1, 0, 0, 0, 254, 255, 1,
		0, 0, 0, 255, 253, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0,
		0, 257, 259, 5, 46, 0, 0, 258, 260, 7, 5, 0, 0, 259, 258, 1, 0, 0, 0, 260,
		261, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 88, 1,
		0, 0, 0, 263, 264, 5, 39, 0, 0, 264, 265, 9, 0, 0, 0, 265, 266, 5, 39,
		0, 0, 266, 90, 1, 0, 0, 0, 267, 269, 7, 6, 0, 0, 268, 267, 1, 0, 0, 0,
		269, 270, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271,
		272, 1, 0, 0, 0, 272, 273, 6, 45, 0, 0, 273, 92, 1, 0, 0, 0, 10, 0, 213,
		220, 229, 237, 245, 250, 255, 261, 270, 1, 6, 0, 0,
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

// calcit2LexerInit initializes any static state used to implement calcit2Lexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// Newcalcit2Lexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Calcit2LexerInit() {
	staticData := &Calcit2LexerLexerStaticData
	staticData.once.Do(calcit2lexerLexerInit)
}

// Newcalcit2Lexer produces a new lexer instance for the optional input antlr.CharStream.
func Newcalcit2Lexer(input antlr.CharStream) *calcit2Lexer {
	Calcit2LexerInit()
	l := new(calcit2Lexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Calcit2LexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "calcit2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// calcit2Lexer tokens.
const (
	calcit2LexerT__0             = 1
	calcit2LexerT__1             = 2
	calcit2LexerGREATER          = 3
	calcit2LexerLESS             = 4
	calcit2LexerPLUS             = 5
	calcit2LexerMINUS            = 6
	calcit2LexerMULTIPLY         = 7
	calcit2LexerDIVIDE           = 8
	calcit2LexerMODULO           = 9
	calcit2LexerNOT              = 10
	calcit2LexerLPAREN           = 11
	calcit2LexerRPAREN           = 12
	calcit2LexerCOLON            = 13
	calcit2LexerEQUALS           = 14
	calcit2LexerPOW              = 15
	calcit2LexerCOLON_EQUALS     = 16
	calcit2LexerSEMICOLON        = 17
	calcit2LexerBIT_AND          = 18
	calcit2LexerBIT_OR           = 19
	calcit2LexerBIT_XOR          = 20
	calcit2LexerBIT_CLEAR        = 21
	calcit2LexerLSHIFT           = 22
	calcit2LexerRSHIFT           = 23
	calcit2LexerLOGICAL_AND      = 24
	calcit2LexerLOGICAL_OR       = 25
	calcit2LexerEQUAL            = 26
	calcit2LexerNOT_EQUAL        = 27
	calcit2LexerLESS_OR_EQUAL    = 28
	calcit2LexerGREATER_OR_EQUAL = 29
	calcit2LexerHEX_TYPE         = 30
	calcit2LexerBIN_TYPE         = 31
	calcit2LexerOCT_TYPE         = 32
	calcit2LexerF_TYPE           = 33
	calcit2LexerI_TYPE           = 34
	calcit2LexerB_TYPE           = 35
	calcit2LexerS_TYPE           = 36
	calcit2LexerP_CAST           = 37
	calcit2LexerFUNCTION_NAME    = 38
	calcit2LexerIDENTIFIER       = 39
	calcit2LexerHEX_LITERAL      = 40
	calcit2LexerBIN_LITERAL      = 41
	calcit2LexerOCT_LITERAL      = 42
	calcit2LexerINTEGER_LITERAL  = 43
	calcit2LexerFLOAT_LITERAL    = 44
	calcit2LexerCHAR_LITERAL     = 45
	calcit2LexerWS               = 46
)
