// Functions for alert throttling.
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

func __thSessionGet(key string, ctx *script.Cntx) []interface{} {
	session := []interface{}{}

	if v := ctx.GetSession(key, session); value.IsList(v) {
		return value.List(v)
	} else {
		ctx.SetSession(key, session)
	}

	return session
}

func __thSessionSet(key string, session []interface{}, data interface{}, ctx *script.Cntx) {
	window := int64(12)

	if val, ok := ctx.GetXInt("__throttle_window__"); ok && val != 0 {
		window = val
	}

	session = append(session, data)
	if len(session) > int(window) {
		session = session[1:]
	}

	ctx.SetSession(key, session)
}

func __thSessionCount(session []interface{}, key string) int64 {
	count := int64(1)

	for _, item := range session {
		if dict, ok := value.AsDict(item); ok {
			if _, ok := dict[key]; ok {
				count++
			}
		}
	}

	return count
}

func __thListToDict(list []interface{}, ctx *script.Cntx) map[string]interface{} {
	listKey, ok := ctx.GetXList("__throttle_keys__")
	if !ok {
		listKey = []interface{}{}
	}

	res, _ := miscListToDict([]interface{}{list, listKey, "_"}, ctx)
	if res == nil {
		return nil
	}

	return value.Dict(res)
}

func __thCheck(key string, list []interface{}, ctx *script.Cntx) []interface{} {
	session, dict, max, min := __thSessionGet(key, ctx), __thListToDict(list, ctx), int64(1), int64(1)

	if val, ok := ctx.GetXInt("__throttle_max__"); ok && val != 0 {
		max = val
	}

	if val, ok := ctx.GetXInt("__throttle_min__"); ok && val != 0 {
		min = val
	}

	listRes := []interface{}{}
	for itemKey, itemValue := range dict {
		count := __thSessionCount(session, itemKey)
		if min <= count && count <= max {
			for _, data := range value.List(itemValue) {
				listRes = append(listRes, data)
			}
		}
	}

	__thSessionSet(key, session, dict, ctx)
	return listRes
}

func throttleCheck(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		return list, nil
	}

	if val := ctx.GetX("__enable_alert_throttle__"); !value.IsTrue(val) {
		return list, nil
	}

	session := ""
	if val := ctx.GetX("__throttle_session__"); val != nil {
		session = value.ToStr(val)
	} else {
		return list, nil
	}

	return __thCheck(session, list, ctx), nil
}
