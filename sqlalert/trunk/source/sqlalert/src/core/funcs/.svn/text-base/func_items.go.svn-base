// Functions for list of dicts.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	"core/value"
	"errors"
	"fmt"
)

// itemMax() - Return the max-item value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func itemMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d(expected 2)", len(args))
		return nil, errors.New(msg)
	}

	argObject, argKey := args[0], value.ToStr(args[1])
	if !value.IsList(argObject) || value.IsFalse(argObject) {
		msg := fmt.Sprintf("arg[0] not LIST or empty: %s", value.ToStr(argObject))
		return nil, errors.New(msg)
	}
	objectList := argObject.([]interface{})

	valueList := []interface{}{}
	for cc, item := range objectList {
		if !value.IsDict(item) {
			msg := fmt.Sprintf("arg[0][%d] '%s' not DICT", cc, value.ToStr(item))
			return nil, errors.New(msg)
		}

		keyValue, ok := item.(map[string]interface{})[argKey]
		if !ok {
			msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", argKey, cc)
			return nil, errors.New(msg)
		}

		switch keyValue.(type) {
		case string, int64, float64, bool:
			valueList = append(valueList, keyValue)

		default:
			msg := fmt.Sprintf("'%s' of args[0]['%s'] not comparable", value.ToStr(keyValue), argKey)
			return nil, errors.New(msg)
		}
	}

	maxIndex, maxValue := 0, valueList[0]
	for cc, item := range valueList[1:] {
		if value.Compare(item, ">", maxValue) {
			maxIndex, maxValue = cc+1, item
		}
	}

	return objectList[maxIndex], nil
}

// itemMin() - Return the min-item value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func itemMin(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d(expected 2)", len(args))
		return nil, errors.New(msg)
	}

	argObject, argKey := args[0], value.ToStr(args[1])
	if !value.IsList(argObject) || value.IsFalse(argObject) {
		msg := fmt.Sprintf("arg[0] not LIST or empty", value.ToStr(argObject))
		return nil, errors.New(msg)
	}
	objectList := argObject.([]interface{})

	valueList := []interface{}{}
	for cc, item := range objectList {
		if !value.IsDict(item) {
			msg := fmt.Sprintf("arg[0][%d] '%s' not DICT", cc, value.ToStr(item))
			return nil, errors.New(msg)
		}

		keyValue, ok := item.(map[string]interface{})[argKey]
		if !ok {
			msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", argKey, cc)
			return nil, errors.New(msg)
		}

		switch keyValue.(type) {
		case string, int64, float64, bool:
			valueList = append(valueList, keyValue)

		default:
			msg := fmt.Sprintf("'%s' of args[0]['%s'] not comparable", value.ToStr(keyValue), argKey)
			return nil, errors.New(msg)
		}
	}

	minIndex, minValue := 0, valueList[0]
	for cc, item := range valueList[1:] {
		if value.Compare(item, "<", minValue) {
			minIndex, minValue = cc+1, item
		}
	}

	return objectList[minIndex], nil
}

// itemMedian() - Return the meidan-item value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func itemMedian(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d(expected 2)", len(args))
		return nil, errors.New(msg)
	}

	args = append(args, int64(50))
	return itemPercentile(args, ctx)
}

// itemPercentile() - Return the percentile-item value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func itemPercentile(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 3 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	return statPercentile(args, ctx)
}

// itemFilter() - Returned the filtered items.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func itemFilter(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	if !value.IsList(args[0]) {
		return nil, errors.New("invalid arguments")
	}

	if value.IsFalse(args[1]) {
		return args[0], nil
	}

	result := []interface{}{}
	subCtx := ctx.Copy()
	for cc, item := range args[0].([]interface{}) {
		if !value.IsDict(item) {
			msg := fmt.Sprintf("arg[0][%d] not DICT: %s", cc, value.ToStr(item))
			return nil, errors.New(msg)
		}

		for k, v := range item.(map[string]interface{}) {
			subCtx.Local[k] = v
		}

		exprValue, err := script.ExecScript([]byte(value.ToStr(args[1])), subCtx)
		if err != nil {
			msg := fmt.Sprintf("'%s' invalid expression: %s", value.ToStr(args[1]), err)
			return args[0], errors.New(msg)
		}

		if value.IsTrue(exprValue) {
			result = append(result, item)
		}
	}

	return result, nil
}

// itemSet() - Returned the values of dict items in a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func itemSet(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 3 {
		msg := fmt.Sprintf("argument mismatch %d (expected 3)", len(args))
		return nil, errors.New(msg)
	}

	if list, ok := value.AsList(args[0]); ok && args[1] != nil && args[2] != nil {
		key := value.ToStr(args[1])
		for _, item := range list {
			if dict, ok := value.AsDict(item); ok {
				dict[key] = args[2]
			}
		}
	}

	return args[0], nil
}

// itemValues() - Returned the values of dict items in a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func itemValues(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	if !value.IsList(args[0]) {
		return nil, errors.New("invalid arguments")
	}

	if value.IsFalse(args[1]) {
		return args[0], nil
	}

	result := []interface{}{}
	for cc, item := range args[0].([]interface{}) {
		if !value.IsDict(item) {
			msg := fmt.Sprintf("arg[0][%d] not DICT: %s", cc, value.ToStr(item))
			return nil, errors.New(msg)
		}

		dict, key := item.(map[string]interface{}), value.ToStr(args[1])
		if dictValue, ok := dict[key]; ok {
			result = append(result, dictValue)
		} else {
			msg := fmt.Sprintf("key '%s' not found in args[0][%d]", value.ToStr(key), cc)
			return nil, errors.New(msg)
		}
	}

	return result, nil
}
