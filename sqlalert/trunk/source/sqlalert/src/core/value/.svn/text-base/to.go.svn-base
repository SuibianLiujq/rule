// File: to.go
//
// This file implements the converting function for values.
//
// Copyright (c) 2017 Yun Li Lai, Ltd, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>
package value

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// ToInt() - Convert @v to INT value.
//
// @v: Value to convert.
func ToInt(v interface{}) (int64, error) {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return Int(v), nil

	case string:
		return strconv.ParseInt(v.(string), 10, 64)
	}

	return Int(0), errors.New("not an INT")
}

// ToFloat() - Convert @v to FLOAT value.
//
// @v: Value to convert.
func ToFloat(v interface{}) (float64, error) {
	switch v.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return Float(v), nil

	case string:
		return strconv.ParseFloat(v.(string), 64)
	}

	return Float(0), errors.New("not a float")
}

// ToStr() Convert @v to STR value.
//
// @v: Value to convert.
func ToStr(v interface{}) string {
	if v == nil {
		return "null"
	}

	switch v.(type) {
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", Int(v))
	case float32, float64:
		return strconv.FormatFloat(Float(v), 'f', -1, 64)
	case string:
		return v.(string)
	case bool:
		if v.(bool) {
			return "true"
		} else {
			return "false"
		}
	case []byte:
		return string(v.([]byte))
	case []interface{}:
		buffer, list := &bytes.Buffer{}, List(v)

		buffer.WriteByte('[')
		if len(list) != 0 {
			buffer.WriteString(ToStr(list[0]))
			for _, item := range list[1:] {
				buffer.WriteByte(',')
				buffer.WriteString(ToStr(item))
			}
		}
		buffer.WriteByte(']')
		return buffer.String()

	case map[string]interface{}:
		buffer, dict, first := &bytes.Buffer{}, Dict(v), true

		buffer.WriteByte('{')
		for key, item := range dict {
			if !first {
				buffer.WriteByte(',')
			}

			first = false
			buffer.WriteString(ToStr(key))
			buffer.WriteByte(':')
			buffer.WriteString(ToStr(item))
		}
		buffer.WriteByte('}')
		return buffer.String()
	}

	return fmt.Sprintf("%v", v)
}

// ToBool() Convert @v to BOOL value.
//
// @v: Value to convert.
func ToBool(v interface{}) (bool, error) {
	return IsTrue(v), nil
}
