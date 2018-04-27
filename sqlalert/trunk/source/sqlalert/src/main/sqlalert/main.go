// Main entry of SQLAlert.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package main

import (
	"core/funcs"
	"core/json"
	"core/logger"
	"core/script"
	"core/sql"
	"es/dsl"
	"fmt"
	"os"
	"path/filepath"
	"settings"
	"tasks/task"
)

// logger_levels - Logger level.
var logger_levels = map[string]int{
	"debug": logger.DEBUG,
	"info":  logger.INFO,
	"warn":  logger.WARN,
	"error": logger.ERROR,
}

// initLogger() - Initialize the global logger.
func initLogger() {
	switch settings.LOGGER {
	case "all", "syslog", "stdout":
		level, ok := logger_levels[settings.LOGGER_LEVEL]
		if !ok {
			msg := fmt.Sprintf("unknown logger level: '%s'", settings.LOGGER_LEVEL)
			msgExit(msg, 1)
		}

		name, program := settings.LOGGER_NAME, filepath.Base(os.Args[0])
		if settings.PROGRAM != "" {
			program = settings.PROGRAM
		}

		logger.Configure(name, settings.LOGGER, program, level)
		return
	}

	msg := fmt.Sprintf("unknown logger output: '%s'", settings.LOGGER)
	msgExit(msg, 1)
}

// testSql() - Test SQL.
//
// @name:  SQL file.
// @check: Only to check SQL grammer.
func testSql(name string, check bool) {
	fileName := settings.GuessScriptFile(name)

	token, err := sql.ParseFile(fileName)
	if err != nil {
		fmt.Printf("ERROR: %s in file '%s'\n", err, fileName)
		return
	}

	if !check {
		ctx := script.NewContext()
		ctx.Funcs = funcs.FuncDict()

		if esDsl, err := dsl.Compile(token, ctx); err == nil {
			fmt.Println("Compiled ES-DSL:")
			fmt.Println("-------------------------------------------")
			json.PPrint(esDsl.Request)
		} else {
			fmt.Printf("ERROR: %s in file\n", fileName)
			return
		}
	}
}

// main() - Main function.
func main() {
	option := parseOption()
	initLogger()

	switch {
	case option.Test != "":
		logger.Debug("test rule")
		task.RunTask(option.Test, option.From, option.To, option.Interval, option.Check)

	default:
		logger.Debug("run task scheduler")
		task.RunSched(settings.GetFileCfgTasks(), option.From, option.To)
	}
}
