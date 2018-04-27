// File: misc.go
//
// This file implements the misc tool functions.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package tools

import ()

// Append() - Append DICT items to a LIST of DICTs.
//
// @data: List of DICTs.
// @ext:  Data to append.
func Append(data []interface{}, ext map[string]interface{}) []interface{} {
	for _, item := range data {
		if dict, ok := item.(map[string]interface{}); ok {
			for k, v := range ext {
				if _, ok := dict[k]; !ok {
					dict[k] = v
				}
			}
		}
	}
	return data
}
