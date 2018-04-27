// File: oper.go
//
// This file implements the operation between two interface value.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package value

import (
	"errors"
	"fmt"
)

// AddInt() - Return result of 'v + INT'.
//
// @v1: Interface value.
// @v2: Value to add with.
func AddInt(v1 interface{}, v2 int64) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		return v1.(int64) + v2, nil
	case float64:
		return v1.(float64) + float64(v2), nil
	case string:
		return fmt.Sprintf("%s%d", v1.(string), v2), nil
	}

	msg := fmt.Sprintf("%s + %s not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// AddFloat() - Return result of 'v + FLOAT'.
//
// @v1: Interface value.
// @v2: Value to add with.
func AddFloat(v1 interface{}, v2 float64) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		return float64(v1.(int64)) + v2, nil
	case float64:
		return v1.(float64) + v2, nil
	case string:
		return fmt.Sprintf("%s%f", v1.(string), v2), nil
	}

	msg := fmt.Sprintf("'%s + %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// AddStr() - Return result of 'v + STR'.
//
// @v1: Interface value.
// @v2: Value to add with.
func AddStr(v1 interface{}, v2 string) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		return fmt.Sprintf("%d%s", v1.(int64), v2), nil
	case float64:
		return fmt.Sprintf("%f%s", v1.(float64), v2), nil
	case string:
		return v1.(string) + v2, nil
	}

	msg := fmt.Sprintf("'%s + %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// AddBool() - Return result of 'v + BOOL'.
//
// @v1: Interface value.
// @v2: Value to add with.
func AddBool(v1 interface{}, v2 bool) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		if v2 {
			return v1.(int64) + 1, nil
		} else {
			return v1, nil
		}
	case float64:
		if v2 {
			return v1.(float64) + 1, nil
		} else {
			return v1, nil
		}
	case string:
		return v1.(string) + ToStr(v2), nil
	}

	msg := fmt.Sprintf("'%s + %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// Add() - Return result of 'v + v'.
//
// @v1: Interface value.
// @v2: Value to add with.
func Add(v1 interface{}, v2 interface{}) (interface{}, error) {
	if v2 == nil {
		return v1, nil
	}

	switch v2.(type) {
	case int64:
		return AddInt(v1, v2.(int64))
	case float64:
		return AddFloat(v1, v2.(float64))
	case string:
		return AddStr(v1, v2.(string))
	}

	msg := fmt.Sprintf("'%s + %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// SubInt() - Return result of 'v - INT'.
//
// @v1: Interface value.
// @v2: Value to sub with.
func SubInt(v1 interface{}, v2 int64) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		return v1.(int64) - v2, nil
	case float64:
		return v1.(float64) - float64(v2), nil
	}

	msg := fmt.Sprintf("'%s - %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// SubFloat() - Return result of 'v - FLOAT'.
//
// @v1: Interface value.
// @v2: Value to sub with.
func SubFloat(v1 interface{}, v2 float64) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		return float64(v1.(int64)) - v2, nil
	case float64:
		return v1.(float64) - v2, nil
	}

	msg := fmt.Sprintf("'%s - %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// SubStr() - Return result of 'v - STR'.
//
// @v1: Interface value.
// @v2: Value to sub with.
func SubStr(v1 interface{}, v2 string) (interface{}, error) {
	msg := fmt.Sprintf("'%s - %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// SubBool() - Return result of 'v - BOOL'.
//
// @v1: Interface value.
// @v2: Value to sub with.
func SubBool(v1 interface{}, v2 bool) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		if v2 {
			return v1.(int64) - 1, nil
		} else {
			return v1, nil
		}
	case float64:
		if v2 {
			return v1.(float64) - 1, nil
		} else {
			return v1, nil
		}
	}

	msg := fmt.Sprintf("'%s - %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// Sub() - Return result of 'v - v'.
//
// @v1: Interface value.
// @v2: Value to sub with.
func Sub(v1 interface{}, v2 interface{}) (interface{}, error) {
	if v2 == nil {
		return v1, nil
	}

	switch v2.(type) {
	case int64:
		return SubInt(v1, v2.(int64))
	case float64:
		return SubFloat(v1, v2.(float64))
	case string:
		return SubStr(v1, v2.(string))
	}

	msg := fmt.Sprintf("'%s - %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// MulInt() - Return result of 'v * INT'.
//
// @v1: Interface value.
// @v2: Value to mul with.
func MulInt(v1 interface{}, v2 int64) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		return v1.(int64) * v2, nil
	case float64:
		return v1.(float64) * float64(v2), nil
	}

	msg := fmt.Sprintf("'%s * %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// MulFloat() - Return result of 'v * FLOAT'.
//
// @v1: Interface value.
// @v2: Value to mul with.
func MulFloat(v1 interface{}, v2 float64) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		return float64(v1.(int64)) * v2, nil
	case float64:
		return v1.(float64) * v2, nil
	}

	msg := fmt.Sprintf("'%s * %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// MulStr() - Return result of 'v * STR'.
//
// @v1: Interface value.
// @v2: Value to mul with.
func MulStr(v1 interface{}, v2 string) (interface{}, error) {
	msg := fmt.Sprintf("'%s * %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// MulBool() - Return result of 'v * BOOL'.
//
// @v1: Interface value.
// @v2: Value to mul with.
func MulBool(v1 interface{}, v2 bool) (interface{}, error) {
	if v1 == nil {
		return v2, nil
	}

	switch v1.(type) {
	case int64:
		if v2 {
			return v1.(int64) + 1, nil
		} else {
			return v1, nil
		}
	case float64:
		if v2 {
			return v1.(float64) + 1, nil
		} else {
			return v1, nil
		}
	}

	msg := fmt.Sprintf("'%s * %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// Mul() - Return result of 'v * v'.
//
// @v1: Interface value.
// @v2: Value to mul with.
func Mul(v1 interface{}, v2 interface{}) (interface{}, error) {
	if v2 == nil {
		return v1, nil
	}

	switch v2.(type) {
	case int64:
		return MulInt(v1, v2.(int64))
	case float64:
		return MulFloat(v1, v2.(float64))
	case string:
		return MulStr(v1, v2.(string))
	}

	msg := fmt.Sprintf("'%s * %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// DivInt() - Return result of 'v / INT'.
//
// @v1: Interface value.
// @v2: Value to div with.
func DivInt(v1 interface{}, v2 int64) (interface{}, error) {
	if v1 == nil {
		return nil, nil
	}

	if v2 == 0 {
		return v1, nil
	}

	switch v1.(type) {
	case int64:
		return float64(v1.(int64)) / float64(v2), nil
	case float64:
		return v1.(float64) / float64(v2), nil
	}

	msg := fmt.Sprintf("'%s / %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// DivFloat() - Return result of 'v / FLOAT'.
//
// @v1: Interface value.
// @v2: Value to div with.
func DivFloat(v1 interface{}, v2 float64) (interface{}, error) {
	if v1 == nil {
		return nil, nil
	}

	if v2 == 0.0 {
		return v1, nil
	}

	switch v1.(type) {
	case int64:
		return float64(v1.(int64)) / v2, nil
	case float64:
		return v1.(float64) / v2, nil
	}

	msg := fmt.Sprintf("'%s / %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// DivStr() - Return result of 'v / STR'.
//
// @v1: Interface value.
// @v2: Value to div with.
func DivStr(v1 interface{}, v2 string) (interface{}, error) {
	msg := fmt.Sprintf("'%s / %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// DivBool() - Return result of 'v / BOOL'.
//
// @v1: Interface value.
// @v2: Value to div with.
func DivBool(v1 interface{}, v2 bool) (interface{}, error) {
	if v1 == nil {
		return nil, nil
	}

	switch v1.(type) {
	case int64, float64:
		return v1, nil
	}

	msg := fmt.Sprintf("'%s / %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// Div() - Return result of 'v / v'.
//
// @v1: Interface value.
// @v2: Value to div with.
func Div(v1 interface{}, v2 interface{}) (interface{}, error) {
	if v2 == nil {
		return v1, nil
	}

	switch v2.(type) {
	case int64:
		return DivInt(v1, v2.(int64))
	case float64:
		return DivFloat(v1, v2.(float64))
	case string:
		return DivStr(v1, v2.(string))
	}

	msg := fmt.Sprintf("'%s / %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// ModInt() - Return result of 'v % INT'.
//
// @v1: Interface value.
// @v2: Value to mod with.
func ModInt(v1 interface{}, v2 int64) (interface{}, error) {
	if v1 == nil {
		return nil, nil
	}

	switch v1.(type) {
	case int64:
		return v1.(int64) % v2, nil
	}

	msg := fmt.Sprintf("'%s % %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// ModFloat() - Return result of 'v % FLOAT'.
//
// @v1: Interface value.
// @v2: Value to mod with.
func ModFloat(v1 interface{}, v2 float64) (interface{}, error) {
	msg := fmt.Sprintf("'%s % %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// ModStr() - Return result of 'v % STR'.
//
// @v1: Interface value.
// @v2: Value to mod with.
func ModStr(v1 interface{}, v2 string) (interface{}, error) {
	msg := fmt.Sprintf("'%s % %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// ModBool() - Return result of 'v % BOOL'.
//
// @v1: Interface value.
// @v2: Value to mod with.
func ModBool(v1 interface{}, v2 bool) (interface{}, error) {
	if v1 == nil {
		return nil, nil
	}

	switch v1.(type) {
	case int64:
		if v2 {
			return v1.(int64) % 1, nil
		} else {
			return v1, nil
		}
	}

	msg := fmt.Sprintf("'%s % %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// Mod() - Return result of 'v % v'.
//
// @v1: Interface value.
// @v2: Value to mod with.
func Mod(v1 interface{}, v2 interface{}) (interface{}, error) {
	if v2 == nil {
		return v1, nil
	}

	switch v2.(type) {
	case int64:
		return ModInt(v1, v2.(int64))
	case float64:
		return ModFloat(v1, v2.(float64))
	case string:
		return ModStr(v1, v2.(string))
	}

	msg := fmt.Sprintf("'%s % %s' not support", TypeStr(v1), TypeStr(v2))
	return nil, errors.New(msg)
}

// Operate() - Return result of 'v opt v'.
//
// @v1: Interface value.
// @v2: Value to operate with.
func Operate(v1 interface{}, opt string, v2 interface{}) (interface{}, error) {
	switch opt {
	case "+":
		return Add(v1, v2)
	case "-":
		return Sub(v1, v2)
	case "*":
		return Mul(v1, v2)
	case "/":
		return Div(v1, v2)
	case "%":
		return Mod(v1, v2)
	}

	msg := fmt.Sprintf("unknown operator '%s'", opt)
	return nil, errors.New(msg)
}
