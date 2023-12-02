package parser

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type CustomCalculatorVisitor struct {
	*Basecalcit2Visitor
}

type Result struct {
	value interface{}
	_type int
}

func (r Result) String() string {
	value_as_str := ""
	switch r._type {
	case calcit2ParserOCT_TYPE:
		value_as_str = strconv.FormatInt(r.value.(int64), 8)
	case calcit2ParserI_TYPE:
		value_as_str = strconv.FormatInt(r.value.(int64), 10)
	case calcit2ParserHEX_TYPE:
		value_as_str = strconv.FormatInt(r.value.(int64), 16)
	case calcit2ParserBIN_TYPE:
		value_as_str = strconv.FormatInt(r.value.(int64), 2)
	case calcit2ParserF_TYPE:
		value_as_str = strconv.FormatFloat(r.value.(float64), 'g', -1, 64)
	case calcit2ParserB_TYPE:
		value_as_str = strconv.FormatBool(r.value.(bool))
	default:
		value_as_str = fmt.Sprintf("%v", r.value)
	}
	return fmt.Sprintf("value: %s, type: %v", value_as_str, r._type)
}

func (r Result) Get() []interface{} {
	return []interface{}{r.value, r._type}
}

func NewCustomCalculatorVisitor() *CustomCalculatorVisitor {
	return &CustomCalculatorVisitor{
		Basecalcit2Visitor: &Basecalcit2Visitor{},
	}
}

func (v *CustomCalculatorVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *CustomCalculatorVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	result := make([]interface{}, 0)
	for _, child := range node.GetChildren() {
		if childNode, ok := child.(antlr.ParseTree); ok {
			tempResult := childNode.Accept(v)
			fmt.Printf("tempResult: %v\n", tempResult)
			if tempResult != nil {
				// if tempResult is a slice, append each element
				if tempResultSlice, ok := tempResult.([]interface{}); ok {
					for _, tempResultSliceElement := range tempResultSlice {
						result = append(result, tempResultSliceElement)
					}
				} else {
					result = append(result, tempResult)
				}
			}
		}
	}

	if len(result) == 1 {
		return result[0]
	}

	return result
}

func (v *CustomCalculatorVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CustomCalculatorVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CustomCalculatorVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CustomCalculatorVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *CustomCalculatorVisitor) VisitBool_expression(ctx *Bool_expressionContext) interface{} {
	if ctx.GetText() == "true" || ctx.GetText() == "false" {
		return Result{value: ctx.GetText() == "true", _type: calcit2ParserB_TYPE}
	}

	raw_op := ctx.GetOp()

	if raw_op == nil && ctx.GetChildCount() > 0 {
		return v.VisitChildren(ctx)
	}

	op := raw_op.GetTokenType()

	raw_left := ctx.GetLeft()
	raw_right := ctx.GetRight()

	if raw_right == nil {
		return nil
	}

	unknown_right := raw_right.Accept(v)
	right := unknown_right.(Result)

	switch op {
	case -1:
		return right
	case calcit2ParserNOT:
		right := right
		return Result{value: !right.value.(bool), _type: right._type}
	case calcit2ParserLOGICAL_AND:
		left := raw_left.Accept(v).(Result)
		right := right
		return Result{value: left.value.(bool) && right.value.(bool), _type: left._type}
	case calcit2ParserLOGICAL_OR:
		left := raw_left.Accept(v).(Result)
		right := right
		return Result{value: left.value.(bool) || right.value.(bool), _type: left._type}
	default:
		return nil
	}
}

func (v *CustomCalculatorVisitor) VisitBool_number_operation(ctx *Bool_number_operationContext) interface{} {
	op := ctx.op.GetTokenType()
	left := ctx.left.Accept(v)
	right := ctx.right.Accept(v)

	switch op {
	case calcit2ParserEQUAL:
		left := left.(Result)
		right := right.(Result)
		leftType := reflect.TypeOf(left.value).String()
		rightType := reflect.TypeOf(right.value).String()
		if leftType == "int64" {
			left.value = float64(left.value.(int64))
		}
		if rightType == "int64" {
			right.value = float64(right.value.(int64))
		}
		return Result{value: left.value == right.value, _type: calcit2LexerB_TYPE}
	case calcit2ParserNOT_EQUAL:
		left := left.(Result)
		right := right.(Result)
		leftType := reflect.TypeOf(left.value).String()
		rightType := reflect.TypeOf(right.value).String()
		if leftType == "int64" {
			left.value = float64(left.value.(int64))
		}
		if rightType == "int64" {
			right.value = float64(right.value.(int64))
		}
		return Result{value: left.value != right.value, _type: calcit2LexerB_TYPE}
	case calcit2ParserGREATER:
		left := left.(Result)
		right := right.(Result)
		leftType := reflect.TypeOf(left.value).String()
		rightType := reflect.TypeOf(right.value).String()
		if leftType == "int64" {
			left.value = float64(left.value.(int64))
		}
		if rightType == "int64" {
			right.value = float64(right.value.(int64))
		}
		return Result{value: left.value.(float64) > right.value.(float64), _type: calcit2LexerB_TYPE}
	case calcit2ParserGREATER_OR_EQUAL:
		left := left.(Result)
		right := right.(Result)
		leftType := reflect.TypeOf(left.value).String()
		rightType := reflect.TypeOf(right.value).String()
		if leftType == "int64" {
			left.value = float64(left.value.(int64))
		}
		if rightType == "int64" {
			right.value = float64(right.value.(int64))
		}
		return Result{value: left.value.(float64) >= right.value.(float64), _type: calcit2LexerB_TYPE}
	case calcit2ParserLESS:
		left := left.(Result)
		right := right.(Result)
		leftType := reflect.TypeOf(left.value).String()
		rightType := reflect.TypeOf(right.value).String()
		if leftType == "int64" {
			left.value = float64(left.value.(int64))
		}
		if rightType == "int64" {
			right.value = float64(right.value.(int64))
		}
		return Result{value: left.value.(float64) < right.value.(float64), _type: calcit2LexerB_TYPE}
	case calcit2ParserLESS_OR_EQUAL:
		left := left.(Result)
		right := right.(Result)
		leftType := reflect.TypeOf(left.value).String()
		rightType := reflect.TypeOf(right.value).String()
		if leftType == "int64" {
			left.value = float64(left.value.(int64))
		}
		if rightType == "int64" {
			right.value = float64(right.value.(int64))
		}
		return Result{value: left.value.(float64) <= right.value.(float64), _type: calcit2LexerB_TYPE}
	default:
		return nil
	}
}

func (v *CustomCalculatorVisitor) VisitInteger_expression(ctx *Integer_expressionContext) interface{} {
	if ctx.OCT_LITERAL() != nil {
		n, _ := strconv.ParseInt(ctx.OCT_LITERAL().GetText(), 8, 64)
		return Result{value: n, _type: calcit2ParserOCT_TYPE}
	} else if ctx.INTEGER_LITERAL() != nil {
		n, _ := strconv.ParseInt(ctx.INTEGER_LITERAL().GetText(), 10, 64)
		return Result{value: n, _type: calcit2ParserI_TYPE}
	} else if ctx.HEX_LITERAL() != nil {
		n, _ := strconv.ParseInt(ctx.HEX_LITERAL().GetText(), 16, 64)
		return Result{value: n, _type: calcit2ParserHEX_TYPE}
	} else if ctx.BIN_LITERAL() != nil {
		n, _ := strconv.ParseInt(ctx.BIN_LITERAL().GetText(), 2, 64)
		return Result{value: n, _type: calcit2ParserBIN_TYPE}
	}

	raw_op := ctx.GetOp()

	if raw_op == nil && ctx.GetChildCount() > 0 {
		return v.VisitChildren(ctx)
	}

	op := raw_op.GetTokenType()

	raw_left := ctx.GetLeft()
	raw_right := ctx.GetRight()

	if raw_right == nil {
		return nil
	}

	unknown_right := raw_right.Accept(v)
	right := unknown_right.(Result)

	switch op {
	case calcit2ParserMINUS:
		if raw_left == nil {
			return Result{value: -right.value.(int64), _type: right._type}
		}
		left := raw_left.Accept(v).(Result)
		//left maintains its type
		return Result{value: left.value.(int64) - right.value.(int64), _type: left._type}
	case calcit2ParserNOT:
		return Result{value: ^right.value.(int64), _type: right._type}
	case calcit2ParserPLUS:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) + right.value.(int64), _type: left._type}
	case calcit2ParserMULTIPLY:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) * right.value.(int64), _type: left._type}
	case calcit2ParserDIVIDE:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) / right.value.(int64), _type: left._type}
	case calcit2ParserMODULO:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) % right.value.(int64), _type: left._type}
	case calcit2ParserPOW:
		left := raw_left.Accept(v).(Result)
		return Result{value: (int64)(math.Pow(float64(left.value.(int64)), float64(right.value.(int64)))), _type: left._type}
	case calcit2ParserBIT_AND:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) & right.value.(int64), _type: left._type}
	case calcit2ParserBIT_OR:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) | right.value.(int64), _type: left._type}
	case calcit2ParserBIT_XOR:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) ^ right.value.(int64), _type: left._type}
	case calcit2ParserLSHIFT:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) << uint64(right.value.(int64)), _type: left._type}
	case calcit2ParserRSHIFT:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) >> uint64(right.value.(int64)), _type: left._type}
	case calcit2ParserBIT_CLEAR:
		left := raw_left.Accept(v).(Result)
		return Result{value: left.value.(int64) &^ right.value.(int64), _type: left._type}
	default:
		return nil
	}
}

func (v *CustomCalculatorVisitor) VisitInteger_cast_expression(ctx *Integer_cast_expressionContext) interface{} {
	raw_right := ctx.GetRight()
	if raw_right == nil {
		return nil
	}
	raw_op := ctx.GetOp()
	if raw_op == nil {
		return nil
	}
	unknown_right := raw_right.Accept(v)
	if unknown_right == nil {
		return nil
	}
	right := unknown_right.(Result)
	op := raw_op.GetTokenType()
	// value could be float64 or int64, so lets cast it to int64
	rightType := reflect.TypeOf(right.value).String()
	if rightType == "float64" {
		right.value = int64(right.value.(float64))
	}
	switch op {
	case calcit2ParserOCT_TYPE:
		return Result{value: right.value, _type: calcit2ParserOCT_TYPE}
	case calcit2ParserI_TYPE:
		return Result{value: right.value, _type: calcit2ParserI_TYPE}
	case calcit2ParserHEX_TYPE:
		return Result{value: right.value, _type: calcit2ParserHEX_TYPE}
	case calcit2ParserBIN_TYPE:
		return Result{value: right.value, _type: calcit2ParserBIN_TYPE}
	default:
		return nil
	}
}

func (v *CustomCalculatorVisitor) VisitInteger_pcast_expression(ctx *Integer_pcast_expressionContext) interface{} {
	right := ctx.GetRight().Accept(v).(Result)
	rightType := right._type
	if rightType == calcit2ParserF_TYPE {
		return Result{value: int64(math.Float64bits(right.value.(float64))), _type: calcit2ParserI_TYPE}
	}
	return right
}

func (v *CustomCalculatorVisitor) VisitFloat_expression(ctx *Float_expressionContext) interface{} {
	if ctx.FLOAT_LITERAL() != nil {
		n, _ := strconv.ParseFloat(ctx.FLOAT_LITERAL().GetText(), 64)
		return Result{value: n, _type: calcit2ParserF_TYPE}
	}

	raw_op := ctx.GetOp()

	if raw_op == nil && ctx.GetChildCount() > 0 {
		return v.VisitChildren(ctx)
	}

	op := raw_op.GetTokenType()

	raw_left := ctx.GetLeft()
	raw_right := ctx.GetRight()

	if raw_right == nil {
		return nil
	}

	right := raw_right.Accept(v)

	switch op {
	case calcit2ParserMINUS:
		if raw_left == nil {
			return Result{value: -right.(Result).value.(float64), _type: calcit2ParserF_TYPE}
		}
		left := raw_left.Accept(v)
		return Result{value: left.(Result).value.(float64) - right.(Result).value.(float64), _type: calcit2ParserF_TYPE}
	case calcit2ParserPLUS:
		left := raw_left.Accept(v)
		return Result{value: left.(Result).value.(float64) + right.(Result).value.(float64), _type: calcit2ParserF_TYPE}
	case calcit2ParserMULTIPLY:
		left := raw_left.Accept(v)
		return Result{value: left.(Result).value.(float64) * right.(Result).value.(float64), _type: calcit2ParserF_TYPE}
	case calcit2ParserDIVIDE:
		left := raw_left.Accept(v)
		return Result{value: left.(Result).value.(float64) / right.(Result).value.(float64), _type: calcit2ParserF_TYPE}
	case calcit2ParserPOW:
		left := raw_left.Accept(v)
		return Result{value: math.Pow(left.(Result).value.(float64), right.(Result).value.(float64)), _type: calcit2ParserF_TYPE}
	default:
		return nil
	}
}

func (v *CustomCalculatorVisitor) VisitFloat_cast_expression(ctx *Float_cast_expressionContext) interface{} {
	left := ctx.right.Accept(v).(Result)
	leftType := reflect.TypeOf(left.value).String()
	if leftType == "int64" {
		left.value = float64(left.value.(int64))
	}
	return Result{value: left.value, _type: calcit2ParserF_TYPE}
}

func (v *CustomCalculatorVisitor) VisitFloat_pcast_expression(ctx *Float_pcast_expressionContext) interface{} {
	right := ctx.GetRight().Accept(v).(Result)
	rightType := right._type
	if rightType != calcit2ParserF_TYPE {
		return Result{value: math.Float64frombits(uint64(right.value.(int64))), _type: calcit2ParserF_TYPE}
	}
	return right
}
