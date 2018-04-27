// Functions for elasticsearch.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/json"
	"core/script"
	"core/tools"
	"core/value"
	"errors"
	"es/client"
	"fmt"
)

// Global variables to search.
var g_host_query = []string{"__es_host_query__", "__es_host__"}
var g_host_insert = []string{"__es_host_insert__", "__es_host__"}
var g_index_regexp = "%[Y|M|D|h|m|s]"
var g_index_fmtstr = map[string]string{
	"%Y": "2006",
	"%M": "01",
	"%D": "02",
	"%h": "15",
	"%m": "04",
	"%s": "05",
	"%f": "999",
	"%z": "-07:00",
}

// elasticQuery() - Query elasticsearch.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func elasticQuery(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	var host, sql interface{}
	var filter string

	switch len(args) {
	case 1:
		host, sql, filter = ctx.GetFirstX(g_host_query), args[0], ""
	case 2:
		host, sql, filter = ctx.GetFirstX(g_host_query), args[0], value.ToStr(args[1])
		if str, ok := host.(string); !ok || str == "" {
			host = ctx.GetFirstX(g_host_query)
		}
	case 3:
		host, sql, filter = args[0], args[1], value.ToStr(args[2])
		if str, ok := host.(string); !ok || str == "" {
			host = ctx.GetFirstX(g_host_query)
		}
	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	if _, ok := host.(string); !ok || value.IsFalse(host) {
		msg := fmt.Sprintf("ES query host '%v' not found or invalid STRING", host)
		return nil, errors.New(msg)
	}

	if _, ok := sql.(string); !ok || value.IsFalse(sql) {
		msg := fmt.Sprintf("SQL '%v' not STRING or empty", sql)
		return nil, errors.New(msg)
	}

	if debug := ctx.GetX("__sql_debug_request__"); value.IsTrue(debug) {
		fmt.Println("----------------------- SQL Request -----------------------")
		fmt.Println(sql)
	}

	inst, err := client.NewSQLClient(sql, ctx)
	if err != nil {
		return nil, err
	}

	v, err := inst.Query(host.(string), ctx)
	if err != nil {
		return nil, err
	}

	if filter == "" || filter == "null" {
		return v, nil
	} else {
		return itemFilter([]interface{}{v, filter}, ctx)
	}
}

// elasticQueryAvgbyNum() - Query elasticsearch and divide number fields with given value.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func elasticQueryAvgbyNum(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	num, filter, ok := interface{}(nil), "", false

	switch len(args) {
	case 3:
		if filter, ok = value.AsStr(args[2]); !ok {
			msg := fmt.Sprintf("arg[2] not a string")
			return nil, errors.New(msg)
		}
		fallthrough

	case 2:
		num = args[1]
		if !value.IsNum(num) {
			msg := fmt.Sprintf("arg[1] not a number")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 3)", len(args))
		return nil, errors.New(msg)
	}

	listRes, err := elasticQuery(args[0:1], ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range value.List(listRes) {
		dictItem := value.Dict(item)
		for itemKey, itemValue := range dictItem {
			if value.IsNum(itemValue) {
				if res, err := value.Div(itemValue, num); err == nil {
					dictItem[itemKey] = res
				}
			}
		}
	}

	if filter != "" && filter != "null" {
		return itemFilter([]interface{}{listRes, filter}, ctx)
	}

	return listRes, err
}

// elasticQueryAvgbyField() - Query elasticsearch and divide number fields with given field.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func elasticQueryAvgbyField(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	str, filter, ok := "", "", false

	switch len(args) {
	case 3:
		if filter, ok = value.AsStr(args[2]); !ok {
			msg := fmt.Sprintf("arg[2] not a string")
			return nil, errors.New(msg)
		}
		fallthrough

	case 2:
		str, ok = value.AsStr(args[1])
		if !ok {
			msg := fmt.Sprintf("arg[1] not a string")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 3)", len(args))
		return nil, errors.New(msg)
	}

	listRes, err := elasticQuery(args[0:1], ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range value.List(listRes) {
		dictItem := value.Dict(item)
		if num, ok := dictItem[str]; ok && value.IsNum(num) {
			for itemKey, itemValue := range dictItem {
				if value.IsNum(itemValue) && itemKey != str {
					if res, err := value.Div(itemValue, num); err == nil {
						dictItem[itemKey] = res
					}
				}
			}
		}
	}

	if filter != "" && filter != "null" {
		return itemFilter([]interface{}{listRes, filter}, ctx)
	}

	return listRes, err
}

// elasticFormatIndex() - Format the index bytes from the context.
//
// @ctx: Script context.
func elasticFormatIndex(ctx *script.Cntx) ([]byte, error) {
	indexValue, ok := ctx.GetXDict("__es_index_alert__")
	if !ok || len(indexValue) == 0 {
		msg := fmt.Sprintf("'__es_index_alert__' not DICT or empty: %s", value.ToStr(indexValue))
		return nil, errors.New(msg)
	}

	indexIndex, ok := indexValue["index"]
	if !ok || !value.IsStr(indexIndex) || value.IsFalse(indexIndex) {
		msg := fmt.Sprintf("index['index'] not found or invalid STRING")
		return nil, errors.New(msg)
	}

	indexType, ok := indexValue["type"]
	if !ok || !value.IsStr(indexType) || value.IsFalse(indexType) {
		msg := fmt.Sprintf("index['type'] not found or invalid STRING")
		return nil, errors.New(msg)
	}

	indexStr := elasticFormatIndexString(indexIndex.(string), ctx)
	index := map[string]interface{}{"index": map[string]interface{}{"_index": indexStr, "_type": indexType.(string)}}
	return json.ToBytes(index)
}

func elasticReplaceFunc(repl string) string {
	if v, ok := g_index_fmtstr[repl]; ok {
		return v
	}
	return repl
}

// elasticFormatIndexString() - Format the index bytes from the context.
//
// @index: String value of index.
// @ctx: Script context.
func elasticFormatIndexString(index string, ctx *script.Cntx) string {
	if now, err := tools.GetTimeNow(ctx); err == nil {
		index = now.ToStr(index)
	}

	return index
}

// elasticAppendData() - Append __alert_append__ items into @data.
//
// @data: Data to append to.
// @ctx:  Script execution context.
func elasticAppendData(data interface{}, ctx *script.Cntx) interface{} {
	list, ok := data.([]interface{})
	if !ok {
		return data
	}

	extData, ok := ctx.GetXDict("__alert_append__")
	if !ok {
		return data
	}

	return tools.Append(list, extData)
}

func __elasticGetFieldFmt(k string, fields map[string]interface{}) string {
	if cfg, ok := fields[k]; ok {
		if dict, ok := value.AsDict(cfg); ok {
			if fmtValue, ok := dict["fmt"]; ok {
				if str, ok := value.AsStr(fmtValue); ok {
					return str
				}
			}
		}
	}

	return ""
}

func __elasticIsFmtFloat(v string) bool {
	switch v {
	case "float":
		fallthrough
	case "percentage":
		fallthrough
	case "pct":
		return true
	}

	return false
}

func __elasticFmtDataItem(data map[string]interface{}, fields map[string]interface{}) {
	for k, v := range data {
		fmtValue := __elasticGetFieldFmt(k, fields)

		if value.IsFloat(v) && !__elasticIsFmtFloat(fmtValue) {
			if intValue, err := value.ToInt(v); err == nil {
				data[k] = intValue
			}
		} else if value.IsStr(v) && fmtValue == "int" {
			if intValue, err := value.ToInt(v); err == nil {
				data[k] = intValue
			}
		}
	}
}

func __elasticFmtData(data interface{}, ctx *script.Cntx) interface{} {
	list, isList := value.AsList(data)
	fields, isDict := ctx.GetXDict("__fields_all__")

	if !isList || !isDict {
		return data
	}

	for _, item := range list {
		dict, ok := value.AsDict(item)
		if !ok {
			continue
		}

		__elasticFmtDataItem(dict, fields)
	}

	return data
}

func __elasticInsert(esClient *client.ESClient, url string, index []byte, data interface{}) {
	v, err := esClient.IndexBulk(url, index, data)
	if err != nil {
		fmt.Println(err)
		fmt.Println(v)
	}
}

// elasticInsert() - Indexing data to elasticsearch.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func elasticInsert(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	var host, data interface{}

	switch len(args) {
	case 1:
		host, data = ctx.GetFirstX(g_host_insert), args[0]
	case 2:
		if host, data = args[0], args[1]; !value.IsStr(host) || value.IsFalse(host) {
			host = ctx.GetFirstX(g_host_insert)
		}
	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	inst, err := client.NewESClient()
	if err != nil {
		return nil, err
	}

	index, err := elasticFormatIndex(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in formatIndex()", err)
		return nil, errors.New(msg)
	}

	data = elasticAppendData(data, ctx)
	data = __elasticFmtData(data, ctx)

	async, _ := ctx.GetXBool("__alert_async__")
	if async {
		go __elasticInsert(inst, "http://"+value.ToStr(host), index, data)
		return nil, err
	}

	return inst.IndexBulk("http://"+value.ToStr(host), index, data)
}
