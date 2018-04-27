// Functions for values in list of dicts.
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

// valueFunction() - Function for valueXxxx().
//
// @name: Low-level function name.
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueFunction(name string, args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d(expected 2)", len(args))
		return nil, errors.New(msg)
	}

	argObject, argKey := args[0], value.ToStr(args[1])
	if !value.IsList(argObject) || value.IsFalse(argObject) {
		msg := fmt.Sprintf("arg[0] not LIST or empty", value.ToStr(argObject))
		return nil, errors.New(msg)
	}

	valueList := []interface{}{}
	for cc, item := range argObject.([]interface{}) {
		if !value.IsDict(item) {
			msg := fmt.Sprintf("arg[0][%d] '%s' not DICT", cc, value.ToStr(item))
			return nil, errors.New(msg)
		}

		keyValue, ok := item.(map[string]interface{})[argKey]
		if !ok {
			msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", argKey, cc)
			return nil, errors.New(msg)
		}

		valueList = append(valueList, keyValue)
	}

	function, ok := g_functions[name]
	if !ok {
		msg := fmt.Sprintf("%s() not found", name)
		return nil, errors.New(msg)
	}

	return function(valueList, ctx)
}

// valueMax() - Return the max-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return valueFunction("max", args, ctx)
}

// funcMinValue() - Return the min-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueMin(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return valueFunction("min", args, ctx)
}

// valueMedian() - Return the median-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueMedian(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return valueFunction("median", args, ctx)
}

// valueAvg() - Return the avg-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueAvg(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return valueFunction("avg", args, ctx)
}

// valueSum() - Return the sum-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueSum(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return valueFunction("sum", args, ctx)
}

// valueVariance() - Return the variance-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueVariance(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return valueFunction("variance", args, ctx)
}

// valueStdev() - Return the stdev-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueStdev(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return valueFunction("stdev", args, ctx)
}

// valuePercentile() - Return the percentile-value value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valuePercentile(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 3 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	item, err := statPercentile(args, ctx)
	if err != nil {
		return nil, err
	}

	return item.(map[string]interface{})[value.ToStr(args[1])], nil
}

// valueMap() - Map values from one to another.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func valueMap(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	list, key, dict, newkey, ok := []interface{}{}, "", map[string]interface{}{}, "", false

	switch len(args) {
	case 4:
		newkey = value.ToStr(args[3])
		fallthrough

	case 3:
		if list, ok = value.AsList(args[0]); !ok {
			msg := fmt.Sprintf("arg[0] not a list")
			return nil, errors.New(msg)
		}

		key = value.ToStr(args[1])
		if dict, ok = value.AsDict(args[2]); !ok {
			msg := fmt.Sprintf("arg[2] not a dict")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 3 or 4)", len(args))
		return nil, errors.New(msg)
	}

	if newkey == "" {
		newkey = key
	}

	for _, item := range list {
		if dictItem, ok := value.AsDict(item); ok {
			if itemValue, ok := dictItem[key]; ok {
				if mapValue, ok := dict[value.ToStr(itemValue)]; ok {
					dictItem[newkey] = mapValue
				}
			}
		}
	}

	return list, nil
}
