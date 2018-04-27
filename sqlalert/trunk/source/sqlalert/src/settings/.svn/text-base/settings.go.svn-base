// Default configuration for SQLAlert.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package settings

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	PROGRAM string = "SQLAlert"
	VERSION string = "1.0.3"
)

var (
	LOGGER       string = "syslog"
	LOGGER_LEVEL string = "warn"
	LOGGER_NAME  string = "SQLAlert"

	PATH_ETC     string = "/usr/local/etc"
	PATH_SCRIPTS string = ""
	PATH_GLOBALS string = "globals"

	CONF_TASKS string = "tasks.json"
)

func PathEtc() string     { return PATH_ETC }
func PathScripts() string { return filepath.Join(PATH_ETC, PATH_SCRIPTS) }
func PathGlobals() string { return filepath.Join(PATH_ETC, PATH_GLOBALS) }

func GetFileEtc(name string) string     { return filepath.Join(PATH_ETC, name) }
func GetFileScripts(name string) string { return filepath.Join(PathScripts(), name) }
func GetFileGlobals(name string) string { return filepath.Join(PathGlobals(), name) }

func GetFileCfgTasks() string {
	if filepath.IsAbs(CONF_TASKS) || strings.HasPrefix(CONF_TASKS, "./") {
		return CONF_TASKS
	}

	return filepath.Join(PATH_ETC, CONF_TASKS)
}

func GuessScriptFile(name string) string {
	if info, err := os.Stat(name); err == nil {
		if !info.IsDir() {
			return name
		}
	}

	return GetFileScripts(name)
}
