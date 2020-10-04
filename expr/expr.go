package expr

import (
	"calc/scanner"
	"strconv"
)

type Expr interface {
	Accept(v Visitor) interface{}
}

type Visitor interface {
	VisitBinaryExpr(expr *Binary) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
}

type Binary struct {
	left     Expr
	operator *scanner.Token
	right    Expr
}

func NewBinary(left Expr, operator *scanner.Token, right Expr) *Binary {
	return &Binary{
		left,
		operator,
		right,
	}
}

func (b *Binary) Accept(v Visitor) interface{} {
	return v.VisitBinaryExpr(b)
}

type Grouping struct {
	expresion Expr
}

func NewGrouping(expression Expr) *Grouping {
	return &Grouping{expression}
}

func (g *Grouping) Accept(v Visitor) interface{} {
	return v.VisitGroupingExpr(g)
}

type Literal struct {
	Number float64
}

func NewLiteral(literal string) *Literal {
	val, _ := strconv.ParseFloat(literal, 64)
	return &Literal{val}
}

func (l *Literal) Accept(v Visitor) interface{} {
	return v.VisitLiteralExpr(l)
}

type Unary struct {
	operator *scanner.Token
	right    Expr
}

func NewUnary(operator *scanner.Token, right Expr) *Unary {
	return &Unary{
		operator,
		right,
	}
}

func (u *Unary) Accept(v Visitor) interface{} {
	return v.VisitUnaryExpr(u)
}
