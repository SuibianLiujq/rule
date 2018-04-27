// File: dsl.go
//
// This file implements the compiler for compiling SQL to DSL.
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

// Dsl - Structure of DSL.
//
// @Index:   Index name.
// @Script:  Script name.
// @Request: Request value.
type Dsl struct {
	Index   string
	Script  string
	Request map[string]interface{}

	UseMetrics   bool
	UseSourceAll bool
	Metrics      []*Metric

	UseBuckets   bool
	WhereBuckets map[string]*Bucket
	Buckets      []*Bucket

	Filter *Filter
	Limits []int64
	Orders []*Order
	Stat   *Stat
}

// Init() - Initailize DSL instance.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Dsl) Init(token sql.Token, ctx *script.Cntx) (*Dsl, error) {
	this.Request = make(map[string]interface{})
	this.WhereBuckets = make(map[string]*Bucket)

	initDefaultValue(ctx)
	this.Script = DEF_SCRIPT

	if token != nil {
		return this.Compile(token, ctx)
	}

	return this, nil
}

// GetUrl() - Returns joined URL.
//
// @host: ES host.
func (this *Dsl) GetUrl(host string) string {
	return "http://" + host + "/" + this.Index
}

// Compile() - Compile SQL to DSL.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Dsl) Compile(token sql.Token, ctx *script.Cntx) (*Dsl, error) {
	if token.Type() != sql.T_STMTS {
		msg := fmt.Sprintf("'%s' not SQL statements", token.Str())
		return nil, errors.New(msg)
	}

	stat, err := StatToken(token, ctx)
	if err != nil {
		return nil, err
	}

	this.Stat = stat
	this.UseBuckets = (len(stat.GroupList) != 0)

	for _, function := range g_compiler_list {
		if err := function(this, ctx); err != nil {
			return nil, err
		}
	}

	return this.Combine(ctx)
}

// Compile() - Compile SQL to METRIC.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Dsl) compileMetric(token sql.Token, ctx *script.Cntx) (*Metric, error) {
	switch token.Type() {
	case sql.T_IDENT, sql.T_STR:
		return metricSource(this, token, ctx)

	case sql.T_FUNC:
		return metricFunc(this, token, ctx)
	}

	return metricScript(this, token, ctx)
}

// compileBucket() - Compile SQL to BUCKET.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Dsl) compileBucket(token sql.Token, ctx *script.Cntx) (*Bucket, error) {
	switch token.Type() {
	case sql.T_IDENT, sql.T_STR:
		if bucket, ok := this.WhereBuckets[token.Str()]; ok {
			return bucket, nil
		}
	}

	return BucketBuild(this, token, ctx)
}

// compileWhereBucket() - Compile SQL (WHERE BY) to BUCKET.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Dsl) compileWhereBucket(token sql.Token, ctx *script.Cntx) (*Bucket, error) {
	return BucketBuild(this, token, ctx)
}

// compileFilter() - Compile SQL to FILTER.
//
// @token: SQL token.
// @ctx:   Script context.
func (this *Dsl) compileFilter(token sql.Token, ctx *script.Cntx) (*Filter, error) {
	return FilterBuild(this, token, ctx)
}

// Compile() - Combine SOURCE/METRICS/BUCKETS/FILTERS to DSL request.
//
// @ctx:   Script context.
func (this *Dsl) Combine(ctx *script.Cntx) (*Dsl, error) {
	if !this.UseMetrics && !this.UseBuckets {
		if err := this.combineSource(ctx); err != nil {
			return nil, err
		}
	} else {
		this.Request["size"] = int64(0)
		this.Request["_source"] = false
		if err := this.combineAggs(ctx); err != nil {
			return nil, err
		}
	}

	if this.Filter == nil {
		this.Request["query"] = FILTER_MATCHALL
	} else {
		this.Request["query"] = this.Filter.DslValue()
	}

	return this, nil
}

// combineSource() - Combine SOURCE.
//
// @ctx:   Script context.
func (this *Dsl) combineSource(ctx *script.Cntx) error {
	sources, scripts := []interface{}{}, map[string]interface{}{}

	for _, metric := range this.Metrics {
		dslValue := metric.DslValue()
		if dslValue == nil {
			continue
		}

		switch metric.Type {
		case METRIC_SOURCE:
			sources = append(sources, dslValue)
		case METRIC_SCRIPT:
			scripts[metric.Name] = dslValue
		default:
			msg := fmt.Sprintf("invalid '%s(%s)' in SELECT", value.ToStr(dslValue), metric.Token.Str())
			return errors.New(msg)
		}
	}

	if len(sources) != 0 {
		this.Request["_source"] = sources
	} else if !this.UseSourceAll {
		this.Request["_source"] = false
	}

	if len(scripts) != 0 {
		this.Request["script_fields"] = scripts
	}

	if len(sources) == 0 && len(scripts) == 0 && !this.UseSourceAll {
		this.Request["size"] = 1
	} else {
		this.Request["size"] = this.Limits[0]
	}

	if this.Orders != nil && len(this.Orders) != 0 {
		orders := []interface{}{}
		for _, item := range this.Orders {
			orders = append(orders, item.Value)
		}
		this.Request["sort"] = orders
	}

	return nil
}

// combineAggs() - Combine aggregation METRICS/BUCKTES.
//
// @ctx:   Script context.
func (this *Dsl) combineAggs(ctx *script.Cntx) error {
	var bucket, lastBucket map[string]interface{}
	var orderValue, orderMetric interface{}
	var orderName string

	if this.UseBuckets && len(this.Buckets) != 0 {
		for cc := len(this.Buckets) - 1; cc >= 0; cc-- {
			orderValue, orderMetric = nil, nil
			if this.Orders != nil {
				if len(this.Orders) == 1 {
					if this.Orders[0] != nil {
						orderValue, orderMetric = this.Orders[0].Value, this.Orders[0].Metric.DslValue()
						orderName = this.Orders[0].Metric.Name
					}
				} else if cc < len(this.Orders) {
					if this.Orders[cc] != nil {
						orderValue, orderMetric = this.Orders[cc].Value, this.Orders[cc].Metric.DslValue()
						orderName = this.Orders[cc].Metric.Name
					}
				}
			}

			item := this.Buckets[cc]
			dslValue := item.DslValue(this.Limits[cc], orderValue)
			if bucket != nil {
				if orderMetric != nil && value.Type(orderMetric) == value.DICT {
					bucket[orderName] = orderMetric
				}

				dslValue["aggs"] = bucket
			}

			if lastBucket == nil {
				lastBucket = dslValue
			}

			bucket = map[string]interface{}{item.Name: dslValue}
		}
	}

	metrics := map[string]interface{}{}
	for _, item := range this.Metrics {
		switch item.Type {
		case METRIC_SCRIPT:
			msg := fmt.Sprintf("script '%s' not allowed in AGG mode", item.Token.Str())
			return errors.New(msg)

		case METRIC_SOURCE:
			if _, ok := this.Stat.GroupDict[item.Selector]; !ok {
				msg := fmt.Sprintf("source '%s' not a item of GROUP BY", item.Token.Str())
				return errors.New(msg)
			}

		default:
			if dslValue := item.DslValue(); dslValue != nil {
				metrics[item.Name] = dslValue
			}
		}
	}

	if lastBucket != nil {
		if len(metrics) != 0 {
			lastBucket["aggs"] = metrics
		}

		this.Request["aggs"] = bucket
	} else {
		if len(metrics) != 0 {
			this.Request["aggs"] = metrics
		}
	}

	return nil
}

// getMetric() - Returns metric of given name.
//
// @name: Name or Alias of the metric.
func (this *Dsl) getMetric(name string) *Metric {
	for _, item := range this.Metrics {
		if item.Name == name {
			return item
		}
	}
	return nil
}

// g_compile_list - Compile function list.
//
// @*Dsl:            Instance of DSL.
// @*script.Cntx: Script context.
var g_compiler_list = []func(*Dsl, *script.Cntx) error{
	compileSelect,
	compileFrom,
	compileWhere,
	compileWhereBy,
	compileGroup,
	compileOrder,
	compileLimit,
	compileHaving,
}

// compileSelect - Compile SELECT token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileSelect(dsl *Dsl, ctx *script.Cntx) error {
	for _, item := range dsl.Stat.SelectList {
		if item.Token.Type() == sql.T_STAR {
			if dsl.UseBuckets {
				return errors.New("'*' in SELECT not allowed in BUCKET-AGG mode")
			}

			if len(dsl.Stat.SelectList) != 1 {
				msg := fmt.Sprintf("more than 1 (%d) items in SELECT", len(dsl.Stat.SelectList))
				return errors.New(msg)
			}

			dsl.UseSourceAll = true
			return nil
		}

		metric, err := dsl.compileMetric(item.Token, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in SELECT", err)
			return errors.New(msg)
		}

		metric.SetAlias(item.Name)
		metric.Token = item.Token
		dsl.Metrics = append(dsl.Metrics, metric)
	}

	return nil
}

// compileFrom - Compile FROM token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileFrom(dsl *Dsl, ctx *script.Cntx) error {
	fromToken, fromIndex := dsl.Stat.From, ""

	if fromToken == nil {
		return errors.New("FROM statement not found")
	}

	switch fromToken.Type() {
	case sql.T_VAR:
		if str, ok := ctx.GetStr(fromToken.(*sql.TokenVar).Value); ok {
			fromIndex = str
		} else {
			msg := fmt.Sprintf("VAR '%s' not found in FROM", fromToken.Str())
			return errors.New(msg)
		}

	default:
		fromIndex = fromToken.Str()
	}

	dsl.Index = fromIndex
	return nil
}

// compileWhere - Compile WHERE token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileWhere(dsl *Dsl, ctx *script.Cntx) error {
	if dsl.Stat.Where != nil {
		filter, err := dsl.compileFilter(dsl.Stat.Where, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in WHERE", err)
			return errors.New(msg)
		}

		dsl.Filter = filter
	}
	return nil
}

// compileWhereBy - Compile WHERE BY token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileWhereBy(dsl *Dsl, ctx *script.Cntx) error {
	for _, item := range dsl.Stat.WhereByDict {
		bucket, err := dsl.compileWhereBucket(item.Token, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in WHERE BY", err)
			return errors.New(msg)
		}

		bucket.Name = item.Name
		dsl.WhereBuckets[bucket.Name] = bucket
	}

	return nil
}

// compileGroup - Compile GROUP BY token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileGroup(dsl *Dsl, ctx *script.Cntx) error {
	for _, item := range dsl.Stat.GroupList {
		bucket, err := dsl.compileBucket(item.Token, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in GROUP BY", err)
			return errors.New(msg)
		}

		bucket.Name = item.Name
		dsl.Buckets = append(dsl.Buckets, bucket)
	}

	return nil
}

// compileOrder - Compile ORDER BY token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileOrder(dsl *Dsl, ctx *script.Cntx) error {
	orderTokens := dsl.Stat.OrderList

	orderList := []*Order{}
	if !dsl.UseMetrics && !dsl.UseBuckets {
		for _, item := range orderTokens {
			token, orderStr := item.(*sql.TokenOrder), DEF_ORDER
			if token.Order != sql.T_ILL {
				orderStr = strings.ToLower(token.Order.Name())
			}

			orderValue := map[string]interface{}{token.Token.Str(): orderStr}
			order := NewOrder(token.Token.Str(), orderStr, orderValue, nil)
			metric := dsl.getMetric(token.Token.Str())

			if metric != nil && metric.Type == METRIC_SCRIPT {
				scriptValue := map[string]interface{}{
					"script": metric.DslValue().(map[string]interface{})["script"],
				}

				scriptValue["type"] = "number"
				scriptValue["order"] = orderStr
				order.Value = map[string]interface{}{"_script": scriptValue}
				order.Metric = metric
			}

			orderList = append(orderList, order)
		}
	} else {
		for _, item := range orderTokens {
			token, orderStr := item.(*sql.TokenOrder), DEF_ORDER

			if token.Token.Str() == "__order_count__" {
				orderList = append(orderList, nil)
			} else {
				metric := dsl.getMetric(token.Token.Str())
				if metric == nil {
					msg := fmt.Sprintf("METRIC '%s' not found in ORDER BY", token.Token.Str())
					return errors.New(msg)
				}

				if token.Order == sql.T_ASC {
					orderStr = "asc"
				}
				orderValue := map[string]interface{}{token.Token.Str(): orderStr}
				order := NewOrder(token.Token.Str(), orderStr, orderValue, metric)
				orderList = append(orderList, order)
			}
		}
	}

	dsl.Orders = orderList
	return nil
}

// compileLimit - Compile LIMIT token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileLimit(dsl *Dsl, ctx *script.Cntx) error {
	stat, limit := dsl.Stat, DEF_LIMITS

	switch len(stat.LimitList) {
	case 1:
		if t, ok := stat.LimitList[0].(*sql.TokenInt); ok {
			limit = t.Value
		}

	case 0:
		dsl.Limits = make([]int64, len(stat.GroupList)+1)
		for cc, _ := range dsl.Limits {
			dsl.Limits[cc] = limit
		}
	}

	dsl.Limits = make([]int64, len(stat.GroupList)+1)
	for cc, _ := range dsl.Limits {
		dsl.Limits[cc] = limit
		if cc < len(stat.LimitList) {
			if t, ok := stat.LimitList[cc].(*sql.TokenInt); ok {
				dsl.Limits[cc] = t.Value
			}
		}
	}

	return nil
}

// compileHaving - Compile HAVING token.
//
// @dsl: DSL instance.
// @ctx: Script context.
func compileHaving(dsl *Dsl, ctx *script.Cntx) error {
	return nil
}

// Compile() - Compile SQL to DSL.
//
// @token: SQL token.
// @ctx:   Script context.
func Compile(token sql.Token, ctx *script.Cntx) (*Dsl, error) {
	return (&Dsl{}).Init(token, ctx)
}
