package task

import (
	"core/logger"
	"core/script"
	"core/sys"
	"core/tools"
	"os"
	"settings"
	"sync"
	"time"
)

// TaskConf - Structure of Task configuration.
//
// @Enable:     Enable the task.
// @RunOnStart: Start the task when the program is up.
// @Interval:   Schedule interval (seconds).
// @File:       File to load.
type TaskConf struct {
	Interval   int64
	RunOnStart bool
	Name       string
	File       string
}

// Task - Structure of Task.
//
// @interval: Schedule interval.
// @name:     Task name.
// @file:     File to load.
// @isActive: Flag of running.
// @nextTime: Next schedule time.
// @runTime:  Running time.
type Task struct {
	interval int64
	name     string
	file     string
	isActive bool
	nextTime int64
	runTime  int64

	token   script.Token
	session sync.Map
}

// Init() - Initialize Task instance.
//
// This function returns the Task instance itself for chain operation.
func (this *Task) Init(now int64, cfg *TaskConf) (*Task, error) {
	this.name = cfg.Name
	this.file = cfg.File
	this.interval = cfg.Interval
	this.isActive = false
	this.nextTime = 0
	this.runTime = 0
	this.token = nil

	if !cfg.RunOnStart {
		this.nextTime = now + this.interval
	}

	return this, nil
}

// Run() - Run task.
//
// @check: Flag to check grammer only.
func (this *Task) Run(check bool, ctx *script.Cntx) {
	this.isActive = true
	ctx.Session = &this.session

	if err := this.__Run(check, ctx); err != nil {
		logger.ErrorFmt("%s in task '%s'", err, this.name)
	}

	this.isActive = false
}

// __Run() - Inner function of Run().
//
// @check: Flag to check grammer only.
func (this *Task) __Run(check bool, ctx *script.Cntx) error {
	nowStr := ""

	if now := ctx.GetX("__now__"); now != nil {
		if now, err := tools.GetTimeNow(ctx); err == nil {
			nowStr = now.ToStr("%Y-%M-%D %h:%m:%s")
		} else {
			logger.Error(err)
		}
	}

	if nowStr == "" {
		logger.InfoFmt("Run script %s", this.file)
	} else {
		logger.InfoFmt("Run script %s (%s)", this.file, nowStr)
	}

	if this.token == nil {
		token, err := script.ParseFile(this.scriptFile(this.file))
		if err != nil {
			return err
		}

		this.token = token
	}

	if !check {
		_, err := script.Exec(this.token, ctx)
		return err
	}

	return nil
}

// scriptFile() - Returns the full path of script file.
//
// @name: File name.
//
// This function returns name itself if name is a valid non-dir file,
// otherwise it returns the full path with the prefix PATH_SCRITPS.
func (this *Task) scriptFile(name string) string {
	if info, err := os.Stat(name); err == nil && !info.IsDir() {
		return name
	}

	return settings.GetFileScripts(name)
}

// parseTime() - Parse 'from' and 'to' time string.
//
// @from: Time string of 'from'.
// @to:   Time string of 'to'.
func parseTime(from, to string) (timeFrom, timeTo *sys.Time) {
	timeFrom, err := sys.ParseTime(from)
	if err != nil {
		logger.ErrorFmt("Invalid 'from': %s", err)
		return
	}

	timeNow := sys.NewTime(0)
	if to != "" {
		if timeTo, err = sys.ParseTime(to); err != nil {
			logger.ErrorFmt("Invalid 'to': %s", err)
			return
		}
	}

	if timeTo == nil || timeTo.Timestamp > timeNow.Timestamp {
		timeTo = timeNow
	}

	return timeFrom, timeTo
}

// runTaskFromLoop() Run task from the specified time in a loop.
//
// @from: Timestamp of 'from'.
// @from: Timestamp of 'to'.
// @interval: Time interval.
func runTaskFromLoop(task *Task, from, to, interval int64, ctx *script.Cntx) {

	for now := from; now <= to; now += interval {
		newCtx := ctx.CopyAll()
		newCtx.Set("__now__", now)

		task.Run(false, newCtx)
		time.Sleep(1 * time.Second)
	}
}

// runTaskFrom() - Run task from the specified time.
//
// @task:     Instance of *Task.
// @from:     Time string of 'from'.
// @to:       Time string of 'to'.
// @interval: String of time interval.
// @ctx:      Script context.
func runTaskFrom(task *Task, check bool, from, to, interval string, ctx *script.Cntx) {
	if check || from == "" {
		task.Run(check, ctx)
		return
	}

	timeFrom, timeTo := parseTime(from, to)
	if timeFrom == nil || timeTo == nil {
		return
	}

	if timeTo.Timestamp < timeFrom.Timestamp {
		logger.ErrorFmt("Time 'to(%s)' is small than 'from(%s)'", to, from)
		return
	}

	timeStep, err := tools.GetTimeInterval(interval)
	if err != nil {
		logger.ErrorFmt("Invalid interval: %s", err)
		return
	}

	runTaskFromLoop(task, timeFrom.Timestamp, timeTo.Timestamp, timeStep, ctx)
}

// RunTask() - Run a task.
//
// @name:  File name of the script.
// @from:  Time 'YYYY-MM-DD hh-mm-ss" to schedule from.
// @to:    Time 'YYYY-MM-DD hh-mm-ss" to schedule to.
//         The default 'to' is now.
// @check: Only to check the grammer.
//
// It only check the script's grammer exclude SQL syntex.
func RunTask(name, from, to, interval string, check bool) {
	now, cfg := time.Now().Unix(), &TaskConf{}

	cfg.Name, cfg.File = "test", name
	inst, err := (&Task{}).Init(now, cfg)
	if err != nil {
		logger.ErrorFmt("Failed to run '%s': %s", name, err)
		return
	}

	err = execGlobals(settings.PathGlobals())
	if err != nil {
		logger.ErrorFmt("Failed to exec global scripts: %s", err)
		return
	}

	runTaskFrom(inst, check, from, to, interval, g_context.CopyAll())
}
