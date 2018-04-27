// Functions for aggregations.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/json"
	"core/script"
	"core/value"
	"errors"
	"fmt"
)

// aggValues() - Aggregate the values to a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggValues(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 3 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	argObject, ok := args[0].([]interface{})
	if !ok {
		msg := fmt.Sprintf("arg[0] not LIST", json.DumpStrAll(argObject))
		return nil, errors.New(msg)
	}

	aggKey, valueKey, aggObject := value.ToStr(args[1]), value.ToStr(args[2]), map[string]interface{}{}
	if value.IsTrue(argObject) {
		for cc, item := range argObject {
			dict, ok := item.(map[string]interface{})
			if !ok {
				msg := fmt.Sprintf("arg[0][%d] '%s' not DICT", cc, value.ToStr(item))
				return nil, errors.New(msg)
			}

			aggItem, ok := dict[aggKey]
			if !ok {
				msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", aggKey, cc)
				return nil, errors.New(msg)
			}
			aggValue := value.ToStr(aggItem)

			keyValue, ok := item.(map[string]interface{})[valueKey]
			if !ok {
				msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", valueKey, cc)
				return nil, errors.New(msg)
			}

			if _, ok := aggObject[aggValue]; ok {
				aggObject[aggValue] = append(aggObject[aggValue].([]interface{}), keyValue)
			} else {
				aggObject[aggValue] = []interface{}{keyValue}
			}
		}
	}

	aggResult := []interface{}{}
	for k, v := range aggObject {
		aggResult = append(aggResult, map[string]interface{}{aggKey: k, valueKey: v})
	}

	return aggResult, nil
}

// aggFunction() - Base function for aggXxx().
//
// @name: Low-level metric function name.
// @args: Values to do aggregation on.
// @ctx:  Script context.
func aggFunction(name string, args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 3 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	argObject, aggKey, argKey := args[0], value.ToStr(args[1]), value.ToStr(args[2])
	if _, ok := argObject.([]interface{}); !ok {
		msg := fmt.Sprintf("arg[0] not LIST or empty", value.ToStr(argObject))
		return nil, errors.New(msg)
	}

	aggResult := []interface{}{}
	if value.IsTrue(argObject) {
		aggObject := map[string]interface{}{}

		for cc, item := range argObject.([]interface{}) {
			dict, ok := item.(map[string]interface{})
			if !ok {
				msg := fmt.Sprintf("arg[0][%d] '%s' not DICT", cc, value.ToStr(item))
				return nil, errors.New(msg)
			}

			aggItem, ok := dict[aggKey]
			if !ok {
				msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", aggKey, cc)
				return nil, errors.New(msg)
			}
			aggValue := value.ToStr(aggItem)

			keyValue, ok := item.(map[string]interface{})[argKey]
			if !ok {
				msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", argKey, cc)
				return nil, errors.New(msg)
			}

			switch keyValue.(type) {
			case int64, float64, bool:
				if dictItem, ok := aggObject[aggValue]; ok {
					aggObject[aggValue] = append(dictItem.([]interface{}), keyValue)
				} else {
					aggObject[aggValue] = []interface{}{keyValue}
				}

			default:
				msg := fmt.Sprintf("args[0]['%s'] uncomparable: %s", argKey, value.ToStr(keyValue))
				return nil, errors.New(msg)
			}
		}

		function, ok := g_functions[name]
		if !ok {
			msg := fmt.Sprintf("%s() not found", name)
			return nil, errors.New(msg)
		}

		for key, item := range aggObject {
			result, err := function(item.([]interface{}), ctx)
			if err != nil {
				return nil, err
			}

			aggResult = append(aggResult, map[string]interface{}{aggKey: key, argKey: result})
		}
	}

	return aggResult, nil
}

// aggMax() - Return the aggregated-max value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return aggFunction("max", args, ctx)
}

// aggMin() - Return the aggregated-min value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggMin(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return aggFunction("min", args, ctx)
}

// aggMedian() - Return the aggregated-median value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggMedian(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return aggFunction("median", args, ctx)
}

// aggAvg() - Return the aggregated-avg value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggAvg(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return aggFunction("avg", args, ctx)
}

// aggSum() - Return the aggregated-sum value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggSum(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return aggFunction("sum", args, ctx)
}

// aggVariance() - Return the aggregated-variance value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggVariance(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return aggFunction("variance", args, ctx)
}

// aggStdev() - Return the aggregated-stdev value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggStdev(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return aggFunction("stdev", args, ctx)
}

// aggPercentile() - Return the aggregated-percentile value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggPercentile(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 4 {
		msg := fmt.Sprintf("argument mismatch %d (expected 4)", len(args))
		return nil, errors.New(msg)
	}

	argObject, aggKey, argKey, argPercent := args[0], value.ToStr(args[1]), value.ToStr(args[2]), args[3]
	if _, ok := argObject.([]interface{}); !ok || value.IsFalse(argObject) {
		msg := fmt.Sprintf("arg[0] not LIST or empty", value.ToStr(argObject))
		return nil, errors.New(msg)
	}

	aggObject := map[string]interface{}{}
	for cc, item := range argObject.([]interface{}) {
		dict, ok := item.(map[string]interface{})
		if !ok {
			msg := fmt.Sprintf("arg[0][%d] '%s' not DICT", cc, value.ToStr(item))
			return nil, errors.New(msg)
		}

		aggItem, ok := dict[aggKey]
		if !ok {
			msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", aggKey, cc)
			return nil, errors.New(msg)
		}
		aggValue := value.ToStr(aggItem)

		keyValue, ok := item.(map[string]interface{})[argKey]
		if !ok {
			msg := fmt.Sprintf("key '%s' not found in arg[0][%d]", argKey, cc)
			return nil, errors.New(msg)
		}

		switch keyValue.(type) {
		case int64, float64, bool:
			if dictItem, ok := aggObject[aggValue]; ok {
				aggObject[aggValue] = append(dictItem.([]interface{}), keyValue)
			} else {
				aggObject[aggValue] = []interface{}{keyValue}
			}

		default:
			msg := fmt.Sprintf("'%s' of args[0]['%s'] not comparable", value.ToStr(keyValue), argKey)
			return nil, errors.New(msg)
		}
	}

	function, ok := g_functions["percentile"]
	if !ok {
		msg := fmt.Sprintf("%s() not found", "percentile")
		return nil, errors.New(msg)
	}

	aggResult := []interface{}{}
	for key, item := range aggObject {
		result, err := function([]interface{}{item, argPercent}, ctx)
		if err != nil {
			return nil, err
		}

		aggResult = append(aggResult, map[string]interface{}{aggKey: key, argKey: result})
	}

	return aggResult, nil
}
