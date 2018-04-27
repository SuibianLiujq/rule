// Functions for workdays.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	"core/tools"
	"core/value"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func workdayCheckDateTimeItem(rangeDict, nowDict map[string]interface{}, nowStr string) bool {
	match := true

	if _, ok := rangeDict["datetime"]; ok {
		for _, item := range rangeDict["datetime"].([]interface{}) {
			if strings.HasPrefix(nowStr, value.ToStr(item)) {
				return true
			}
		}

		match = false
	}

	for key, item := range nowDict {
		if _, ok := rangeDict[key]; ok {
			if value.In(item, rangeDict[key]) {
				return false
			}
		}
	}

	return match
}

func workdayCheckDateTime(rangeList []interface{}, ctx *script.Cntx) (interface{}, error) {
	now, err := tools.GetTimeNow(ctx)
	if err != nil {
		return true, err
	}

	nowDict, nowStr := map[string]interface{}{}, now.Time.Format("2006-01-02 15:04:05")
	if now.Time.Hour() != 0 {
		nowDict["hours"] = int64(now.Time.Hour())
	}

	if now.Time.Weekday() != 0 {
		nowDict["week_days"] = int64(now.Time.Weekday())
	}

	if now.Time.Day() != 0 {
		nowDict["month_days"] = int64(now.Time.Day())
	}

	if now.Time.Month() != 0 {
		nowDict["months"] = int64(now.Time.Month())
	}

	firstOn, nowOn := true, true
	for _, rangeItem := range rangeList {
		rangeDict := rangeItem.(map[string]interface{})

		dateType := value.ToStr(rangeDict["type"])
		result := workdayCheckDateTimeItem(rangeDict, nowDict, nowStr)

		if dateType == "on" {
			if firstOn {
				nowOn = result
			} else if result {
				nowOn = true
			}
			firstOn = false
		} else if result {
			nowOn = false
		}
	}

	return nowOn, nil
}

func workdayStrToList(str string, ctx *script.Cntx) ([]interface{}, error) {
	strList := strings.Fields(value.ToStr(str))

	if len(strList) != 3 || strList[1] != "to" {
		msg := fmt.Sprintf("invalid range '%s'", value.ToStr(str))
		return nil, errors.New(msg)
	}

	fromValue, err := strconv.ParseInt(strList[0], 10, 64)
	if err != nil {
		msg := fmt.Sprintf("left value '%s' not INTEGER in '%s'", strList[0], value.ToStr(str))
		return nil, errors.New(msg)
	}

	toValue, err := strconv.ParseInt(strList[2], 10, 64)
	if err != nil {
		msg := fmt.Sprintf("right value '%s' not INTEGER in '%s'", strList[2], value.ToStr(str))
		return nil, errors.New(msg)
	}

	if fromValue > toValue {
		fromValue, toValue = toValue, fromValue
	}

	numList := []interface{}{}
	for cc := fromValue; cc <= toValue; cc++ {
		numList = append(numList, int64(cc))
	}

	return numList, nil
}

func workdayBuildRangeItem(v interface{}, ctx *script.Cntx) ([]interface{}, error) {
	numDict, list := map[int64]bool{}, []interface{}{}

	if listValue, ok := v.([]interface{}); ok {
		list = listValue
	} else {
		list = append(list, v)
	}

	for _, item := range list {
		switch item.(type) {
		case int64:
			numDict[item.(int64)] = true

		case string:
			listValue, err := workdayStrToList(item.(string), ctx)
			if err != nil {
				return nil, err
			}

			for _, listItem := range listValue {
				numDict[listItem.(int64)] = true
			}

		default:
			msg := fmt.Sprintf("'%s' not INTEGER or STRING", value.ToStr(item))
			return nil, errors.New(msg)
		}
	}

	numList := []interface{}{}
	if len(numDict) != 0 {
		for key, _ := range numDict {
			numList = append(numList, key)
		}
	}

	return numList, nil
}

func workdayBuildRangeList(list []interface{}, ctx *script.Cntx) ([]interface{}, error) {
	resultList := []interface{}{}
	for cc, dateItem := range list {
		dateDict, ok := dateItem.(map[string]interface{})
		if !ok {
			msg := fmt.Sprintf("'%s' not a DICT", value.ToStr(dateItem))
			return nil, errors.New(msg)
		}

		dateType := dateDict["type"]
		if dateType != "on" && dateType != "off" {
			msg := fmt.Sprintf("type '%s' not 'on or off' in item[%d]", dateType, cc)
			return nil, errors.New(msg)
		}

		resultItem := map[string]interface{}{"type": dateType}
		for rangeKey, rangeItem := range dateDict {
			switch rangeKey {
			case "type":
			case "datetime":
				listValue := []interface{}{}
				if rangeValue, ok := rangeItem.([]interface{}); ok {
					listValue = rangeValue
				} else {
					listValue = append(listValue, rangeItem)
				}

				for _, dtItem := range listValue {
					if value.IsStr(dtItem) {
						msg := fmt.Sprintf("datetime '%s' not STRING in item[%d]", value.ToStr(dtItem), cc)
						return nil, errors.New(msg)
					}
				}
				resultItem[rangeKey] = listValue

			case "months", "month_days", "week_days", "hours":
				listValue, err := workdayBuildRangeItem(rangeItem, ctx)
				if err != nil {
					msg := fmt.Sprintf("%s in item[%d]", err, cc)
					return nil, errors.New(msg)
				}

				resultItem[rangeKey] = listValue

			default:
				msg := fmt.Sprintf("range key '%s' not support in item[%d]", value.ToStr(rangeKey), cc)
				return nil, errors.New(msg)
			}
		}

		resultList = append(resultList, resultItem)
	}

	return resultList, nil
}

func funcWorkdayCheckDateTime(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	var workdayList []interface{}

	switch len(args) {
	case 0:
		workdayValue := ctx.GetX("__workdays__")
		if value.IsFalse(workdayValue) {
			return true, nil
		}

		listValue, ok := workdayValue.([]interface{})
		if !ok {
			msg := fmt.Sprintf("'__workdays__' not a LIST: %s", value.ToStr(workdayValue))
			return true, errors.New(msg)
		}
		workdayList = listValue
	case 1:
		if listValue, ok := args[0].([]interface{}); ok {
			workdayList = listValue
		} else {
			msg := fmt.Sprintf("args[0] '%s' not LIST", value.ToStr(args[0]))
			return true, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected <= 1 )", len(args))
		return nil, errors.New(msg)
	}

	rangeList, err := workdayBuildRangeList(workdayList, ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in '__workdays__'", err)
		return true, errors.New(msg)
	}

	result, err := workdayCheckDateTime(rangeList, ctx)
	return result, err
}
