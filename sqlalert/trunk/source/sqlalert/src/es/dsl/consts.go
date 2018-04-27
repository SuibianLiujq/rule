// File: consts.go
//
// This file defines constant and default values for DSL.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package dsl

import (
	"core/script"
)

// DEF_LIMITS - Default limit-size of each buckets.
//
// If there's no buckets this value is the limitation of ES _sources.
// Script context value '__es_bucket_size__' will change this value.
var DEF_LIMITS int64 = 5

// DEF_SCRIPT - Default script to use.
//
// All script will build depends on the value. Script context value
// '__es_script__' will change this value.
var DEF_SCRIPT string = "groovy"

// DEF_ORDER
//
// Default sort order.
// '__es_order__' will change this value.
var DEF_ORDER string = "desc"

// initDefaultValue() - Initialize the default values.
//
// @ctx: Script context.
func initDefaultValue(ctx *script.Cntx) {
	if v, ok := ctx.GetXStr("__es_script__"); ok {
		DEF_SCRIPT = v
	}

	if v, ok := ctx.GetXInt("__es_bucket_size__"); ok {
		DEF_LIMITS = v
	}

	if v, ok := ctx.GetXStr("__es_order__"); ok {
		DEF_ORDER = v
	}
}
