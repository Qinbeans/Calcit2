package main

import (
	"calcit2/parser"
	"fmt"
	"testing"

	"github.com/antlr4-go/antlr/v4"
)

func TestVisitorInteger(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("-2 + 2")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: 0, type: 34" {
		t.Errorf("Expected 0, got %d", result)
		return
	}
}

func TestVisitorBoolean(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("true")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: true, type: 35" {
		t.Errorf("Expected true, got %s", result)
		return
	}
}

func TestVisitorFloat(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("-2.5 + 2.5")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: 0, type: 33" {
		t.Errorf("Expected 0, got %s", result)
		return
	}
}

func TestVisitorParen(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("-(2 + 2)")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: -4, type: 34" {
		t.Errorf("Expected -4, got %d", result)
		return
	}
}

func TestVisitorBooleanComplex(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("true && false || true")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: true, type: 35" {
		t.Errorf("Expected true, got %v", result)
		return
	}
}

func TestVisitorBasicPcast(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("pcast(2.5)")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: 4612811918334230528, type: 34" {
		t.Errorf("Expected 4612811918334230528, got %v", result)
		return
	}
}

func TestVisitorComplexPcast(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("pcast(pcast(2.5) / (16 >> 3))")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree)
	fmt.Printf("result: %v\n", result)
	if result.(parser.Result).String() != "value: 1.6781266645200465e-154, type: 33" {
		t.Errorf("Expected 1.6781266645200465e-154, got %v", result)
		return
	}
}

func TestVisitorBooleanComplex2(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("true && (2 > 1 || 3 < 2)")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()

	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)

	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: true, type: 35" {
		t.Errorf("Expected true, got %v", result)
		return
	}
}

func TestVisitorBooleanComplex3(t *testing.T) {
	// Setup the input
	// Gives no shits about comparing floats to ints
	is := antlr.NewInputStream("true && (2.0 == 2) && (3 != 2) && (3 >= 2) && (3 <= 3) && (3 > 2) && (2 < 3) && (2 <= 2) && (2 >= 2) && (2 < 3) && (2 > 1) && (2 != 3) && (2 == 2)")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)
	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: true, type: 35" {
		t.Errorf("Expected true, got %v", result)
		return
	}
}

func TestVisitorBooleanComplex4(t *testing.T) {
	// Setup the input
	// Check complex arithmetic with boolean operators
	is := antlr.NewInputStream("2 << 3 > 3")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// Create the Parser
	p := parser.Newcalcit2Parser(stream)
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(parser.NewTreeShapeListener(), tree)
	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree).(parser.Result)
	if result.String() != "value: true, type: 35" {
		t.Errorf("Expected true, got %v", result)
		return
	}
}

func TestVisitorIntToOct(t *testing.T) {
	// Setup the input
	is := antlr.NewInputStream("oct(10)")

	// Create the Lexer
	lexer := parser.Newcalcit2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.Newcalcit2Parser(stream)

	tree := p.Program()
	visitor := parser.NewCustomCalculatorVisitor()
	result := visitor.Visit(tree)
	fmt.Printf("result: %v\n", result)
	if result.(parser.Result).String() != "value: 12, type: 32" {
		t.Errorf("Expected 12, got %v", result)
		return
	}
}
