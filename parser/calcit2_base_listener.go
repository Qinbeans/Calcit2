// Code generated from calcit2.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // calcit2

import "github.com/antlr4-go/antlr/v4"

// Basecalcit2Listener is a complete listener for a parse tree produced by calcit2Parser.
type Basecalcit2Listener struct{}

var _ calcit2Listener = &Basecalcit2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basecalcit2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basecalcit2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basecalcit2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basecalcit2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *Basecalcit2Listener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *Basecalcit2Listener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *Basecalcit2Listener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *Basecalcit2Listener) ExitStatement(ctx *StatementContext) {}

// EnterFunction_assignment is called when production function_assignment is entered.
func (s *Basecalcit2Listener) EnterFunction_assignment(ctx *Function_assignmentContext) {}

// ExitFunction_assignment is called when production function_assignment is exited.
func (s *Basecalcit2Listener) ExitFunction_assignment(ctx *Function_assignmentContext) {}

// EnterExpression is called when production expression is entered.
func (s *Basecalcit2Listener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *Basecalcit2Listener) ExitExpression(ctx *ExpressionContext) {}

// EnterNumber is called when production number is entered.
func (s *Basecalcit2Listener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *Basecalcit2Listener) ExitNumber(ctx *NumberContext) {}

// EnterAny is called when production any is entered.
func (s *Basecalcit2Listener) EnterAny(ctx *AnyContext) {}

// ExitAny is called when production any is exited.
func (s *Basecalcit2Listener) ExitAny(ctx *AnyContext) {}

// EnterInteger_expression is called when production integer_expression is entered.
func (s *Basecalcit2Listener) EnterInteger_expression(ctx *Integer_expressionContext) {}

// ExitInteger_expression is called when production integer_expression is exited.
func (s *Basecalcit2Listener) ExitInteger_expression(ctx *Integer_expressionContext) {}

// EnterInteger_cast_expression is called when production integer_cast_expression is entered.
func (s *Basecalcit2Listener) EnterInteger_cast_expression(ctx *Integer_cast_expressionContext) {}

// ExitInteger_cast_expression is called when production integer_cast_expression is exited.
func (s *Basecalcit2Listener) ExitInteger_cast_expression(ctx *Integer_cast_expressionContext) {}

// EnterInteger_pcast_expression is called when production integer_pcast_expression is entered.
func (s *Basecalcit2Listener) EnterInteger_pcast_expression(ctx *Integer_pcast_expressionContext) {}

// ExitInteger_pcast_expression is called when production integer_pcast_expression is exited.
func (s *Basecalcit2Listener) ExitInteger_pcast_expression(ctx *Integer_pcast_expressionContext) {}

// EnterFloat_expression is called when production float_expression is entered.
func (s *Basecalcit2Listener) EnterFloat_expression(ctx *Float_expressionContext) {}

// ExitFloat_expression is called when production float_expression is exited.
func (s *Basecalcit2Listener) ExitFloat_expression(ctx *Float_expressionContext) {}

// EnterFloat_cast_expression is called when production float_cast_expression is entered.
func (s *Basecalcit2Listener) EnterFloat_cast_expression(ctx *Float_cast_expressionContext) {}

// ExitFloat_cast_expression is called when production float_cast_expression is exited.
func (s *Basecalcit2Listener) ExitFloat_cast_expression(ctx *Float_cast_expressionContext) {}

// EnterFloat_pcast_expression is called when production float_pcast_expression is entered.
func (s *Basecalcit2Listener) EnterFloat_pcast_expression(ctx *Float_pcast_expressionContext) {}

// ExitFloat_pcast_expression is called when production float_pcast_expression is exited.
func (s *Basecalcit2Listener) ExitFloat_pcast_expression(ctx *Float_pcast_expressionContext) {}

// EnterBool_expression is called when production bool_expression is entered.
func (s *Basecalcit2Listener) EnterBool_expression(ctx *Bool_expressionContext) {}

// ExitBool_expression is called when production bool_expression is exited.
func (s *Basecalcit2Listener) ExitBool_expression(ctx *Bool_expressionContext) {}

// EnterBool_cast_expression is called when production bool_cast_expression is entered.
func (s *Basecalcit2Listener) EnterBool_cast_expression(ctx *Bool_cast_expressionContext) {}

// ExitBool_cast_expression is called when production bool_cast_expression is exited.
func (s *Basecalcit2Listener) ExitBool_cast_expression(ctx *Bool_cast_expressionContext) {}

// EnterBool_number_operation is called when production bool_number_operation is entered.
func (s *Basecalcit2Listener) EnterBool_number_operation(ctx *Bool_number_operationContext) {}

// ExitBool_number_operation is called when production bool_number_operation is exited.
func (s *Basecalcit2Listener) ExitBool_number_operation(ctx *Bool_number_operationContext) {}
