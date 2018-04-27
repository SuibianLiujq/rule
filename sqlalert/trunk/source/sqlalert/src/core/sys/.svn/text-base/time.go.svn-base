// File: time.go
//
// This file implements the Time() and corresponding functions.
//
// Copyright (C) 2017 YUN Li Lai, Nanjiing, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package sys

import (
	"regexp"
	"time"
)

const (
	TIME_FMTSTR = "2006-01-02 15:04:05.999-07:00"
)

// Time format string.
var g_regexp = regexp.MustCompile("%[Y|M|D|h|m|s|f|z]")
var g_index_fmtstr = map[string]string{
	"%Y": "2006",
	"%M": "01",
	"%D": "02",
	"%h": "15",
	"%m": "04",
	"%s": "05",
	"%f": "999",
	"%z": "-07:00",
}

// Time() - Structure of Time.
//
// @Now: Timestamp of now in milliseconds.
type Time struct {
	Timestamp int64
	Time      time.Time
}

// Init() - Initialize Time instance.
//
// @now: Timestamp in milliseconds.
//
// This function returns the instance itself for chain operation.
func (this *Time) Init(now int64) *Time {
	if now > 0 {
		this.Timestamp = now
		this.Time = time.Unix(now/1000, (now%1000)*1000000)
	} else {
		this.Time = time.Now()
		this.Timestamp = this.Time.UnixNano() / 1000000
	}

	return this
}

// Init() - Initialize Time instance.
//
// @t: Instance of time.Time.
//
// This function returns the instance itself for chain operation.
func (this *Time) InitTime(t time.Time) *Time {
	this.Time = t
	this.Timestamp = this.Time.UnixNano() / 1000000
	return this
}

// Str() - Convert time to formatted string.
//
// This function use TIME_FMTSTR as formatter.
func (this *Time) Str() string {
	return this.ToStr(TIME_FMTSTR)
}

// ToStr() - Convert time to specified format string.
//
// @fmtStr: Foramt string to format the time.
func (this *Time) ToStr(fmtStr string) string {
	return this.Time.Format(Formatter(fmtStr))
}

// Functions to return the date and time.
func (this *Time) Years() int64     { return int64(this.Time.Year()) }
func (this *Time) YearDays() int64  { return int64(this.Time.YearDay()) }
func (this *Time) Months() int64    { return int64(this.Time.Month()) }
func (this *Time) MonthDays() int64 { return int64(this.Time.Day()) }
func (this *Time) Hours() int64     { return int64(this.Time.Hour()) }
func (this *Time) Minutes() int64   { return int64(this.Time.Minute()) }
func (this *Time) Seconds() int64   { return int64(this.Time.Second()) }
func (this *Time) WeekDays() int64 {
	days := int64(this.Time.Weekday())
	if days == 0 {
		days = 7
	}
	return days
}

func __get_years(t *Time) int64      { return t.Years() }
func __get_year_days(t *Time) int64  { return t.YearDays() }
func __get_months(t *Time) int64     { return t.Months() }
func __get_month_days(t *Time) int64 { return t.MonthDays() }
func __get_week_days(t *Time) int64  { return t.WeekDays() }
func __get_hours(t *Time) int64      { return t.Hours() }
func __get_minutes(t *Time) int64    { return t.Minutes() }
func __get_seconds(t *Time) int64    { return t.Seconds() }

var g_value_funcs = map[string]func(*Time) int64{
	"years":      __get_years,
	"year_days":  __get_year_days,
	"months":     __get_months,
	"month_days": __get_month_days,
	"week_days":  __get_week_days,
	"hours":      __get_hours,
	"minutes":    __get_minutes,
	"seconds":    __get_seconds,
}

// Get() - Returns the Date or Time depends on the given name.
//
// @name: Name of the value to return.
func (this *Time) Get(name string) int64 {
	if function, ok := g_value_funcs[name]; ok {
		return function(this)
	}

	return int64(0)
}

// TimeZone() - Returns the timezone string.
func (this *Time) TimeZone() string {
	return this.Time.Format("-07:00")
}

// NewTime() - Create Time instance.
//
// @now: Timestamp in milliseconds.
//
// Time instance use @now as NOW timestamp if @now is not 0,
// otherwise it use system (time.Now()) to obtains the timestamp.
func NewTime(now int64) *Time {
	return (&Time{}).Init(now)
}

// ParseTime() - Parse time with default formatter.
//
// @str: Time string in "%Y-%M-%H %h:%m:%s.%f" format.
func ParseTime(str string) (*Time, error) {
	res, err := time.Parse(TIME_FMTSTR, str+time.Now().Format("-07:00"))
	if err != nil {
		return nil, err
	}

	return (&Time{}).InitTime(res), nil
}

// ParseFmtTime() - Parse time with specified formatter.
//
// @fmtStr: String like "%Y-%M-%H %h:%m:%s.%f%z".
// @str:    Time string.
func ParseFmtTime(timeStr, fmtStr string) (*Time, error) {
	fmtStr = Formatter(fmtStr)

	res, err := time.Parse(fmtStr, timeStr)
	if err != nil {
		return nil, err
	}

	return (&Time{}).InitTime(res), nil
}

// __replaceFormatter() - Replace the time formatter.
//
// @repl: Replace string.
//
// This function is called by regexp module.
func __replaceFormatter(repl string) string {
	if item, ok := g_index_fmtstr[repl]; ok {
		return item
	}
	return repl
}

// Formatter() - Translate formatter to system-like format.
//
// @raw: Foramt string like "%Y-%M-%H %h:%m:%s.%f%z".
func Formatter(raw string) string {
	return g_regexp.ReplaceAllStringFunc(raw, __replaceFormatter)
}
