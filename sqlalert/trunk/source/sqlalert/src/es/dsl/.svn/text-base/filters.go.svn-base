// Compile SQL tokens to DSL filters.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package dsl

import (
	"core/script"
	"core/sql"
	"core/value"
	"errors"
	"fmt"
	"strings"
)

// Static filter 'match_all'.
var FILTER_MATCHALL = map[string]interface{}{"match_all": map[string]interface{}{}}

// Filter Type.
type FilterType int

// All filter types.
const (
	FILTER_SCRIPT FilterType = iota
	FILTER_TERM
	FILTER_TERMS
	FILTER_RANGE
	FILTER_MUST
	FILTER_MUSTNOT
	FILTER_SHOULD
	FILTER_LOGICAL
	FILTER_QUERY
)

// Names of DSL filters.
var g_filter_names = [...]string{
	FILTER_SCRIPT:  "script",
	FILTER_TERM:    "term",
	FILTER_TERMS:   "terms",
	FILTER_RANGE:   "range",
	FILTER_MUST:    "must",
	FILTER_MUSTNOT: "must_not",
	FILTER_SHOULD:  "should",
	FILTER_QUERY:   "query_string",
}

// String() - Returns name of filter type.
//
// This function returns the string name of filter type.
func (this FilterType) Str() string { return g_filter_names[this] }

// Filter - Structure of filter.
//
// @Type:  Filter type.
// @Value: Filter value.
type Filter struct {
	Type  FilterType
	Value interface{}
}

// Init() - Initialize Filter instance.
//
// @t: Filter type.
// @v: Filter value.
//
// This function return Filter instance itself for chain operation.
func (this *Filter) Init(t FilterType, v interface{}) (*Filter, error) {
	this.Type, this.Value = t, v
	return this, nil
}

// DslValue() - Returns the DSL structure.
func (this *Filter) DslValue() interface{} {
	var v interface{}

	switch this.Value.(type) {
	case *Filter:
		v = this.Value.(*Filter).DslValue()

	case []*Filter:
		list := []interface{}{}
		for _, filter := range this.Value.([]*Filter) {
			list = append(list, filter.DslValue())
		}
		v = list

	default:
		v = this.Value
	}

	v = map[string]interface{}{this.Type.Str(): v}
	switch this.Type {
	case FILTER_MUST, FILTER_MUSTNOT, FILTER_SHOULD:
		v = map[string]interface{}{"bool": v}
	}

	return v
}

// filterScript() - Compile SQL token to ES DSL script filter.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterScript(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	var sv *ScriptValue
	var err error

	switch token.Type() {
	case sql.T_COMP, sql.T_COND:
		if sv, err = ScriptBuild(dsl, token, ctx); err != nil {
			return nil, err
		}

	default:
		msg := fmt.Sprintf("'%s' not filter", token.Str())
		return nil, errors.New(msg)
	}

	return (&Filter{}).Init(FILTER_SCRIPT, sv.AsField(dsl.Script))
}

// filterIn() - Compile SQL token to ES DSL terms filter.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterIn(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	t := token.(*sql.TokenIn)

	if t.Key.Type() != sql.T_IDENT && t.Key.Type() != sql.T_STR {
		msg := fmt.Sprintf("key '%s' not STR", t.Key.Str())
		return nil, errors.New(msg)
	}

	key := t.Key.Str()
	object, err := execToken(t.Object, ctx)
	if err != nil {
		return nil, err
	}

	if !value.IsList(object) {
		msg := fmt.Sprintf("%s (%s) not LIST", value.ToStr(object), t.Object.Str())
		return nil, errors.New(msg)
	}

	filter := map[string]interface{}{key: object}
	return (&Filter{}).Init(FILTER_TERMS, filter)
}

// filterTerm() - Compile SQL token to ES DSL term filter.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterTerm(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	t := token.(*sql.TokenComp)

	switch {
	case t.Left.Type() == sql.T_IDENT && t.Right.Type() != sql.T_IDENT:
		fallthrough
	case t.Left.Type() != sql.T_IDENT && t.Right.Type() == sql.T_IDENT:
		leftToken, rightToken := t.Left, t.Right
		if rightToken.Type() == sql.T_IDENT {
			leftToken, rightToken = t.Right, t.Left
		}

		leftValue := leftToken.(*sql.TokenIdent).Value
		rightValue, err := ScriptBuild(dsl, rightToken, ctx)
		if err != nil {
			return nil, err
		}

		if rightValue.Type == SCRIPT_VALUE {
			filter := map[string]interface{}{leftValue: rightValue.Value}
			return (&Filter{}).Init(FILTER_TERM, filter)
		}
	}

	return filterScript(dsl, token, ctx)
}

// Reversed operator of compare operators.
var g_operator_reverses = map[sql.TokenType]sql.TokenType{
	sql.T_EQ: sql.T_NE,
	sql.T_NE: sql.T_EQ,
	sql.T_LT: sql.T_GT,
	sql.T_LE: sql.T_GE,
	sql.T_GT: sql.T_LT,
	sql.T_GE: sql.T_LE,
}

// Names of compare operators in ES range filter.
var g_operator_names = map[sql.TokenType]string{
	sql.T_EQ: "eq",
	sql.T_NE: "ne",
	sql.T_LT: "lt",
	sql.T_LE: "lte",
	sql.T_GT: "gt",
	sql.T_GE: "gte",
}

func filterRange(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	t := token.(*sql.TokenComp)

	switch {
	case t.Left.Type() == sql.T_IDENT && t.Right.Type() != sql.T_IDENT:
		fallthrough
	case t.Left.Type() != sql.T_IDENT && t.Right.Type() == sql.T_IDENT:
		operator := t.Operator

		leftToken, rightToken := t.Left, t.Right
		if rightToken.Type() == sql.T_IDENT {
			leftToken, rightToken = t.Right, t.Left

			if opt, ok := g_operator_reverses[operator]; !ok {
				msg := fmt.Sprintf("invalid operator '%s' in %s", operator.Name(), token.Str())
				return nil, errors.New(msg)
			} else {
				operator = opt
			}
		}

		leftValue := leftToken.(*sql.TokenIdent).Value
		rightValue, err := ScriptBuild(dsl, rightToken, ctx)
		if err != nil {
			return nil, err
		}

		if rightValue.Type == SCRIPT_VALUE {
			optName, ok := g_operator_names[operator]
			if !ok {
				msg := fmt.Sprintf("invalid operator '%s' in %s", t.Operator.Name(), token.Str())
				return nil, errors.New(msg)
			}

			filter := map[string]interface{}{leftValue: map[string]interface{}{optName: rightValue.Value}}
			return (&Filter{}).Init(FILTER_RANGE, filter)
		}
	}

	return filterScript(dsl, token, ctx)
}

// filterComparison() - Compile SQL COMPAIRSON token to ES DSL filter.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterComparison(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	t := token.(*sql.TokenComp)

	switch t.Operator {
	case sql.T_EQ, sql.T_NE:
		filter, err := filterTerm(dsl, token, ctx)
		if err != nil {
			return nil, err
		}

		if t.Operator == sql.T_NE {
			filter, _ = (&Filter{}).Init(FILTER_MUSTNOT, filter)
		}

		return filter, nil

	case sql.T_LT, sql.T_LE, sql.T_GT, sql.T_GE:
		return filterRange(dsl, token, ctx)
	}

	msg := fmt.Sprintf("invalid operator in '%s'", token.Str())
	return nil, errors.New(msg)
}

// filterLogical() - Compile SQL token LOGICAL to ES DSL filter.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterLogical(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	t := token.(*sql.TokenLogical)

	var left *Filter = nil
	if t.Left != nil {
		if filter, err := FilterBuild(dsl, t.Left, ctx); err == nil {
			left = filter
		} else {
			return nil, err
		}
	}

	var right *Filter = nil
	if t.Right != nil {
		if filter, err := FilterBuild(dsl, t.Right, ctx); err == nil {
			right = filter
		} else {
			return nil, err
		}
	}

	switch t.Operator {
	case sql.T_AND:
		return (&Filter{}).Init(FILTER_MUST, []*Filter{left, right})

	case sql.T_OR:
		return (&Filter{}).Init(FILTER_SHOULD, []*Filter{left, right})

	case sql.T_NOT:
		return (&Filter{}).Init(FILTER_MUSTNOT, []*Filter{right})
	}

	msg := fmt.Sprintf("invalid operator in '%s'", token.Str())
	return nil, errors.New(msg)
}

// filterFuncLast() - Compile LAST(...) function.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFuncLast(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Filter, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("last() argument mismatch %d/1", len(args))
		return nil, errors.New(msg)
	}

	interval, err := evalTokenTimeInterval(dsl, args[0], ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last()", err)
		return nil, errors.New(msg)
	}

	timeNow, err := evalTimeNow(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last()", err)
		return nil, errors.New(msg)
	}

	timeRange := map[string]interface{}{}
	timeRange["gte"] = timeNow - interval
	timeRange["lte"] = timeNow
	timeRange["format"] = string("epoch_millis")

	filter := map[string]interface{}{"@timestamp": timeRange}
	return &Filter{Type: FILTER_RANGE, Value: filter}, nil
}

// filterFuncLastDays() - Compile LAST_DAYS(...) function.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFuncLastDays(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Filter, error) {
	var dayToken, intervalToken sql.Token

	switch len(args) {
	case 1:
		dayToken, intervalToken = args[0], nil
	case 2:
		dayToken, intervalToken = args[0], args[1]
	default:
		msg := fmt.Sprintf("last() argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	if dayToken.Type() != sql.T_INT {
		msg := fmt.Sprintf("'%s' not a INTEGER number token", dayToken.Str())
		return nil, errors.New(msg)
	}

	days := dayToken.(*sql.TokenInt).Value
	if intervalToken == nil {
		interval := int64(24 * 60 * 60 * 1000)

		timeNow, err := evalTimeNowDay(ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in last()", err)
			return nil, errors.New(msg)
		}

		timeRange := map[string]interface{}{}
		timeRange["gte"] = timeNow - (interval * days)
		timeRange["lte"] = timeNow
		timeRange["format"] = string("epoch_millis")

		filter := map[string]interface{}{"@timestamp": timeRange}
		return &Filter{Type: FILTER_RANGE, Value: filter}, nil
	}

	interval, err := evalTokenTimeInterval(dsl, intervalToken, ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last()", err)
		return nil, errors.New(msg)
	}

	timeNow, err := evalTimeNow(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last()", err)
		return nil, errors.New(msg)
	}

	filterList, usPerDay := []*Filter{}, int64(24*60*60*1000)
	for cc := int64(1); cc < days+1; cc++ {
		timeRange := map[string]interface{}{}
		timeRange["gte"] = timeNow - (cc * usPerDay) - interval
		timeRange["lte"] = timeNow - (cc * usPerDay)
		timeRange["format"] = string("epoch_millis")

		filter, _ := (&Filter{}).Init(FILTER_RANGE, map[string]interface{}{"@timestamp": timeRange})
		filterList = append(filterList, filter)
	}

	return (&Filter{}).Init(FILTER_SHOULD, filterList)
}

// filterFuncLastWorkdays() - Compile last_workdays(...) function.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFuncLastWorkdays(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Filter, error) {
	var dayToken, intervalToken sql.Token

	switch len(args) {
	case 1:
		dayToken, intervalToken = args[0], nil
	case 2:
		dayToken, intervalToken = args[0], args[1]
	default:
		msg := fmt.Sprintf("last_workdays() argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	if dayToken.Type() != sql.T_INT {
		msg := fmt.Sprintf("'%s' not a INTEGER number token", dayToken.Str())
		return nil, errors.New(msg)
	}

	days := dayToken.(*sql.TokenInt).Value
	if intervalToken == nil {
		interval := int64(24 * 60 * 60 * 1000)

		timeNow, err := evalTimeNowDay(ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in last_workdays()", err)
			return nil, errors.New(msg)
		}

		gotDays, filterList := int64(0), []*Filter{}
		for gotDays < days {
			timeNow = timeNow - interval

			if evalCheckDatetime(timeNow, ctx) {
				timeRange := map[string]interface{}{}
				timeRange["gt"] = timeNow
				timeRange["lt"] = timeNow + interval
				timeRange["format"] = string("epoch_millis")

				filter, _ := (&Filter{}).Init(FILTER_RANGE, map[string]interface{}{"@timestamp": timeRange})
				filterList = append(filterList, filter)
				gotDays += 1
			}
		}

		return (&Filter{}).Init(FILTER_SHOULD, filterList)
	}

	interval, err := evalTokenTimeInterval(dsl, intervalToken, ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last_workdays()", err)
		return nil, errors.New(msg)
	}

	timeNow, err := evalTimeNow(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last_workdays()", err)
		return nil, errors.New(msg)
	}

	timeNowDay, err := evalTimeNowDay(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last_workdays()", err)
		return nil, errors.New(msg)
	}

	timeNowInterval := timeNow - timeNowDay
	gotDays, filterList, usPerDay := int64(0), []*Filter{}, int64(24*60*60*1000)
	for gotDays < days {
		timeNowDay = timeNowDay - usPerDay

		if evalCheckDatetime(timeNowDay, ctx) {
			timeRange := map[string]interface{}{}
			timeRange["gte"] = timeNowDay + timeNowInterval - interval
			timeRange["lte"] = timeNowDay + +timeNowInterval
			timeRange["format"] = string("epoch_millis")

			filter, _ := (&Filter{}).Init(FILTER_RANGE, map[string]interface{}{"@timestamp": timeRange})
			filterList = append(filterList, filter)
			gotDays += 1
		}
	}

	return (&Filter{}).Init(FILTER_SHOULD, filterList)
}

// filterFuncLastWeekdays() - Compile last_weekdays(...) function.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFuncLastWeekdays(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Filter, error) {
	var dayToken, intervalToken sql.Token

	switch len(args) {
	case 1:
		dayToken, intervalToken = args[0], nil
	case 2:
		dayToken, intervalToken = args[0], args[1]
	default:
		msg := fmt.Sprintf("last_weekdays() argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	if dayToken.Type() != sql.T_INT {
		msg := fmt.Sprintf("'%s' not a INTEGER number token", dayToken.Str())
		return nil, errors.New(msg)
	}

	days := dayToken.(*sql.TokenInt).Value
	if intervalToken == nil {
		interval := int64(7 * 24 * 60 * 60 * 1000)

		timeNow, err := evalTimeNowDay(ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in last_weekdays()", err)
			return nil, errors.New(msg)
		}

		timeRange := map[string]interface{}{}
		timeRange["gte"] = timeNow - (interval * days)
		timeRange["lte"] = timeNow
		timeRange["format"] = string("epoch_millis")

		filter := map[string]interface{}{"@timestamp": timeRange}
		return &Filter{Type: FILTER_RANGE, Value: filter}, nil
	}

	interval, err := evalTokenTimeInterval(dsl, intervalToken, ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last_weekdays()", err)
		return nil, errors.New(msg)
	}

	timeNow, err := evalTimeNow(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in last_weekdays()", err)
		return nil, errors.New(msg)
	}

	filterList, usPerDay := []*Filter{}, int64(7*24*60*60*1000)
	for cc := int64(1); cc < days+1; cc++ {
		timeRange := map[string]interface{}{}
		timeRange["gte"] = timeNow - (cc * usPerDay) - interval
		timeRange["lte"] = timeNow - (cc * usPerDay)
		timeRange["format"] = string("epoch_millis")

		filter, _ := (&Filter{}).Init(FILTER_RANGE, map[string]interface{}{"@timestamp": timeRange})
		filterList = append(filterList, filter)
	}

	return (&Filter{}).Init(FILTER_SHOULD, filterList)
}

// __buildIPRangeStr() - Build query_string of IP ranges.
//
// @field:  Filed name.
// @ranges: List of IP ranges.
func __buildIPRangeStr(field string, ranges []interface{}) (interface{}, error) {
	rangeList := []string{}

	for _, item := range ranges {
		switch item.(type) {
		case []interface{}:
			list := value.List(item)
			if len(list) != 2 || !(value.IsStr(list[0]) && value.IsStr(list[1])) {
				msg := fmt.Sprintf("%s invalid range", value.ToStr(item))
				return nil, errors.New(msg)
			}

			rangeList = append(rangeList, fmt.Sprintf("%s: [%s TO %s]", field, list[0], list[1]))

		case string:
			strValue := value.Str(item)

			list := strings.Fields(strValue)
			if len(list) == 3 && (list[1] == "TO" || list[1] == "to") {
				rangeList = append(rangeList, fmt.Sprintf("%s: [%s TO %s]", field, list[0], list[2]))
			} else {
				from, to, err := evalIPRangeStr(strValue)
				if err != nil {
					return nil, err
				} else {
					rangeList = append(rangeList, fmt.Sprintf("%s: [%s TO %s]", field, from, to))
				}
			}
		}
	}

	return strings.Join(rangeList, " OR "), nil
}

// filterFuncIPRange() - Compile ip_range(...) function.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFuncIPRange(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Filter, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("ip_range() argument mismatch %d (expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	fieldValue, err := execToken(args[0], ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in ip_range()", err)
		return nil, errors.New(msg)
	}

	fieldStr, ok := value.AsStr(fieldValue)
	if !ok {
		msg := fmt.Sprintf("args[0] not a STR value: %s in ip_range()", args[0].Str())
		return nil, errors.New(msg)
	}

	ranges := []interface{}{}
	for _, token := range args[1:] {
		if itemValue, err := execToken(token, ctx); err != nil {
			msg := fmt.Sprintf("%s in ip_range()", err)
			return nil, errors.New(msg)
		} else {
			ranges = append(ranges, itemValue)
		}
	}

	queryStr, err := __buildIPRangeStr(fieldStr, ranges)
	if err != nil {
		msg := fmt.Sprintf("%s in ip_range()", err)
		return nil, errors.New(msg)
	}

	return (&Filter{}).Init(FILTER_QUERY, map[string]interface{}{"query": queryStr})
}

// filterFuncIPRanges() - Compile ip_ranges(...) function.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFuncIPRanges(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Filter, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("ip_range() argument mismatch %d (expected >= 1)", len(args))
		return nil, errors.New(msg)
	}

	fieldValue, err := execToken(args[0], ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in ip_range()", err)
		return nil, errors.New(msg)
	}

	fieldStr, ok := value.AsStr(fieldValue)
	if !ok {
		msg := fmt.Sprintf("args[0] not a STR value: %s in ip_range()", args[0].Str())
		return nil, errors.New(msg)
	}

	ranges := []interface{}{}
	for _, token := range args[1:] {
		itemValue, err := execToken(token, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in ip_range()", err)
			return nil, errors.New(msg)
		}

		list, ok := value.AsList(itemValue)
		if !ok {
			msg := fmt.Sprintf("%s not a list", value.ToStr(itemValue))
			return nil, errors.New(msg)
		}

		for _, rangeValue := range list {
			ranges = append(ranges, rangeValue)
		}
	}

	queryStr, err := __buildIPRangeStr(fieldStr, ranges)
	if err != nil {
		msg := fmt.Sprintf("%s in ip_range()", err)
		return nil, errors.New(msg)
	}

	return (&Filter{}).Init(FILTER_QUERY, map[string]interface{}{"query": queryStr})
}

// filterFuncQueryString() - Compile query_string(...) function.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFuncQueryString(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Filter, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("query_string() argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	tokenValue, err := execToken(args[0], ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in query_string()", err)
		return nil, errors.New(msg)
	}

	if !value.IsStr(tokenValue) {
		msg := fmt.Sprintf("args[0] not a STR value: %s in query_string()", args[0].Str())
		return nil, errors.New(msg)
	}

	return (&Filter{}).Init(FILTER_QUERY, map[string]interface{}{"query": tokenValue})
}

// All supported filter functions.
var g_filter_functions = map[string]func(*Dsl, []sql.Token, *script.Cntx) (*Filter, error){
	"last":          filterFuncLast,
	"last_days":     filterFuncLastDays,
	"last_workdays": filterFuncLastWorkdays,
	"last_weeks":    filterFuncLastWeekdays,
	"ip_range":      filterFuncIPRange,
	"ip_ranges":     filterFuncIPRanges,
	"query_string":  filterFuncQueryString,
}

// filterFunction() - Compile SQL FUNCTION token to DSL filters.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func filterFunction(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	t := token.(*sql.TokenFunc)

	name := strings.ToLower(t.Name)
	if function, ok := g_filter_functions[name]; ok {
		return function(dsl, t.List, ctx)
	}

	if filter, err := filterScript(dsl, token, ctx); err == nil {
		return filter, nil
	}

	msg := fmt.Sprintf("filter function '%s' not support", name)
	return nil, errors.New(msg)
}

// FilterBuild() - Compile SQL token to ES DSL filter.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func FilterBuild(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Filter, error) {
	switch token.Type() {
	case sql.T_COMP:
		return filterComparison(dsl, token, ctx)

	case sql.T_LOGICAL:
		return filterLogical(dsl, token, ctx)

	case sql.T_FUNC:
		return filterFunction(dsl, token, ctx)

	case sql.T_COND:
		return filterScript(dsl, token, ctx)

	case sql.T_IN:
		return filterIn(dsl, token, ctx)
	}

	msg := fmt.Sprintf("unknown filter '%s'", token.Str())
	return nil, errors.New(msg)
}
