// File: stats.go
//
// This file implements the statistic function of SQL tokens.
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

// Stat - Structure of Stat
//
// @SelectXxx: List/Dict of SELECT items.
// @GroupXxx:  List/Dict of GROUP BY items.
// @From:      FROM item.
// @Where:     WHERE item.
// @OrderList: List of ORDER items.
// @LimitList: List of LIMIT items.
// @Having:    HAVING item.
type Stat struct {
	SelectList  []*Field
	SelectDict  map[string]*Field
	GroupList   []*Field
	GroupDict   map[string]*Field
	From        sql.Token
	Where       sql.Token
	WhereByDict map[string]*Field
	OrderList   []sql.Token
	LimitList   []sql.Token
	Having      sql.Token
}

// Init() - Initialize Stat structure.
//
// @token: SQL token.
// @ctx:   Script context.
//
// This function returns the State instance itself for chain operation.
// It calls this.Scan() method to scan SQL tokens if @token is not null.
func (this *Stat) Init(token sql.Token, ctx *script.Cntx) (*Stat, error) {
	this.SelectDict = make(map[string]*Field)
	this.GroupDict = make(map[string]*Field)
	this.WhereByDict = make(map[string]*Field)

	if token != nil {
		return this.Scan(token, ctx)
	}

	return this, nil
}

// Scan() - Scan the SQL tokens.
//
// @token: SQL token.
// @ctx:   Script context.
//
// This function scan the SQL tokens and do some statistics.
func (this *Stat) Scan(token sql.Token, ctx *script.Cntx) (*Stat, error) {
	t := token.(*sql.TokenStmts)

	err := this.scanSelect(t.Select, ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in SELECT", err)
		return nil, errors.New(msg)
	}

	for _, item := range t.List {
		switch item.Type() {
		case sql.T_FROM:
			err = this.scanFrom(item, ctx)
		case sql.T_WHERE:
			err = this.scanWhere(item, ctx)
		case sql.T_WHEREBY:
			err = this.scanWhereBy(item, ctx)
		case sql.T_GROUPBY:
			err = this.scanGroup(item, ctx)
		case sql.T_ORDERBY:
			err = this.scanOrder(item, ctx)
		case sql.T_LIMIT:
			err = this.scanLimit(item, ctx)
		case sql.T_HAVING:
			err = this.scanHaving(item, ctx)
		default:
			msg := fmt.Sprintf("'%s' not SQL statement", item.Str())
			return nil, errors.New(msg)
		}

		if err != nil {
			msg := fmt.Sprintf("%s in %s", item.Type().Name())
			return nil, errors.New(msg)
		}
	}

	return this, nil
}

// scanSelect() - Scan SELECT token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanSelect(token sql.Token, ctx *script.Cntx) error {
	t := token.(*sql.TokenSelect)

	for _, item := range t.List {
		field, err := NewField(item, ctx)
		if err != nil {
			return err
		}

		this.SelectList = append(this.SelectList, field)
		this.SelectDict[field.Name] = field
	}

	return nil
}

// scanFrom() - Scan FROM token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanFrom(token sql.Token, ctx *script.Cntx) error {
	t := token.(*sql.TokenFrom)

	if len(t.List) != 1 {
		msg := fmt.Sprintf("more than one (%d) token", len(t.List))
		return errors.New(msg)
	}

	switch t.List[0].Type() {
	case sql.T_IDENT, sql.T_STR, sql.T_VAR:
		this.From = t.List[0]

	default:
		msg := fmt.Sprintf("invalid token '%s'", t.List[0].Str())
		return errors.New(msg)
	}

	return nil
}

// scanWhere() - Scan WHERE token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanWhere(token sql.Token, ctx *script.Cntx) error {
	this.Where = token.(*sql.TokenWhere).Token
	return nil
}

// scanWhereBy() - Scan WHERE BY token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanWhereBy(token sql.Token, ctx *script.Cntx) error {
	t := token.(*sql.TokenWhereBy)

	for _, item := range t.List {
		field, err := NewField(item, ctx)
		if err != nil {
			return err
		}

		this.WhereByDict[field.Name] = field
	}

	return nil
}

// scanGroup() - Scan GROUP BY token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanGroup(token sql.Token, ctx *script.Cntx) error {
	t := token.(*sql.TokenGroupBy)

	for _, item := range t.List {
		field, err := NewField(item, ctx)
		if err != nil {
			return err
		}

		this.GroupList = append(this.GroupList, field)
		this.GroupDict[field.Name] = field
	}

	return nil
}

// scanOrder() - Scan ORDER BY token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanOrder(token sql.Token, ctx *script.Cntx) error {
	t := token.(*sql.TokenOrderBy)

	for _, item := range t.List {
		if item.Type() != sql.T_ORDER {
			msg := fmt.Sprintf("invalid token '%s'", item.Str())
			return errors.New(msg)
		}
	}

	this.OrderList = t.List
	return nil
}

// scanLimit() - Scan LIMIT token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanLimit(token sql.Token, ctx *script.Cntx) error {
	t := token.(*sql.TokenLimit)

	for _, item := range t.List {
		switch item.Type() {
		case sql.T_INT, sql.T_VAR:
		default:
			msg := fmt.Sprintf("invalid token '%s'", item.Str())
			return errors.New(msg)
		}
	}

	this.LimitList = t.List
	return nil
}

// scanHaving() - Scan HAVING token.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Stat) scanHaving(token sql.Token, ctx *script.Cntx) error {
	this.Having = token.(*sql.TokenHaving).Token
	return nil
}

// StatToken() - Scan SQL tokens and returns Stat instance.
//
// @token: SQL token.
// @ctx:   Script context.
func StatToken(token sql.Token, ctx *script.Cntx) (*Stat, error) {
	return (&Stat{}).Init(token, ctx)
}
