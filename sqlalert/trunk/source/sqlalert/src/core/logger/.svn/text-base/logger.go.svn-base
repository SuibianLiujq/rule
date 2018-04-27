// Write log messages to STDOUT and/or SYSLOG.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	DEBUG = iota
	INFO  = iota
	WARN  = iota
	ERROR = iota
)

const (
	__LOG_DEFAULT__ = WARN
)

type Logger struct {
	Program string
	Name    string
	Level   int

	EnableStdout bool
	EnableSyslog bool
}

// Init() - Initialize the logger object with the given module name.
//
// @name: Logger name.
//
// This function do following thing(s):
//   (1) Initialize the default log level to be LOG_WARN.
//   (2) Initialize the syslog if enabled.
func (this *Logger) Init(name string) bool {
	this.Name = name
	this.SetLevel(__LOG_DEFAULT__)
	this.Enable("stdout", true)
	this.SetProgram(os.Args[0])
	return true
}

// SetName() - Set the name of logger.
//
// @name: Logger name.
func (this *Logger) SetName(name string) {
	this.Name = name
}

// SetLevel() - Set the level of logger.
//
// @level: One of LOG_DEBUG, LOG_INFO, LOG_WARN, LOG_ERROR.
//
// The sequence of log level is LOG_DEBUG < LOG_INFO < LOG_WARN < LOG_ERROR.
// Message with level lower than the level set by user (default 'LOG_WARN')
// will not be output.
func (this *Logger) SetLevel(level int) {
	if level >= DEBUG && level <= ERROR {
		this.Level = level
	}
}

// SetProgram() - Set the program name.
//
// @name - Program name used by logger.
//         Default value is os.Args[0].
//
// This function will do nothing if name is empty.
func (this *Logger) SetProgram(name string) {
	if len(name) != 0 {
		this.Program = filepath.Base(name)
	}
}

// Enable() - Enable or disable the output.
//
// @name:   Output name, one of 'stdout', 'syslog' and 'all'.
// @enable: Action flag, enable or disable.
func (this *Logger) Enable(name string, enable bool) {
	switch strings.ToLower(name) {
	case "syslog":
		this.EnableSyslog = enable
	case "stdout":
		this.EnableStdout = enable
	case "all":
		this.EnableStdout = enable
		this.EnableSyslog = enable
	default:
		fmt.Fprintf(os.Stderr, "WARN: unknown logger output name '%v'\n", name)
	}
}

// Debug() - Write the debug message.
//
// @args: Messages to be write, can be any value.
func (this *Logger) Debug(args ...interface{}) {
	if this.Level > DEBUG {
		return
	}

	this.__WriteMessage(DEBUG, "DBG", false, "", args...)
}

// Info() - Write the info message.
//
// @args: Messages to be write, can be any value.
func (this *Logger) Info(args ...interface{}) {
	if this.Level > INFO {
		return
	}

	this.__WriteMessage(INFO, "INF", false, "", args...)
}

// Warn() - Write the warnning message.
//
// @args: Messages to be write, can be any value.
func (this *Logger) Warn(args ...interface{}) {
	if this.Level > WARN {
		return
	}

	this.__WriteMessage(WARN, "WAR", false, "", args...)
}

// Error() - Write the error message.
//
// @args: Messages to be write, can be any value.
func (this *Logger) Error(args ...interface{}) {
	if this.Level > ERROR {
		return
	}

	this.__WriteMessage(ERROR, "ERR", false, "", args...)
}

// DebugFmt() - Formatting verison of Debug().
//
// @args: Messages to be write, can be any value.
func (this *Logger) DebugFmt(fmt string, args ...interface{}) {
	if this.Level > DEBUG {
		return
	}

	this.__WriteMessage(DEBUG, "DBG", true, fmt, args...)
}

// InfoFmt() - Formatting verison of Info().
//
// @args: Messages to be write, can be any value.
func (this *Logger) InfoFmt(fmt string, args ...interface{}) {
	if this.Level > INFO {
		return
	}

	this.__WriteMessage(INFO, "INF", true, fmt, args...)
}

// WarnFmt() - Formatting verison of Warn().
//
// @args: Messages to be write, can be any value.
func (this *Logger) WarnFmt(fmt string, args ...interface{}) {
	if this.Level > WARN {
		return
	}

	this.__WriteMessage(WARN, "WAR", true, fmt, args...)
}

// ErrorFmt() - Formatting verison of Error().
//
// @args: Messages to be write, can be any value.
func (this *Logger) ErrorFmt(fmt string, args ...interface{}) {
	if this.Level > ERROR {
		return
	}

	this.__WriteMessage(ERROR, "ERR", true, fmt, args...)
}

// __WriteMessage() - Write messages to stdout and syslog.
//
// @level:     Priority of syslog priority.
// @levelName: Name of priority.
// @useFmt:    Flag of formatting verson output.
// @fmtStr:    Format string.
// @msgs:      Messages to be write.
func (this *Logger) __WriteMessage(level int, levelName string, useFmt bool, fmtStr string, msgs ...interface{}) {
	msg := ""
	if useFmt {
		msg = fmt.Sprintf(fmtStr, msgs...)
	} else {
		for _, value := range msgs {
			msg = fmt.Sprintf("%s %v", msg, value)
		}
	}

	msg = strings.TrimSpace(msg)
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	if this.EnableStdout {
		WriteStdout(timeStr, this.Program, levelName, msg)
	}

	if this.EnableSyslog {
		WriteSyslog(level, this.Program, levelName, msg)
	}
}

// New() - Create a new Logger object.
//
// @name: Logger name.
//
// Create a new Logger object with name and returns the address. Value nil
// will be returned if failed to initialize the logger.
func NewLogger(name string) *Logger {
	logger := new(Logger)

	if logger.Init(name) {
		return logger
	}

	return nil
}
