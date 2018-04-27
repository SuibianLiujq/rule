// File: token.go
//
// This file defines SQL tokens.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package sql

import (
	"fmt"
)

// TokenType - Type of SQL token.
type TokenType int

// All SQL tokens.
const (
	T_ILL TokenType = iota

	T_IDENT
	T_VAR
	T_STR
	T_INT
	T_FLOAT
	T_BOOL
	T_NULL
	T_LIST
	T_PAIR
	T_DICT

	T_ADD
	T_SUB
	T_MUL
	T_DIV
	T_MOD

	T_EQ
	T_NE
	T_LT
	T_LE
	T_GT
	T_GE

	T_AND
	T_OR
	T_NOT

	T_FUNC
	T_COND
	T_OPER
	T_COMP
	T_LOGICAL
	T_INDEX
	T_IN

	T_STMTS
	T_SELECT
	T_FROM
	T_WHERE
	T_WHEREBY
	T_GROUPBY
	T_ORDER
	T_ORDERBY
	T_LIMIT
	T_HAVING
	T_AS
	T_STAR
	T_UNIQUE
	T_NUMUNIT
	T_DESC
	T_ASC
)

// Names of all SQL tokens.
var g_token_names = [...]string{
	T_ILL: "ILLEGAL",

	T_IDENT: "IDENT",
	T_VAR:   "VAR",
	T_STR:   "STR",
	T_INT:   "INT",
	T_FLOAT: "FLOAT",
	T_BOOL:  "BOOL",
	T_NULL:  "NULL",
	T_LIST:  "LIST",
	T_PAIR:  "PAIR",
	T_DICT:  "DICT",

	T_ADD: "+",
	T_SUB: "-",
	T_MUL: "*",
	T_DIV: "/",
	T_MOD: "%",

	T_EQ: "==",
	T_NE: "!=",
	T_LT: "<",
	T_LE: "<=",
	T_GT: ">",
	T_GE: ">=",

	T_AND: "&&",
	T_OR:  "||",
	T_NOT: "!",

	T_FUNC:    "FUNC",
	T_COND:    "COND",
	T_OPER:    "OPER",
	T_COMP:    "COMP",
	T_LOGICAL: "LOGICAL",
	T_INDEX:   "INDEX",
	T_IN:      "IN",

	T_STMTS:   "STMTS",
	T_SELECT:  "SELECT",
	T_FROM:    "FROM",
	T_WHERE:   "WHERE",
	T_GROUPBY: "GROUP BY",
	T_WHEREBY: "WHERE BY",
	T_ORDER:   "ORDER",
	T_ORDERBY: "ORDER BY",
	T_LIMIT:   "LIMIT",
	T_HAVING:  "HAVING",
	T_AS:      "AS",
	T_STAR:    "STAR",
	T_UNIQUE:  "UNIQUE",
	T_NUMUNIT: "NUMUNIT",
	T_DESC:    "DESC",
	T_ASC:     "ASC",
}

// Name() - Returns name of token type.
func (this TokenType) Name() string {
	if 0 <= this && this <= TokenType(len(g_token_names)) {
		return g_token_names[this]
	}

	return fmt.Sprintf("unknown (%d)", this)
}
