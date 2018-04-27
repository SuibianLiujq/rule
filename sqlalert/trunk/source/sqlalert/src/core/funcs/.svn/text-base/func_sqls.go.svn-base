// Functions for build SQL statements dynamically.
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

func __sqlMakeCondSingle(k, v interface{}, ctx *script.Cntx) string {
	strKey, strValue := value.ToStr(k), value.ToStr(v)
	if value.IsStr(v) {
		strValue = "'" + strValue + "'"
	}

	if val := ctx.GetX("__enable_ip_to_sip_dip__"); value.IsTrue(val) && strKey == "ip" {
		return "sip = " + strValue + " OR dip = " + strValue
	}

	if val := ctx.GetX("__enable_port_to_sport_dport__"); value.IsTrue(val) && strKey == "port" {
		return "sport = " + strValue + " OR dport = " + strValue
	}

	return value.ToStr(k) + " = " + strValue
}

func __sqlMakeCondItem(label, idx, k, v interface{}, ctx *script.Cntx) string {
	if !value.IsList(v) {
		return __sqlMakeCondSingle(k, v, ctx)
	}

	strLabel, strKey := value.ToStr(label), value.ToStr(k)
	name := "__sql_expr_" + strLabel + "_" + strKey + "_" + value.ToStr(idx) + "__"

	ctx.SetGlobal(name, v)
	if dict, ok := ctx.GetXDict("__fields_ipfields__"); ok {
		if val, ok := dict[strKey]; ok && value.IsTrue(val) {
			if val := ctx.GetX("__enable_ip_to_sip_dip__"); value.IsTrue(val) && strKey == "ip" {
				return "(ip_ranges(sip, $(" + name + ")) OR ip_ranges(dip, $(" + name + ")))"
			} else {
				return "ip_ranges(" + strKey + ", $(" + name + "))"
			}
		}
	}

	if val := ctx.GetX("__enable_port_to_sport_dport__"); value.IsTrue(val) && strKey == "port" {
		return "(sport IN $(" + name + ") OR dport IN $(" + name + "))"
	}

	return strKey + " IN $(" + name + ")"
}

func __sqlMakeCond(label interface{}, listCond []interface{}, ctx *script.Cntx) string {
	listFilter := []interface{}{}

	for cc, item := range listCond {
		if value.IsStr(item) {
			listFilter = append(listFilter, item)
			continue
		}

		if value.IsDict(item) {
			list := []interface{}{}
			for k, v := range value.Dict(item) {
				list = append(list, __sqlMakeCondItem(label, cc, k, v, ctx))
			}

			if len(list) > 0 {
				res, _ := miscJoin([]interface{}{list, " AND "}, ctx)
				listFilter = append(listFilter, res)
			}
		}
	}

	if len(listFilter) == 0 {
		return ""
	} else if len(listFilter) == 1 {
		return value.ToStr(listFilter[0])
	}

	res, _ := miscJoin([]interface{}{listFilter, ") OR ("}, ctx)
	return "(" + value.ToStr(res) + ")"
}

func __sqlMakeWhitelist(ctx *script.Cntx) string {
	if list, ok := ctx.GetXList("__whitelist__"); ok {
		if str := __sqlMakeCond("whtelist", list, ctx); str != "" {
			return "NOT (" + str + ")"
		}
	}

	return ""
}

func __sqlMakeBlacklist(ctx *script.Cntx) string {
	if list, ok := ctx.GetXList("__blacklist__"); ok {
		return __sqlMakeCond("blacklist", list, ctx)
	}

	return ""
}

func __sqlMakeWhiteBlackList(ctx *script.Cntx) string {
	whitelist, blacklist := __sqlMakeWhitelist(ctx), __sqlMakeBlacklist(ctx)

	if whitelist != "" && blacklist != "" {
		return "(" + whitelist + ") OR (" + blacklist + ")"
	}

	if whitelist != "" {
		return whitelist
	}

	return blacklist
}

func __sqlMakeSelect(listRes []interface{}, ctx *script.Cntx) ([]interface{}, error) {
	list := ctx.GetX("__sql_select__")
	if list == nil {
		msg := fmt.Sprintf("variable '__sql_select__' not found")
		return nil, errors.New(msg)
	}

	if !value.IsList(list) {
		list = []interface{}{list}
	}

	str, _ := miscJoin([]interface{}{list, ","}, ctx)
	return append(listRes, "SELECT "+value.ToStr(str)), nil
}

func __sqlMakeFrom(listRes []interface{}, ctx *script.Cntx) ([]interface{}, error) {
	str, ok := ctx.GetXStr("__sql_from__")
	if !ok {
		msg := fmt.Sprintf("variable '__sql_from__' not a string")
		return nil, errors.New(msg)
	}

	return append(listRes, "FROM `"+str+"`"), nil
}

func __sqlMakeWhere(listRes []interface{}, ctx *script.Cntx) ([]interface{}, error) {
	listWhere, ok := []interface{}{}, false
	if val := ctx.GetX("__sql_where__"); val != nil {
		if listWhere, ok = value.AsList(val); !ok {
			listWhere = append(listWhere, val)
		}
	}

	if val := ctx.GetX("__sql_filter__"); val != nil {
		if list, ok := value.AsList(val); ok {
			tmp, _ := miscJoin([]interface{}{list, ") AND ("}, ctx)
			listWhere = append(listWhere, "("+value.Str(tmp)+")")
		} else {
			listWhere = append(listWhere, val)
		}
	}

	if len(value.List(listWhere)) == 0 {
		msg := fmt.Sprintf("variable '__sql_where__' is an empty list")
		return nil, errors.New(msg)
	}

	listWhere, whiteblack := value.CopyList(listWhere), __sqlMakeWhiteBlackList(ctx)
	if whiteblack != "" {
		listWhere = append(listWhere, whiteblack)
	}

	if len(listWhere) == 1 {
		return append(listRes, "WHERE "+value.ToStr(listWhere[0])), nil
	}

	res, _ := miscJoin([]interface{}{listWhere, ") AND ("}, ctx)
	return append(listRes, "WHERE ("+value.ToStr(res)+")"), nil
}

func __sqlMakeGroup(listRes []interface{}, ctx *script.Cntx) ([]interface{}, error) {
	val := ctx.GetX("__sql_group__")

	if str, ok := value.AsStr(val); ok {
		return append(listRes, "GROUP BY "+str), nil
	} else if value.IsList(val) {
		res, _ := miscJoin([]interface{}{val, ", "}, ctx)
		return append(listRes, "GROUP BY "+value.ToStr(res)), nil
	}

	return listRes, nil
}

func __sqlMakeLimit(listRes []interface{}, ctx *script.Cntx) ([]interface{}, error) {
	val := ctx.GetX("__sql_limit__")

	if str, ok := value.AsStr(val); ok {
		return append(listRes, "LIMIT "+str), nil
	} else if value.IsList(val) {
		res, _ := miscJoin([]interface{}{val, ", "}, ctx)
		return append(listRes, "LIMIT "+value.ToStr(res)), nil
	}

	return listRes, nil
}

func __sqlMakeOrder(listRes []interface{}, ctx *script.Cntx) ([]interface{}, error) {
	val := ctx.GetX("__sql_order__")

	if str, ok := value.AsStr(val); ok {
		return append(listRes, "ORDER BY "+str), nil
	} else if value.IsList(val) {
		res, _ := miscJoin([]interface{}{val, ", "}, ctx)
		return append(listRes, "ORDER BY "+value.ToStr(res)), nil
	}

	return listRes, nil
}

type __SQL_FUNC func([]interface{}, *script.Cntx) ([]interface{}, error)

var __g_sql_maker []__SQL_FUNC = []__SQL_FUNC{
	__sqlMakeSelect,
	__sqlMakeFrom,
	__sqlMakeWhere,
	__sqlMakeGroup,
	__sqlMakeLimit,
	__sqlMakeOrder,
}

func sqlMake(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	list, err := []interface{}{}, error(nil)

	for _, maker := range __g_sql_maker {
		if list, err = maker(list, ctx); err != nil {
			return nil, err
		}
	}

	return miscJoin([]interface{}{list, " "}, ctx)
}
