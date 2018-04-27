// File: stdout.go
//
// Write log messages to SYSLOG on windows.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
//
// +build windows

package logger

// WriteSyslog() - Write messages to SYSLOG.
//
// @level:     Log level.
// @progName:  Program name.
// @levelName: Level name.
// @msg:       Message to write.
//
// This function do nothing for windowns.
func WriteSyslog(level int, progName, levelName, msg string) {
}
