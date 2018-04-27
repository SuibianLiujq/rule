// File: types.go
//
// This file defines the base type of golang values. Users can check
// value type more convenient.
//
// Copyright (c) 2017 Yun Li Lai, Ltd, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>
package value

import (
	"fmt"
)

// ValueType - Type value of all golang types.
type ValueType int

// All supported types.
const (
	UNKNOWN ValueType = iota
	INT
	FLOAT
	STR
	BOOL
	NULL
	LIST
	DICT
)

// g_type_names - Names of all type.
var g_type_names = [...]string{
	UNKNOWN: "UNKNOWN",
	INT:     "INT",
	FLOAT:   "FLOAT",
	STR:     "STR",
	BOOL:    "BOOL",
	NULL:    "NULL",
	LIST:    "LIST",
	DICT:    "DICT",
}

// Str() - Return name of type.
func (this ValueType) Str() string {
	return g_type_names[this]
}

// IsBasic() - Test whether a type is basic type.
func (this ValueType) IsBasic() bool {
	switch this {
	case INT, FLOAT, STR, BOOL:
		return true
	default:
		return false
	}
}

// TypeStr()/Type() - Returns type STR/VALUE of given value.
//
// @v: Value to test.
//
// This function returns UNKNOWN for all unknown type.
func TypeStr(v interface{}) string { return Type(v).Str() }
func Type(v interface{}) ValueType {
	if v == nil {
		return NULL
	}

	switch v.(type) {
	case int, int8, int16, int32, int64:
		return INT
	case float32, float64:
		return FLOAT
	case string:
		return STR
	case bool:
		return BOOL
	case []interface{}:
		return LIST
	case map[string]interface{}:
		return DICT
	default:
		return UNKNOWN
	}
}

// Int() - Create INT value with given value.
//
// @v: Value to use.
func Int(v interface{}) int64 {
	switch v.(type) {
	case int:
		return int64(v.(int))
	case int8:
		return int64(v.(int8))
	case int16:
		return int64(v.(int16))
	case int32:
		return int64(v.(int32))
	case int64:
		return int64(v.(int64))
	case float32, float64:
		return int64(Float(v))
	case bool:
		if v.(bool) {
			return int64(1)
		} else {
			return int64(0)
		}
	default:
		return int64(0)
	}
}

// Float() - Create FLOAT value with given value.
//
// @v: Value to use.
func Float(v interface{}) float64 {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return float64(Int(v))
	case float32:
		return float64(v.(float32))
	case float64:
		return v.(float64)
	case bool:
		if v.(bool) {
			return float64(1)
		} else {
			return float64(0)
		}
	default:
		return float64(0)
	}
}

// Bool() - Create BOOL value with given value.
//
// @v: Value to use.
func Bool(v interface{}) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return Int(v) == 0
	case float32, float64:
		return Float(v) == 0.0
	case bool:
		return v.(bool)
	case string:
		return v.(string) == "true"
	default:
		return false
	}
}

// Str() - Create STR value with given value.
//
// @v: Value to use.
func Str(v interface{}) string {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", Int(v))
	case float32, float64:
		return fmt.Sprintf("%f", Float(v))
	case string:
		return v.(string)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// List() - Convert a value to LIST.
//
// @v: Value to use.
func List(v interface{}) []interface{} {
	switch v.(type) {
	case []interface{}:
		return v.([]interface{})
	default:
		return []interface{}{}
	}
}

// Dict() - Convert a value to DICT.
//
// @v: Value to use.
func Dict(v interface{}) map[string]interface{} {
	switch v.(type) {
	case map[string]interface{}:
		return v.(map[string]interface{})
	default:
		return nil
	}
}

// Len() - Returns the length of LIST/DICT.
//
// @v: Value to test.
//
// This function returns the length if the given value is LIST/DICT,
// and it returns 0 for all other types of value.
func Len(v interface{}) int64 {
	switch v.(type) {
	case []interface{}:
		return int64(len(List(v)))
	case map[string]interface{}:
		return int64(len(Dict(v)))
	case string:
		return int64(len(Str(v)))
	default:
		return int64(0)
	}
}

// Copy() - Deeply copy.
//
// @v: Value to copy.
func Copy(v interface{}) interface{} {
	switch v.(type) {
	case []interface{}:
		list := []interface{}{}
		for _, item := range List(v) {
			list = append(list, Copy(item))
		}
		return list

	case map[string]interface{}:
		dict := map[string]interface{}{}
		for key, item := range Dict(v) {
			dict[key] = Copy(item)
		}
		return dict

	default:
		return v
	}
}

// CopyList() - Deeply copy a LIST value.
//
// @v: Value to copy.
func CopyList(v interface{}) []interface{} {
	if IsList(v) {
		return List(Copy(v))
	}

	return []interface{}{}
}

// Copy() - Deeply copy a DICT value.
//
// @v: Value to copy.
func CopyDict(v interface{}) map[string]interface{} {
	if IsDict(v) {
		return Dict(Copy(v))
	}

	return map[string]interface{}{}
}
