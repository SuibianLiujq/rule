// Functions to format number.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"core/script"
	"core/value"
	"errors"
	"fmt"
)

// fmtTime() - Format time interval to readable string.
//
// @timeValue: Time interval value.
func __fmtTime(us int64) string {
	s, ms, us := us/1000000, (us/1000)%1000, us%1000

	var text string
	if s > 0 {
		if ms > 0 {
			text = fmt.Sprintf("%d.%.3ds", s%60, ms)
		} else {
			text = fmt.Sprintf("%ds", s%60)
		}

		m := s / 60
		if m > 0 {
			text = fmt.Sprintf("%dm %s", m%60, text)

			h := m / 60
			if h > 0 {
				text = fmt.Sprintf("%dh %s", h%24, text)

				d := h / 24
				if d > 0 {
					text = fmt.Sprintf("%dd %s", d, text)
				}
			}
		}
	} else if ms > 0 {
		if us > 0 {
			text = fmt.Sprintf("%d.%dms", ms, us)
		} else {
			text = fmt.Sprintf("%dms", ms)
		}
	} else {
		text = fmt.Sprintf("%dus", us)
	}

	return text
}

// fmtTime() - Format time interval.
//
// @args: Function arguments
// @ctx:  Script context.
func fmtTime(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	var argTime interface{}

	timeUnit := "ms"
	switch len(args) {
	case 1:
		argTime = args[0]
	case 2:
		argTime, timeUnit = args[0], value.ToStr(args[1])
	default:
		msg := fmt.Sprintf("argument mismatch %d(expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	var timeValue float64
	switch argTime.(type) {
	case int64:
		timeValue = float64(argTime.(int64))
	case float64:
		timeValue = argTime.(float64)
	default:
		msg := fmt.Sprintf("arg[0] not NUMBER", value.ToStr(argTime))
		return nil, errors.New(msg)
	}

	switch timeUnit {
	case "s":
		timeValue *= 1000000.0
	case "ms":
		timeValue *= 1000.0
	case "us":
	default:
		msg := fmt.Sprintf("time unit '%s' unsupport", timeUnit)
		return nil, errors.New(msg)
	}

	return __fmtTime(int64(timeValue)), nil
}

var g_byte_units = []struct {
	unitByte string
	unitBit  string
	value    int64
}{
	{"EB", "Eb", int64(1024 * 1024 * 1024 * 1024 * 1024 * 1024)},
	{"PB", "Pb", int64(1024 * 1024 * 1024 * 1024 * 1024)},
	{"TB", "Tb", int64(1024 * 1024 * 1024 * 1024)},
	{"GB", "Gb", int64(1024 * 1024 * 1024)},
	{"MB", "Mb", int64(1024 * 1024)},
	{"KB", "Kb", int64(1024)},
}

// fmtByte() - Format byte value to readable string.
//
// @byteValue: Byte value.
func __fmtByte(byteValue int64, name string) string {
	floatValue := float64(byteValue)

	for _, item := range g_byte_units {
		resValue := floatValue / float64(item.value)
		if resValue >= 1.0 {
			if name == "byte" {
				return fmt.Sprintf("%.1f %s", resValue, item.unitByte)
			} else {
				return fmt.Sprintf("%.1f %s", resValue, item.unitBit)
			}
		}
	}

	if name == "byte" {
		return fmt.Sprintf("%d B", byteValue)
	} else {
		return fmt.Sprintf("%d b", byteValue)
	}
}

// fmtBytes() - Format bytes.
//
// @args: Function arguments
// @ctx:  Script context.
func fmtBytes(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	byteValue := int64(0)
	switch args[0].(type) {
	case int64:
		byteValue = args[0].(int64)
	case float64:
		byteValue = int64(args[0].(float64))

	default:
		msg := fmt.Sprintf("arg[0] not NUMBER", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	return __fmtByte(byteValue, "byte"), nil
}

// funcFmtBits() - Format bits.
//
// @args: Function arguments
// @ctx:  Script context.
func fmtBits(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	bitValue := int64(0)
	switch args[0].(type) {
	case int64:
		bitValue = args[0].(int64)
	case float64:
		bitValue = int64(args[0].(float64))

	default:
		msg := fmt.Sprintf("arg[0] not NUMBER", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	//	bitValue /= 8
	return __fmtByte(bitValue, "bit"), nil
}

// fmtPercentage() - Format percentage.
//
// @args: Function arguments
// @ctx:  Script context.
func fmtPercentage(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	floatValue := float64(0.0)
	switch args[0].(type) {
	case int64:
		floatValue = float64(args[0].(int64))
	case float64:
		floatValue = args[0].(float64)

	default:
		msg := fmt.Sprintf("arg[0] not NUMBER", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	return fmt.Sprintf("%.2f%%", floatValue*100), nil
}

// fmtInt() - Format INT.
//
// @args: Function arguments
// @ctx:  Script context.
func fmtInt(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	if intValue, err := value.ToInt(args[0]); err == nil {
		return fmt.Sprintf("%d", intValue), nil
	}

	return value.ToStr(args[0]), nil
}

// fmtFloat() - Format FLOAT.
//
// @args: Function arguments
// @ctx:  Script context.
func fmtFloat(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	num := value.Int(2)
	if len(args) >= 2 {
		if intValue, ok := value.AsInt(args[1]); ok {
			num = intValue
		}
	}

	if floatValue, err := value.ToFloat(args[0]); err == nil {
		fmtStr := fmt.Sprintf("%%.%df", num)
		return fmt.Sprintf(fmtStr, floatValue), nil
	}

	return value.ToStr(args[0]), nil
}

// fmtStr() - Format STR.
//
// @args: Function arguments
// @ctx:  Script context.
func fmtStr(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("argument mismatch %d(expected 1)", len(args))
		return nil, errors.New(msg)
	}

	num := value.Int(2)
	if len(args) >= 2 {
		if intValue, ok := value.AsInt(args[1]); ok {
			num = intValue
		}
	}

	fmtStr := fmt.Sprintf("%%%d.s", num)
	return fmt.Sprintf(fmtStr, value.ToStr(args[0])), nil
}
