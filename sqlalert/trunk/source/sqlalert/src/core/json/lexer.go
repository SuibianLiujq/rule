// File: lexer.go
//
// This file implements the JSON lexer.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package json

import (
	"core/scanner"
	"fmt"
	"strings"
)

// All JSON keywords.
var g_keywords = map[string]struct {
	ytype int
	value interface{}
}{
	"true":  {Y_BOOL, true},
	"false": {Y_BOOL, false},
	"null":  {Y_NULL, nil},
}

// Lexer - JSON lexer.
type Lexer struct {
	scanner *scanner.Scanner
	value   interface{}
	errMsg  string
}

// Init() - Initialize the Lexer instance.
//
// @src: Input byte stream.
func (this *Lexer) Init(src []byte, name string) (*Lexer, error) {
	s, err := scanner.New(src, name)
	if err != nil {
		return nil, err
	}

	this.scanner, this.value, this.errMsg = s, nil, ""
	return this, nil
}

// Lex() - Scan input bytes and return an token.
//
// @lval: Argument of lexer (given by goyacc).
func (this *Lexer) Lex(lval *yySymType) int {
	if this.scanner.GotEOF {
		return 0
	}

	token := Y_ERR
	if !this.scanner.GotError {
		switch t, v := this.scanner.Scan(); t {
		case scanner.T_EOF:
			token = 0
		case scanner.T_ILL:
			if this.scanner.GotError {
				this.Error(this.scanner.Error)
			}
		default:
			token = this.translate(t, v, lval)
		}
	}

	return token
}

// Error() - Error handler.
//
// @msg: Error message.
//
// This function handle the goyacc errors.
func (this *Lexer) Error(msg string) {
	msg = fmt.Sprintf("%s around line %d", msg, this.scanner.LineNum)
	this.errMsg = msg
}

// translate() - Translate scanner tokens to JSON yacc tokens.
//
// @t:    Scanner token type.
// @v:    Scanner token value.
// @lval: Argument of lexer (given by goyacc).
func (this *Lexer) translate(t int, v interface{}, lval *yySymType) int {
	switch t {
	case scanner.T_IDENT:
		key := strings.ToLower(v.(string))
		if item, ok := g_keywords[key]; ok {
			lval.value = item.value
			return item.ytype
		}
		return Y_ERR

	case scanner.T_STR:
		lval.value = v
		return Y_STR

	case scanner.T_INT:
		lval.value = v
		return Y_INT

	case scanner.T_FLOAT:
		lval.value = v
		return Y_FLOAT

	case '-', ',', ':', '[', ']', '{', '}':
		lval.value = v
		return t
	}

	lval.value = nil
	return Y_ERR
}

// setValue() - Set JSON object to Lexer.
//
// @lexer: Instance of Lexer.
// @v:     JSON object value.
func setValue(lexer interface{}, v interface{}) {
	lexer.(*Lexer).value = v
}
