// Task scheduler.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package task

import (
	"core/funcs"
	"core/json"
	"core/logger"
	"core/script"
	"core/value"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"settings"
	"time"
)

// SCHED_TIME - The very low-level timer in seconds.
//
// Scheduler will sleep SCHED_TIME seconds before the next runing
// of all tasks.
const SCHED_TIME time.Duration = time.Duration(1) * time.Second

// g_tasks - Task list.
//
// Scheduler will run tasks in g_tasks list one-by-one in each loop.
var g_tasks = []*Task{}

// g_ctx - The global context of all scripts.
//
// Scheduler load and execute the 'etc/globals/*' scripts to generate
// the global context. While schedule the tasks, scheduler copies the
// context for each task every time.
//
// Scripts in 'etc/globals' will be load and execute only once.
var g_context = script.NewContext()

// parseConfigure() - Parse the configuration item of task.
//
// @name: Task name.
// @cfgDict: Configuration item.
func parseConfigure(name string, cfgDict map[string]interface{}) ([]*TaskConf, error) {
	cfg := &TaskConf{}

	cfg.RunOnStart = false
	if cfgValue, ok := cfgDict["runOnStart"]; ok {
		if value.IsTrue(cfgValue) {
			cfg.RunOnStart = true
		}
	}

	if cfgValue, ok := cfgDict["interval"]; !ok {
		msg := fmt.Sprintf("%s.interval not found", name)
		return nil, errors.New(msg)
	} else if intValue, ok := cfgValue.(int64); !ok || intValue <= 0 {
		msg := fmt.Sprintf("%s.interval not INT or <= 0", name)
		return nil, errors.New(msg)
	} else {
		cfg.Interval = int64(intValue)
	}

	files := []string{}
	if cfgValue, ok := cfgDict["file"]; !ok || value.IsFalse(cfgValue) {
		msg := fmt.Sprintf("%s.file is not found or empty", name)
		return nil, errors.New(msg)
	} else {
		switch cfgValue.(type) {
		case string:
			files = append(files, cfgValue.(string))

		case []interface{}:
			for cc, item := range cfgValue.([]interface{}) {
				if str, ok := item.(string); !ok || str == "" {
					msg := fmt.Sprintf("%s.file[%d] is not a STRING or empty", name, cc)
					return nil, errors.New(msg)
				}

				files = append(files, item.(string))
			}
		}
	}

	list := []*TaskConf{}
	for _, item := range files {
		list = append(list, &TaskConf{
			RunOnStart: cfg.RunOnStart,
			Interval:   cfg.Interval,
			Name:       item + " (" + name + ")",
			File:       item,
		})
	}

	return list, nil
}

// loadConfigure() - Load the tasks configuration.
//
// @name: File name (full path).
func loadConfigure(name string) ([]*TaskConf, error) {
	cfg, err := json.ParseFile(name)
	if err != nil {
		msg := fmt.Sprintf("Failed to load '%s': %s", name, err)
		return nil, errors.New(msg)
	}

	cfgDict, ok := cfg.(map[string]interface{})
	if !ok {
		msg := fmt.Sprintf("Content of '%s' not a DICT", name)
		return nil, errors.New(msg)
	}

	list := []*TaskConf{}
	for key, item := range cfgDict {
		if itemDict, ok := item.(map[string]interface{}); ok {
			if enable, ok := itemDict["enable"]; ok && value.IsTrue(enable) {
				if cfgList, err := parseConfigure(key, itemDict); err == nil {
					for _, cfgItem := range cfgList {
						list = append(list, cfgItem)
					}
				} else {
					return nil, err
				}
			}
		} else {
			msg := fmt.Sprintf("'%s' in '%s' not a DICT", key, name)
			return nil, errors.New(msg)
		}
	}

	return list, nil
}

// execGlobals() - Executes the global scripts under the given folder.
//
// @path: The 'globals' path (full path).
func execGlobals(path string) error {
	g_context.Funcs = funcs.FuncDict()
	g_context.CtxGlobal = g_context
	g_context.Set("__etc_scripts__", settings.PathScripts())

	if dirs, err := ioutil.ReadDir(path); err == nil {
		for _, file := range dirs {
			if file.IsDir() {
				err := execGlobals(filepath.Join(path, file.Name()))
				if err != nil {
					return err
				}
			} else if filepath.Ext(file.Name()) == ".rule" {
				_, err := script.ExecFile(filepath.Join(path, file.Name()), g_context)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// loadTasks() - Load tasks from the configuration file.
//
// @name: Configuration file path (full path).
func loadTasks(name string) error {
	cfgTasks, err := loadConfigure(name)
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	for _, item := range cfgTasks {
		if task, err := (&Task{}).Init(now, item); err == nil {
			logger.Info("Append task:", task.name, task.interval)
			g_tasks = append(g_tasks, task)
		} else {
			return err
		}
	}

	if len(g_tasks) == 0 {
		return errors.New("task list is empty")
	}

	return execGlobals(settings.PathGlobals())
}

func runSchedFromLoop(ch chan bool, task *Task, from, to int64) {
	runTaskFromLoop(task, from, to, task.interval*1000, g_context)
	ch <- true
}

// runSchedFrom() - Run to schedule the tasks from the specified time.
//
// @from: Time 'YYYY-MM-DD hh-mm-ss" to schedule from.
// @to:   Time 'YYYY-MM-DD hh-mm-ss" to schedule to.
//        The default 'to' is now.
func runSchedFrom(from, to string) error {
	timeFrom, timeTo := parseTime(from, to)
	if timeFrom == nil || timeTo == nil {
		return nil
	}

	if timeTo.Timestamp < timeFrom.Timestamp {
		logger.ErrorFmt("Time 'to(%s)' is small than 'from(%s)'", to, from)
		return nil
	}

	ch := make(chan bool, len(g_tasks))
	for _, taskItem := range g_tasks {
		go runSchedFromLoop(ch, taskItem, timeFrom.Timestamp, timeTo.Timestamp)
	}

	for cc := 0; cc < len(g_tasks); cc++ {
		<-ch
	}

	return nil
}

// runSchedLoop() - Run to schedule the tasks in a loop.
func runSchedLoop() error {
	for {
		time.Sleep(SCHED_TIME)

		now := time.Now().Unix()
		for _, taskItem := range g_tasks {
			if taskItem.nextTime > now {
				continue
			}

			if taskItem.isActive {
				taskItem.runTime += taskItem.interval
				logger.WarnFmt("Task '%s' is still running after %d seconds", taskItem.name, taskItem.runTime)
			} else {
				taskItem.runTime = 0
				go taskItem.Run(false, g_context.CopyAll())
			}

			taskItem.nextTime = now + taskItem.interval
		}
	}

	return nil
}

// RunSched() - Run to schedule the tasks.
//
// @from: Time 'YYYY-MM-DD hh-mm-ss" to schedule from.
// @to:   Time 'YYYY-MM-DD hh-mm-ss" to schedule to.
//        The default 'to' is now.
// @name: Name of configuration file.
func RunSched(name, from, to string) {
	if err := loadTasks(name); err != nil {
		logger.ErrorFmt("Failed to load tasks: %s", err)
		return
	}

	if from != "" {
		if err := runSchedFrom(from, to); err != nil {
			logger.ErrorFmt("Failed to run scheduler from '%s': %s", from, err)
		}
		return
	}

	if err := runSchedLoop(); err != nil {
		logger.ErrorFmt("Failed to run scheduler loop: %s", err)
		return
	}
}
