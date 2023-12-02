// Code generated from calcit2.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // calcit2

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

type calcit2Parser struct {
	*antlr.BaseParser
}

var Calcit2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcit2ParserInit() {
	staticData := &Calcit2ParserStaticData
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
		"program", "statement", "function_assignment", "expression", "number",
		"any", "integer_expression", "integer_cast_expression", "integer_pcast_expression",
		"float_expression", "float_cast_expression", "float_pcast_expression",
		"bool_expression", "bool_cast_expression", "bool_number_operation",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 170, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 5, 0,
		32, 8, 0, 10, 0, 12, 0, 35, 9, 0, 1, 1, 1, 1, 3, 1, 39, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 51, 8, 3, 1, 4,
		1, 4, 3, 4, 55, 8, 4, 1, 5, 1, 5, 3, 5, 59, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 73, 8, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6,
		87, 8, 6, 10, 6, 12, 6, 90, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 3, 9, 112, 8, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 5, 9, 123, 8, 9, 10, 9, 12, 9, 126, 9, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 148, 8, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 156, 8, 12, 10, 12, 12, 12, 159,
		9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 0, 3, 12, 18, 24, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 0, 11, 1, 0, 39, 43, 1, 0, 7, 9, 1, 0, 5, 6, 1, 0, 18, 23, 1, 0,
		30, 34, 2, 0, 39, 39, 44, 44, 1, 0, 7, 8, 1, 0, 1, 2, 1, 0, 26, 27, 1,
		0, 24, 25, 2, 0, 3, 4, 26, 29, 182, 0, 33, 1, 0, 0, 0, 2, 38, 1, 0, 0,
		0, 4, 40, 1, 0, 0, 0, 6, 50, 1, 0, 0, 0, 8, 54, 1, 0, 0, 0, 10, 58, 1,
		0, 0, 0, 12, 72, 1, 0, 0, 0, 14, 91, 1, 0, 0, 0, 16, 96, 1, 0, 0, 0, 18,
		111, 1, 0, 0, 0, 20, 127, 1, 0, 0, 0, 22, 132, 1, 0, 0, 0, 24, 147, 1,
		0, 0, 0, 26, 160, 1, 0, 0, 0, 28, 165, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0,
		31, 30, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1,
		0, 0, 0, 34, 1, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 36, 39, 3, 4, 2, 0, 37,
		39, 3, 6, 3, 0, 38, 36, 1, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 3, 1, 0, 0,
		0, 40, 41, 5, 38, 0, 0, 41, 42, 5, 11, 0, 0, 42, 43, 5, 39, 0, 0, 43, 44,
		5, 12, 0, 0, 44, 45, 5, 14, 0, 0, 45, 46, 3, 6, 3, 0, 46, 5, 1, 0, 0, 0,
		47, 51, 3, 12, 6, 0, 48, 51, 3, 18, 9, 0, 49, 51, 3, 24, 12, 0, 50, 47,
		1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 7, 1, 0, 0, 0,
		52, 55, 3, 12, 6, 0, 53, 55, 3, 18, 9, 0, 54, 52, 1, 0, 0, 0, 54, 53, 1,
		0, 0, 0, 55, 9, 1, 0, 0, 0, 56, 59, 3, 8, 4, 0, 57, 59, 3, 24, 12, 0, 58,
		56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59, 11, 1, 0, 0, 0, 60, 61, 6, 6, -1,
		0, 61, 62, 5, 11, 0, 0, 62, 63, 3, 12, 6, 0, 63, 64, 5, 12, 0, 0, 64, 73,
		1, 0, 0, 0, 65, 73, 3, 14, 7, 0, 66, 73, 3, 16, 8, 0, 67, 68, 5, 6, 0,
		0, 68, 73, 3, 12, 6, 7, 69, 70, 5, 10, 0, 0, 70, 73, 3, 12, 6, 6, 71, 73,
		7, 0, 0, 0, 72, 60, 1, 0, 0, 0, 72, 65, 1, 0, 0, 0, 72, 66, 1, 0, 0, 0,
		72, 67, 1, 0, 0, 0, 72, 69, 1, 0, 0, 0, 72, 71, 1, 0, 0, 0, 73, 88, 1,
		0, 0, 0, 74, 75, 10, 5, 0, 0, 75, 76, 5, 15, 0, 0, 76, 87, 3, 12, 6, 6,
		77, 78, 10, 4, 0, 0, 78, 79, 7, 1, 0, 0, 79, 87, 3, 12, 6, 5, 80, 81, 10,
		3, 0, 0, 81, 82, 7, 2, 0, 0, 82, 87, 3, 12, 6, 4, 83, 84, 10, 2, 0, 0,
		84, 85, 7, 3, 0, 0, 85, 87, 3, 12, 6, 3, 86, 74, 1, 0, 0, 0, 86, 77, 1,
		0, 0, 0, 86, 80, 1, 0, 0, 0, 86, 83, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88,
		86, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 13, 1, 0, 0, 0, 90, 88, 1, 0, 0,
		0, 91, 92, 7, 4, 0, 0, 92, 93, 5, 11, 0, 0, 93, 94, 3, 8, 4, 0, 94, 95,
		5, 12, 0, 0, 95, 15, 1, 0, 0, 0, 96, 97, 5, 37, 0, 0, 97, 98, 5, 11, 0,
		0, 98, 99, 3, 18, 9, 0, 99, 100, 5, 12, 0, 0, 100, 17, 1, 0, 0, 0, 101,
		102, 6, 9, -1, 0, 102, 103, 5, 11, 0, 0, 103, 104, 3, 18, 9, 0, 104, 105,
		5, 12, 0, 0, 105, 112, 1, 0, 0, 0, 106, 112, 3, 20, 10, 0, 107, 112, 3,
		22, 11, 0, 108, 109, 5, 6, 0, 0, 109, 112, 3, 18, 9, 5, 110, 112, 7, 5,
		0, 0, 111, 101, 1, 0, 0, 0, 111, 106, 1, 0, 0, 0, 111, 107, 1, 0, 0, 0,
		111, 108, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 124, 1, 0, 0, 0, 113,
		114, 10, 4, 0, 0, 114, 115, 5, 15, 0, 0, 115, 123, 3, 18, 9, 5, 116, 117,
		10, 3, 0, 0, 117, 118, 7, 6, 0, 0, 118, 123, 3, 18, 9, 4, 119, 120, 10,
		2, 0, 0, 120, 121, 7, 2, 0, 0, 121, 123, 3, 18, 9, 3, 122, 113, 1, 0, 0,
		0, 122, 116, 1, 0, 0, 0, 122, 119, 1, 0, 0, 0, 123, 126, 1, 0, 0, 0, 124,
		122, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 19, 1, 0, 0, 0, 126, 124, 1,
		0, 0, 0, 127, 128, 5, 33, 0, 0, 128, 129, 5, 11, 0, 0, 129, 130, 3, 8,
		4, 0, 130, 131, 5, 12, 0, 0, 131, 21, 1, 0, 0, 0, 132, 133, 5, 37, 0, 0,
		133, 134, 5, 11, 0, 0, 134, 135, 3, 12, 6, 0, 135, 136, 5, 12, 0, 0, 136,
		23, 1, 0, 0, 0, 137, 138, 6, 12, -1, 0, 138, 139, 5, 11, 0, 0, 139, 140,
		3, 24, 12, 0, 140, 141, 5, 12, 0, 0, 141, 148, 1, 0, 0, 0, 142, 148, 3,
		26, 13, 0, 143, 148, 3, 28, 14, 0, 144, 145, 5, 10, 0, 0, 145, 148, 3,
		24, 12, 2, 146, 148, 7, 7, 0, 0, 147, 137, 1, 0, 0, 0, 147, 142, 1, 0,
		0, 0, 147, 143, 1, 0, 0, 0, 147, 144, 1, 0, 0, 0, 147, 146, 1, 0, 0, 0,
		148, 157, 1, 0, 0, 0, 149, 150, 10, 4, 0, 0, 150, 151, 7, 8, 0, 0, 151,
		156, 3, 24, 12, 5, 152, 153, 10, 3, 0, 0, 153, 154, 7, 9, 0, 0, 154, 156,
		3, 24, 12, 4, 155, 149, 1, 0, 0, 0, 155, 152, 1, 0, 0, 0, 156, 159, 1,
		0, 0, 0, 157, 155, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 25, 1, 0, 0,
		0, 159, 157, 1, 0, 0, 0, 160, 161, 5, 35, 0, 0, 161, 162, 5, 11, 0, 0,
		162, 163, 3, 10, 5, 0, 163, 164, 5, 12, 0, 0, 164, 27, 1, 0, 0, 0, 165,
		166, 3, 8, 4, 0, 166, 167, 7, 10, 0, 0, 167, 168, 3, 8, 4, 0, 168, 29,
		1, 0, 0, 0, 14, 33, 38, 50, 54, 58, 72, 86, 88, 111, 122, 124, 147, 155,
		157,
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

// calcit2ParserInit initializes any static state used to implement calcit2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// Newcalcit2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Calcit2ParserInit() {
	staticData := &Calcit2ParserStaticData
	staticData.once.Do(calcit2ParserInit)
}

// Newcalcit2Parser produces a new parser instance for the optional input antlr.TokenStream.
func Newcalcit2Parser(input antlr.TokenStream) *calcit2Parser {
	Calcit2ParserInit()
	this := new(calcit2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Calcit2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "calcit2.g4"

	return this
}

// calcit2Parser tokens.
const (
	calcit2ParserEOF              = antlr.TokenEOF
	calcit2ParserT__0             = 1
	calcit2ParserT__1             = 2
	calcit2ParserGREATER          = 3
	calcit2ParserLESS             = 4
	calcit2ParserPLUS             = 5
	calcit2ParserMINUS            = 6
	calcit2ParserMULTIPLY         = 7
	calcit2ParserDIVIDE           = 8
	calcit2ParserMODULO           = 9
	calcit2ParserNOT              = 10
	calcit2ParserLPAREN           = 11
	calcit2ParserRPAREN           = 12
	calcit2ParserCOLON            = 13
	calcit2ParserEQUALS           = 14
	calcit2ParserPOW              = 15
	calcit2ParserCOLON_EQUALS     = 16
	calcit2ParserSEMICOLON        = 17
	calcit2ParserBIT_AND          = 18
	calcit2ParserBIT_OR           = 19
	calcit2ParserBIT_XOR          = 20
	calcit2ParserBIT_CLEAR        = 21
	calcit2ParserLSHIFT           = 22
	calcit2ParserRSHIFT           = 23
	calcit2ParserLOGICAL_AND      = 24
	calcit2ParserLOGICAL_OR       = 25
	calcit2ParserEQUAL            = 26
	calcit2ParserNOT_EQUAL        = 27
	calcit2ParserLESS_OR_EQUAL    = 28
	calcit2ParserGREATER_OR_EQUAL = 29
	calcit2ParserHEX_TYPE         = 30
	calcit2ParserBIN_TYPE         = 31
	calcit2ParserOCT_TYPE         = 32
	calcit2ParserF_TYPE           = 33
	calcit2ParserI_TYPE           = 34
	calcit2ParserB_TYPE           = 35
	calcit2ParserS_TYPE           = 36
	calcit2ParserP_CAST           = 37
	calcit2ParserFUNCTION_NAME    = 38
	calcit2ParserIDENTIFIER       = 39
	calcit2ParserHEX_LITERAL      = 40
	calcit2ParserBIN_LITERAL      = 41
	calcit2ParserOCT_LITERAL      = 42
	calcit2ParserINTEGER_LITERAL  = 43
	calcit2ParserFLOAT_LITERAL    = 44
	calcit2ParserCHAR_LITERAL     = 45
	calcit2ParserWS               = 46
)

// calcit2Parser rules.
const (
	calcit2ParserRULE_program                  = 0
	calcit2ParserRULE_statement                = 1
	calcit2ParserRULE_function_assignment      = 2
	calcit2ParserRULE_expression               = 3
	calcit2ParserRULE_number                   = 4
	calcit2ParserRULE_any                      = 5
	calcit2ParserRULE_integer_expression       = 6
	calcit2ParserRULE_integer_cast_expression  = 7
	calcit2ParserRULE_integer_pcast_expression = 8
	calcit2ParserRULE_float_expression         = 9
	calcit2ParserRULE_float_cast_expression    = 10
	calcit2ParserRULE_float_pcast_expression   = 11
	calcit2ParserRULE_bool_expression          = 12
	calcit2ParserRULE_bool_cast_expression     = 13
	calcit2ParserRULE_bool_number_operation    = 14
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

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
	p.RuleIndex = calcit2ParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
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

func (s *ProgramContext) Statement(i int) IStatementContext {
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

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, calcit2ParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35114578873414) != 0 {
		{
			p.SetState(30)
			p.Statement()
		}

		p.SetState(35)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Function_assignment() IFunction_assignmentContext
	Expression() IExpressionContext

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
	p.RuleIndex = calcit2ParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Function_assignment() IFunction_assignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_assignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_assignmentContext)
}

func (s *StatementContext) Expression() IExpressionContext {
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

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, calcit2ParserRULE_statement)
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case calcit2ParserFUNCTION_NAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(36)
			p.Function_assignment()
		}

	case calcit2ParserT__0, calcit2ParserT__1, calcit2ParserMINUS, calcit2ParserNOT, calcit2ParserLPAREN, calcit2ParserHEX_TYPE, calcit2ParserBIN_TYPE, calcit2ParserOCT_TYPE, calcit2ParserF_TYPE, calcit2ParserI_TYPE, calcit2ParserB_TYPE, calcit2ParserP_CAST, calcit2ParserIDENTIFIER, calcit2ParserHEX_LITERAL, calcit2ParserBIN_LITERAL, calcit2ParserOCT_LITERAL, calcit2ParserINTEGER_LITERAL, calcit2ParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Expression()
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

// IFunction_assignmentContext is an interface to support dynamic dispatch.
type IFunction_assignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION_NAME() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	EQUALS() antlr.TerminalNode
	Expression() IExpressionContext

	// IsFunction_assignmentContext differentiates from other interfaces.
	IsFunction_assignmentContext()
}

type Function_assignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_assignmentContext() *Function_assignmentContext {
	var p = new(Function_assignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_function_assignment
	return p
}

func InitEmptyFunction_assignmentContext(p *Function_assignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_function_assignment
}

func (*Function_assignmentContext) IsFunction_assignmentContext() {}

func NewFunction_assignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_assignmentContext {
	var p = new(Function_assignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_function_assignment

	return p
}

func (s *Function_assignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_assignmentContext) FUNCTION_NAME() antlr.TerminalNode {
	return s.GetToken(calcit2ParserFUNCTION_NAME, 0)
}

func (s *Function_assignmentContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Function_assignmentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(calcit2ParserIDENTIFIER, 0)
}

func (s *Function_assignmentContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Function_assignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(calcit2ParserEQUALS, 0)
}

func (s *Function_assignmentContext) Expression() IExpressionContext {
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

func (s *Function_assignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_assignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_assignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterFunction_assignment(s)
	}
}

func (s *Function_assignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitFunction_assignment(s)
	}
}

func (s *Function_assignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitFunction_assignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Function_assignment() (localctx IFunction_assignmentContext) {
	localctx = NewFunction_assignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, calcit2ParserRULE_function_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(calcit2ParserFUNCTION_NAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(calcit2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(42)
		p.Match(calcit2ParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(43)
		p.Match(calcit2ParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(calcit2ParserEQUALS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Expression()
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
	Integer_expression() IInteger_expressionContext
	Float_expression() IFloat_expressionContext
	Bool_expression() IBool_expressionContext

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
	p.RuleIndex = calcit2ParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Integer_expression() IInteger_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInteger_expressionContext)
}

func (s *ExpressionContext) Float_expression() IFloat_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_expressionContext)
}

func (s *ExpressionContext) Bool_expression() IBool_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, calcit2ParserRULE_expression)
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(47)
			p.integer_expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(48)
			p.float_expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)
			p.bool_expression(0)
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

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Integer_expression() IInteger_expressionContext
	Float_expression() IFloat_expressionContext

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_number
	return p
}

func InitEmptyNumberContext(p *NumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_number
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) Integer_expression() IInteger_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInteger_expressionContext)
}

func (s *NumberContext) Float_expression() IFloat_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_expressionContext)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, calcit2ParserRULE_number)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.integer_expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.float_expression(0)
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

// IAnyContext is an interface to support dynamic dispatch.
type IAnyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() INumberContext
	Bool_expression() IBool_expressionContext

	// IsAnyContext differentiates from other interfaces.
	IsAnyContext()
}

type AnyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyContext() *AnyContext {
	var p = new(AnyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_any
	return p
}

func InitEmptyAnyContext(p *AnyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_any
}

func (*AnyContext) IsAnyContext() {}

func NewAnyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyContext {
	var p = new(AnyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_any

	return p
}

func (s *AnyContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *AnyContext) Bool_expression() IBool_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_expressionContext)
}

func (s *AnyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AnyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterAny(s)
	}
}

func (s *AnyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitAny(s)
	}
}

func (s *AnyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitAny(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Any() (localctx IAnyContext) {
	localctx = NewAnyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, calcit2ParserRULE_any)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Number()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.bool_expression(0)
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

// IInteger_expressionContext is an interface to support dynamic dispatch.
type IInteger_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IInteger_expressionContext

	// GetRight returns the right rule contexts.
	GetRight() IInteger_expressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IInteger_expressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IInteger_expressionContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllInteger_expression() []IInteger_expressionContext
	Integer_expression(i int) IInteger_expressionContext
	Integer_cast_expression() IInteger_cast_expressionContext
	Integer_pcast_expression() IInteger_pcast_expressionContext
	MINUS() antlr.TerminalNode
	NOT() antlr.TerminalNode
	INTEGER_LITERAL() antlr.TerminalNode
	HEX_LITERAL() antlr.TerminalNode
	BIN_LITERAL() antlr.TerminalNode
	OCT_LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	POW() antlr.TerminalNode
	MULTIPLY() antlr.TerminalNode
	DIVIDE() antlr.TerminalNode
	MODULO() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	BIT_AND() antlr.TerminalNode
	BIT_OR() antlr.TerminalNode
	BIT_XOR() antlr.TerminalNode
	BIT_CLEAR() antlr.TerminalNode
	LSHIFT() antlr.TerminalNode
	RSHIFT() antlr.TerminalNode

	// IsInteger_expressionContext differentiates from other interfaces.
	IsInteger_expressionContext()
}

type Integer_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IInteger_expressionContext
	right  IInteger_expressionContext
	op     antlr.Token
}

func NewEmptyInteger_expressionContext() *Integer_expressionContext {
	var p = new(Integer_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_integer_expression
	return p
}

func InitEmptyInteger_expressionContext(p *Integer_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_integer_expression
}

func (*Integer_expressionContext) IsInteger_expressionContext() {}

func NewInteger_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_expressionContext {
	var p = new(Integer_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_integer_expression

	return p
}

func (s *Integer_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Integer_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Integer_expressionContext) GetLeft() IInteger_expressionContext { return s.left }

func (s *Integer_expressionContext) GetRight() IInteger_expressionContext { return s.right }

func (s *Integer_expressionContext) SetLeft(v IInteger_expressionContext) { s.left = v }

func (s *Integer_expressionContext) SetRight(v IInteger_expressionContext) { s.right = v }

func (s *Integer_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Integer_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Integer_expressionContext) AllInteger_expression() []IInteger_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInteger_expressionContext); ok {
			len++
		}
	}

	tst := make([]IInteger_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInteger_expressionContext); ok {
			tst[i] = t.(IInteger_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Integer_expressionContext) Integer_expression(i int) IInteger_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_expressionContext); ok {
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

	return t.(IInteger_expressionContext)
}

func (s *Integer_expressionContext) Integer_cast_expression() IInteger_cast_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_cast_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInteger_cast_expressionContext)
}

func (s *Integer_expressionContext) Integer_pcast_expression() IInteger_pcast_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_pcast_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInteger_pcast_expressionContext)
}

func (s *Integer_expressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(calcit2ParserMINUS, 0)
}

func (s *Integer_expressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(calcit2ParserNOT, 0)
}

func (s *Integer_expressionContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserINTEGER_LITERAL, 0)
}

func (s *Integer_expressionContext) HEX_LITERAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserHEX_LITERAL, 0)
}

func (s *Integer_expressionContext) BIN_LITERAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserBIN_LITERAL, 0)
}

func (s *Integer_expressionContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserOCT_LITERAL, 0)
}

func (s *Integer_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(calcit2ParserIDENTIFIER, 0)
}

func (s *Integer_expressionContext) POW() antlr.TerminalNode {
	return s.GetToken(calcit2ParserPOW, 0)
}

func (s *Integer_expressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(calcit2ParserMULTIPLY, 0)
}

func (s *Integer_expressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserDIVIDE, 0)
}

func (s *Integer_expressionContext) MODULO() antlr.TerminalNode {
	return s.GetToken(calcit2ParserMODULO, 0)
}

func (s *Integer_expressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(calcit2ParserPLUS, 0)
}

func (s *Integer_expressionContext) BIT_AND() antlr.TerminalNode {
	return s.GetToken(calcit2ParserBIT_AND, 0)
}

func (s *Integer_expressionContext) BIT_OR() antlr.TerminalNode {
	return s.GetToken(calcit2ParserBIT_OR, 0)
}

func (s *Integer_expressionContext) BIT_XOR() antlr.TerminalNode {
	return s.GetToken(calcit2ParserBIT_XOR, 0)
}

func (s *Integer_expressionContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(calcit2ParserBIT_CLEAR, 0)
}

func (s *Integer_expressionContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLSHIFT, 0)
}

func (s *Integer_expressionContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRSHIFT, 0)
}

func (s *Integer_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterInteger_expression(s)
	}
}

func (s *Integer_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitInteger_expression(s)
	}
}

func (s *Integer_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitInteger_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Integer_expression() (localctx IInteger_expressionContext) {
	return p.integer_expression(0)
}

func (p *calcit2Parser) integer_expression(_p int) (localctx IInteger_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewInteger_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IInteger_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, calcit2ParserRULE_integer_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case calcit2ParserLPAREN:
		{
			p.SetState(61)
			p.Match(calcit2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)

			var _x = p.integer_expression(0)

			localctx.(*Integer_expressionContext).right = _x
		}
		{
			p.SetState(63)
			p.Match(calcit2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case calcit2ParserHEX_TYPE, calcit2ParserBIN_TYPE, calcit2ParserOCT_TYPE, calcit2ParserF_TYPE, calcit2ParserI_TYPE:
		{
			p.SetState(65)
			p.Integer_cast_expression()
		}

	case calcit2ParserP_CAST:
		{
			p.SetState(66)
			p.Integer_pcast_expression()
		}

	case calcit2ParserMINUS:
		{
			p.SetState(67)

			var _m = p.Match(calcit2ParserMINUS)

			localctx.(*Integer_expressionContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)

			var _x = p.integer_expression(7)

			localctx.(*Integer_expressionContext).right = _x
		}

	case calcit2ParserNOT:
		{
			p.SetState(69)

			var _m = p.Match(calcit2ParserNOT)

			localctx.(*Integer_expressionContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)

			var _x = p.integer_expression(6)

			localctx.(*Integer_expressionContext).right = _x
		}

	case calcit2ParserIDENTIFIER, calcit2ParserHEX_LITERAL, calcit2ParserBIN_LITERAL, calcit2ParserOCT_LITERAL, calcit2ParserINTEGER_LITERAL:
		{
			p.SetState(71)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17042430230528) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(86)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
			case 1:
				localctx = NewInteger_expressionContext(p, _parentctx, _parentState)
				localctx.(*Integer_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_integer_expression)
				p.SetState(74)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(75)

					var _m = p.Match(calcit2ParserPOW)

					localctx.(*Integer_expressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(76)

					var _x = p.integer_expression(6)

					localctx.(*Integer_expressionContext).right = _x
				}

			case 2:
				localctx = NewInteger_expressionContext(p, _parentctx, _parentState)
				localctx.(*Integer_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_integer_expression)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(78)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Integer_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Integer_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(79)

					var _x = p.integer_expression(5)

					localctx.(*Integer_expressionContext).right = _x
				}

			case 3:
				localctx = NewInteger_expressionContext(p, _parentctx, _parentState)
				localctx.(*Integer_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_integer_expression)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(81)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Integer_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == calcit2ParserPLUS || _la == calcit2ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Integer_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(82)

					var _x = p.integer_expression(4)

					localctx.(*Integer_expressionContext).right = _x
				}

			case 4:
				localctx = NewInteger_expressionContext(p, _parentctx, _parentState)
				localctx.(*Integer_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_integer_expression)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(84)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Integer_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&16515072) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Integer_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(85)

					var _x = p.integer_expression(3)

					localctx.(*Integer_expressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInteger_cast_expressionContext is an interface to support dynamic dispatch.
type IInteger_cast_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() INumberContext

	// SetRight sets the right rule contexts.
	SetRight(INumberContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	Number() INumberContext
	HEX_TYPE() antlr.TerminalNode
	BIN_TYPE() antlr.TerminalNode
	OCT_TYPE() antlr.TerminalNode
	I_TYPE() antlr.TerminalNode
	F_TYPE() antlr.TerminalNode

	// IsInteger_cast_expressionContext differentiates from other interfaces.
	IsInteger_cast_expressionContext()
}

type Integer_cast_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	right  INumberContext
}

func NewEmptyInteger_cast_expressionContext() *Integer_cast_expressionContext {
	var p = new(Integer_cast_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_integer_cast_expression
	return p
}

func InitEmptyInteger_cast_expressionContext(p *Integer_cast_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_integer_cast_expression
}

func (*Integer_cast_expressionContext) IsInteger_cast_expressionContext() {}

func NewInteger_cast_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_cast_expressionContext {
	var p = new(Integer_cast_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_integer_cast_expression

	return p
}

func (s *Integer_cast_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_cast_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Integer_cast_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Integer_cast_expressionContext) GetRight() INumberContext { return s.right }

func (s *Integer_cast_expressionContext) SetRight(v INumberContext) { s.right = v }

func (s *Integer_cast_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Integer_cast_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Integer_cast_expressionContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *Integer_cast_expressionContext) HEX_TYPE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserHEX_TYPE, 0)
}

func (s *Integer_cast_expressionContext) BIN_TYPE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserBIN_TYPE, 0)
}

func (s *Integer_cast_expressionContext) OCT_TYPE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserOCT_TYPE, 0)
}

func (s *Integer_cast_expressionContext) I_TYPE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserI_TYPE, 0)
}

func (s *Integer_cast_expressionContext) F_TYPE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserF_TYPE, 0)
}

func (s *Integer_cast_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_cast_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_cast_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterInteger_cast_expression(s)
	}
}

func (s *Integer_cast_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitInteger_cast_expression(s)
	}
}

func (s *Integer_cast_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitInteger_cast_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Integer_cast_expression() (localctx IInteger_cast_expressionContext) {
	localctx = NewInteger_cast_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, calcit2ParserRULE_integer_cast_expression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Integer_cast_expressionContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33285996544) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Integer_cast_expressionContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(92)
		p.Match(calcit2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)

		var _x = p.Number()

		localctx.(*Integer_cast_expressionContext).right = _x
	}
	{
		p.SetState(94)
		p.Match(calcit2ParserRPAREN)
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

// IInteger_pcast_expressionContext is an interface to support dynamic dispatch.
type IInteger_pcast_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() IFloat_expressionContext

	// SetRight sets the right rule contexts.
	SetRight(IFloat_expressionContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	P_CAST() antlr.TerminalNode
	Float_expression() IFloat_expressionContext

	// IsInteger_pcast_expressionContext differentiates from other interfaces.
	IsInteger_pcast_expressionContext()
}

type Integer_pcast_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	right  IFloat_expressionContext
}

func NewEmptyInteger_pcast_expressionContext() *Integer_pcast_expressionContext {
	var p = new(Integer_pcast_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_integer_pcast_expression
	return p
}

func InitEmptyInteger_pcast_expressionContext(p *Integer_pcast_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_integer_pcast_expression
}

func (*Integer_pcast_expressionContext) IsInteger_pcast_expressionContext() {}

func NewInteger_pcast_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_pcast_expressionContext {
	var p = new(Integer_pcast_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_integer_pcast_expression

	return p
}

func (s *Integer_pcast_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_pcast_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Integer_pcast_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Integer_pcast_expressionContext) GetRight() IFloat_expressionContext { return s.right }

func (s *Integer_pcast_expressionContext) SetRight(v IFloat_expressionContext) { s.right = v }

func (s *Integer_pcast_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Integer_pcast_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Integer_pcast_expressionContext) P_CAST() antlr.TerminalNode {
	return s.GetToken(calcit2ParserP_CAST, 0)
}

func (s *Integer_pcast_expressionContext) Float_expression() IFloat_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_expressionContext)
}

func (s *Integer_pcast_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_pcast_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_pcast_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterInteger_pcast_expression(s)
	}
}

func (s *Integer_pcast_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitInteger_pcast_expression(s)
	}
}

func (s *Integer_pcast_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitInteger_pcast_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Integer_pcast_expression() (localctx IInteger_pcast_expressionContext) {
	localctx = NewInteger_pcast_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, calcit2ParserRULE_integer_pcast_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)

		var _m = p.Match(calcit2ParserP_CAST)

		localctx.(*Integer_pcast_expressionContext).op = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.Match(calcit2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)

		var _x = p.float_expression(0)

		localctx.(*Integer_pcast_expressionContext).right = _x
	}
	{
		p.SetState(99)
		p.Match(calcit2ParserRPAREN)
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

// IFloat_expressionContext is an interface to support dynamic dispatch.
type IFloat_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IFloat_expressionContext

	// GetRight returns the right rule contexts.
	GetRight() IFloat_expressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IFloat_expressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IFloat_expressionContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllFloat_expression() []IFloat_expressionContext
	Float_expression(i int) IFloat_expressionContext
	Float_cast_expression() IFloat_cast_expressionContext
	Float_pcast_expression() IFloat_pcast_expressionContext
	MINUS() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	POW() antlr.TerminalNode
	MULTIPLY() antlr.TerminalNode
	DIVIDE() antlr.TerminalNode
	PLUS() antlr.TerminalNode

	// IsFloat_expressionContext differentiates from other interfaces.
	IsFloat_expressionContext()
}

type Float_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IFloat_expressionContext
	right  IFloat_expressionContext
	op     antlr.Token
}

func NewEmptyFloat_expressionContext() *Float_expressionContext {
	var p = new(Float_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_float_expression
	return p
}

func InitEmptyFloat_expressionContext(p *Float_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_float_expression
}

func (*Float_expressionContext) IsFloat_expressionContext() {}

func NewFloat_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_expressionContext {
	var p = new(Float_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_float_expression

	return p
}

func (s *Float_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Float_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Float_expressionContext) GetLeft() IFloat_expressionContext { return s.left }

func (s *Float_expressionContext) GetRight() IFloat_expressionContext { return s.right }

func (s *Float_expressionContext) SetLeft(v IFloat_expressionContext) { s.left = v }

func (s *Float_expressionContext) SetRight(v IFloat_expressionContext) { s.right = v }

func (s *Float_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Float_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Float_expressionContext) AllFloat_expression() []IFloat_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFloat_expressionContext); ok {
			len++
		}
	}

	tst := make([]IFloat_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFloat_expressionContext); ok {
			tst[i] = t.(IFloat_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Float_expressionContext) Float_expression(i int) IFloat_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_expressionContext); ok {
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

	return t.(IFloat_expressionContext)
}

func (s *Float_expressionContext) Float_cast_expression() IFloat_cast_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_cast_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_cast_expressionContext)
}

func (s *Float_expressionContext) Float_pcast_expression() IFloat_pcast_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_pcast_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_pcast_expressionContext)
}

func (s *Float_expressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(calcit2ParserMINUS, 0)
}

func (s *Float_expressionContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserFLOAT_LITERAL, 0)
}

func (s *Float_expressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(calcit2ParserIDENTIFIER, 0)
}

func (s *Float_expressionContext) POW() antlr.TerminalNode {
	return s.GetToken(calcit2ParserPOW, 0)
}

func (s *Float_expressionContext) MULTIPLY() antlr.TerminalNode {
	return s.GetToken(calcit2ParserMULTIPLY, 0)
}

func (s *Float_expressionContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserDIVIDE, 0)
}

func (s *Float_expressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(calcit2ParserPLUS, 0)
}

func (s *Float_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterFloat_expression(s)
	}
}

func (s *Float_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitFloat_expression(s)
	}
}

func (s *Float_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitFloat_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Float_expression() (localctx IFloat_expressionContext) {
	return p.float_expression(0)
}

func (p *calcit2Parser) float_expression(_p int) (localctx IFloat_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewFloat_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IFloat_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, calcit2ParserRULE_float_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case calcit2ParserLPAREN:
		{
			p.SetState(102)
			p.Match(calcit2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)

			var _x = p.float_expression(0)

			localctx.(*Float_expressionContext).right = _x
		}
		{
			p.SetState(104)
			p.Match(calcit2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case calcit2ParserF_TYPE:
		{
			p.SetState(106)
			p.Float_cast_expression()
		}

	case calcit2ParserP_CAST:
		{
			p.SetState(107)
			p.Float_pcast_expression()
		}

	case calcit2ParserMINUS:
		{
			p.SetState(108)

			var _m = p.Match(calcit2ParserMINUS)

			localctx.(*Float_expressionContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)

			var _x = p.float_expression(5)

			localctx.(*Float_expressionContext).right = _x
		}

	case calcit2ParserIDENTIFIER, calcit2ParserFLOAT_LITERAL:
		{
			p.SetState(110)
			_la = p.GetTokenStream().LA(1)

			if !(_la == calcit2ParserIDENTIFIER || _la == calcit2ParserFLOAT_LITERAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				localctx.(*Float_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_float_expression)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(114)

					var _m = p.Match(calcit2ParserPOW)

					localctx.(*Float_expressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(115)

					var _x = p.float_expression(5)

					localctx.(*Float_expressionContext).right = _x
				}

			case 2:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				localctx.(*Float_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_float_expression)
				p.SetState(116)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(117)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Float_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == calcit2ParserMULTIPLY || _la == calcit2ParserDIVIDE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Float_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(118)

					var _x = p.float_expression(4)

					localctx.(*Float_expressionContext).right = _x
				}

			case 3:
				localctx = NewFloat_expressionContext(p, _parentctx, _parentState)
				localctx.(*Float_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_float_expression)
				p.SetState(119)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(120)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Float_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == calcit2ParserPLUS || _la == calcit2ParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Float_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(121)

					var _x = p.float_expression(3)

					localctx.(*Float_expressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFloat_cast_expressionContext is an interface to support dynamic dispatch.
type IFloat_cast_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() INumberContext

	// SetRight sets the right rule contexts.
	SetRight(INumberContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	F_TYPE() antlr.TerminalNode
	Number() INumberContext

	// IsFloat_cast_expressionContext differentiates from other interfaces.
	IsFloat_cast_expressionContext()
}

type Float_cast_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	right  INumberContext
}

func NewEmptyFloat_cast_expressionContext() *Float_cast_expressionContext {
	var p = new(Float_cast_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_float_cast_expression
	return p
}

func InitEmptyFloat_cast_expressionContext(p *Float_cast_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_float_cast_expression
}

func (*Float_cast_expressionContext) IsFloat_cast_expressionContext() {}

func NewFloat_cast_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_cast_expressionContext {
	var p = new(Float_cast_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_float_cast_expression

	return p
}

func (s *Float_cast_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_cast_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Float_cast_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Float_cast_expressionContext) GetRight() INumberContext { return s.right }

func (s *Float_cast_expressionContext) SetRight(v INumberContext) { s.right = v }

func (s *Float_cast_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Float_cast_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Float_cast_expressionContext) F_TYPE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserF_TYPE, 0)
}

func (s *Float_cast_expressionContext) Number() INumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *Float_cast_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_cast_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_cast_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterFloat_cast_expression(s)
	}
}

func (s *Float_cast_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitFloat_cast_expression(s)
	}
}

func (s *Float_cast_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitFloat_cast_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Float_cast_expression() (localctx IFloat_cast_expressionContext) {
	localctx = NewFloat_cast_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, calcit2ParserRULE_float_cast_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)

		var _m = p.Match(calcit2ParserF_TYPE)

		localctx.(*Float_cast_expressionContext).op = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.Match(calcit2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)

		var _x = p.Number()

		localctx.(*Float_cast_expressionContext).right = _x
	}
	{
		p.SetState(130)
		p.Match(calcit2ParserRPAREN)
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

// IFloat_pcast_expressionContext is an interface to support dynamic dispatch.
type IFloat_pcast_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() IInteger_expressionContext

	// SetRight sets the right rule contexts.
	SetRight(IInteger_expressionContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	P_CAST() antlr.TerminalNode
	Integer_expression() IInteger_expressionContext

	// IsFloat_pcast_expressionContext differentiates from other interfaces.
	IsFloat_pcast_expressionContext()
}

type Float_pcast_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	right  IInteger_expressionContext
}

func NewEmptyFloat_pcast_expressionContext() *Float_pcast_expressionContext {
	var p = new(Float_pcast_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_float_pcast_expression
	return p
}

func InitEmptyFloat_pcast_expressionContext(p *Float_pcast_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_float_pcast_expression
}

func (*Float_pcast_expressionContext) IsFloat_pcast_expressionContext() {}

func NewFloat_pcast_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_pcast_expressionContext {
	var p = new(Float_pcast_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_float_pcast_expression

	return p
}

func (s *Float_pcast_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_pcast_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Float_pcast_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Float_pcast_expressionContext) GetRight() IInteger_expressionContext { return s.right }

func (s *Float_pcast_expressionContext) SetRight(v IInteger_expressionContext) { s.right = v }

func (s *Float_pcast_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Float_pcast_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Float_pcast_expressionContext) P_CAST() antlr.TerminalNode {
	return s.GetToken(calcit2ParserP_CAST, 0)
}

func (s *Float_pcast_expressionContext) Integer_expression() IInteger_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInteger_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInteger_expressionContext)
}

func (s *Float_pcast_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_pcast_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_pcast_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterFloat_pcast_expression(s)
	}
}

func (s *Float_pcast_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitFloat_pcast_expression(s)
	}
}

func (s *Float_pcast_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitFloat_pcast_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Float_pcast_expression() (localctx IFloat_pcast_expressionContext) {
	localctx = NewFloat_pcast_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, calcit2ParserRULE_float_pcast_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)

		var _m = p.Match(calcit2ParserP_CAST)

		localctx.(*Float_pcast_expressionContext).op = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.Match(calcit2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)

		var _x = p.integer_expression(0)

		localctx.(*Float_pcast_expressionContext).right = _x
	}
	{
		p.SetState(135)
		p.Match(calcit2ParserRPAREN)
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

// IBool_expressionContext is an interface to support dynamic dispatch.
type IBool_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IBool_expressionContext

	// GetRight returns the right rule contexts.
	GetRight() IBool_expressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IBool_expressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IBool_expressionContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllBool_expression() []IBool_expressionContext
	Bool_expression(i int) IBool_expressionContext
	Bool_cast_expression() IBool_cast_expressionContext
	Bool_number_operation() IBool_number_operationContext
	NOT() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LOGICAL_AND() antlr.TerminalNode
	LOGICAL_OR() antlr.TerminalNode

	// IsBool_expressionContext differentiates from other interfaces.
	IsBool_expressionContext()
}

type Bool_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IBool_expressionContext
	right  IBool_expressionContext
	op     antlr.Token
}

func NewEmptyBool_expressionContext() *Bool_expressionContext {
	var p = new(Bool_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_bool_expression
	return p
}

func InitEmptyBool_expressionContext(p *Bool_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_bool_expression
}

func (*Bool_expressionContext) IsBool_expressionContext() {}

func NewBool_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_expressionContext {
	var p = new(Bool_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_bool_expression

	return p
}

func (s *Bool_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Bool_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Bool_expressionContext) GetLeft() IBool_expressionContext { return s.left }

func (s *Bool_expressionContext) GetRight() IBool_expressionContext { return s.right }

func (s *Bool_expressionContext) SetLeft(v IBool_expressionContext) { s.left = v }

func (s *Bool_expressionContext) SetRight(v IBool_expressionContext) { s.right = v }

func (s *Bool_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Bool_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Bool_expressionContext) AllBool_expression() []IBool_expressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBool_expressionContext); ok {
			len++
		}
	}

	tst := make([]IBool_expressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBool_expressionContext); ok {
			tst[i] = t.(IBool_expressionContext)
			i++
		}
	}

	return tst
}

func (s *Bool_expressionContext) Bool_expression(i int) IBool_expressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_expressionContext); ok {
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

	return t.(IBool_expressionContext)
}

func (s *Bool_expressionContext) Bool_cast_expression() IBool_cast_expressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_cast_expressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_cast_expressionContext)
}

func (s *Bool_expressionContext) Bool_number_operation() IBool_number_operationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_number_operationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_number_operationContext)
}

func (s *Bool_expressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(calcit2ParserNOT, 0)
}

func (s *Bool_expressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserEQUAL, 0)
}

func (s *Bool_expressionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserNOT_EQUAL, 0)
}

func (s *Bool_expressionContext) LOGICAL_AND() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLOGICAL_AND, 0)
}

func (s *Bool_expressionContext) LOGICAL_OR() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLOGICAL_OR, 0)
}

func (s *Bool_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterBool_expression(s)
	}
}

func (s *Bool_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitBool_expression(s)
	}
}

func (s *Bool_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitBool_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Bool_expression() (localctx IBool_expressionContext) {
	return p.bool_expression(0)
}

func (p *calcit2Parser) bool_expression(_p int) (localctx IBool_expressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBool_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBool_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, calcit2ParserRULE_bool_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(138)
			p.Match(calcit2ParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)

			var _x = p.bool_expression(0)

			localctx.(*Bool_expressionContext).right = _x
		}
		{
			p.SetState(140)
			p.Match(calcit2ParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(142)
			p.Bool_cast_expression()
		}

	case 3:
		{
			p.SetState(143)
			p.Bool_number_operation()
		}

	case 4:
		{
			p.SetState(144)

			var _m = p.Match(calcit2ParserNOT)

			localctx.(*Bool_expressionContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)

			var _x = p.bool_expression(2)

			localctx.(*Bool_expressionContext).right = _x
		}

	case 5:
		{
			p.SetState(146)
			_la = p.GetTokenStream().LA(1)

			if !(_la == calcit2ParserT__0 || _la == calcit2ParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBool_expressionContext(p, _parentctx, _parentState)
				localctx.(*Bool_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_bool_expression)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(150)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Bool_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == calcit2ParserEQUAL || _la == calcit2ParserNOT_EQUAL) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Bool_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(151)

					var _x = p.bool_expression(5)

					localctx.(*Bool_expressionContext).right = _x
				}

			case 2:
				localctx = NewBool_expressionContext(p, _parentctx, _parentState)
				localctx.(*Bool_expressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, calcit2ParserRULE_bool_expression)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(153)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Bool_expressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == calcit2ParserLOGICAL_AND || _la == calcit2ParserLOGICAL_OR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Bool_expressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(154)

					var _x = p.bool_expression(4)

					localctx.(*Bool_expressionContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBool_cast_expressionContext is an interface to support dynamic dispatch.
type IBool_cast_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetRight returns the right rule contexts.
	GetRight() IAnyContext

	// SetRight sets the right rule contexts.
	SetRight(IAnyContext)

	// Getter signatures
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	B_TYPE() antlr.TerminalNode
	Any() IAnyContext

	// IsBool_cast_expressionContext differentiates from other interfaces.
	IsBool_cast_expressionContext()
}

type Bool_cast_expressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     antlr.Token
	right  IAnyContext
}

func NewEmptyBool_cast_expressionContext() *Bool_cast_expressionContext {
	var p = new(Bool_cast_expressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_bool_cast_expression
	return p
}

func InitEmptyBool_cast_expressionContext(p *Bool_cast_expressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_bool_cast_expression
}

func (*Bool_cast_expressionContext) IsBool_cast_expressionContext() {}

func NewBool_cast_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_cast_expressionContext {
	var p = new(Bool_cast_expressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_bool_cast_expression

	return p
}

func (s *Bool_cast_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_cast_expressionContext) GetOp() antlr.Token { return s.op }

func (s *Bool_cast_expressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *Bool_cast_expressionContext) GetRight() IAnyContext { return s.right }

func (s *Bool_cast_expressionContext) SetRight(v IAnyContext) { s.right = v }

func (s *Bool_cast_expressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLPAREN, 0)
}

func (s *Bool_cast_expressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(calcit2ParserRPAREN, 0)
}

func (s *Bool_cast_expressionContext) B_TYPE() antlr.TerminalNode {
	return s.GetToken(calcit2ParserB_TYPE, 0)
}

func (s *Bool_cast_expressionContext) Any() IAnyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAnyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAnyContext)
}

func (s *Bool_cast_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_cast_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_cast_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterBool_cast_expression(s)
	}
}

func (s *Bool_cast_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitBool_cast_expression(s)
	}
}

func (s *Bool_cast_expressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitBool_cast_expression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Bool_cast_expression() (localctx IBool_cast_expressionContext) {
	localctx = NewBool_cast_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, calcit2ParserRULE_bool_cast_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)

		var _m = p.Match(calcit2ParserB_TYPE)

		localctx.(*Bool_cast_expressionContext).op = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.Match(calcit2ParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(162)

		var _x = p.Any()

		localctx.(*Bool_cast_expressionContext).right = _x
	}
	{
		p.SetState(163)
		p.Match(calcit2ParserRPAREN)
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

// IBool_number_operationContext is an interface to support dynamic dispatch.
type IBool_number_operationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() INumberContext

	// GetRight returns the right rule contexts.
	GetRight() INumberContext

	// SetLeft sets the left rule contexts.
	SetLeft(INumberContext)

	// SetRight sets the right rule contexts.
	SetRight(INumberContext)

	// Getter signatures
	AllNumber() []INumberContext
	Number(i int) INumberContext
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LESS() antlr.TerminalNode
	GREATER() antlr.TerminalNode
	LESS_OR_EQUAL() antlr.TerminalNode
	GREATER_OR_EQUAL() antlr.TerminalNode

	// IsBool_number_operationContext differentiates from other interfaces.
	IsBool_number_operationContext()
}

type Bool_number_operationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	left   INumberContext
	op     antlr.Token
	right  INumberContext
}

func NewEmptyBool_number_operationContext() *Bool_number_operationContext {
	var p = new(Bool_number_operationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_bool_number_operation
	return p
}

func InitEmptyBool_number_operationContext(p *Bool_number_operationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = calcit2ParserRULE_bool_number_operation
}

func (*Bool_number_operationContext) IsBool_number_operationContext() {}

func NewBool_number_operationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_number_operationContext {
	var p = new(Bool_number_operationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = calcit2ParserRULE_bool_number_operation

	return p
}

func (s *Bool_number_operationContext) GetParser() antlr.Parser { return s.parser }

func (s *Bool_number_operationContext) GetOp() antlr.Token { return s.op }

func (s *Bool_number_operationContext) SetOp(v antlr.Token) { s.op = v }

func (s *Bool_number_operationContext) GetLeft() INumberContext { return s.left }

func (s *Bool_number_operationContext) GetRight() INumberContext { return s.right }

func (s *Bool_number_operationContext) SetLeft(v INumberContext) { s.left = v }

func (s *Bool_number_operationContext) SetRight(v INumberContext) { s.right = v }

func (s *Bool_number_operationContext) AllNumber() []INumberContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumberContext); ok {
			len++
		}
	}

	tst := make([]INumberContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumberContext); ok {
			tst[i] = t.(INumberContext)
			i++
		}
	}

	return tst
}

func (s *Bool_number_operationContext) Number(i int) INumberContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberContext); ok {
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

	return t.(INumberContext)
}

func (s *Bool_number_operationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserEQUAL, 0)
}

func (s *Bool_number_operationContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserNOT_EQUAL, 0)
}

func (s *Bool_number_operationContext) LESS() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLESS, 0)
}

func (s *Bool_number_operationContext) GREATER() antlr.TerminalNode {
	return s.GetToken(calcit2ParserGREATER, 0)
}

func (s *Bool_number_operationContext) LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserLESS_OR_EQUAL, 0)
}

func (s *Bool_number_operationContext) GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(calcit2ParserGREATER_OR_EQUAL, 0)
}

func (s *Bool_number_operationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_number_operationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_number_operationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.EnterBool_number_operation(s)
	}
}

func (s *Bool_number_operationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(calcit2Listener); ok {
		listenerT.ExitBool_number_operation(s)
	}
}

func (s *Bool_number_operationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case calcit2Visitor:
		return t.VisitBool_number_operation(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *calcit2Parser) Bool_number_operation() (localctx IBool_number_operationContext) {
	localctx = NewBool_number_operationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, calcit2ParserRULE_bool_number_operation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)

		var _x = p.Number()

		localctx.(*Bool_number_operationContext).left = _x
	}
	{
		p.SetState(166)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Bool_number_operationContext).op = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1006632984) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Bool_number_operationContext).op = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(167)

		var _x = p.Number()

		localctx.(*Bool_number_operationContext).right = _x
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

func (p *calcit2Parser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *Integer_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Integer_expressionContext)
		}
		return p.Integer_expression_Sempred(t, predIndex)

	case 9:
		var t *Float_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Float_expressionContext)
		}
		return p.Float_expression_Sempred(t, predIndex)

	case 12:
		var t *Bool_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Bool_expressionContext)
		}
		return p.Bool_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *calcit2Parser) Integer_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *calcit2Parser) Float_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *calcit2Parser) Bool_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 7:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
