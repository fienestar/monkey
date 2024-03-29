package ast

import (
	"bytes"
	"monkey/token"
	"monkey/types"
)

type Node interface {
	TokenLiteral() types.InputString
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() types.InputString {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return types.InputString{}
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()                  {}
func (ls *LetStatement) TokenLiteral() types.InputString { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(ls.TokenLiteral()) + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value types.InputString
}

func (i *Identifier) expressionNode()                 {}
func (i *Identifier) TokenLiteral() types.InputString { return i.Token.Literal }
func (i *Identifier) String() string {
	return string(i.Value)
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()                  {}
func (rs *ReturnStatement) TokenLiteral() types.InputString { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(string(rs.TokenLiteral()) + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()                  {}
func (es *ExpressionStatement) TokenLiteral() types.InputString { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}
