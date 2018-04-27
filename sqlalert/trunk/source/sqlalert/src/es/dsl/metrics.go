// File: metrics.go
//
// This file implements compilter of DSL metrics.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package dsl

import (
	"core/script"
	"core/sql"
	"core/value"
	"errors"
	"fmt"
	"strings"
)

// Metric type.
type MetricType int

// All metric types.
const (
	METRIC_SOURCE MetricType = iota
	METRIC_SCRIPT
	METRIC_COUNT
	METRIC_COUNT_BUCKET
	METRIC_VALUE_COUNT
	METRIC_CARDINALITY
	METRIC_SUM
	METRIC_MIN
	METRIC_MAX
	METRIC_AVG
	METRIC_SQUARES
	METRIC_STDEV
	METRIC_STDEV_UPPER
	METRIC_STDEV_LOWER
	METRIC_VARIANCE
)

// g_metrics - Defines fo DSL metrics.
//
// @Name:  Name of DSL metric.
// @value: Name of DSL response value.
var g_metrics = [...]struct {
	Name, Value, ExtValue string
}{
	METRIC_SOURCE:       {"", "value", ""},
	METRIC_SCRIPT:       {"", "value", ""},
	METRIC_COUNT:        {"", "value", ""},
	METRIC_COUNT_BUCKET: {"", "doc_count", ""},
	METRIC_VALUE_COUNT:  {"value_count", "value", ""},
	METRIC_CARDINALITY:  {"cardinality", "value", ""},
	METRIC_SUM:          {"sum", "value", ""},
	METRIC_MIN:          {"min", "value", ""},
	METRIC_MAX:          {"max", "value", ""},
	METRIC_AVG:          {"avg", "value", ""},
	METRIC_SQUARES:      {"extended_stats", "sum_of_squares", ""},
	METRIC_STDEV:        {"extended_stats", "std_deviation", ""},
	METRIC_STDEV_UPPER:  {"extended_stats", "std_deviation_bounds", "upper"},
	METRIC_STDEV_LOWER:  {"extended_stats", "std_deviation_bounds", "lower"},
	METRIC_VARIANCE:     {"extended_stats", "variance", ""},
}

func (this MetricType) Name() string     { return g_metrics[this].Name }
func (this MetricType) Value() string    { return g_metrics[this].Value }
func (this MetricType) ExtValue() string { return g_metrics[this].ExtValue }

// Metric - Structure of ES Metric.
//
// @Type:     Metric type.
// @Name:     Metric name.
// @Selector: Metric value selector.
// @Value:    Metric value.
// @Token:    SQL token.
type Metric struct {
	Type     MetricType
	Name     string
	Selector string
	Value    interface{}
	Token    sql.Token
}

// Init() - Initialize Metric instance.
//
// @t: Metric type.
// @n: Metric name.
// @s: Metric value selector.
// @v: Metric value.
//
// This function returns the Metric instance itself for chain operation.
func (this *Metric) Init(t MetricType, n, s string, v interface{}) (*Metric, error) {
	this.Type, this.Name, this.Selector, this.Value = t, n, s, v
	return this, nil
}

// DslValue() - Returns the DSL value.
func (this *Metric) DslValue() interface{} {
	if this.Value == nil {
		return nil
	}

	switch this.Type {
	case METRIC_SOURCE, METRIC_SCRIPT:
		return this.Value
	}

	return map[string]interface{}{this.Type.Name(): this.Value}
}

// SetAlias() - Set alias name for metric.
//
// @name: Metric alias name.
//
// This function also changes the selector if it is empty.
func (this *Metric) SetAlias(name string) {
	this.Name = name
	if this.Selector == "" {
		this.Selector = name
	}
}

// SelectValue() - Select response value.
//
// @v: Response value.
func (this *Metric) SelectValue(v interface{}) (interface{}, error) {
	switch this.Type {
	case METRIC_SCRIPT:
		if list, ok := v.([]interface{}); ok && len(list) != 0 {
			return list[0], nil
		}

		msg := fmt.Sprintf("'%s' not LIST or empty", value.ToStr(v))
		return nil, errors.New(msg)

	case METRIC_COUNT, METRIC_SOURCE:
		return v, nil
	}

	dict, ok := v.(map[string]interface{})
	if !ok {
		msg := fmt.Sprintf("'%s' not a DICT", value.ToStr(v))
		return nil, errors.New(msg)
	}

	metricValue, ok := dict[this.Type.Value()]
	if !ok {
		msg := fmt.Sprintf("'%s' not found in '%s'", this.Type.Value(), value.ToStr(v))
		return nil, errors.New(msg)
	}

	metricDict, ok := metricValue.(map[string]interface{})
	if ok && this.Type.ExtValue() != "" {
		if extValue, ok := metricDict[this.Type.ExtValue()]; ok {
			return extValue, nil
		}
	}

	return metricValue, nil
}

// metricSource - Compile SOURCE metric.
//
// @dsl:   DSL instance.
// @token: SQL token.
// @ctx:   Script context.
func metricSource(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Metric, error) {
	name, stat := token.Str(), dsl.Stat

	if dsl.UseBuckets {
		if _, ok := stat.GroupDict[name]; !ok {
			msg := fmt.Sprintf("'%s' not in GROUP BY", name)
			return nil, errors.New(msg)
		}
	}

	return (&Metric{}).Init(METRIC_SOURCE, name, name, name)
}

// metricScript - Compile SCRIPT metric.
//
// @dsl:   DSL instance.
// @token: SQL token.
// @ctx:   Script context.
func metricScript(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Metric, error) {
	sv, err := ScriptBuild(dsl, token, ctx)
	if err != nil {
		return nil, err
	}

	if sv.Type == SCRIPT_VALUE {
		if str, ok := sv.Value.(string); ok {
			return (&Metric{}).Init(METRIC_SOURCE, str, "", str)
		} else {
			msg := fmt.Sprint("%s (%s) not STR", value.ToStr(sv.Value), token.Str())
			return nil, errors.New(msg)
		}
	}

	return (&Metric{}).Init(METRIC_SCRIPT, "", "", sv.AsField(dsl.Script))
}

// metricFuncCount() - Compile SQL token 'count(...)' to ES metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncCount(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch (%d/1)", len(args))
		return nil, errors.New(msg)
	}

	switch token := args[0]; token.Type() {
	case sql.T_STAR:
		return (&Metric{}).Init(METRIC_COUNT, "", "count(*)", nil)

	case sql.T_UNIQUE:
		return metricFuncCardinality(dsl, args, ctx)

	case sql.T_IDENT, sql.T_STR:
		strValue := token.Str()
		if _, ok := dsl.Stat.GroupDict[strValue]; !ok {
			return metricFuncValueCount(dsl, args, ctx)
		}

		return (&Metric{}).Init(METRIC_COUNT_BUCKET, strValue, strValue, nil)
	}

	return metricFuncValueCount(dsl, args, ctx)
}

// metricFuncStats() - Compile stats metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncStats(mtype MetricType, dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	if len(args) != 1 {
		msg := fmt.Sprint("argument mismatch (%d/1)", len(args))
		return nil, errors.New(msg)
	}
	token := args[0]
	if token.Type() == sql.T_UNIQUE {
		token = token.(*sql.TokenUnique).Token
	}
	dsl.UseMetrics = true

	sv, err := ScriptBuild(dsl, token, ctx)
	if err != nil {
		return nil, err
	}

	switch sv.Type {
	case SCRIPT_FIELD, SCRIPT_VALUE:
		if str, ok := sv.Value.(string); ok {
			n, v := str, map[string]interface{}{"field": str}
			return (&Metric{}).Init(mtype, n, "", v)
		}

	default:
		n, v := sv.Value.(string), sv.AsField(dsl.Script)
		return (&Metric{}).Init(mtype, n, "", v)
	}

	msg := fmt.Sprintf("invalid meric '%s'", token.Str())
	return nil, errors.New(msg)
}

// metricFuncSum() - Compile SQL token 'sum(...)' to ES metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncSum(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_SUM, dsl, args, ctx)
}

// metricFuncMax() - Compile SQL token 'max(...)' to ES metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncMax(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_MAX, dsl, args, ctx)
}

// metricFuncMin() - Compile SQL token 'min(...)' to ES metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncMin(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_MIN, dsl, args, ctx)
}

// metricFuncAvg() - Compile SQL token 'avg(...)' to ES metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncAvg(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_AVG, dsl, args, ctx)
}

// metricFuncCardinality() - Compile SQL token 'cardinality(...)' to ES metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncCardinality(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_CARDINALITY, dsl, args, ctx)
}

// metricFuncValueCount() - Compile SQL token 'value_count(...)' to ES metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncValueCount(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_VALUE_COUNT, dsl, args, ctx)
}

// metricFuncSquares() - Compile SQL token 'sum_of_squares(...) metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncSquares(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_SQUARES, dsl, args, ctx)
}

// metricFuncStdev() - Compile SQL token 'stdev(...) metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncStdev(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_STDEV, dsl, args, ctx)
}

// metricFuncStdevUpper() - Compile SQL token 'stdev_upper(...) metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncStdevUpper(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_STDEV_UPPER, dsl, args, ctx)
}

// metricFuncStdevLower() - Compile SQL token 'stdev_lower(...) metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncStdevLower(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_STDEV_LOWER, dsl, args, ctx)
}

// metricFuncVariance() - Compile SQL token 'variance(...) metrics.
//
// @dsl:   Instance of DSL.
// @token: SQL token.
// @ctx:   Script context.
func metricFuncVariance(dsl *Dsl, args []sql.Token, ctx *script.Cntx) (*Metric, error) {
	return metricFuncStats(METRIC_VARIANCE, dsl, args, ctx)
}

// Metric functions.
var g_metric_funcs = map[string]func(*Dsl, []sql.Token, *script.Cntx) (*Metric, error){
	"count":          metricFuncCount,
	"sum":            metricFuncSum,
	"max":            metricFuncMax,
	"min":            metricFuncMin,
	"avg":            metricFuncAvg,
	"cardinality":    metricFuncCardinality,
	"value_count":    metricFuncValueCount,
	"sum_of_squares": metricFuncSquares,
	"stdev":          metricFuncStdev,
	"stdev_upper":    metricFuncStdevUpper,
	"stdev_lower":    metricFuncStdevLower,
	"variance":       metricFuncVariance,
}

// metricFunc - Compile FUNC metric.
//
// @dsl:   DSL instance.
// @token: SQL token.
// @ctx:   Script context.
func metricFunc(dsl *Dsl, token sql.Token, ctx *script.Cntx) (*Metric, error) {
	t := token.(*sql.TokenFunc)

	name := strings.ToLower(t.Name)
	if function, ok := g_metric_funcs[name]; ok {
		if metric, err := function(dsl, t.List, ctx); err == nil {
			return metric, nil
		} else {
			msg := fmt.Sprintf("%s in '%s()'", err, name)
			return nil, errors.New(msg)
		}
	}

	if metric, err := metricScript(dsl, token, ctx); err == nil {
		return metric, nil
	}

	msg := fmt.Sprintf("metric func '%s' not support", name)
	return nil, errors.New(msg)
}
