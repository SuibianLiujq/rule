// Functions for aggregations.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	"core/value"
	"errors"
	"fmt"
	"strings"
)

func __aggItemsMakeKey(dict map[string]interface{}, keys []interface{}) string {
	list := make([]string, 0, len(keys))

	for _, item := range keys {
		if dictItem, ok := dict[value.ToStr(item)]; ok {
			list = append(list, value.ToStr(dictItem))
		}
	}

	return strings.Join(list, "_")
}

func __aggItems(aggList, aggKeys []interface{}) (interface{}, error) {
	if len(aggKeys) == 0 {
		return map[string]interface{}{"": aggList}, nil
	}

	aggDict := map[string]interface{}{}

	for cc, item := range aggList {
		if !value.IsDict(item) {
			return nil, errors.New(fmt.Sprintf("item[%d] not a dict", cc))
		}

		key := __aggItemsMakeKey(value.Dict(item), aggKeys)
		if _, ok := aggDict[key]; ok {
			aggDict[key] = append(value.List(aggDict[key]), item)
		} else {
			aggDict[key] = []interface{}{item}
		}
	}

	return aggDict, nil
}

func aggItems(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	argList, argKeys, ok := []interface{}(nil), []interface{}(nil), false

	switch len(args) {
	case 2:
		if value.IsList(args[1]) {
			argKeys = value.List(args[1])
		} else if value.IsStr(args[1]) {
			argKeys = []interface{}{args[1]}
		} else {
			return nil, errors.New(fmt.Sprintf("args[1] not list or str"))
		}
		fallthrough

	case 1:
		if argList, ok = value.AsList(args[0]); !ok {
			return nil, errors.New(fmt.Sprintf("args[0] not a list"))
		}

	default:
		return nil, errors.New(fmt.Sprintf("arg mismatch %d (expected 1 or 2", len(args)))
	}

	return __aggItems(argList, argKeys)
}

func __aggItemsFunc(args []interface{}, name string, ctx *script.Cntx) (interface{}, error) {
	argList, argSort, argKeys, ok := []interface{}(nil), "", []interface{}(nil), false

	switch len(args) {
	case 3:
		if args[2] != nil {
			if value.IsList(args[2]) {
				argKeys = value.List(args[2])
			} else if value.IsStr(args[2]) {
				argKeys = []interface{}{args[2]}
			} else {
				return nil, errors.New(fmt.Sprintf("args[2] not list or str"))
			}
		}
		fallthrough

	case 2:
		argSort = value.ToStr(args[1])
		if argList, ok = value.AsList(args[0]); !ok {
			return nil, errors.New(fmt.Sprintf("args[0] not a list"))
		}

	default:
		return nil, errors.New(fmt.Sprintf("arg mismatch %d (expected 2 or 3", len(args)))
	}

	aggDict, err := __aggItems(argList, argKeys)
	if err != nil {
		return nil, err
	}

	resList := []interface{}{}
	for _, list := range value.Dict(aggDict) {
		function, ok := g_functions["item_"+name]
		if !ok {
			msg := fmt.Sprintf("agg function '%s' not found", name)
			return nil, errors.New(msg)
		}

		aggItem, err := function([]interface{}{list, argSort}, ctx)
		if err != nil {
			return nil, err
		}

		resList = append(resList, aggItem)
	}

	return resList, nil
}

func aggItemsMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return __aggItemsFunc(args, "max", ctx)
}

// aggItems() - Aggregate the items to a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggItemsOld(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d(expected 2)", len(args))
		return nil, errors.New(msg)
	}

	argObject, ok := args[0].([]interface{})
	if !ok {
		msg := fmt.Sprintf("arg[0] not LIST", value.ToStr(argObject))
		return nil, errors.New(msg)
	}

	aggKey, aggObject := value.ToStr(args[1]), map[string]interface{}{}
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

			if _, ok := aggObject[aggValue]; ok {
				aggObject[aggValue] = append(aggObject[aggValue].([]interface{}), item)
			} else {
				aggObject[aggValue] = []interface{}{item}
			}
		}
	}

	return aggObject, nil
}

// aggItemsMax() - Returns each max item of aggregation result.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func aggItemsMaxOld(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 3 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		msg := fmt.Sprintf("arg[0] not a string")
		return nil, errors.New(msg)
	}

	aggRes, err := aggItems([]interface{}{list, args[1]}, ctx)
	if err != nil {
		return nil, err
	}

	listRes := []interface{}{}
	for _, item := range value.Dict(aggRes) {
		maxItem, err := itemMax([]interface{}{item, args[2]}, ctx)
		if err == nil {
			listRes = append(listRes, maxItem)
		} else {
			return nil, err
		}
	}

	return listRes, nil
}

// agg2ItemsMax() - Returns each max item of aggregation result.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func agg2ItemsMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 4 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		msg := fmt.Sprintf("arg[0] not a string")
		return nil, errors.New(msg)
	}

	aggRes, err := aggItems([]interface{}{list, args[1]}, ctx)
	if err != nil {
		return nil, err
	}

	listRes := []interface{}{}
	for _, item1 := range value.Dict(aggRes) {
		aggRes1, err := aggItems([]interface{}{item1, args[2]}, ctx)
		if err != nil {
			return nil, err
		}

		for _, item2 := range value.Dict(aggRes1) {
			maxItem, err := itemMax([]interface{}{item2, args[3]}, ctx)
			if err == nil {
				listRes = append(listRes, maxItem)
			} else {
				return nil, err
			}
		}
	}

	return listRes, nil
}

// agg3ItemsMax() - Returns each max item of aggregation result.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func agg3ItemsMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 5 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		msg := fmt.Sprintf("arg[0] not a string")
		return nil, errors.New(msg)
	}

	aggRes, err := aggItems([]interface{}{list, args[1]}, ctx)
	if err != nil {
		return nil, err
	}

	listRes := []interface{}{}
	for _, item1 := range value.Dict(aggRes) {
		aggRes1, err := aggItems([]interface{}{item1, args[2]}, ctx)
		if err != nil {
			return nil, err
		}

		for _, item2 := range value.Dict(aggRes1) {
			aggRes2, err := aggItems([]interface{}{item2, args[3]}, ctx)
			if err != nil {
				return nil, err
			}

			for _, item3 := range value.Dict(aggRes2) {
				maxItem, err := itemMax([]interface{}{item3, args[4]}, ctx)
				if err == nil {
					listRes = append(listRes, maxItem)
				} else {
					return nil, err
				}
			}
		}
	}

	return listRes, nil
}

// agg4ItemsMax() - Returns each max item of aggregation result.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func agg4ItemsMax(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 6 {
		msg := fmt.Sprintf("argument mismatch %d(expected 3)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		msg := fmt.Sprintf("arg[0] not a string")
		return nil, errors.New(msg)
	}

	aggRes, err := aggItems([]interface{}{list, args[1]}, ctx)
	if err != nil {
		return nil, err
	}

	listRes := []interface{}{}
	for _, item1 := range value.Dict(aggRes) {
		aggRes1, err := aggItems([]interface{}{item1, args[2]}, ctx)
		if err != nil {
			return nil, err
		}

		for _, item2 := range value.Dict(aggRes1) {
			aggRes2, err := aggItems([]interface{}{item2, args[3]}, ctx)
			if err != nil {
				return nil, err
			}

			for _, item3 := range value.Dict(aggRes2) {
				aggRes3, err := aggItems([]interface{}{item3, args[4]}, ctx)
				if err != nil {
					return nil, err
				}

				for _, item4 := range value.Dict(aggRes3) {
					maxItem, err := itemMax([]interface{}{item4, args[5]}, ctx)
					if err == nil {
						listRes = append(listRes, maxItem)
					} else {
						return nil, err
					}
				}
			}
		}
	}

	return listRes, nil
}
