// Code generated from calcit2.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // calcit2

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by calcit2Parser.
type calcit2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by calcit2Parser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by calcit2Parser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by calcit2Parser#function_assignment.
	VisitFunction_assignment(ctx *Function_assignmentContext) interface{}

	// Visit a parse tree produced by calcit2Parser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by calcit2Parser#any.
	VisitAny(ctx *AnyContext) interface{}

	// Visit a parse tree produced by calcit2Parser#integer_expression.
	VisitInteger_expression(ctx *Integer_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#integer_cast_expression.
	VisitInteger_cast_expression(ctx *Integer_cast_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#integer_pcast_expression.
	VisitInteger_pcast_expression(ctx *Integer_pcast_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#float_expression.
	VisitFloat_expression(ctx *Float_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#float_cast_expression.
	VisitFloat_cast_expression(ctx *Float_cast_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#float_pcast_expression.
	VisitFloat_pcast_expression(ctx *Float_pcast_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#bool_expression.
	VisitBool_expression(ctx *Bool_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#bool_cast_expression.
	VisitBool_cast_expression(ctx *Bool_cast_expressionContext) interface{}

	// Visit a parse tree produced by calcit2Parser#bool_number_operation.
	VisitBool_number_operation(ctx *Bool_number_operationContext) interface{}
}
