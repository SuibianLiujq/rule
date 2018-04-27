// File: utils.go
//
// This file implements the system interface functions
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package sys

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile() - Read all file content.
//
// @name: String of file name.
func ReadFile(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		msg := fmt.Sprintf("%s in %s", err, name)
		return nil, errors.New(msg)
	}

	return content, nil
}
