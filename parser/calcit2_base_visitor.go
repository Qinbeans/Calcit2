// Code generated from calcit2.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // calcit2

import "github.com/antlr4-go/antlr/v4"

type Basecalcit2Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *Basecalcit2Visitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitFunction_assignment(ctx *Function_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitAny(ctx *AnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitInteger_expression(ctx *Integer_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitInteger_cast_expression(ctx *Integer_cast_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitInteger_pcast_expression(ctx *Integer_pcast_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitFloat_expression(ctx *Float_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitFloat_cast_expression(ctx *Float_cast_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitFloat_pcast_expression(ctx *Float_pcast_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitBool_expression(ctx *Bool_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitBool_cast_expression(ctx *Bool_cast_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *Basecalcit2Visitor) VisitBool_number_operation(ctx *Bool_number_operationContext) interface{} {
	return v.VisitChildren(ctx)
}
