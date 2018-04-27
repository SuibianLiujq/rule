// File: fields.go
//
// This file implements SQL fileds.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package dsl

import (
	"core/script"
	"core/sql"
	"errors"
	"fmt"
)

// Filed - Structure of fields.
//
// @Name:  Field name.
// @Token: SQL token.
type Field struct {
	Name  string
	Token sql.Token
}

// Init() - Initialize the Field instance.
//
// @name:  Field name.
// @token: SQL token.
func (this *Field) Init(name string, token sql.Token) *Field {
	this.Name, this.Token = name, token
	return this
}

// NewField() - Create filed with SQL token.
//
// @token: SQL token.
// @ctx:   Script context.
func NewField(token sql.Token, ctx *script.Cntx) (*Field, error) {
	switch token.Type() {
	case sql.T_IDENT, sql.T_STR, sql.T_STAR:
		return (&Field{}).Init(token.Str(), token), nil

	case sql.T_AS:
		t := token.(*sql.TokenAs)
		return (&Field{}).Init(t.Name, t.Token), nil
	}

	msg := fmt.Sprintf("'%s' needs an alias", token.Str())
	return nil, errors.New(msg)
}

type Order struct {
	Name, Order string
	Value       interface{}
	Metric      *Metric
}

func NewOrder(n, o string, v interface{}, m *Metric) *Order {
	return &Order{Name: n, Order: o, Value: v, Metric: m}
}
