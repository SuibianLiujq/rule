// File: comp.go
//
// This file implements 'in' operate.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package value

func inListInt(v interface{}, o []int) bool {
	intValue, ok := AsInt(v)
	if !ok {
		return false
	}

	for _, item := range o {
		if intValue == Int(item) {
			return true
		}
	}

	return false
}

func inListInt8(v interface{}, o []int8) bool {
	intValue, ok := AsInt(v)
	if !ok {
		return false
	}

	for _, item := range o {
		if intValue == Int(item) {
			return true
		}
	}

	return false
}

func inListInt16(v interface{}, o []int16) bool {
	intValue, ok := AsInt(v)
	if !ok {
		return false
	}

	for _, item := range o {
		if intValue == Int(item) {
			return true
		}
	}

	return false
}

func inListInt32(v interface{}, o []int32) bool {
	intValue, ok := AsInt(v)
	if !ok {
		return false
	}

	for _, item := range o {
		if intValue == Int(item) {
			return true
		}
	}

	return false
}

func inListInt64(v interface{}, o []int64) bool {
	intValue, ok := AsInt(v)
	if !ok {
		return false
	}

	for _, item := range o {
		if intValue == Int(item) {
			return true
		}
	}

	return false
}

func inListStr(v interface{}, o []string) bool {
	strValue, ok := AsStr(v)
	if !ok {
		return false
	}

	for _, item := range o {
		if strValue == item {
			return true
		}
	}

	return false
}

// inList() - Test whether @v is a element of LIST @o.
//
// @v: Value to test.
// @o: List object to check.
func inList(v interface{}, o []interface{}) bool {
	for _, item := range o {
		if Compare(v, "==", item) {
			return true
		}
	}
	return false
}

func inDictInt(v string, o map[string]int) bool {
	for key, _ := range o {
		if v == key {
			return true
		}
	}
	return false
}

func inDictInt8(v string, o map[string]int8) bool {
	for key, _ := range o {
		if v == key {
			return true
		}
	}
	return false
}

func inDictInt16(v string, o map[string]int16) bool {
	for key, _ := range o {
		if v == key {
			return true
		}
	}
	return false
}

func inDictInt32(v string, o map[string]int32) bool {
	for key, _ := range o {
		if v == key {
			return true
		}
	}
	return false
}

func inDictInt64(v string, o map[string]int64) bool {
	for key, _ := range o {
		if v == key {
			return true
		}
	}
	return false
}

func inDictStr(v string, o map[string]string) bool {
	for key, _ := range o {
		if v == key {
			return true
		}
	}
	return false
}

// inDict() - Test whether @v is a element of DICT @o.
//
// @v: Value to test.
// @o: Dict object to check.
//
// This function only checks the keys in the DICT.
func inDict(v string, o map[string]interface{}) bool {
	for key, _ := range o {
		if v == key {
			return true
		}
	}
	return false
}

// In() - Test whether @v is a element of @o.
//
// @v: Value to test.
// @o: Object to check.
func In(v, o interface{}) bool {
	if v == nil || o == nil {
		return false
	}

	switch o.(type) {
	case []int:
		return inListInt(v, o.([]int))
	case []int8:
		return inListInt8(v, o.([]int8))
	case []int16:
		return inListInt16(v, o.([]int16))
	case []int32:
		return inListInt32(v, o.([]int32))
	case []int64:
		return inListInt64(v, o.([]int64))
	case []string:
		return inListStr(v, o.([]string))
	case []interface{}:
		return inList(v, o.([]interface{}))

	case map[string]int:
		return inDictInt(ToStr(v), o.(map[string]int))
	case map[string]int8:
		return inDictInt8(ToStr(v), o.(map[string]int8))
	case map[string]int16:
		return inDictInt16(ToStr(v), o.(map[string]int16))
	case map[string]int32:
		return inDictInt32(ToStr(v), o.(map[string]int32))
	case map[string]int64:
		return inDictInt64(ToStr(v), o.(map[string]int64))
	case map[string]string:
		return inDictStr(ToStr(v), o.(map[string]string))
	case map[string]interface{}:
		return inDict(ToStr(v), o.(map[string]interface{}))
	}

	return false
}
