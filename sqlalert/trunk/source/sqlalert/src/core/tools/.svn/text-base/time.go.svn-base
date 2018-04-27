package tools

import (
	"core/script"
	"core/sys"
	"core/value"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// g_time_interval_regexp - Reguler expression of time interval.
//
// Example: 5m -> five minutes.
var g_time_interval_regexp = regexp.MustCompile(`(\d+)\s*(\w+)`)

// g_time_unit_value - Value of time unit.
//
// The referrence unit of time unit is ms.
var g_time_unit_value = map[string]int64{
	"ms":      1,
	"s":       1 * 1000,
	"sec":     1 * 1000,
	"secs":    1 * 1000,
	"second":  1 * 1000,
	"seconds": 1 * 1000,
	"m":       60 * 1000,
	"min":     60 * 1000,
	"mins":    60 * 1000,
	"minute":  60 * 1000,
	"minutes": 60 * 1000,
	"h":       60 * 60 * 1000,
	"hour":    60 * 60 * 1000,
	"hours":   60 * 60 * 1000,
	"d":       24 * 60 * 60 * 1000,
	"day":     24 * 60 * 60 * 1000,
	"days":    24 * 60 * 60 * 1000,
}

// GetTimeNow() - Return sys.Time instance of now.
//
// @ctx: Script context.
//
// This function try to parse ctx.__now__ (with ctx.__now_fmt__) as
// now time instance if ctx.__now__ is set otherwise it use system
// timestamp as now time.
func GetTimeNow(ctx *script.Cntx) (nowTime *sys.Time, err error) {
	var nowVal, nowFmt interface{}
	if ctx != nil {
		nowVal = ctx.GetX("__now__")
		nowFmt = ctx.GetX("__now_fmt__")
	}

	switch value.Type(nowVal) {
	case value.INT:
		nowTime, err = sys.NewTime(value.Int(nowVal)), nil

	case value.STR:
		switch nowStr := value.Str(nowVal); {
		case nowFmt == nil:
			nowTime, err = sys.ParseTime(nowStr)

		case value.Type(nowFmt) == value.STR:
			nowTime, err = sys.ParseFmtTime(nowStr, value.Str(nowFmt))
		}

	default:
		nowTime, err = sys.NewTime(0), nil
	}

	return nowTime, err
}

// GetTime() - Return sys.Time with given @now value.
//
// @now: Value of now.
// @ctx: Script context.
func GetTime(now interface{}, ctx *script.Cntx) (nowTime *sys.Time, err error) {
	if now == nil && ctx != nil {
		return GetTimeNow(ctx)
	}

	switch value.Type(now) {
	case value.INT:
		nowTime, err = sys.NewTime(value.Int(now)), nil

	case value.STR:
		nowTime, err = sys.ParseTime(value.Str(now))

	default:
		nowTime, err = sys.NewTime(0), nil
	}

	return nowTime, err
}

func GetTimeUnitValue(unit string) (int64, error) {
	if v, ok := g_time_unit_value[unit]; ok {
		return v, nil
	}

	msg := fmt.Sprintf("invalid time unit '%s'", unit)
	return 0, errors.New(msg)
}

// GetTimeInterval() - Returns the ms number the the interval string.
//
// @text: Interval string.
func GetTimeInterval(text string) (int64, error) {
	strList := g_time_interval_regexp.FindStringSubmatch(text)
	if len(strList) < 3 {
		msg := fmt.Sprintf("invalid time interval '%s'", text)
		return 0, errors.New(msg)
	}

	unitValue, err := GetTimeUnitValue(strList[2])
	if err != nil {
		msg := fmt.Sprintf("%s in %s", err, text)
		return 0, errors.New(msg)
	}

	num, err := strconv.ParseInt(strList[1], 10, 64)
	if err != nil {
		msg := fmt.Sprintf("%s in %s", err, text)
		return 0, errors.New(msg)
	}

	return num * unitValue, nil
}
