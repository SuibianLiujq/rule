// Functions for math statistics.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	"core/value"
	"errors"
	"fmt"
	"math"
)

// funcMax() - Return the max value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) == 0 {
		msg := fmt.Sprintf("argument mismatch %d (expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	for cc, argItem := range args {
		switch argItem.(type) {
		case []interface{}:
			for i, listItem := range argItem.([]interface{}) {
				switch listItem.(type) {
				case string, int64, float64, bool:
					list = append(list, listItem)

				default:
					msg := fmt.Sprintf("'%s' of args[%d][%d] not comparable", value.ToStr(listItem), cc, i)
					return nil, errors.New(msg)
				}
			}
		case string, int64, float64, bool:
			list = append(list, argItem)

		default:
			msg := fmt.Sprintf("'%s' of args[%d] not comparable", value.ToStr(argItem), cc)
			return nil, errors.New(msg)

		}
	}

	maxValue := list[0]
	for _, item := range list[1:] {
		if value.Compare(item, ">", maxValue) {
			maxValue = item
		}
	}

	return maxValue, nil
}

// funcMin() - Return the min value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statMin(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) == 0 {
		msg := fmt.Sprintf("argument mismatch %d(expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	for cc, argItem := range args {
		switch argItem.(type) {
		case []interface{}:
			for i, listItem := range argItem.([]interface{}) {
				switch listItem.(type) {
				case string, int64, float64, bool:
					list = append(list, listItem)

				default:
					msg := fmt.Sprintf("'%s' of args[%d][%d] not comparable", value.ToStr(listItem), cc, i)
					return nil, errors.New(msg)
				}
			}
		case string, int64, float64, bool:
			list = append(list, argItem)

		default:
			msg := fmt.Sprintf("'%s' of args[%d] not comparable", value.ToStr(argItem), cc)
			return nil, errors.New(msg)

		}
	}

	minValue := list[0]
	for _, item := range list[1:] {
		if value.Compare(item, "<", minValue) {
			minValue = item
		}
	}

	return minValue, nil
}

// funcMedian() - Return the median value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statMedian(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) == 0 {
		msg := fmt.Sprintf("argument mismatch %d(expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	for cc, argItem := range args {
		switch argItem.(type) {
		case []interface{}:
			for i, listItem := range argItem.([]interface{}) {
				switch listItem.(type) {
				case string, int64, float64, bool:
					list = append(list, listItem)

				default:
					msg := fmt.Sprintf("'%s' of args[%d][%d] not comparable", value.ToStr(listItem), cc, i)
					return nil, errors.New(msg)
				}
			}
		case string, int64, float64, bool:
			list = append(list, argItem)

		default:
			msg := fmt.Sprintf("'%s' of args[%d] not comparable", value.ToStr(argItem), cc)
			return nil, errors.New(msg)

		}
	}

	return statPercentile([]interface{}{list, int64(50)}, ctx)
}

// funcAvg() - Return the avg value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statAvg(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) == 0 {
		msg := fmt.Sprintf("argument mismatch %d(expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	for cc, argItem := range args {
		switch argItem.(type) {
		case []interface{}:
			for i, listItem := range argItem.([]interface{}) {
				switch listItem.(type) {
				case int64, float64, bool:
					list = append(list, listItem)

				default:
					msg := fmt.Sprintf("'%s' of args[%d][%d] not comparable", value.ToStr(listItem), cc, i)
					return nil, errors.New(msg)
				}
			}
		case int64, float64, bool:
			list = append(list, argItem)

		default:
			msg := fmt.Sprintf("'%s' of args[%d] not comparable", value.ToStr(argItem), cc)
			return nil, errors.New(msg)

		}
	}

	sumValue := list[0]
	for _, item := range list[1:] {
		value, err := value.Add(sumValue, item)
		if err != nil {
			return nil, err
		}
		sumValue = value
	}

	return value.Div(sumValue, int64(len(list)))
}

// funcSum() - Return the sum value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statSum(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) == 0 {
		msg := fmt.Sprintf("argument mismatch %d(expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	for cc, argItem := range args {
		switch argItem.(type) {
		case []interface{}:
			for i, listItem := range argItem.([]interface{}) {
				switch listItem.(type) {
				case int64, float64, bool:
					list = append(list, listItem)

				default:
					msg := fmt.Sprintf("'%s' of args[%d][%d] not comparable", value.ToStr(listItem), cc, i)
					return nil, errors.New(msg)
				}
			}
		case int64, float64, bool:
			list = append(list, argItem)

		default:
			msg := fmt.Sprintf("'%s' of args[%d] not comparable", value.ToStr(argItem), cc)
			return nil, errors.New(msg)

		}
	}

	sumValue := list[0]
	for _, item := range list[1:] {
		value, err := value.Add(sumValue, item)
		if err != nil {
			return nil, err
		}
		sumValue = value
	}

	return sumValue, nil
}

// funcVariance() - Return the variance value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statVariance(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) == 0 {
		msg := fmt.Sprintf("argument mismatch %d(expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	for cc, argItem := range args {
		switch argItem.(type) {
		case []interface{}:
			for i, listItem := range argItem.([]interface{}) {
				switch listItem.(type) {
				case int64, float64, bool:
					list = append(list, listItem)

				default:
					msg := fmt.Sprintf("'%s' of args[%d][%d] not comparable", value.ToStr(listItem), cc, i)
					return nil, errors.New(msg)
				}
			}
		case int64, float64, bool:
			list = append(list, argItem)

		default:
			msg := fmt.Sprintf("'%s' of args[%d] not comparable", value.ToStr(argItem), cc)
			return nil, errors.New(msg)
		}
	}

	sumValue := list[0]
	for _, item := range list[1:] {
		value, err := value.Add(sumValue, item)
		if err != nil {
			return nil, err
		}
		sumValue = value
	}

	avgValue, _ := value.Div(sumValue, int64(len(list)))
	sumValue = nil

	for _, item := range list {
		subValue, err := value.Sub(item, avgValue)
		if err != nil {
			return nil, err
		}

		sqrtValue, err := value.Mul(subValue, subValue)
		if err != nil {
			return nil, err
		}

		if sumValue == nil {
			sumValue = sqrtValue
		} else {
			sumValue, _ = value.Add(sumValue, sqrtValue)
		}
	}

	v, err := value.Div(sumValue, float64(len(list)))
	return v, err
}

// funcStdev() - Return the stdev value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statStdev(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) == 0 {
		msg := fmt.Sprintf("argument mismatch %d(expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	for cc, argItem := range args {
		switch argItem.(type) {
		case []interface{}:
			for i, listItem := range argItem.([]interface{}) {
				switch listItem.(type) {
				case int64, float64, bool:
					list = append(list, listItem)

				default:
					msg := fmt.Sprintf("'%s' of args[%d][%d] not comparable", value.ToStr(listItem), cc, i)
					return nil, errors.New(msg)
				}
			}
		case int64, float64, bool:
			list = append(list, argItem)

		default:
			msg := fmt.Sprintf("'%s' of args[%d] not comparable", value.ToStr(argItem), cc)
			return nil, errors.New(msg)

		}
	}

	sumValue := list[0]
	for _, item := range list[1:] {
		value, err := value.Add(sumValue, item)
		if err != nil {
			return nil, err
		}
		sumValue = value
	}

	avgValue, _ := value.Div(sumValue, int64(len(list)))
	sumValue = nil

	for _, item := range list {
		subValue, err := value.Sub(item, avgValue)
		if err != nil {
			return nil, err
		}

		sqrtValue, err := value.Mul(subValue, subValue)
		if err != nil {
			return nil, err
		}

		if sumValue == nil {
			sumValue = sqrtValue
		} else {
			sumValue, _ = value.Add(sumValue, sqrtValue)
		}
	}

	v, _ := value.Div(sumValue, float64(len(list)))
	if value.IsInt(v) {
		return float64(math.Sqrt(float64(v.(int64)))), nil
	}

	return math.Sqrt(v.(float64)), nil
}

// funcPercentile() - Return the percentile value of given arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func statPercentile(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 3)", len(args))
		return nil, errors.New(msg)
	}

	argObject, argKey, argPercent := args[0], string(""), args[1]
	if len(args) == 3 {
		argObject, argKey, argPercent = args[0], value.ToStr(args[1]), args[2]
	}

	objectList, ok := argObject.([]interface{})
	if !ok || value.IsFalse(objectList) {
		msg := fmt.Sprintf("%s of arg[0] not LIST or empty", value.ToStr(argObject))
		return nil, errors.New(msg)
	}

	if !value.IsInt(argPercent) {
		msg := fmt.Sprintf("%s of arg[1] not INTEGER", value.ToStr(argPercent))
		return nil, errors.New(msg)
	}

	if len(args) == 3 {
		for cc, item := range objectList {
			if !value.IsDict(item) {
				msg := fmt.Sprintf("'%s' of args[0]['%d'] not DICT", value.ToStr(item), cc)
				return nil, errors.New(msg)
			}
		}
	} else {
		for cc, item := range objectList {
			switch item.(type) {
			case string, int64, float64, bool:
			default:
				msg := fmt.Sprintf("'%s' of args[0]['%d'] not comparable", value.ToStr(item), cc)
				return nil, errors.New(msg)
			}
		}
	}

	sortList := value.CopySort(objectList, argKey, "asc")
	percentile := argPercent.(int64)

	if percentile < 0 || percentile > 100 {
		msg := fmt.Sprintf("invalid percentile of arg[2] (expected 0~100)", value.ToStr(argPercent))
		return nil, errors.New(msg)
	}

	index := (int64(len(sortList)-1) * percentile) / 100
	return sortList[index], nil
}
