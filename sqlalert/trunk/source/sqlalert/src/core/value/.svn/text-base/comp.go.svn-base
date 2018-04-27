// File: comp.go
//
// This file implements the comparison between two interface value.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package value

import (
	"fmt"
	"strconv"
)

// CmpIntInt() - Cmpare INT and INT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpIntInt(v1 int64, opt string, v2 int64) bool {
	switch opt {
	case "<":
		return v1 < v2
	case ">":
		return v1 > v2
	case "<=":
		return v1 <= v2
	case ">=":
		return v1 >= v2
	case "==":
		return v1 == v2
	case "!=":
		return v1 != v2
	}

	return false
}

// CmpIntFloat() - Cmpare INT and FLOAT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpIntFloat(v1 int64, opt string, v2 float64) bool {
	return CmpFloatFloat(float64(v1), opt, v2)
}

// CmpIntStr() - Cmpare INT and STR.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpIntStr(v1 int64, opt string, v2 string) bool {
	if v, err := strconv.ParseInt(v2, 10, 64); err == nil {
		return CmpIntInt(v1, opt, v)
	}

	return CmpStrStr(fmt.Sprintf("%d", v1), opt, v2)
}

// CmpIntBool() - Cmpare INT and BOOL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpIntBool(v1 int64, opt string, v2 bool) bool {
	if v2 {
		return CmpIntInt(v1, opt, 1)
	}

	return CmpIntInt(v1, opt, 0)
}

// CmpIntNull() - Cmpare INT and NULL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpIntNull(v1 int64, opt string, v2 interface{}) bool {
	switch opt {
	case "<":
		return false
	case ">":
		return true
	case "<=":
		return false
	case ">=":
		return true
	case "==":
		return false
	case "!=":
		return true
	}

	return false
}

// CmpInt() - Cmpare v and INT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpInt(v1 interface{}, opt string, v2 int64) bool {
	if v1 == nil {
		return CmpNullInt(v1, opt, v2)
	}

	switch v1.(type) {
	case int64:
		return CmpIntInt(v1.(int64), opt, v2)
	case float64:
		return CmpFloatInt(v1.(float64), opt, v2)
	case string:
		return CmpStrInt(v1.(string), opt, v2)
	case bool:
		return CmpBoolInt(v1.(bool), opt, v2)
	}

	if opt == "!=" {
		return true
	}
	return false
}

// CmpFloatInt() - Cmpare FLOAT and INT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpFloatInt(v1 float64, opt string, v2 int64) bool {
	return CmpFloatFloat(v1, opt, float64(v2))
}

// CmpIntInt() - Cmpare FLOAT and FLOAT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpFloatFloat(v1 float64, opt string, v2 float64) bool {
	switch opt {
	case "<":
		return v1 < v2
	case ">":
		return v1 > v2
	case "<=":
		return v1 <= v2
	case ">=":
		return v1 >= v2
	case "==":
		return v1 == v2
	case "!=":
		return v1 != v2
	}

	return false
}

// CmpIntStr() - Cmpare INT and STRING.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpFloatStr(v1 float64, opt string, v2 string) bool {
	if v, err := strconv.ParseFloat(v2, 64); err == nil {
		return CmpFloatFloat(v1, opt, v)
	}

	return CmpStrStr(fmt.Sprintf("%f", v1), opt, v2)
}

// CmpFloatBool() - Cmpare FLOAT and BOOL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpFloatBool(v1 float64, opt string, v2 bool) bool {
	if v2 {
		return CmpFloatFloat(v1, opt, 1.0)
	}

	return CmpFloatFloat(v1, opt, 0.0)
}

// CmpFloatNull() - Cmpare FLOAT and NULL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpFloatNull(v1 float64, opt string, v2 interface{}) bool {
	switch opt {
	case "<":
		return false
	case ">":
		return true
	case "<=":
		return false
	case ">=":
		return true
	case "==":
		return false
	case "!=":
		return true
	}

	return false
}

// CmpFloat() - Cmpare v and FLOAT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpFloat(v1 interface{}, opt string, v2 float64) bool {
	if v1 == nil {
		return CmpNullFloat(v1, opt, v2)
	}

	switch v1.(type) {
	case int64:
		return CmpIntFloat(v1.(int64), opt, v2)
	case float64:
		return CmpFloatFloat(v1.(float64), opt, v2)
	case string:
		return CmpStrFloat(v1.(string), opt, v2)
	}

	if opt == "!=" {
		return true
	}
	return false
}

// CmpStrInt() - Cmpare STRING and INT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpStrInt(v1 string, opt string, v2 int64) bool {
	if v, err := strconv.ParseInt(v1, 10, 64); err == nil {
		return CmpIntInt(v, opt, v2)
	}

	return CmpStrStr(v1, opt, fmt.Sprintf("%d", v2))
}

// CmpStrFloat() - Cmpare STR and FLOAT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpStrFloat(v1 string, opt string, v2 float64) bool {
	if v, err := strconv.ParseFloat(v1, 64); err == nil {
		return CmpFloatFloat(v, opt, v2)
	}

	return CmpStrStr(v1, opt, fmt.Sprintf("%f", v2))
}

// CmpStrStr() - Cmpare STR and STR.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpStrStr(v1 string, opt string, v2 string) bool {
	switch opt {
	case "<":
		return v1 < v2
	case ">":
		return v1 > v2
	case "<=":
		return v1 <= v2
	case ">=":
		return v1 >= v2
	case "==":
		return v1 == v2
	case "!=":
		return v1 != v2
	}

	return false
}

// CmpStrBool() - Cmpare STR and BOOL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpStrBool(v1 string, opt string, v2 bool) bool {
	if v, err := strconv.ParseInt(v1, 10, 64); err == nil {
		if v2 {
			return CmpIntInt(v, opt, 1)
		} else {
			return CmpIntInt(v, opt, 0)
		}
	}

	return CmpStrStr(v1, opt, ToStr(v2))
}

// CmpStrNull() - Cmpare STR and NULL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpStrNull(v1 string, opt string, v2 interface{}) bool {
	switch opt {
	case "<":
		return false
	case ">":
		return true
	case "<=":
		return false
	case ">=":
		return true
	case "==":
		return false
	case "!=":
		return true
	}

	return false
}

// CmpStr() - Cmpare v and STR.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpStr(v1 interface{}, opt string, v2 string) bool {
	if v1 == nil {
		return CmpNullStr(v1, opt, v2)
	}

	switch v1.(type) {
	case int64:
		return CmpIntStr(v1.(int64), opt, v2)
	case float64:
		return CmpFloatStr(v1.(float64), opt, v2)
	case string:
		return CmpStrStr(v1.(string), opt, v2)
	}

	if opt == "!=" {
		return true
	}
	return false
}

// CmpBoolInt() - Cmpare BOOL and INT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpBoolInt(v1 bool, opt string, v2 int64) bool {
	if v1 {
		return CmpIntInt(1, opt, v2)
	}

	return CmpIntInt(0, opt, v2)
}

// CmpBoolFloat() - Cmpare BOOL and FLOAT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpBoolFloat(v1 bool, opt string, v2 float64) bool {
	if v1 {
		return CmpFloatFloat(1.0, opt, v2)
	}

	return CmpFloatFloat(0.0, opt, v2)
}

// CmpBoolStr() - Cmpare BOOL and STR.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpBoolStr(v1 bool, opt string, v2 string) bool {
	if v, err := strconv.ParseInt(v2, 10, 64); err == nil {
		if v1 {
			return CmpIntInt(1, opt, v)
		} else {
			return CmpIntInt(0, opt, v)
		}
	}

	return CmpStrStr(ToStr(v1), opt, v2)
}

// CmpBoolBool() - Cmpare BOOL and BOOL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpBoolBool(v1 bool, opt string, v2 bool) bool {
	i1 := int64(0)
	if v1 {
		i1 = 1
	}

	i2 := int64(0)
	if v2 {
		i2 = 1
	}

	return CmpIntInt(i1, opt, i2)
}

// CmpBoolNull() - Cmpare BOOL and NULL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpBoolNull(v1 bool, opt string, v2 interface{}) bool {
	switch opt {
	case "<":
		return false
	case ">":
		return true
	case "<=":
		return false
	case ">=":
		return true
	case "==":
		return false
	case "!=":
		return true
	}

	return false
}

// CmpBool() - Cmpare v and BOOL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpBool(v1 interface{}, opt string, v2 bool) bool {
	if v1 == nil {
		return CmpNullBool(v1, opt, v2)
	}

	switch v1.(type) {
	case int64:
		return CmpIntBool(v1.(int64), opt, v2)
	case float64:
		return CmpFloatBool(v1.(float64), opt, v2)
	case string:
		return CmpStrBool(v1.(string), opt, v2)
	}

	if opt == "!=" {
		return true
	}
	return false
}

// CmpNullInt() - Cmpare NULL and INT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpNullInt(v1 interface{}, opt string, v2 int64) bool {
	return CmpIntInt(0, opt, v2)
}

// CmpNullFloat() - Cmpare NULL and FLOAT.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpNullFloat(v1 interface{}, opt string, v2 float64) bool {
	return CmpFloatFloat(0.0, opt, v2)
}

// CmpNullStr() - Cmpare NULL and STR.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpNullStr(v1 interface{}, opt string, v2 string) bool {
	return CmpStrStr("", opt, v2)
}

// CmpNullStr() - Cmpare NULL and STR.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpNullBool(v1 interface{}, opt string, v2 bool) bool {
	return CmpBoolBool(false, opt, v2)
}

// CmpNullStr() - Cmpare NULL and STR.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpNullNull(v1 interface{}, opt string, v2 interface{}) bool {
	switch opt {
	case "<":
		return false
	case ">":
		return false
	case "<=":
		return true
	case ">=":
		return true
	case "==":
		return true
	case "!=":
		return false
	}

	return false
}

// CmpBool() - Cmpare v and BOOL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func CmpNull(v1 interface{}, opt string, v2 interface{}) bool {
	if v1 == nil {
		return CmpNullNull(v1, opt, v2)
	}

	switch v1.(type) {
	case int64:
		return CmpIntNull(v1.(int64), opt, v2)
	case float64:
		return CmpFloatNull(v1.(float64), opt, v2)
	case string:
		return CmpStrNull(v1.(string), opt, v2)
	}

	if opt == "!=" {
		return true
	}
	return false
}

// CmpBool() - Cmpare v and BOOL.
//
// @v1:  Left value to compare.
// @opt: String value of operator.
// @v2:  Right value to compare.
func Compare(v1 interface{}, opt string, v2 interface{}) bool {
	if v2 == nil {
		return CmpNull(v1, opt, v2)
	}

	switch v2.(type) {
	case int64:
		return CmpInt(v1, opt, v2.(int64))
	case float64:
		return CmpFloat(v1, opt, v2.(float64))
	case string:
		return CmpStr(v1, opt, v2.(string))
	case bool:
		return CmpBool(v1, opt, v2.(bool))
	}

	if opt == "!=" {
		return true
	}

	return false
}
