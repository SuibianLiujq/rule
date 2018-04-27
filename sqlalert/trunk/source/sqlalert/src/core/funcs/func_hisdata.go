// Functions for history data lernning.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	//	"core/value"
	"errors"
	"fmt"
)

func hisCheck(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return args[0], nil
}
