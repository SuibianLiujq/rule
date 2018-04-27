// File: token.go
//
// This file defines SCRIPT tokens.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package script

import (
	"fmt"
)

// TokenType - Type of SCRIPT token.
type TokenType int

// All SCRIPT tokens.
const (
	T_ILL TokenType = iota

	T_IDENT
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

	T_INC
	T_DEC

	T_FUNC
	T_COND
	T_OPER
	T_COMP
	T_LOGICAL
	T_INDEX
	T_IN

	T_STMTS
	T_EXPR
	T_ASSIGN
	T_IF
	T_IF_ELSE
	T_ELSEIF
	T_ELSE
	T_FOR
	T_FOR_ITER
	T_FORIN
	T_FORIN_ITER
	T_CONTINUE
	T_BREAK
	T_RETURN
	T_DEFINE
	T_INCLUDE
	T_IMPORT
)

// Names of all SCRIPT tokens.
var g_token_names = [...]string{
	T_ILL: "ILLEGAL",

	T_IDENT: "IDENT",
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

	T_INC: "++",
	T_DEC: "--",

	T_FUNC:    "FUNC",
	T_COND:    "COND",
	T_OPER:    "OPER",
	T_COMP:    "COMP",
	T_LOGICAL: "LOGICAL",
	T_INDEX:   "INDEX",
	T_IN:      "IN",

	T_STMTS:      "STMTS",
	T_EXPR:       "EXPR",
	T_ASSIGN:     "ASSIGN",
	T_IF:         "IF",
	T_IF_ELSE:    "IF_ELSE",
	T_ELSEIF:     "ELSEIF",
	T_ELSE:       "ELSE",
	T_FOR:        "FOR",
	T_FOR_ITER:   "FOR_ITER",
	T_FORIN:      "FORIN",
	T_FORIN_ITER: "FORIN_ITER",
	T_CONTINUE:   "CONTINUE",
	T_BREAK:      "BREAK",
	T_RETURN:     "RETURN",
	T_DEFINE:     "DEFINE",
	T_INCLUDE:    "INCLUDE",
	T_IMPORT:     "IMPORT",
}

// Name() - Returns name of token type.
func (this TokenType) Name() string {
	if 0 <= this && this <= TokenType(len(g_token_names)) {
		return g_token_names[this]
	}

	return fmt.Sprintf("unknown (%d)", this)
}
