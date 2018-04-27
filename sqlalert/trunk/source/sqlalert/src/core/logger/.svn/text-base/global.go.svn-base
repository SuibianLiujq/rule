// File: global.go
//
// This file implements the global logger instance. All module
// using logger package shared the global logger instance.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package logger

import (
	"os"
)

// _logger - Global logger instance.
//
// See Logger class for more information.
var _logger = NewLogger(os.Args[0])

// Configure() - Configure global logger.
//
// @name:    Global logger name.
// @enable:  Global features to enable (STDOUT/SYSLOG).
// @program: Global program name.
// @level:   Global logger level.
//
// This function is unsafe in multi-routines.
func Configure(name string, enable string, program string, level int) {
	_logger.SetName(name)
	_logger.Enable(enable, true)
	_logger.SetProgram(program)
	_logger.SetLevel(level)
}

// Functions to write log messages.
func Debug(args ...interface{}) { _logger.Debug(args...) }
func Info(args ...interface{})  { _logger.Info(args...) }
func Warn(args ...interface{})  { _logger.Warn(args...) }
func Error(args ...interface{}) { _logger.Error(args...) }

// Functions to write formatted log messages.
func DebugFmt(fmt string, args ...interface{}) { _logger.DebugFmt(fmt, args...) }
func InfoFmt(fmt string, args ...interface{})  { _logger.InfoFmt(fmt, args...) }
func WarnFmt(fmt string, args ...interface{})  { _logger.WarnFmt(fmt, args...) }
func ErrorFmt(fmt string, args ...interface{}) { _logger.ErrorFmt(fmt, args...) }
