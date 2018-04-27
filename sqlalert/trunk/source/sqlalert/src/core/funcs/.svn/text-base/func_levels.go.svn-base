// Functions for alert levels.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	"core/value"
)

func __levelItem(args []interface{}, ctx *script.Cntx) map[string]interface{} {
	level := ""

	if len(args) > 0 {
		level = value.ToStr(args[0])
	}

	if level == "" {
		if v := ctx.GetX("__level__"); v != nil {
			level = value.ToStr(v)
		} else {
			level = "INFO"
		}
	}

	if dictLevels, ok := ctx.GetXDict("__levels_all__"); ok {
		if item, ok := dictLevels[level]; ok {
			if val, ok := value.AsDict(item); ok {
				dict := value.CopyDict(val)
				if _, ok := dict["name"]; !ok {
					dict["name"] = level
				}

				if _, ok := dict["desc"]; !ok {
					dict["desc"] = level
				}

				return dict
			}
		}
	}

	return map[string]interface{}{"name": level, "desc": level}
}

func __levelHighlight(val, color interface{}) interface{} {
	if val == nil || color == nil {
		return val
	}

	return "<font color=\"" + value.ToStr(color) + "\">" + value.ToStr(val) + "</font>"
}

func levelName(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return __levelItem(args, ctx)["name"], nil
}

func levelNameHighlight(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	item := __levelItem(args, ctx)
	color, _ := item["color"]

	return __levelHighlight(item["name"], color), nil
}

func levelDesc(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	return __levelItem(args, ctx)["desc"], nil
}

func levelDescHighlight(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	item := __levelItem(args, ctx)
	color, _ := item["color"]

	return __levelHighlight(item["desc"], color), nil
}

func levelNameDesc(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	item := __levelItem(args, ctx)
	return value.ToStr(item["name"]) + " (" + value.ToStr(item["desc"]) + ")", nil
}

func levelNameDescHighlight(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	item := __levelItem(args, ctx)
	color, _ := item["color"]

	return __levelHighlight(value.ToStr(item["name"])+" ("+value.ToStr(item["desc"])+")", color), nil
}
