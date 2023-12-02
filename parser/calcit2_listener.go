// Code generated from calcit2.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // calcit2

import "github.com/antlr4-go/antlr/v4"

// calcit2Listener is a complete listener for a parse tree produced by calcit2Parser.
type calcit2Listener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterFunction_assignment is called when entering the function_assignment production.
	EnterFunction_assignment(c *Function_assignmentContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterAny is called when entering the any production.
	EnterAny(c *AnyContext)

	// EnterInteger_expression is called when entering the integer_expression production.
	EnterInteger_expression(c *Integer_expressionContext)

	// EnterInteger_cast_expression is called when entering the integer_cast_expression production.
	EnterInteger_cast_expression(c *Integer_cast_expressionContext)

	// EnterInteger_pcast_expression is called when entering the integer_pcast_expression production.
	EnterInteger_pcast_expression(c *Integer_pcast_expressionContext)

	// EnterFloat_expression is called when entering the float_expression production.
	EnterFloat_expression(c *Float_expressionContext)

	// EnterFloat_cast_expression is called when entering the float_cast_expression production.
	EnterFloat_cast_expression(c *Float_cast_expressionContext)

	// EnterFloat_pcast_expression is called when entering the float_pcast_expression production.
	EnterFloat_pcast_expression(c *Float_pcast_expressionContext)

	// EnterBool_expression is called when entering the bool_expression production.
	EnterBool_expression(c *Bool_expressionContext)

	// EnterBool_cast_expression is called when entering the bool_cast_expression production.
	EnterBool_cast_expression(c *Bool_cast_expressionContext)

	// EnterBool_number_operation is called when entering the bool_number_operation production.
	EnterBool_number_operation(c *Bool_number_operationContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitFunction_assignment is called when exiting the function_assignment production.
	ExitFunction_assignment(c *Function_assignmentContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitAny is called when exiting the any production.
	ExitAny(c *AnyContext)

	// ExitInteger_expression is called when exiting the integer_expression production.
	ExitInteger_expression(c *Integer_expressionContext)

	// ExitInteger_cast_expression is called when exiting the integer_cast_expression production.
	ExitInteger_cast_expression(c *Integer_cast_expressionContext)

	// ExitInteger_pcast_expression is called when exiting the integer_pcast_expression production.
	ExitInteger_pcast_expression(c *Integer_pcast_expressionContext)

	// ExitFloat_expression is called when exiting the float_expression production.
	ExitFloat_expression(c *Float_expressionContext)

	// ExitFloat_cast_expression is called when exiting the float_cast_expression production.
	ExitFloat_cast_expression(c *Float_cast_expressionContext)

	// ExitFloat_pcast_expression is called when exiting the float_pcast_expression production.
	ExitFloat_pcast_expression(c *Float_pcast_expressionContext)

	// ExitBool_expression is called when exiting the bool_expression production.
	ExitBool_expression(c *Bool_expressionContext)

	// ExitBool_cast_expression is called when exiting the bool_cast_expression production.
	ExitBool_cast_expression(c *Bool_cast_expressionContext)

	// ExitBool_number_operation is called when exiting the bool_number_operation production.
	ExitBool_number_operation(c *Bool_number_operationContext)
}
