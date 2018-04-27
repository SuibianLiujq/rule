// File: ast.go
//
// This file implements the Abstract Syntax Tree of SCRIPT.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package script

import (
	"fmt"
)

// Token - Interface class of all SCRIPT tokens.
//
// @Type():  Returns the type of token.
// @Label(): Returns the token's label.
// @Str():   Returns the string value of token.
// @Src():  Returns the line number.
type Token interface {
	Type() TokenType
	Label() string
	Str() string
	FmtStr() string
	Src() string
}

// TokenIdent - IDENT token.
//
// @Value: String value if IDENT name.
// @line:  Line number.
type TokenIdent struct {
	Value string
	src   string
}

func (t *TokenIdent) Init(v interface{}, s string) Token { t.Value, t.src = v.(string), s; return t }
func (t *TokenIdent) Type() TokenType                    { return T_IDENT }
func (t *TokenIdent) Src() string                        { return t.src }
func (t *TokenIdent) Label() string                      { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenIdent) Str() string                        { return t.Value }
func (t *TokenIdent) FmtStr() string                     { return t.Str() }

// TokenStr - STR token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenStr struct {
	Value string
	src   string
}

func (t *TokenStr) Init(v interface{}, s string) Token { t.Value, t.src = v.(string), s; return t }
func (t *TokenStr) Type() TokenType                    { return T_STR }
func (t *TokenStr) Src() string                        { return t.src }
func (t *TokenStr) Label() string                      { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenStr) Str() string                        { return t.Value }
func (t *TokenStr) FmtStr() string                     { return fmt.Sprintf("'%s'", t.Value) }

// TokenInt - INT token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenInt struct {
	Value int64
	src   string
}

func (t *TokenInt) Init(v interface{}, s string) Token { t.Value, t.src = v.(int64), s; return t }
func (t *TokenInt) Type() TokenType                    { return T_INT }
func (t *TokenInt) Src() string                        { return t.src }
func (t *TokenInt) Label() string                      { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenInt) Str() string                        { return fmt.Sprintf("%d", t.Value) }
func (t *TokenInt) FmtStr() string                     { return t.Str() }

// TokenFloat - FLOAT token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenFloat struct {
	Value float64
	src   string
}

func (t *TokenFloat) Init(v interface{}, s string) Token { t.Value, t.src = v.(float64), s; return t }
func (t *TokenFloat) Type() TokenType                    { return T_FLOAT }
func (t *TokenFloat) Src() string                        { return t.src }
func (t *TokenFloat) Label() string                      { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenFloat) Str() string                        { return fmt.Sprintf("%v", t.Value) }
func (t *TokenFloat) FmtStr() string                     { return t.Str() }

// TokenBool - BOOL token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenBool struct {
	Value bool
	src   string
}

func (t *TokenBool) Init(v interface{}, s string) Token { t.Value, t.src = v.(bool), s; return t }
func (t *TokenBool) Type() TokenType                    { return T_BOOL }
func (t *TokenBool) Src() string                        { return t.src }
func (t *TokenBool) Label() string                      { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenBool) Str() string                        { return fmt.Sprintf("%v", t.Value) }
func (t *TokenBool) FmtStr() string                     { return t.Str() }

// TokenNull - NULL token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenNull struct {
	src string
}

func (t *TokenNull) Init(s string) Token { t.src = s; return t }
func (t *TokenNull) Type() TokenType     { return T_NULL }
func (t *TokenNull) Src() string         { return t.src }
func (t *TokenNull) Label() string       { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenNull) Str() string         { return "null" }
func (t *TokenNull) FmtStr() string      { return t.Str() }

// TokenList - LIST token.
//
// @List: List of tokens.
type TokenList struct {
	List []Token
}

func (t *TokenList) Init(list []Token) Token { t.List = list; return t }
func (t *TokenList) Type() TokenType         { return T_LIST }
func (t *TokenList) Src() string             { return "" }
func (t *TokenList) Label() string           { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenList) Str() string             { return "[...]" }
func (t *TokenList) FmtStr() string          { return t.Str() }

// TokenPair - PAIR token.
//
// @Key:   Key token of the pair.
// @Value: Value token of the pair.
type TokenPair struct {
	Key, Value Token
}

func (t *TokenPair) Init(k, v Token) Token { t.Key, t.Value = k, v; return t }
func (t *TokenPair) Type() TokenType       { return T_PAIR }
func (t *TokenPair) Src() string           { return t.Key.Src() }
func (t *TokenPair) Label() string         { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenPair) Str() string           { return t.Key.FmtStr() + " : " + t.Value.FmtStr() }
func (t *TokenPair) FmtStr() string        { return t.Str() }

// TokenDict - DICT token.
//
// @List: List of tokens.
type TokenDict struct {
	List []Token
}

func (t *TokenDict) Init(list []Token) Token { t.List = list; return t }
func (t *TokenDict) Type() TokenType         { return T_DICT }
func (t *TokenDict) Src() string             { return "" }
func (t *TokenDict) Label() string           { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenDict) Str() string             { return "{...}" }
func (t *TokenDict) FmtStr() string          { return t.Str() }

// TokenFunc - FUNC token.
//
// @Name: Function name.
// @List: List of tokens.
// @line: Line number.
type TokenFunc struct {
	Name string
	List []Token
	src  string
}

func (t *TokenFunc) Init(name Token, list []Token) Token {
	t.Name, t.List, t.src = name.Str(), list, name.Src()
	return t
}
func (t *TokenFunc) Type() TokenType { return T_FUNC }
func (t *TokenFunc) Src() string     { return t.src }
func (t *TokenFunc) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenFunc) Str() string     { return t.Name + "()" }
func (t *TokenFunc) FmtStr() string  { return t.Str() }

// TokenCond - COND token.
//
// @Left, Middle, Right: Three tokens of condition expression.
type TokenCond struct {
	Left, Middle, Right Token
}

func (t *TokenCond) Init(left, middle, right Token) Token {
	t.Left, t.Middle, t.Right = left, middle, right
	return t
}
func (t *TokenCond) Type() TokenType { return T_COND }
func (t *TokenCond) Src() string     { return t.Left.Src() }
func (t *TokenCond) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenCond) Str() string     { return t.Left.Str() + " ? " + t.Middle.Str() + " : " + t.Right.Str() }
func (t *TokenCond) FmtStr() string {
	return t.Left.FmtStr() + " ? " + t.Middle.FmtStr() + " : " + t.Right.FmtStr()
}

// TokenOper - OPER token.
//
// @Left, Right: Left && Right token.
// @Operator:    Operator (token type).
// @line:        Line number.
type TokenOper struct {
	Left, Right Token
	Operator    TokenType
	src         string
}

func (t *TokenOper) Init(left Token, opt TokenType, right Token) Token {
	t.Left, t.Right, t.Operator = left, right, opt

	if t.Left != nil {
		t.src = t.Left.Src()
	} else if t.Right != nil {
		t.src = t.Right.Src()
	}

	return t
}
func (t *TokenOper) Type() TokenType { return T_OPER }
func (t *TokenOper) Src() string     { return t.src }
func (t *TokenOper) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenOper) Str() string {
	v := ""
	if t.Left != nil {
		v += t.Left.Str()
	}

	v += " " + t.Operator.Name() + " "
	if t.Right != nil {
		v += t.Right.Str()
	}

	return v
}

func (t *TokenOper) FmtStr() string {
	v := ""
	if t.Left != nil {
		v += t.Left.FmtStr()
	}

	v += " " + t.Operator.Name() + " "
	if t.Right != nil {
		v += t.Right.FmtStr()
	}

	return v
}

// TokenComp - COMP token.
//
// @Left, Right: Left && Right token.
// @Operator:    Operator (token type).
type TokenComp struct {
	Left, Right Token
	Operator    TokenType
}

func (t *TokenComp) Init(left Token, opt TokenType, right Token) Token {
	t.Left, t.Right, t.Operator = left, right, opt
	return t
}
func (t *TokenComp) Type() TokenType { return T_COMP }
func (t *TokenComp) Src() string     { return t.Left.Src() }
func (t *TokenComp) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenComp) Str() string     { return t.Left.Str() + " " + t.Operator.Name() + " " + t.Right.Str() }
func (t *TokenComp) FmtStr() string {
	return t.Left.FmtStr() + " " + t.Operator.Name() + " " + t.Right.FmtStr()
}

// TokenLogical - LOGICAL token.
//
// @Left, Right: Left && Right token.
// @Operator:    Operator (token type).
// @line:        Line number.
type TokenLogical struct {
	Left, Right Token
	Operator    TokenType
	src         string
}

func (t *TokenLogical) Init(left Token, opt TokenType, right Token) Token {
	t.Left, t.Right, t.Operator = left, right, opt

	if t.Left != nil {
		t.src = t.Left.Src()
	} else if t.Right != nil {
		t.src = t.Right.Src()
	}

	return t
}
func (t *TokenLogical) Type() TokenType { return T_LOGICAL }
func (t *TokenLogical) Src() string     { return t.src }
func (t *TokenLogical) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenLogical) Str() string {
	v := ""
	if t.Left != nil {
		v += t.Left.Str()
	}

	v += t.Operator.Name()
	if t.Right != nil {
		v += t.Right.Str()
	}

	return v
}

func (t *TokenLogical) FmtStr() string {
	v := ""
	if t.Left != nil {
		v += t.Left.FmtStr()
	}

	v += " " + t.Operator.Name() + " "
	if t.Right != nil {
		v += t.Right.FmtStr()
	}

	return v
}

// TokenIndex - INDEX token.
//
// @Left, Right: Left && Right token.
// @Operator:    Operator (token type).
// @line:        Line number.
type TokenIndex struct {
	Object, Key Token
}

func (t *TokenIndex) Init(obj, key Token) Token { t.Object, t.Key = obj, key; return t }
func (t *TokenIndex) Type() TokenType           { return T_INDEX }
func (t *TokenIndex) Src() string               { return t.Object.Src() }
func (t *TokenIndex) Label() string             { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenIndex) Str() string               { return t.Object.Str() + "[" + t.Key.Str() + "]" }
func (t *TokenIndex) FmtStr() string            { return t.Object.FmtStr() + "[" + t.Key.FmtStr() + "]" }

// TokenIn - IN token.
//
// @Left, Right: Left && Right token.
// @Operator:    Operator (token type).
// @line:        Line number.
type TokenIn struct {
	Object, Key Token
}

func (t *TokenIn) Init(key, obj Token) Token { t.Object, t.Key = obj, key; return t }
func (t *TokenIn) Type() TokenType           { return T_IN }
func (t *TokenIn) Src() string               { return t.Key.Src() }
func (t *TokenIn) Label() string             { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenIn) Str() string               { return t.Key.Str() + " in " + t.Object.Str() }
func (t *TokenIn) FmtStr() string            { return t.Key.FmtStr() + " in " + t.Object.FmtStr() }

// TokenStmts - STMTS token.
//
// @List: List of tokens.
type TokenStmts struct {
	List []Token
}

func (t *TokenStmts) Init(list []Token) Token { t.List = list; return t }
func (t *TokenStmts) Type() TokenType         { return T_STMTS }
func (t *TokenStmts) Src() string             { return "" }
func (t *TokenStmts) Label() string           { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenStmts) Str() string             { return "...; ...;" }
func (t *TokenStmts) FmtStr() string          { return t.Str() }

// TokenExpr - EXPR token.
//
// @Token: Expr token.
type TokenExpr struct {
	Token Token
}

func (t *TokenExpr) Init(token Token) Token { t.Token = token; return t }
func (t *TokenExpr) Type() TokenType        { return T_EXPR }
func (t *TokenExpr) Src() string            { return t.Token.Src() }
func (t *TokenExpr) Label() string          { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenExpr) Str() string            { return t.Token.Str() }
func (t *TokenExpr) FmtStr() string         { return t.Token.FmtStr() }

// TokenAssign - ASSIGN token.
//
// @Token: Expr token.
type TokenAssign struct {
	Left, Right Token
}

func (t *TokenAssign) Init(l, r Token) Token { t.Left, t.Right = l, r; return t }
func (t *TokenAssign) Type() TokenType       { return T_ASSIGN }
func (t *TokenAssign) Src() string           { return t.Left.Src() }
func (t *TokenAssign) Label() string         { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenAssign) Str() string           { return t.Left.Str() + " = " + t.Right.Str() }
func (t *TokenAssign) FmtStr() string        { return t.Left.FmtStr() + " = " + t.Right.FmtStr() }

// TokenIfElse - IF_ELSE token.
//
// @If:     IF token.
// @ElseIf: ELSE_IF token.
// @Else:   ELSE token.
type TokenIfElse struct {
	If, Else Token
	ElseIf   []Token
}

func (t *TokenIfElse) Init(i Token, ei []Token, e Token) Token {
	t.If, t.ElseIf, t.Else = i, ei, e
	return t
}
func (t *TokenIfElse) Type() TokenType { return T_IF_ELSE }
func (t *TokenIfElse) Src() string     { return t.If.Src() }
func (t *TokenIfElse) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenIfElse) FmtStr() string  { return t.Str() }
func (t *TokenIfElse) Str() string {
	v := t.If.Str()
	if t.ElseIf != nil && len(t.ElseIf) != 0 {
		v += " else if {...}"
	}

	if t.Else != nil {
		v += " " + t.Else.Str()
	}
	return v
}

// TokenIf - IF token.
//
// @Cond: Condition token.
// @List: List of tokens.
type TokenIf struct {
	Cond Token
	List []Token
}

func (t *TokenIf) Init(c Token, l []Token) Token { t.Cond, t.List = c, l; return t }
func (t *TokenIf) Type() TokenType               { return T_IF }
func (t *TokenIf) Src() string                   { return t.Cond.Src() }
func (t *TokenIf) Label() string                 { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenIf) Str() string                   { return "if (" + t.Cond.Str() + ") {...}" }
func (t *TokenIf) FmtStr() string                { return "if (" + t.Cond.FmtStr() + ") {...}" }

// TokenElseIf - ELSEIF token.
//
// @Cond: Condition token.
// @List: List of tokens.
type TokenElseIf struct {
	Cond Token
	List []Token
}

func (t *TokenElseIf) Init(c Token, l []Token) Token { t.Cond, t.List = c, l; return t }
func (t *TokenElseIf) Type() TokenType               { return T_ELSEIF }
func (t *TokenElseIf) Src() string                   { return t.Cond.Src() }
func (t *TokenElseIf) Label() string                 { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenElseIf) Str() string                   { return "else if (" + t.Cond.Str() + ") {...}" }
func (t *TokenElseIf) FmtStr() string                { return "else if (" + t.Cond.FmtStr() + ") {...}" }

// TokenElse - ELSE token.
//
// @List: List of tokens.
type TokenElse struct {
	List []Token
}

func (t *TokenElse) Init(list []Token) Token { t.List = list; return t }
func (t *TokenElse) Type() TokenType         { return T_ELSE }
func (t *TokenElse) Src() string             { return "" }
func (t *TokenElse) Label() string           { return fmt.Sprintf("%s, %s", t.Str(), t.Src()) }
func (t *TokenElse) Str() string             { return "else {...}" }
func (t *TokenElse) FmtStr() string          { return t.Str() }

// TokenForIter - FOR_ITER token.
//
// @Start: Intialization token.
// @Cond:  Condition token.
// @Next:  Next token.
type TokenForIter struct {
	Start, Cond, Next Token
}

func (t *TokenForIter) Init(s, c, n Token) Token { t.Start, t.Cond, t.Next = s, c, n; return t }
func (t *TokenForIter) Type() TokenType          { return T_FOR_ITER }
func (t *TokenForIter) Src() string              { return "" }
func (t *TokenForIter) Label() string            { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenForIter) Str() string              { return "(...; ...; ...)" }
func (t *TokenForIter) FmtStr() string           { return t.Str() }

// TokenFor - FOR token.
//
// @Iter: FOR iteration token.
// @List: List of statement tokens.
type TokenFor struct {
	Iter Token
	List []Token
}

func (t *TokenFor) Init(iter Token, list []Token) Token { t.Iter, t.List = iter, list; return t }
func (t *TokenFor) Type() TokenType                     { return T_FOR }
func (t *TokenFor) Src() string                         { return t.Iter.Src() }
func (t *TokenFor) Label() string                       { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenFor) Str() string                         { return "for " + t.Iter.Str() + " {...}" }
func (t *TokenFor) FmtStr() string                      { return "for " + t.Iter.FmtStr() + " {...}" }

// TokenForinIter - FORIN iteration token.
//
// @Key:    Key token.
// @Value:  Value token.
// @Object: Object token.
type TokenForinIter struct {
	Key, Value, Object Token
}

func (t *TokenForinIter) Init(k, v, o Token) Token { t.Key, t.Value, t.Object = k, v, o; return t }
func (t *TokenForinIter) Type() TokenType          { return T_FORIN_ITER }
func (t *TokenForinIter) Src() string              { return "" }
func (t *TokenForinIter) Label() string            { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenForinIter) Str() string              { return "(x, y in z)" }
func (t *TokenForinIter) FmtStr() string           { return t.Str() }

// TokenForin - FORIN token.
//
// @Iter: FORIN iteration token.
// @List: List of statement tokens.
type TokenForin struct {
	Iter Token
	List []Token
}

func (t *TokenForin) Init(iter Token, list []Token) Token { t.Iter, t.List = iter, list; return t }
func (t *TokenForin) Type() TokenType                     { return T_FORIN }
func (t *TokenForin) Src() string                         { return "" }
func (t *TokenForin) Label() string                       { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenForin) Str() string                         { return "(x, y in z)" }
func (t *TokenForin) FmtStr() string                      { return t.Str() }

// TokenContinue - CONTINUE token.
type TokenContinue struct{}

func (t *TokenContinue) Init() Token     { return t }
func (t *TokenContinue) Type() TokenType { return T_CONTINUE }
func (t *TokenContinue) Src() string     { return "" }
func (t *TokenContinue) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenContinue) Str() string     { return "continue" }
func (t *TokenContinue) FmtStr() string  { return t.Str() }

// TokenBreak - BREAK token.
type TokenBreak struct{}

func (t *TokenBreak) Init() Token     { return t }
func (t *TokenBreak) Type() TokenType { return T_BREAK }
func (t *TokenBreak) Src() string     { return "" }
func (t *TokenBreak) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenBreak) Str() string     { return "break" }
func (t *TokenBreak) FmtStr() string  { return t.Str() }

// TokenReturn - RETURN token.
type TokenReturn struct{ Token Token }

func (t *TokenReturn) Init(token Token) Token { t.Token = token; return t }
func (t *TokenReturn) Type() TokenType        { return T_RETURN }
func (t *TokenReturn) Src() string            { return "" }
func (t *TokenReturn) Label() string          { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenReturn) Str() string            { return "return" }
func (t *TokenReturn) FmtStr() string         { return t.Str() }

// TokenDefine - DEFINE token.
//
// @Name:  Name of definition.
// @Args:  Names of arguments.
// @Stmts: Statements.
type TokenDefine struct {
	Name  string
	Args  []string
	Stmts []Token
	src   string
}

func (t *TokenDefine) Init(name Token, args, list []Token) Token {
	t.Name, t.Stmts, t.src = name.Str(), list, name.Src()
	if args != nil {
		for _, item := range args {
			t.Args = append(t.Args, item.Str())
		}
	}
	return t
}
func (t *TokenDefine) Type() TokenType { return T_DEFINE }
func (t *TokenDefine) Src() string     { return t.src }
func (t *TokenDefine) Label() string   { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenDefine) Str() string     { return "def " + t.Name + " {...}" }
func (t *TokenDefine) FmtStr() string  { return t.Str() }

// TokenInclude - INCLUDE token.
//
// @Name:  Name of definition.
// @src:  Names of arguments.
type TokenInclude struct {
	Name Token
	src  string
}

func (t *TokenInclude) Init(name Token) Token { t.Name, t.src = name, name.Src(); return t }
func (t *TokenInclude) Type() TokenType       { return T_INCLUDE }
func (t *TokenInclude) Src() string           { return t.src }
func (t *TokenInclude) Label() string         { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenInclude) Str() string           { return "include '" + t.Name.Src() + "'" }
func (t *TokenInclude) FmtStr() string        { return t.Str() }

// TokenImport - IMPORT token.
//
// @Name:  Name of definition.
// @src:  Names of arguments.
type TokenImport struct {
	Name Token
	src  string
}

func (t *TokenImport) Init(name Token) Token { t.Name, t.src = name, name.Src(); return t }
func (t *TokenImport) Type() TokenType       { return T_IMPORT }
func (t *TokenImport) Src() string           { return t.src }
func (t *TokenImport) Label() string         { return fmt.Sprintf("%s, %s", t.FmtStr(), t.Src()) }
func (t *TokenImport) Str() string           { return "import '" + t.Name.Str() + "'" }
func (t *TokenImport) FmtStr() string        { return t.Str() }
