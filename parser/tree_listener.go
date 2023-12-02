package parser

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*Basecalcit2Listener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}
