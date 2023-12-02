package main

import (
	"context"
	"errors"
	"reflect"
	"strconv"

	"calcit2/parser"

	"github.com/antlr4-go/antlr/v4"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Parse(formula string) (string, error) {
	// Setup the input
	is := antlr.NewInputStream(formula)

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	visitor := parser.NewCustomCalculatorVisitor()
	unknown_result := visitor.Visit(tree)

	if unknown_result == nil {
		return "", errors.New("unknown error, result is nil")
	}

	// check type of unknown_result
	unknown_type := reflect.TypeOf(unknown_result)
	if unknown_type.String() != "parser.Result" {
		return "", errors.New("unknown error, syntax error")
	}

	result := unknown_result.(parser.Result).Get()

	switch result[1] {
	case 30: // hex
		return strconv.FormatInt(result[0].(int64), 16), nil
	case 31: // bin
		return strconv.FormatInt(result[0].(int64), 2), nil
	case 32: // oct
		return strconv.FormatInt(result[0].(int64), 8), nil
	case 33: // float64
		return strconv.FormatFloat(result[0].(float64), 'g', -1, 64), nil
	case 34: // int64
		return strconv.FormatInt(result[0].(int64), 10), nil
	case 35: // bool
		return strconv.FormatBool(result[0].(bool)), nil
	case 36: // string
		return result[0].(string), nil
	default:
		return "", errors.New("unknown type")
	}
}
