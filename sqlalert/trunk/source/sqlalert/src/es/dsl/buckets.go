// File: buckets.go
//
// This file implements the compiler of DSL buckets.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package dsl

import (
	"core/script"
	"core/sql"
	"core/sys"
	"core/value"
	"errors"
	"fmt"
	"strings"
)

// Bucket type.
type BucketType int

// All bucket types.
const (
	BUCKET_TERMS BucketType = iota
	BUCKET_RANGE
	BUCKET_IP_RANGE
	BUCKET_DATE_RANGE
	BUCKET_HISTOGRAM
	BUCKET_DATE_HISTOGRAM
	BUCKET_FILTERS
)

// Names of DSL buckets.
var g_bucket_names = [...]string{
	BUCKET_TERMS:          "terms",
	BUCKET_RANGE:          "range",
	BUCKET_IP_RANGE:       "ip_range",
	BUCKET_DATE_RANGE:     "date_range",
	BUCKET_HISTOGRAM:      "histogram",
	BUCKET_DATE_HISTOGRAM: "date_histogram",
	BUCKET_FILTERS:        "filters",
}

var g_bucket_value_keys = []string{
	"key_as_string",
	"key",
}

// String() - Returns name of bucket type.
//
// This function returns the string name of bucket type.
func (this BucketType) Str() string { return g_bucket_names[this] }

// Bucket - Structure of Bucket.
//
// @Type:
// @Name:
// @Value:
type Bucket struct {
	Type  BucketType
	Name  string
	Value map[string]interface{}
}

// Init() Initialize Bucket instance.
//
// @t: Bucket type.
// @n: Bucket name.
// @v: Bucket value.
//
// This function returns Bucket instance itself for chain operation.
func (this *Bucket) Init(t BucketType, n string, v map[string]interface{}) (*Bucket, error) {
	this.Type, this.Name, this.Value = t, n, v
	return this, nil
}

// DslValue() - Returns the DSL value of the bucket.
func (this *Bucket) DslValue(limit int64, order interface{}) map[string]interface{} {
	switch this.Type {
	case BUCKET_TERMS:
		this.Value["size"] = limit
	}

	if order != nil {
		this.Value["order"] = order
	}

	return map[string]interface{}{this.Type.Str(): this.Value}
}

// SelectValue() - Select response value.
//
// @v: Response value.
func (this *Bucket) SelectValue(v interface{}) (interface{}, error) {
	dict, ok := value.AsDict(v)
	if !ok {
		msg := fmt.Sprintf("'%s' not a DICT", value.ToStr(v))
		return nil, errors.New(msg)
	}

	for _, item := range g_bucket_value_keys {
		if itemValue, ok := dict[item]; ok {
			return value.ToStr(itemValue), nil
		}
	}

	switch this.Type {
	case BUCKET_RANGE, BUCKET_IP_RANGE, BUCKET_DATE_RANGE:
		from, okFrom := dict["from"]
		to, okTo := dict["to"]

		if !okFrom && !okTo {
			msg := fmt.Sprintf("'from && to' not found in '%s'", value.ToStr(dict))
			return nil, errors.New(msg)
		}

		keyStr := "*"
		if from != nil {
			keyStr = value.ToStr(from)
		}
		keyStr += " to "

		if to != nil {
			keyStr += value.ToStr(to)
		} else {
			keyStr += "*"
		}

		return keyStr, nil
	}

	return v, nil
}

// bucketTerms() - Compile SQL token to ES terms bucket.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func bucketTerms(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Bucket, error) {
	sv, err := ScriptBuild(dsl, token, ctx)
	if err != nil {
		return nil, err
	}

	switch sv.Type {
	case SCRIPT_FIELD, SCRIPT_VALUE:
		if str, ok := sv.Value.(string); ok {
			n, v := str, map[string]interface{}{"field": str}
			return (&Bucket{}).Init(BUCKET_TERMS, n, v)
		}

	default:
		n, v := sv.ScriptStr(), sv.AsField(dsl.Script)
		return (&Bucket{}).Init(BUCKET_TERMS, n, v)
	}

	msg := fmt.Sprintf("invalid bucket '%s'", token.Str())
	return nil, errors.New(msg)
}

// bucketFuncTerms() - Compile SQL token 'terms(field)' to ES terms bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFuncTerms(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch (%d/1)", len(args))
		return nil, errors.New(msg)
	}

	bucket, err := bucketTerms(dsl, args[0], ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in terms()", err)
		return nil, errors.New(msg)
	}

	return bucket, nil
}

// bucketRange() - Compile SQL token 'xx_range(field, [], [], ...)' to ES range bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketRange(btype BucketType, dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	if len(args) < 2 {
		msg := fmt.Sprintf("arguments mismatch %d (expected > 2)", len(args))
		return nil, errors.New(msg)
	}

	fieldToken, rangeList := args[0], args[1:]
	fieldValue, err := ScriptBuild(dsl, fieldToken, ctx)
	if err != nil {
		msg := fmt.Sprintf("invalid field: %s", err)
		return nil, errors.New(msg)
	}

	rangesValue, err := evalTokenRangeList(dsl, rangeList, ctx)
	if err != nil {
		msg := fmt.Sprintf("invalid ranges, %s", err)
		return nil, errors.New(msg)
	}

	switch fieldValue.Type {
	case SCRIPT_FIELD, SCRIPT_VALUE:
		if str, ok := fieldValue.Value.(string); ok {
			n, v := str, map[string]interface{}{"field": str, "ranges": rangesValue}
			return (&Bucket{}).Init(btype, n, v)
		}

	default:
		n := fieldValue.ScriptStr()
		v := map[string]interface{}{"script": fieldValue.AsFieldPure(dsl.Script), "ranges": rangesValue}
		return (&Bucket{}).Init(btype, n, v)
	}

	msg := fmt.Sprintf("invalid field '%s'", fieldToken.Str())
	return nil, errors.New(msg)
}

// bucketFuncRange() - Compile SQL token 'range(field)' to ES range bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFuncRange(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	return bucketRange(BUCKET_RANGE, dsl, args, ctx)
}

// bucketFuncIPRange() - Compile SQL token 'ip_range(field)' to ES terms bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFuncIPRange(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	return bucketRange(BUCKET_IP_RANGE, dsl, args, ctx)
}

// bucketFuncDateRange() - Compile SQL token 'date_range(field)' to ES range bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFuncDateRange(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	numTokens := 0
	for _, item := range args {
		if item.Type() != sql.T_LIST {
			numTokens++
			continue
		}
		break
	}

	format, rangeArgs := "", []sql.Token{}
	switch numTokens {
	case 0:
		rangeArgs = append(rangeArgs, (&sql.TokenStr{}).Init("@timestamp"))
		rangeArgs = append(rangeArgs, args...)

	case 1:
		if args[0].Type() != sql.T_STR {
			msg := fmt.Sprintf("'format' argument not a string value: %s", args[1].Str())
			return nil, errors.New(msg)
		}

		format = args[0].(*sql.TokenStr).Value
		rangeArgs = append(rangeArgs, (&sql.TokenStr{}).Init("@timestamp"))
		rangeArgs = append(rangeArgs, args[1:]...)

	case 2:
		if args[1].Type() != sql.T_STR {
			msg := fmt.Sprintf("'format' argument not a string value: %s", args[1].Str())
			return nil, errors.New(msg)
		}

		format = args[1].(*sql.TokenStr).Value
		rangeArgs = append(rangeArgs, args[0])
		rangeArgs = append(rangeArgs, args[2:]...)

	default:
		msg := fmt.Sprintf("too many argument (%d) before range arguments", numTokens)
		return nil, errors.New(msg)
	}

	bucket, err := bucketRange(BUCKET_DATE_RANGE, dsl, rangeArgs, ctx)
	if err != nil {
		msg := fmt.Sprintf("%s in 'date_range()'", err)
		return nil, errors.New(msg)
	}

	if format != "" {
		bucket.Value["format"] = format
	}

	bucket.Value["time_zone"] = sys.NewTime(0).TimeZone()
	return bucket, nil
}

// bucketHistogram() - Compile SQL token 'xx_histogram(field, [], [], ...)' to ES histogram bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketHistogram(btype BucketType, dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("arguments mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	fieldToken, intervalToken := args[0], args[1]
	fieldValue, err := ScriptBuild(dsl, fieldToken, ctx)
	if err != nil {
		msg := fmt.Sprintf("invalid field: %s", err)
		return nil, errors.New(msg)
	}

	var intervalValue string
	if intervalToken.Type() == sql.T_NUMUNIT {
		t := intervalToken.(*sql.TokenNumUnit)
		intervalValue = t.Num.Str() + t.Unit.Str()
	} else {
		intervalValue = intervalToken.Str()
	}

	switch fieldValue.Type {
	case SCRIPT_FIELD, SCRIPT_VALUE:
		if str, ok := fieldValue.Value.(string); ok {
			n, v := str, map[string]interface{}{"field": str, "interval": intervalValue}
			return (&Bucket{}).Init(btype, n, v)
		}

	default:
		n := fieldValue.ScriptStr()
		v := map[string]interface{}{"script": fieldValue.AsFieldPure(dsl.Script), "interval": intervalValue}
		return (&Bucket{}).Init(btype, n, v)
	}

	msg := fmt.Sprintf("invalid field '%s'", fieldToken.Str())
	return nil, errors.New(msg)
}

// bucketFuncHistogram() - Compile SQL token 'histogram(field, [], [], ...)' to ES histogram bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFuncHistogram(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	return bucketHistogram(BUCKET_HISTOGRAM, dsl, args, ctx)
}

// bucketFuncDateHistogram() - Compile SQL token 'date_histogram(field, [], [], ...)' to ES histogram bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFuncDateHistogram(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	var histogramArgs []sql.Token
	var formatToken sql.Token

	switch len(args) {
	case 1:
		histogramArgs = append(histogramArgs, (&sql.TokenStr{}).Init("@timestamp"))
		histogramArgs = append(histogramArgs, args[0])

	case 2:
		if args[0].Type() == sql.T_NUMUNIT {
			histogramArgs = append(histogramArgs, (&sql.TokenStr{}).Init("@timestamp"))
			histogramArgs = append(histogramArgs, args[0])
			formatToken = args[1]
		} else if args[1].Type() == sql.T_NUMUNIT {
			histogramArgs = args
		} else {
			msg := fmt.Sprintf("invalid arguments '%s' or '%s'", args[0].Str(), args[1].Str())
			return nil, errors.New(msg)
		}

	case 3:
		histogramArgs, formatToken = args[0:2], args[2]

	default:
		msg := fmt.Sprintf("arguments mismatch %d (expected 1, 2 or 3)", len(args))
		return nil, errors.New(msg)
	}

	bucket, err := bucketHistogram(BUCKET_DATE_HISTOGRAM, dsl, histogramArgs, ctx)
	if err != nil {
		return nil, err
	}

	bucket.Value["time_zone"] = "+08:00"
	if formatToken != nil {
		if formatToken.Type() != sql.T_STR {
			msg := fmt.Sprintf("'format' argument not a string value: %s", formatToken.Str())
			return nil, errors.New(msg)
		}

		bucket.Value["format"] = formatToken.(*sql.TokenStr).Value
	}

	return bucket, nil
}

// bucketFilter() - Compile SQL token 'filter(expr)' to ES filter bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFilter(dsl *Dsl, token sql.Token, ctx *script.Cntx) (interface{}, error) {
	filter, err := FilterBuild(dsl, token, ctx)
	if err != nil {
		return nil, err
	}

	dslValue := filter.DslValue()
	if _, ok := dslValue.(map[string]interface{}); !ok {
		return nil, err
	}

	return dslValue, nil
}

// bucketFuncFilters() - Compile SQL token 'filters(expr)' to ES filters bucket.
//
// @dsl:  ES DSL object.
// @args: Function arguments.
// @ctx:  Script context.
func bucketFuncFilters(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Bucket, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("arguments mismatch %d (expected >1)", len(args))
		return nil, errors.New(msg)
	}

	filters := map[string]interface{}{}
	for cc, item := range args {
		if item.Type() != sql.T_AS {
			msg := fmt.Sprintf("the %d argument '%s' not AS token", cc, item.Str())
			return nil, errors.New(msg)
		}
		t := item.(*sql.TokenAs)

		filter, err := bucketFilter(dsl, t.Token, ctx)
		if err != nil {
			msg := fmt.Sprintf("the %d filter '%s' %s", cc, item.Str())
			return nil, errors.New(msg)
		}

		filters[t.Name] = filter
	}

	return (&Bucket{}).Init(BUCKET_FILTERS, "", map[string]interface{}{"filters": filters})
}

// All supported bucket functions.
var g_bucket_functions = map[string]func(*Dsl, []sql.Token, *script.Cntx) (*Bucket, error){
	"terms":          bucketFuncTerms,
	"range":          bucketFuncRange,
	"ip_range":       bucketFuncIPRange,
	"date_range":     bucketFuncDateRange,
	"histogram":      bucketFuncHistogram,
	"date_histogram": bucketFuncDateHistogram,
	"filters":        bucketFuncFilters,
}

// BucketBuild() - Compile SQL token to ES DSL bucket.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func BucketBuild(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Bucket, error) {
	if token.Type() == sql.T_FUNC {
		t := token.(*sql.TokenFunc)

		name := strings.ToLower(t.Name)
		if function, ok := g_bucket_functions[name]; ok {
			if bucket, err := function(dsl, t.List, ctx); err == nil {
				return bucket, nil
			} else {
				msg := fmt.Sprintf("%s in '%s()'", err, name)
				return nil, errors.New(msg)
			}
		}
	}

	if bucket, err := bucketTerms(dsl, token, ctx); err == nil {
		return bucket, nil
	}

	msg := fmt.Sprintf("invalid bucket '%s'", token.Str())
	return nil, errors.New(msg)
}
