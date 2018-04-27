// File syslog_linux.go
//
// This file implements the Logger to write messages to SYSLOG.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
//
// +build linux

package logger

import (
	"fmt"
	"log/syslog"
	"strings"
)

var g_levels = [...]syslog.Priority{
	DEBUG: syslog.LOG_DEBUG,
	INFO:  syslog.LOG_INFO,
	WARN:  syslog.LOG_WARNING,
	ERROR: syslog.LOG_ERR,
}

// WriteSyslog() - Write messages to SYSLOG.
//
// @level:     Log level.
// @progName:  Program name.
// @levelName: Level name.
// @msg:       Message to write.
func WriteSyslog(level int, progName, levelName, msg string) {
	sysLogger, err := syslog.New(syslog.LOG_LOCAL0|g_levels[level], progName)
	if err == nil {
		sysLogger.Err(fmt.Sprintf("[%s] %v\n", levelName, msg))
		sysLogger.Close()
	}
}

func WriteSyslogExt(level int, progName, levelName, msg string) {
	sysLogger, err := syslog.New(syslog.LOG_LOCAL0|g_levels[level], progName)
	if err == nil {
		msgList := strings.Split(msg, "\n")
		for _, item := range msgList {
			sysLogger.Err(fmt.Sprintf("[%s] %v\n", levelName, item))
		}
		sysLogger.Close()
	}
}
