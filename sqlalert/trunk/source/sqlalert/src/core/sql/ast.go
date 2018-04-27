// File: ast.go
//
// This file implements the Abstract Syntax Tree of SQL.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package sql

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
	Str() string
}

// TokenIdent - IDENT token.
//
// @v: String value if IDENT name.
type TokenIdent struct {
	Value string
}

func (t *TokenIdent) Init(v interface{}) Token { t.Value = v.(string); return t }
func (t *TokenIdent) Type() TokenType          { return T_IDENT }
func (t *TokenIdent) Str() string              { return t.Value }

// TokenVar - VAR token.
//
// @token: Token of VAR.
type TokenVar struct {
	Double bool
	Value  string
}

func (t *TokenVar) Init(token Token, double bool) Token { t.Value = token.Str(); return t }
func (t *TokenVar) Type() TokenType                     { return T_VAR }
func (t *TokenVar) Str() string {
	if !t.Double {
		return "$(" + t.Value + ")"
	} else {
		return "$$(" + t.Value + ")"
	}
}

// TokenStr - STR token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenStr struct {
	Value string
}

func (t *TokenStr) Init(v interface{}) Token { t.Value = v.(string); return t }
func (t *TokenStr) Type() TokenType          { return T_STR }
func (t *TokenStr) Str() string              { return t.Value }

// TokenInt - INT token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenInt struct {
	Value int64
}

func (t *TokenInt) Init(v interface{}) Token { t.Value = v.(int64); return t }
func (t *TokenInt) Type() TokenType          { return T_INT }
func (t *TokenInt) Str() string              { return fmt.Sprintf("%d", t.Value) }

// TokenFloat - FLOAT token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenFloat struct {
	Value float64
}

func (t *TokenFloat) Init(v interface{}) Token { t.Value = v.(float64); return t }
func (t *TokenFloat) Type() TokenType          { return T_FLOAT }
func (t *TokenFloat) Str() string              { return fmt.Sprintf("%f", t.Value) }

// TokenBool - BOOL token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenBool struct {
	Value bool
}

func (t *TokenBool) Init(v interface{}) Token { t.Value = v.(bool); return t }
func (t *TokenBool) Type() TokenType          { return T_BOOL }
func (t *TokenBool) Str() string              { return boolStr(t.Value) }

// TokenNull - NULL token.
//
// @Value: Value of token.
// @line:  Line number.
type TokenNull struct{}

func (t *TokenNull) Init() Token     { return t }
func (t *TokenNull) Type() TokenType { return T_NULL }
func (t *TokenNull) Str() string     { return "null" }

// TokenList - LIST token.
//
// @List: List of tokens.
type TokenList struct {
	List []Token
}

func (t *TokenList) Init(list []Token) Token { t.List = list; return t }
func (t *TokenList) Type() TokenType         { return T_LIST }
func (t *TokenList) Str() string             { return "[" + join(t.List, ",") + "]" }

// TokenPair - PAIR token.
//
// @Key:   Key token of the pair.
// @Value: Value token of the pair.
type TokenPair struct {
	Key, Value Token
}

func (t *TokenPair) Init(k, v Token) Token { t.Key, t.Value = k, v; return t }
func (t *TokenPair) Type() TokenType       { return T_PAIR }
func (t *TokenPair) Str() string           { return t.Key.Str() + " : " + t.Value.Str() }

// TokenDict - DICT token.
//
// @List: List of tokens.
type TokenDict struct {
	List []Token
}

func (t *TokenDict) Init(list []Token) Token { t.List = list; return t }
func (t *TokenDict) Type() TokenType         { return T_DICT }
func (t *TokenDict) Str() string             { return "{" + join(t.List, ",") + "}" }

// TokenFunc - FUNC token.
//
// @Name: Function name.
// @List: List of tokens.
// @line: Line number.
type TokenFunc struct {
	Name string
	List []Token
}

func (t *TokenFunc) Init(n Token, l []Token) Token { t.Name, t.List = n.Str(), l; return t }
func (t *TokenFunc) Type() TokenType               { return T_FUNC }
func (t *TokenFunc) Str() string                   { return t.Name + "(" + join(t.List, ",") + ")" }

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
func (t *TokenCond) Str() string     { return t.Left.Str() + " ? " + t.Middle.Str() + " : " + t.Right.Str() }

// TokenOper - OPER token.
//
// @Left, Right: Left && Right token.
// @Operator:    Operator (token type).
// @line:        Line number.
type TokenOper struct {
	Left, Right Token
	Operator    TokenType
}

func (t *TokenOper) Init(left Token, opt TokenType, right Token) Token {
	t.Left, t.Right, t.Operator = left, right, opt
	return t
}
func (t *TokenOper) Type() TokenType { return T_OPER }
func (t *TokenOper) Str() string     { return t.Left.Str() + " " + t.Operator.Name() + " " + t.Right.Str() }

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
func (t *TokenComp) Str() string     { return t.Left.Str() + " " + t.Operator.Name() + " " + t.Right.Str() }

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
	return t
}
func (t *TokenLogical) Type() TokenType { return T_LOGICAL }
func (t *TokenLogical) Str() (v string) {
	if t.Left != nil {
		v = t.Left.Str()
	}

	v += t.Operator.Name()
	if t.Right != nil {
		v += t.Right.Str()
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
func (t *TokenIndex) Str() string               { return t.Object.Str() + "['" + t.Key.Str() + "']" }

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
func (t *TokenIn) Str() string               { return t.Key.Str() + " in " + t.Object.Str() }

// TokenStmts - STMTS token.
//
// @Select: SELECT token.
// @List:   List of tokens.
type TokenStmts struct {
	Select Token
	List   []Token
}

func (t *TokenStmts) Init(s Token, l []Token) Token { t.Select, t.List = s, l; return t }
func (t *TokenStmts) Type() TokenType               { return T_STMTS }
func (t *TokenStmts) Str() string                   { return t.Select.Str() + " " + join(t.List, " ") }

// TokenSelect - SELECT token.
//
// @List: List of tokens.
type TokenSelect struct {
	List []Token
}

func (t *TokenSelect) Init(list []Token) Token { t.List = list; return t }
func (t *TokenSelect) Type() TokenType         { return T_SELECT }
func (t *TokenSelect) Str() string             { return "SELECT " + join(t.List, ",") }

// TokenFrom - FROM token.
//
// @List: List of tokens.
type TokenFrom struct {
	List []Token
}

func (t *TokenFrom) Init(list []Token) Token { t.List = list; return t }
func (t *TokenFrom) Type() TokenType         { return T_FROM }
func (t *TokenFrom) Str() string             { return "FROM " + join(t.List, ",") }

// TokenWhere - WHERE token.
//
// @List: List of tokens.
type TokenWhere struct {
	Token Token
}

func (t *TokenWhere) Init(token Token) Token { t.Token = token; return t }
func (t *TokenWhere) Type() TokenType        { return T_WHERE }
func (t *TokenWhere) Str() string            { return "WHERE " + t.Token.Str() }

// TokenWhereBy - GROUP BY token.
//
// @List: List of tokens.
type TokenWhereBy struct {
	List []Token
}

func (t *TokenWhereBy) Init(list []Token) Token { t.List = list; return t }
func (t *TokenWhereBy) Type() TokenType         { return T_WHEREBY }
func (t *TokenWhereBy) Str() string             { return "WHERE BY " + join(t.List, ",") }

// TokenGroupBy - GROUP BY token.
//
// @List: List of tokens.
type TokenGroupBy struct {
	List []Token
}

func (t *TokenGroupBy) Init(list []Token) Token { t.List = list; return t }
func (t *TokenGroupBy) Type() TokenType         { return T_GROUPBY }
func (t *TokenGroupBy) Str() string             { return "GROUP BY " + join(t.List, ",") }

// TokenOrder - ORDER token.
//
// @List: List of tokens.
type TokenOrder struct {
	Token Token
	Order TokenType
}

func (t *TokenOrder) Init(tkn Token, odr TokenType) Token { t.Token, t.Order = tkn, odr; return t }
func (t *TokenOrder) Type() TokenType                     { return T_ORDER }
func (t *TokenOrder) Str() string                         { return t.Token.Str() + " " + t.Order.Name() }

// TokenOrderBy - ORDER BY token.
//
// @List: List of tokens.
type TokenOrderBy struct {
	List []Token
}

func (t *TokenOrderBy) Init(list []Token) Token { t.List = list; return t }
func (t *TokenOrderBy) Type() TokenType         { return T_ORDERBY }
func (t *TokenOrderBy) Str() string             { return "ORDER BY " + join(t.List, ",") }

// TokenLimit - LIMIT token.
//
// @List: List of tokens.
type TokenLimit struct {
	List []Token
}

func (t *TokenLimit) Init(list []Token) Token { t.List = list; return t }
func (t *TokenLimit) Type() TokenType         { return T_LIMIT }
func (t *TokenLimit) Str() string             { return "LIMIT " + join(t.List, ",") }

// TokenHaving - HAVING token.
//
// @List: List of tokens.
type TokenHaving struct {
	Token Token
}

func (t *TokenHaving) Init(token Token) Token { t.Token = token; return t }
func (t *TokenHaving) Type() TokenType        { return T_HAVING }
func (t *TokenHaving) Str() string            { return "HAVING " + t.Token.Str() }

// TokenStar - STAR token.
//
// @List: List of tokens.
type TokenStar struct{}

func (t *TokenStar) Init() Token     { return t }
func (t *TokenStar) Type() TokenType { return T_STAR }
func (t *TokenStar) Str() string     { return "*" }

// TokenStar - STAR token.
//
// @List: List of tokens.
type TokenAs struct {
	Token Token
	Name  string
}

func (t *TokenAs) Init(tkn, n Token) Token { t.Token, t.Name = tkn, n.Str(); return t }
func (t *TokenAs) Type() TokenType         { return T_AS }
func (t *TokenAs) Str() string             { return t.Token.Str() + " AS " + t.Name }

// TokenUnique - UNIQUE token.
//
// @List: List of tokens.
type TokenUnique struct {
	Token Token
}

func (t *TokenUnique) Init(token Token) Token { t.Token = token; return t }
func (t *TokenUnique) Type() TokenType        { return T_UNIQUE }
func (t *TokenUnique) Str() string            { return "UNIQUE " + t.Token.Str() }

// TokenNumUnit - NUMUNIT token.
//
// @List: List of tokens.
type TokenNumUnit struct {
	Num, Unit Token
}

func (t *TokenNumUnit) Init(n, u Token) Token { t.Num, t.Unit = n, u; return t }
func (t *TokenNumUnit) Type() TokenType       { return T_NUMUNIT }
func (t *TokenNumUnit) Str() string           { return t.Num.Str() + " " + t.Unit.Str() }

// join() - Join a list of tokens as string.
//
// @list: List of tokens.
// @sp:   Spliter string.
func join(list []Token, sp string) string {
	if len(list) == 0 {
		return ""
	}

	value := list[0].Str()
	for _, item := range list[1:] {
		value += sp + " " + item.Str()
	}

	return value
}

// boolStr() - Returns string value of a bool.
//
// @v: Bool value to test.
func boolStr(v bool) string {
	if v {
		return "true"
	}
	return "false"
}
