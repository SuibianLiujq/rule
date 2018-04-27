// File: as.go
//
// This file implements the type asserting functions.
//
// Copyright (c) 2017 Yun Li Lai, Ltd, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>
package value

// AsInt() - Returns the value as INT.
//
// @v: Value to test.
func AsInt(v interface{}) (int64, bool) {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return Int(v), true
	default:
		return Int(0), false
	}
}

// AsFloat() - Returns the value as FLOAT.
//
// @v: Value to test.
func AsFloat(v interface{}) (float64, bool) {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return Float(v), true
	default:
		return Float(0), false
	}
}

// AsBool() - Returns the value as BOOL.
//
// @v: Value to test.
func AsBool(v interface{}) (bool, bool) {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64, bool:
		return Bool(v), true
	default:
		return false, false
	}
}

// AsStr() - Returns the value as STR.
//
// @v: Value to test.
func AsStr(v interface{}) (string, bool) {
	str, ok := v.(string)
	return str, ok
}

// AsList() - Returns the value as LIST.
//
// @v: Value to test.
func AsList(v interface{}) ([]interface{}, bool) {
	list, ok := v.([]interface{})
	return list, ok
}

// AsDict() - Returns the value as DICT.
//
// @v: Value to test.
func AsDict(v interface{}) (map[string]interface{}, bool) {
	dict, ok := v.(map[string]interface{})
	return dict, ok
}
