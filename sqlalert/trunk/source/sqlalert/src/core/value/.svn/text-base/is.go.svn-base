// File: is.go
//
// This file implements the checking functions of value type.
//
// Copyright (c) 2017 Yun Li Lai, Ltd, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>
package value

// IsXxx() - Type checking functions.
//
// @v: Value to test.
func IsInt(v interface{}) bool      { return Type(v) == INT }
func IsFloat(v interface{}) bool    { return Type(v) == FLOAT }
func IsBool(v interface{}) bool     { return Type(v) == BOOL }
func IsStr(v interface{}) bool      { return Type(v) == STR }
func IsList(v interface{}) bool     { return Type(v) == LIST }
func IsDict(v interface{}) bool     { return Type(v) == DICT }
func IsNum(v interface{}) bool      { return IsInt(v) || IsFloat(v) }
func IsIterable(v interface{}) bool { return IsList(v) || IsDict(v) }

// IsTrue()/IsFalse() - Test whether a value is TRUE/FALSE.
//
// @v: Value to test.
//
// This function test the real value of the given value. For example,
// it return false if a given value is INT and its value is 0.
func IsTrue(v interface{}) bool { return !IsFalse(v) }
func IsFalse(v interface{}) bool {
	switch {
	case IsInt(v):
		return Int(v) == Int(0)
	case IsFloat(v):
		return Float(v) == Float(0)
	case IsStr(v):
		return Str(v) == "false" || Str(v) == ""
	case IsBool(v):
		return !Bool(v)
	case IsList(v) || IsDict(v):
		return Len(v) == 0
	default:
		return v == nil
	}
}
