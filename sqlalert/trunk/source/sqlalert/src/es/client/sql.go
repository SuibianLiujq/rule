// Simple SQL-ES client base on RESTful interface.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package client

import (
	"core/json"
	"core/script"
	"core/sql"
	"core/value"
	"errors"
	"es/dsl"
	"fmt"
)

// SQLClient - Structure of SQL-ES client.
type SQLClient struct {
	esClient *ESClient
	esDsl    *dsl.Dsl
}

// Init() - Initialize SQLClient instance with given SQL text.
//
// @text: Byte stream of SQL text.
// @ctx:  Script context.
//
// This function return SQLClient instance itself for chain operation.
func (this *SQLClient) Init(text []byte, ctx *script.Cntx) (cli *SQLClient, err error) {
	if text == nil || len(text) == 0 {
		return this.InitToken(nil, ctx)
	}

	token, err := sql.Parse(text)
	if err != nil {
		msg := fmt.Sprintf("SQL: '%s' %s", string(text), err)
		return nil, errors.New(msg)
	}

	return this.InitToken(token, ctx)
}

// Init() - Initialize SQLClient instance with given SQL token.
//
// @text: Byte stream of SQL text.
// @ctx:  Script context.
//
// This function return SQLClient instance itself for chain operation.
func (this *SQLClient) InitToken(token sql.Token, ctx *script.Cntx) (cli *SQLClient, err error) {
	this.esClient, err = (&ESClient{}).Init()
	if err != nil {
		return nil, err
	}

	timeout := ctx.GetX("__es_timeout__")
	if intValue, ok := timeout.(int64); ok && intValue > 0 {
		this.esClient.SetTimeout(intValue)
	}

	if token != nil {
		this.esDsl, err = dsl.Compile(token, ctx)
		if err != nil {
			return nil, err
		}
	}

	return this, nil
}

// Query() - Query url using inner DSL.
//
// @host: ES host.
// @ctx:  Script context.
func (this *SQLClient) Query(host string, ctx *script.Cntx) (interface{}, error) {
	if this.esClient == nil || this.esDsl == nil {
		return nil, errors.New(fmt.Sprintf("not initialized"))
	}

	if debug := ctx.GetX("__es_debug_request__"); value.IsTrue(debug) {
		fmt.Println("----------------------- Request -----------------------")
		fmt.Println(this.esDsl.GetUrl(host))
		json.PPrint(this.esDsl.Request)
	}

	esRsp, err := this.esClient.Search(this.esDsl.GetUrl(host), this.esDsl.Request)
	if err != nil {
		return nil, err
	}

	if debug := ctx.GetX("__es_debug_response__"); value.IsTrue(debug) {
		fmt.Println("----------------------- Response -----------------------")
		json.PPrint(esRsp)
	}

	result, err := this.parseResponse(esRsp)
	if debug := ctx.GetX("__sql_debug_response__"); value.IsTrue(debug) {
		fmt.Println("----------------------- SQL Result -----------------------")
		fmt.Println("length =", len(result))
		json.PPrint(result)
	}

	return result, err
}

// parseResponse() - Parse ES response into list.
//
// @rsp: ES response in interface{} type.
func (this *SQLClient) parseResponse(rsp interface{}) (list []interface{}, err error) {
	rspDict, ok := rsp.(map[string]interface{})
	if !ok {
		msg := fmt.Sprintf("response not dict: %s", json.DumpStrAll(rsp))
		return nil, errors.New(msg)
	}

	total, hits, aggs, err := this.splitResponse(rspDict)
	if err != nil {
		return nil, err
	}

	if !(this.esDsl.UseMetrics || this.esDsl.UseBuckets) {
		if list, err = this.parseHits(hits); err != nil {
			msg := fmt.Sprintf("%s in 'hits.hits'", err)
			return nil, errors.New(msg)
		}
	} else {
		if list, err = this.parseAggs(aggs); err != nil {
			msg := fmt.Sprintf("%s in 'aggregations'", err)
			return nil, errors.New(msg)
		}
	}

	if !this.esDsl.UseSourceAll {
		for _, item := range list {
			if itemDict, ok := item.(map[string]interface{}); ok {
				itemDict["count(*)"] = total
			} else {
				msg := fmt.Sprintf("result item not a DICT: %s", json.DumpStrAll(item))
				return nil, errors.New(msg)
			}
		}

		if list, err = this.selectValue(list); err == nil {
			if len(this.esDsl.Orders) == 1 && this.esDsl.Orders[0] != nil {
				value.Sort(list, this.esDsl.Orders[0].Name, this.esDsl.Orders[0].Order)
			}
		}
	}

	return list, nil
}

// parseHits() - Parse 'hits' response.
//
// @hits: ES 'hits' response.
func (this *SQLClient) parseHits(hits []interface{}) ([]interface{}, error) {
	list := []interface{}{}

	for cc, item := range hits {
		itemDict, ok := item.(map[string]interface{})
		if !ok {
			msg := fmt.Sprintf("hits[%d] not a DICT: %s", cc, json.DumpStrAll(item))
			return nil, errors.New(msg)
		}

		listItem := map[string]interface{}{}

		if source, ok := itemDict["_source"]; ok {
			sourceDict, ok := source.(map[string]interface{})
			if !ok {
				msg := fmt.Sprintf("hits[%d]['_source'] not a DICT: %s", cc, json.DumpStrAll(source))
				return nil, errors.New(msg)
			}

			for k, v := range sourceDict {
				listItem[k] = v
			}
		}

		if fields, ok := itemDict["fields"]; ok {
			if fieldsDict, ok := fields.(map[string]interface{}); !ok {
				msg := fmt.Sprintf("hits[%d]['fieds'] not a DICT: %s", cc, json.DumpStrAll(fields))
				return nil, errors.New(msg)
			} else {
				for k, v := range fieldsDict {
					listItem[k] = v
				}
			}
		}

		list = append(list, listItem)
	}

	return list, nil
}

// parseAggs() - Parse 'aggregations' response.
//
// @aggs: ES 'aggregations' response.
func (this *SQLClient) parseAggs(aggs map[string]interface{}) ([]interface{}, error) {
	key, buckets := this.getAggBuckets(aggs)
	if buckets == nil {
		return this.parseAggMetrics(aggs)
	}

	list := []interface{}{}
	switch buckets.(type) {
	case []interface{}:
		for cc, item := range buckets.([]interface{}) {
			itemDict, ok := item.(map[string]interface{})
			if !ok {
				msg := fmt.Sprintf("bucket item %s[%d] not a DICT: %s", key, cc, json.DumpStrAll(item))
				return nil, errors.New(msg)
			}

			bucketValue, err := this.getAggBucketValue(itemDict)
			if err != nil {
				msg := fmt.Sprintf("%s in bucket item %s[%d]: %s", key, cc, json.DumpStrAll(item))
				return nil, errors.New(msg)
			}

			subList, err := this.parseAggBucket(key, bucketValue, itemDict)
			if err != nil {
				return nil, err
			}

			for _, subItem := range subList {
				list = append(list, subItem)
			}
		}

	case map[string]interface{}:
		for k, item := range buckets.(map[string]interface{}) {
			itemDict, ok := item.(map[string]interface{})
			if !ok {
				msg := fmt.Sprintf("bucket item %s[%s] not a DICT: %s", key, k, json.DumpStrAll(item))
				return nil, errors.New(msg)
			}

			bucketValue := map[string]interface{}{"key": k}
			subList, err := this.parseAggBucket(key, bucketValue, itemDict)
			if err != nil {
				return nil, err
			}

			for _, subItem := range subList {
				list = append(list, subItem)
			}
		}

	default:
		msg := fmt.Sprintf("invalid buckets: %s", json.DumpStrAll(buckets))
		return nil, errors.New(msg)
	}

	return list, nil
}

// parseAggMetrics() - Parse metrics value from rsp.aggs.
//
// @aggs: Part of response.
func (this *SQLClient) parseAggMetrics(aggs map[string]interface{}) ([]interface{}, error) {
	metricValue := map[string]interface{}{}
	for k, item := range aggs {
		if _, ok := item.(map[string]interface{}); ok {
			metricValue[k] = item
		}
	}

	if len(metricValue) != 0 {
		return []interface{}{metricValue}, nil
	}

	return []interface{}{}, nil
}

// parseAggBucket() - Parse the value of bucket itself from rsp.aggs.
//
// @k:    Key of bucket.
// @v:    Value of bucket.
// @aggs: Part of response.
func (this *SQLClient) parseAggBucket(k string, v, aggs map[string]interface{}) ([]interface{}, error) {
	list, err := this.parseAggs(aggs)
	if err != nil {
		return nil, err
	}

	for _, item := range list {
		item.(map[string]interface{})[k] = v
	}

	if len(list) == 0 {
		return []interface{}{map[string]interface{}{k: v}}, nil
	}

	return list, nil
}

// getAggBuckets() - Returns the 'buckets' item of the response.
//
// @aggs: Part of response.
func (this *SQLClient) getAggBuckets(aggs map[string]interface{}) (string, interface{}) {
	for key, item := range aggs {
		if itemDict, ok := item.(map[string]interface{}); ok {
			if buckets, ok := itemDict["buckets"]; ok {
				return key, buckets
			}
		}
	}

	return "", nil
}

// getAggBucketValue() - Returns the value of bucket.
//
// @aggs: Part of response.
func (this *SQLClient) getAggBucketValue(aggs map[string]interface{}) (map[string]interface{}, error) {
	v := map[string]interface{}{}

	for key, item := range aggs {
		if !value.IsDict(item) {
			v[key] = item
		}
	}

	return v, nil
}

// selectValue() - Select values for cached DSL metrics and buckets.
//
// @rspList: List values of ES response.
func (this *SQLClient) selectValue(rspList []interface{}) ([]interface{}, error) {
	list := []interface{}{}

	for _, rspItem := range rspList {
		rspDict, ok := rspItem.(map[string]interface{})
		if !ok {
			continue
		}

		listItem := map[string]interface{}{}
		if err := this.selectBuckets(rspDict, listItem); err != nil {
			continue
		}

		if err := this.selectMetrics(rspDict, listItem); err != nil {
			continue
		}

		if len(listItem) != 0 {
			list = append(list, listItem)
		}
	}

	return list, nil
}

// selectMetrics() - Select value for metrics.
//
// @rspDict: Response value in map[string]interface{} type.
// @result:  A DICT to contaian selcted value.
func (this *SQLClient) selectMetrics(rspDict map[string]interface{}, result map[string]interface{}) error {
	for _, item := range this.esDsl.Metrics {
		if item.Selector == "" {
			msg := fmt.Sprintf("metrics.%s.Selector is empty", item.Name)
			return errors.New(msg)
		}

		if item.Type != dsl.METRIC_COUNT_BUCKET {
			if _, ok := this.esDsl.Stat.GroupDict[item.Selector]; ok {
				continue
			}
		}

		rspValue, ok := rspDict[item.Selector]
		if !ok {
			msg := fmt.Sprintf("no response for metrics.%s.%s", item.Name, item.Selector)
			return errors.New(msg)
		}

		v, err := item.SelectValue(rspValue)
		if err != nil {
			msg := fmt.Sprintf("%s in metrics.%s.%s", err, item.Name, item.Selector)
			return errors.New(msg)
		}

		result[item.Name] = v
	}

	return nil
}

// selectBuckets() - Select value for buckets.
//
// @rspDict: Response value in map[string]interface{} type.
// @result:  A DICT to contaian selcted value.
func (this *SQLClient) selectBuckets(rspDict map[string]interface{}, result map[string]interface{}) error {
	for _, item := range this.esDsl.Buckets {
		rspValue, ok := rspDict[item.Name]
		if !ok {
			msg := fmt.Sprintf("no response for buckets.%s", item.Name)
			return errors.New(msg)
		}

		v, err := item.SelectValue(rspValue)
		if err != nil {
			msg := fmt.Sprintf("%s in buckets.%s", err, item.Name)
			return errors.New(msg)
		}

		result[item.Name] = v
	}

	return nil
}

// splitResponse() - Split ES response into three parts: total (count), hits, aggs.
//
// @rsp: Response value in map[string]interface{} type.
func (this *SQLClient) splitResponse(rsp map[string]interface{}) (total interface{}, hits []interface{}, aggs map[string]interface{}, err error) {
	if hitsValue, ok := rsp["hits"]; !ok {
		err = errors.New(fmt.Sprintf("'hits' not found in: %s", json.DumpStrAll(rsp)))
		return
	} else {
		if hitsDict, ok := hitsValue.(map[string]interface{}); !ok {
			err = errors.New(fmt.Sprintf("'hits' not a DICT: %s", json.DumpStrAll(rsp)))
			return
		} else {
			if total, ok = hitsDict["total"]; !ok {
				err = errors.New(fmt.Sprintf("'hits.total' not found in: %s", json.DumpStrAll(hitsDict)))
				return
			}

			if hitsValue, ok = hitsDict["hits"]; !ok {
				err = errors.New(fmt.Sprintf("'hits.hits' not found in: %s", json.DumpStrAll(hitsDict)))
				return
			} else {
				if hits, ok = hitsValue.([]interface{}); !ok {
					err = errors.New(fmt.Sprintf("'hits.hits' not a LIST: %s", json.DumpStrAll(hitsValue)))
					return
				}
			}
		}
	}

	if aggsValue, ok := rsp["aggregations"]; ok {
		if aggs, ok = aggsValue.(map[string]interface{}); !ok {
			err = errors.New(fmt.Sprintf("'aggregations' not a DICT: %s", json.DumpStrAll(aggsValue)))
			return
		}
	}

	return
}

// NewSQLClient() - Create SQLClient instance.
//
// @sqlValue: SQL statements value.
//            Type of this argument must be []byte or sql.Token.
func NewSQLClient(sqlValue interface{}, ctx *script.Cntx) (*SQLClient, error) {
	switch sqlValue.(type) {
	case string:
		return (&SQLClient{}).Init([]byte(sqlValue.(string)), ctx)

	case []byte:
		return (&SQLClient{}).Init(sqlValue.([]byte), ctx)

	case sql.Token:
		return (&SQLClient{}).InitToken(sqlValue.(sql.Token), ctx)
	}

	msg := fmt.Sprintf("only []byte or sql.Token support, got: %s", sqlValue)
	return nil, errors.New(msg)
}
