// File: helper.go
//
// This file implements helper functions to use SCRIPT parser.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package script

import (
	"core/sys"
	"errors"
)

// Parse() - Parse the @src bytes.
//
// @src: The source byte stream.
//
// This function parse the input byte stream and returns the
// parsed token or error.
func Parse(src []byte, name string) (token Token, err error) {
	lexer, err := (&Lexer{}).Init(src, name)
	if err != nil {
		return nil, err
	}

	if yyParse(lexer) != 0 {
		msg := lexer.errMsg
		if lexer.scanner.GotError {
			msg = lexer.scanner.Error
		}

		return nil, errors.New(msg)
	}

	return lexer.token, nil
}

// ParseFile() - Read file and parse file content into tokens.
//
// @name: File name.
//
// This function read @name file's content and parse it into
// script tokens. It returns the top-level token or error.
func ParseFile(name string) (token Token, err error) {
	content, err := sys.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return Parse(content, name)
}

// ExecScript() - Execute the script text.
//
// @text: Script text.
// @ctx:  Script execution context.
func ExecScript(text []byte, ctx *Cntx) (interface{}, error) {
	token, err := Parse(text, "")
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		ctx = NewContext()
	}

	return Exec(token, ctx)
}

// ExecFile() - Execute the script file.
//
// @text: Script text.
// @ctx:  Script execution context.
func ExecFile(name string, ctx *Cntx) (interface{}, error) {
	token, err := ParseFile(name)
	if err != nil {
		return nil, err
	}

	if ctx == nil {
		ctx = NewContext()
	}

	return Exec(token, ctx)
}
