// File: scripts.go
//
// This file implements compilter DSL scripts.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package dsl

import (
	"core/script"
	"core/sql"
	"core/value"
	"errors"
	"fmt"
	"strings"
	"sync"
)

// Script type.
type ScriptType int

// All supported script type.
//
// @SCRIPT_FIELD:    A field doc['name'].value string.
// @SCRIPT_VALUE:    A value.
// @SCRIPT_S_SCRIPT: A single value script like doc['name'].value.length().
// @SCRIPT_M_SCRIPT: A multi-value (ecpression) string.
const (
	SCRIPT_FIELD ScriptType = iota
	SCRIPT_VALUE
	SCRIPT_S_SCRIPT
	SCRIPT_M_SCRIPT
)

// All supported script operators.
var g_script_oprators = map[sql.TokenType]string{
	sql.T_ADD: " + ",
	sql.T_SUB: " - ",
	sql.T_MUL: " * ",
	sql.T_DIV: " / ",
	sql.T_MOD: " % ",
	sql.T_EQ:  " == ",
	sql.T_NE:  " != ",
	sql.T_LT:  " < ",
	sql.T_LE:  " <= ",
	sql.T_GT:  " > ",
	sql.T_GE:  " >= ",
	sql.T_AND: " && ",
	sql.T_OR:  " || ",
	sql.T_NOT: " ! ",
}

// ScriptValue - Structure of ScriptValue.
//
// @Type:  Script type.
// @Value: Script value.
type ScriptValue struct {
	Type  ScriptType
	Value interface{}
}

// Init() - Initialize ScriptValue.
//
// @t: Script type.
// @v: Script value.
//
// This function return this ScriptValue instance itself for chain operation.
func (this *ScriptValue) Init(t ScriptType, v interface{}) (*ScriptValue, error) {
	this.Type, this.Value = t, v
	return this, nil
}

// String() - Returns string value of script value.
func (this *ScriptValue) Str() string {
	switch this.Value.(type) {
	case string:
		return this.Value.(string)
	}

	return value.ToStr(this.Value)
}

// ScriptString() - Returns script formatted string.
func (this *ScriptValue) ScriptStr() string {
	switch this.Type {
	case SCRIPT_FIELD:
		//		return "_source." + this.Str()
		return "doc['" + this.Str() + "'].value"

	case SCRIPT_VALUE:
		if value.IsStr(this.Value) {
			return "'" + this.Str() + "'"
		}
	}

	return this.Str()
}

// AsField() - Returns script value as ES script field.
func (this *ScriptValue) AsField(lang string) map[string]interface{} {
	return map[string]interface{}{"script": this.AsFieldPure(lang)}
}

// AsFieldPure() - Returns script value as ES script field.
//                 This function reutrns the pure script value without 'script' key.
func (this *ScriptValue) AsFieldPure(lang string) map[string]interface{} {
	return map[string]interface{}{"inline": this.ScriptStr(), "lang": lang}
}

// Operate() - Return script value of 'this operator other'.
//
// @opt:   Operator type (SQL token type).
// @other: Another script value.
func (this *ScriptValue) Operate(opt sql.TokenType, other *ScriptValue) (*ScriptValue, error) {
	if this.Type == SCRIPT_VALUE && other.Type == SCRIPT_VALUE {
		if this.Type != other.Type {
			msg := fmt.Sprintf("invalid operator '%s'", opt.Name())
			return nil, errors.New(msg)
		}

		v, err := value.Operate(this.Value, opt.Name(), other.Value)
		if err != nil {
			return nil, err
		}

		return (&ScriptValue{}).Init(SCRIPT_VALUE, v)
	}

	return this.ConcateOpteration(opt, other)
}

// Compare() - Return script value of 'this compare other'.
//
// @opt:   Operator type (SQL token type).
// @other: Another script value.
func (this *ScriptValue) Compare(opt sql.TokenType, other *ScriptValue) (*ScriptValue, error) {
	if this.Type == SCRIPT_VALUE && other.Type == SCRIPT_VALUE {
		if this.Type != other.Type {
			msg := fmt.Sprintf("invalid operator '%s'", opt.Name())
			return nil, errors.New(msg)
		}

		v := value.Compare(this.Value, opt.Name(), other.Value)
		return (&ScriptValue{}).Init(SCRIPT_VALUE, v)
	}

	return this.ConcateOpteration(opt, other)
}

// ConcateOpterator() - Concate another ScriptValue with given opterator.
//
// @opt:   Operator type (SQL token type).
// @other: Another script value.
func (this *ScriptValue) ConcateOpteration(opt sql.TokenType, other *ScriptValue) (*ScriptValue, error) {
	optStr, ok := g_script_oprators[opt]
	if !ok {
		msg := fmt.Sprintf("invalid operator '%s'", opt.Name())
		return nil, errors.New(msg)
	}

	var v string
	if this.Type == SCRIPT_M_SCRIPT {
		v = "(" + this.ScriptStr() + ")"
	} else {
		v = this.ScriptStr()
	}

	v += optStr
	if other.Type == SCRIPT_M_SCRIPT {
		v += "(" + other.ScriptStr() + ")"
	} else {
		v += other.ScriptStr()
	}

	return (&ScriptValue{}).Init(SCRIPT_M_SCRIPT, v)
}

// scriptOperation() - Build SQL OPERATION token as scripts.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func scriptOperation(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	t := token.(*sql.TokenOper)

	left, err := ScriptBuild(dsl, t.Left, ctx)
	if err != nil {
		return nil, err
	}

	right, err := ScriptBuild(dsl, t.Right, ctx)
	if err != nil {
		return nil, err
	}

	return left.Operate(t.Operator, right)
}

// scriptComparison() - Build SQL COMPARISON token as scripts.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func scriptComparison(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	t := token.(*sql.TokenComp)

	left, err := ScriptBuild(dsl, t.Left, ctx)
	if err != nil {
		return nil, err
	}

	right, err := ScriptBuild(dsl, t.Right, ctx)
	if err != nil {
		return nil, err
	}

	return left.Compare(t.Operator, right)
}

// scriptLogical() - Build SQL LOGICAL token as scripts.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func scriptLogical(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	t := token.(*sql.TokenLogical)

	var left *ScriptValue = nil
	if t.Left != nil {
		sv, err := ScriptBuild(dsl, t.Left, ctx)
		if err != nil {
			return nil, err
		} else {
			left = sv
		}
	}

	var right *ScriptValue = nil
	if t.Right != nil {
		sv, err := ScriptBuild(dsl, t.Right, ctx)
		if err != nil {
			return nil, err
		} else {
			right = sv
		}
	}

	switch t.Operator {
	case sql.T_AND:
		if left.Type == SCRIPT_VALUE && right.Type == SCRIPT_VALUE {
			v := (value.IsTrue(left.Value) && value.IsTrue(right.Value))
			return (&ScriptValue{}).Init(SCRIPT_VALUE, v)
		}

		return left.ConcateOpteration(t.Operator, right)

	case sql.T_OR:
		if left.Type == SCRIPT_VALUE && right.Type == SCRIPT_VALUE {
			v := (value.IsTrue(left.Value) || value.IsTrue(right.Value))
			return (&ScriptValue{}).Init(SCRIPT_VALUE, v)
		}

		return left.ConcateOpteration(t.Operator, right)

	case sql.T_NOT:
		if right.Type == SCRIPT_VALUE {
			v := value.IsFalse(right.Value)
			return (&ScriptValue{}).Init(SCRIPT_VALUE, v)
		}

		expStr := g_script_oprators[t.Operator]
		if right.Type == SCRIPT_M_SCRIPT {
			expStr += "(" + right.ScriptStr() + ")"
		} else {
			expStr += right.ScriptStr()
		}

		return (&ScriptValue{}).Init(SCRIPT_S_SCRIPT, expStr)
	}

	msg := fmt.Sprintf("invalid operator '%s' in '%s'", t.Operator.Name(), token.Str())
	return nil, errors.New(msg)
}

// scriptCondition() - Build SQL CONDITION token as scripts.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func scriptCondition(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	t := token.(*sql.TokenCond)

	left, err := ScriptBuild(dsl, t.Left, ctx)
	if err != nil {
		return nil, err
	}

	middle, err := ScriptBuild(dsl, t.Middle, ctx)
	if err != nil {
		return nil, err
	}

	right, err := ScriptBuild(dsl, t.Right, ctx)
	if err != nil {
		return nil, err
	}

	if left.Type == SCRIPT_VALUE {
		if value.IsTrue(left.Value) {
			return middle, nil
		} else {
			return right, nil
		}
	}

	var v string
	if left.Type == SCRIPT_M_SCRIPT {
		v = "(" + left.ScriptStr() + ")"
	} else {
		v = left.ScriptStr()
	}
	v += " ? "

	if middle.Type == SCRIPT_M_SCRIPT {
		v += "(" + middle.ScriptStr() + ")"
	} else {
		v += middle.ScriptStr()
	}
	v += " : "

	if right.Type == SCRIPT_M_SCRIPT {
		v += "(" + right.ScriptStr() + ")"
	} else {
		v += right.ScriptStr()
	}

	return (&ScriptValue{}).Init(SCRIPT_M_SCRIPT, v)
}

// scriptFuncLen() - Build SQL token 'len(field)' ES script.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func scriptFuncLen(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("len() argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	sv, err := ScriptBuild(dsl, args[0], ctx)
	if err != nil {
		msg := fmt.Sprintf("%s' in script len()", err)
		return nil, errors.New(msg)
	}

	switch sv.Type {
	case SCRIPT_FIELD:
		return (&ScriptValue{}).Init(SCRIPT_S_SCRIPT, sv.ScriptStr()+".length()")

	case SCRIPT_VALUE:
		if str, ok := sv.Value.(string); ok {
			return (&ScriptValue{}).Init(SCRIPT_VALUE, int64(len(str)))
		}
	}

	msg := fmt.Sprintf("invalid argument '%s' in len()", args[0].Str())
	return nil, errors.New(msg)
}

// scriptFuncScript() - Build SQL token 'script(str)' ES script.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func scriptFuncScript(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("len() argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	switch args[0].Type() {
	case sql.T_STR, sql.T_VAR:
		sv, err := ScriptBuild(dsl, args[0], ctx)
		if err != nil {
			msg := fmt.Sprintf("%s' in script len()", err)
			return nil, errors.New(msg)
		}

		return (&ScriptValue{}).Init(SCRIPT_M_SCRIPT, sv.Str())
	}

	msg := fmt.Sprintf("'%s' not STR or VAR in script()", args[0].Str())
	return nil, errors.New(msg)
}

// Script functions and it's initialization function.
var g_script_functions_once sync.Once
var g_script_functions map[string]func(*Dsl, []sql.Token, *script.Cntx) (*ScriptValue, error)

func init_script_functions() {
	g_script_functions = map[string]func(*Dsl, []sql.Token, *script.Cntx) (*ScriptValue, error){
		"len":    scriptFuncLen,
		"script": scriptFuncScript,
	}
}

// scriptFunction() - Build SQL token as ES script.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func scriptFunction(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	g_script_functions_once.Do(init_script_functions)

	t := token.(*sql.TokenFunc)
	name := strings.ToLower(t.Name)

	if fn, ok := g_script_functions[name]; ok {
		return fn(dsl, t.List, ctx)
	}

	msg := fmt.Sprintf("script function '%s' not support", name)
	return nil, errors.New(msg)
}

// ScriptBuild() - Build SQL tokens as scripts.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func ScriptBuild(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*ScriptValue, error) {
	switch token.Type() {
	case sql.T_VAR:
		varValue := ctx.Get(token.(*sql.TokenVar).Value)
		if varValue == nil {
			msg := fmt.Sprintf("'%s' not found or empty", token.Str())
			return nil, errors.New(msg)
		}
		return (&ScriptValue{}).Init(SCRIPT_VALUE, varValue)
	case sql.T_IDENT:
		return (&ScriptValue{}).Init(SCRIPT_FIELD, token.Str())
	case sql.T_STR:
		return (&ScriptValue{}).Init(SCRIPT_VALUE, token.(*sql.TokenStr).Value)
	case sql.T_INT:
		return (&ScriptValue{}).Init(SCRIPT_VALUE, token.(*sql.TokenInt).Value)
	case sql.T_FLOAT:
		return (&ScriptValue{}).Init(SCRIPT_VALUE, token.(*sql.TokenFloat).Value)
	case sql.T_BOOL:
		return (&ScriptValue{}).Init(SCRIPT_VALUE, token.(*sql.TokenBool).Value)
	case sql.T_NULL:
		return (&ScriptValue{}).Init(SCRIPT_VALUE, token.Str())
	case sql.T_OPER:
		return scriptOperation(dsl, token, ctx)
	case sql.T_COMP:
		return scriptComparison(dsl, token, ctx)
	case sql.T_COND:
		return scriptCondition(dsl, token, ctx)
	case sql.T_FUNC:
		return scriptFunction(dsl, token, ctx)
	case sql.T_LOGICAL:
		return scriptLogical(dsl, token, ctx)
	}

	msg := fmt.Sprintf("unsupport script '%s'", token.Str())
	return nil, errors.New(msg)
}
