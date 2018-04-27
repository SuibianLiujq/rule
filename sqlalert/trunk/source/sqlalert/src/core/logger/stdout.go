// File: stdout.go
//
// Write log messages to STDOUT.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.

package logger

import (
	"fmt"
	"strings"
)

func WriteStdout(timeStr, progName, levelName, msg string) {
	fmt.Printf("%v %v: [%s] %v\n", timeStr, progName, levelName, msg)
}

func WriteStdoutExt(timeStr, progName, levelName, msg string) {
	msgList := strings.Split(msg, "\n")
	for _, item := range msgList {
		fmt.Printf("%v %v: [%s] %v\n", timeStr, progName, levelName, item)
	}
}
