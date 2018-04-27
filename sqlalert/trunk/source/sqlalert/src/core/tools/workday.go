// File: workday.go
//
// This file implements the workday check function.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package tools

import (
	"core/script"
	"core/sys"
	"core/value"
	"errors"
	"fmt"
	"strings"
)

// DTItem - Sructure of datetime item.
//
// @Type:     String value of 'on/off'.
// @DateTime: String value of time.
// @Lists:    A dict of number list.
type DTItem struct {
	Type     string
	DateTime []string
	Lists    map[string][]int64
}

// Init() - Initialize the DTItem instance.
func (this *DTItem) Init() *DTItem {
	this.Lists = make(map[string][]int64)
	return this
}

// DTNow - Struct of DTNow.
//
// @TimeStrList: List of time string value.
// @NumDict:     A dict of numbers.
type DTNow struct {
	TimeStrList []string
	NumDict     map[string]int64
}

// g_dt_fmtstr - All support time formatter.
var g_dt_fmtstr = []string{
	"%Y-%M-%D %h:%m:%s",
	"%M-%D %h:%m:%s",
	"%D %h:%m:%s",
	"%h:%m:%s",
}

// Key is the name of the datetime fields.
// Value is the min value of the datetime fields.
var g_dt_match_names = map[string]int64{
	"months":     1,
	"month_days": 1,
	"week_days":  1,
	"hours":      0,
	"minutes":    0,
}

// DTParseRange() - Parse string to range.
//
// @v: String like '1 to 7'.
//
// This function returns from, to, err three values.
func DTParseRange(v string) (int64, int64, error) {
	strList := strings.Fields(v)
	if len(strList) == 3 && strings.ToLower(strList[1]) == "to" {
		if from, err := value.ToInt(strList[0]); err == nil {
			if to, err := value.ToInt(strList[2]); err == nil {
				return from, to, nil
			}
		}
	}

	msg := fmt.Sprintf("invalid range '%s'", v)
	return 0, 0, errors.New(msg)
}

// DTConvertNumList() - Convert to list of number.
//
// @v: Value to convert.
func DTConvertNumList(v interface{}) ([]int64, error) {
	numDict := map[int64]bool{}

	if !value.IsList(v) {
		v = []interface{}{v}
	}

	for _, item := range value.List(v) {
		switch {
		case value.IsInt(item):
			numDict[value.Int(item)] = true

		case value.IsStr(item):
			from, to, err := DTParseRange(value.Str(item))
			if err != nil {
				msg := fmt.Sprintf("invalid range '%s'", value.Str(item))
				return nil, errors.New(msg)
			}

			for cc := from; cc <= to; cc++ {
				numDict[cc] = true
			}

		default:
			msg := fmt.Sprintf("invalid range '%s'", value.ToStr(item))
			return nil, errors.New(msg)
		}
	}

	numList := []int64{}
	for key, _ := range numDict {
		numList = append(numList, key)
	}

	return numList, nil
}

// DTConvertType() - Check the type of datetime item.
//
// @v: Value to test.
func DTConvertType(v interface{}) (string, error) {
	str, ok := value.AsStr(v)
	if !ok {
		return "", errors.New("type not a STR")
	}

	switch str {
	case "on", "off":
		return str, nil
	}

	msg := fmt.Sprintf("invalid type '%s'", str)
	return "", errors.New(msg)
}

// DTConvertTimeStr() - Convert 'datetime' to a list of string.
//
// @v: Value to convert.
func DTConvertTimeStr(v interface{}) ([]string, error) {
	strList := []string{}

	if !value.IsList(v) {
		v = []interface{}{v}
	}

	for _, item := range value.List(v) {
		str, ok := value.AsStr(item)
		if !ok {
			msg := fmt.Sprintf("'%s' not STR", value.ToStr(item))
			return nil, errors.New(msg)
		}

		strList = append(strList, str)
	}

	return strList, nil
}

// DTConvertItem() - Convert datetime item.
//
// @dict: Dict value of datetime item.
func DTConvertItem(dict map[string]interface{}) (*DTItem, error) {
	dtItem := (&DTItem{}).Init()

	for key, item := range dict {
		switch key {
		case "type":
			str, err := DTConvertType(item)
			if err != nil {
				return nil, err

			}
			dtItem.Type = str

		case "datetime":
			strList, err := DTConvertTimeStr(item)
			if err != nil {
				return nil, err
			}
			dtItem.DateTime = strList

		default:
			if value.In(key, g_dt_match_names) {
				numList, err := DTConvertNumList(item)
				if err != nil {
					msg := fmt.Sprintf("%s %s", key, err)
					return nil, errors.New(msg)
				}
				dtItem.Lists[key] = numList
			} else {
				msg := fmt.Sprintf("unsupport key '%s'", key)
				return nil, errors.New(msg)
			}
		}
	}

	return dtItem, nil
}

// DTConvertList() - Convert confiugred datetime list to a readable format.
//
// @v: Value to convert.
func DTConvertList(v interface{}) ([]*DTItem, error) {
	if !value.IsList(v) {
		return nil, errors.New("not a LIST")
	}

	dtList := []*DTItem{}
	for cc, listItem := range value.List(v) {
		dict, ok := value.AsDict(listItem)
		if !ok {
			msg := fmt.Sprintf("item[%d] not a DICT", cc)
			return nil, errors.New(msg)
		}

		dtItem, err := DTConvertItem(dict)
		if err != nil {
			msg := fmt.Sprintf("%s in item[%d]", err, cc)
			return nil, errors.New(msg)
		}

		dtList = append(dtList, dtItem)
	}

	return dtList, nil
}

// DTConvertNow() - Convert now value to DTNow instance.
//
// @now: Timestamp of now.
func DTConvertNow(now *sys.Time) (*DTNow, error) {
	dtNow := &DTNow{}

	for _, item := range g_dt_fmtstr {
		dtNow.TimeStrList = append(dtNow.TimeStrList, now.ToStr(item))
	}

	dtNow.NumDict = make(map[string]int64)
	for key, item := range g_dt_match_names {
		intValue := now.Get(key)
		if intValue >= item {
			dtNow.NumDict[key] = intValue
		}
	}

	return dtNow, nil
}

// CheckTimeStr() - Test whether time str in str list of now.
//
// @str:     Time string.
// @strList: List of string.
func DTCheckTimeStr(str string, strList []string) bool {
	for _, item := range strList {
		if strings.HasPrefix(item, str) {
			return true
		}
	}
	return false
}

// DTNumIn() - Test whether a num in a list.
//
// @num:     Value to check.
// @numList: List of number.
func DTNumIn(num int64, numList []int64) bool {
	for _, item := range numList {
		if num == item {
			return true
		}
	}
	return false
}

// DTCheckItem() - Test whether now is in item of datetime list.
//
// @dtNow:  Instance of DTNow.
// @dtItem: Instance of DTItem.
func DTCheckItem(dtNow *DTNow, dtItem *DTItem) bool {
	match := true

	if len(dtItem.DateTime) != 0 {
		for _, item := range dtItem.DateTime {
			if DTCheckTimeStr(item, dtNow.TimeStrList) {
				return true
			}
		}

		match = false
	}

	for key, item := range dtNow.NumDict {
		if numList, ok := dtItem.Lists[key]; ok {
			if !DTNumIn(item, numList) {
				return false
			}
		}
	}

	return match
}

// DTCheckList() - Check whether now is 'on' in the datetime list.
//
// @now:    Timestamp of now.
// @dtList: List of DTItem instance.
func DTCheckList(now *sys.Time, dtList []*DTItem) (bool, error) {
	dtNow, err := DTConvertNow(now)
	if err != nil {
		msg := fmt.Sprintf("invalid now '%d'", now)
		return false, errors.New(msg)
	}

	firstOn, nowOn := true, true
	for _, item := range dtList {
		result := DTCheckItem(dtNow, item)
		if item.Type == "on" {
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

// DTCheckNow() - Check whether now is 'on' in the datetime list.
//
// @now: Timestamp of now.
// @v:   Configuration of __workdays__.
func DTCheckNow(now, cfg interface{}) (bool, error) {
	dtList, err := DTConvertList(cfg)
	if err != nil {
		return false, err
	}

	nowTime, ok := now.(*sys.Time)
	if !ok {
		if nowTime, err = GetTime(now, nil); err != nil {
			return false, err
		}
	}

	return DTCheckList(nowTime, dtList)
}

// DTCheck() - Check whether 'now' is 'on' in the datetime list.
//
// @v:   Configuration of __workdays__.
// @ctx: Script context.
func DTCheck(now, cfg interface{}, ctx *script.Cntx) (bool, error) {
	nowTime, err := GetTime(now, ctx)
	if err != nil {
		return false, err
	}

	if cfg == nil {
		if res, ok := ctx.GetXList("__workdays__"); !ok {
			return true, nil
		} else {
			cfg = res
		}
	}

	return DTCheckNow(nowTime, cfg)
}
