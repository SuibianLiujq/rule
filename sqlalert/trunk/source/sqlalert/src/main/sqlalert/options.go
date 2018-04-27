// Parse command-line options.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package main

import (
	"core/options"
	"fmt"
	"os"
	"settings"
)

// Option - Structure of command-line options.
//
// See the tag after the members for more information.
type Option struct {
	Help        options.Help `options:"-h, --help,         desc='Show this help'"`
	Version     bool         `options:"-v, --version,      desc='Show version string'"`
	Etc         string       `options:"-e, --etc,          desc='Cofiguration folder path'"`
	Logger      string       `options:"-l, --logger,       desc='Logger output name: none/all/stdout/syslog'"`
	LoggerLevel string       `options:"    --logger-level, desc='Logger level: debug/info/warn/error'"`
	Check       bool         `options:"-c, --check,        desc='Check the grammer only'"`
	Test        string       `options:"-t, --test,         desc='Rule name or FILE to test'"`
	From        string       `options:"    --from,         desc='Time to to run tasks from (YYYY-MM-DD hh-mm-ss)'"`
	To          string       `options:"    --to,           desc='Time to run tasks to (YYYY-MM-DD hh-mm-ss)'"`
	Interval    string       `options:"    --interval,     desc='Time interval for --from and --to options <num>[d|h|m|s]'"`
	Addr        string       `options:"    --listen-addr,  desc='IP address to listen'"`
	Port        string       `options:"    --listen-port,  desc='Port to listen'"`
	Task        string       `options:"    --task,         desc='Task list to run (filepath)'"`
}

// DESC - Description string of SQLAlert.
const DESC string = `
    SQLAlert is an alert engine base-on ES (ElasticSearch) and SQL.
    It execute RDL (Rule Description Language) script to query ES server
    using SQL and output the alert messages.
    
    Please read the white book of SQLAlert for more information.
`

// printHelp() - Print help message to stdout.
func helpMsg() {
	fmt.Println(DESC)
	options.PrintHelp()
	fmt.Println("Version:", settings.VERSION)
	fmt.Println("Bug report: lidan.zhang@clearcouds-global.com\n")
}

// msgExit() - Print error message and exit.
//
// This function prints @msg to stdout and calls to os.Exit(1) to exit
// the program.
func msgExit(msg string, code int) {
	if code == 0 {
		fmt.Println(msg)
	} else {
		fmt.Println("\nERROR:", msg)
		helpMsg()
	}

	os.Exit(code)
}

// parseCmdline() - Parse the command-line.
//
// This function parses command-line options and returns all of the options.
// If the command-line option is not correct this function will print usage
// to stdout and call os.Exit(1) function to exit.
func parseCmdline() *Option {
	option := &Option{
		Help:        false,
		Version:     false,
		Etc:         settings.PATH_ETC,
		Logger:      settings.LOGGER,
		LoggerLevel: settings.LOGGER_LEVEL,
	}

	if err := options.Parse(option); err != nil {
		if err != options.ErrHelpRequest {
			fmt.Println("\nERROR:", err)
		}

		helpMsg()
		os.Exit(1)
	}

	return option
}

// parseOption() - Parse command line options.
//
// This function call parseCmdline() to parse command-line options and
// do first-step checking of the options.
func parseOption() *Option {
	option := parseCmdline()

	if option.Version {
		msg := fmt.Sprintf("version: %s", settings.VERSION)
		msgExit(msg, 0)
	}

	if option.Etc != "" {
		settings.PATH_ETC = option.Etc
	}

	if stat, err := os.Stat(option.Etc); err != nil {
		msg := fmt.Sprintf("Invalid '-e/--etc' option: %s", err)
		msgExit(msg, 1)
	} else {
		if !stat.IsDir() {
			msg := fmt.Sprintf("Invalid '-e/--etc' option: %s", option.Etc)
			msgExit(msg, 1)
		}
	}

	switch option.Logger {
	case "none", "all", "stdout", "syslog":
		settings.LOGGER = option.Logger
	default:
		msg := fmt.Sprintf("Invalid '-l/--logger' option: %s", option.Logger)
		msgExit(msg, 1)
	}

	switch option.LoggerLevel {
	case "debug", "info", "warn", "error":
		settings.LOGGER_LEVEL = option.LoggerLevel
	default:
		msg := fmt.Sprintf("Invalid '--logger-level' option: %s", option.LoggerLevel)
		msgExit(msg, 1)
	}

	if option.Check && (option.Test == "") {
		msg := "option -c/--check only used with option -t/--test"
		msgExit(msg, 1)
	}

	if option.To != "" && option.From == "" {
		msg := "option '--to' only used with '--from'"
		msgExit(msg, 1)
	}

	if option.From != "" && option.Test != "" && option.Interval == "" {
		msg := "option '--interval' not set"
		msgExit(msg, 1)
	}

	if option.Task != "" {
		settings.CONF_TASKS = option.Task
	}

	return option
}
