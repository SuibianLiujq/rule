// Evaluate value of <NUMBER UNIT> token.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package dsl

import (
	"core/script"
	"core/sql"
	"core/tools"
	"core/value"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

// g_unitvalue_time - All supported time unit.
//
// The referrence unit of time unit is ms.
var g_unit_base_time = map[string]int64{
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

// g_unit_base_byte - All supported bytes unit.
//
// The referrence unit of time unit is ms.
var g_unit_base_byte = map[string]int64{
	"b":  1,
	"kb": 1 * 1024,
	"mb": 1 * 1024 * 1024,
	"gb": 1 * 1024 * 1024 * 1024,
	"tb": 1 * 1024 * 1024 * 1024 * 1024,
	"pb": 1 * 1024 * 1024 * 1024 * 1024 * 1024,
}

// evalTimeUnitValue() - Returns base value of time unit.
//
// @unit: Time unit name.
func evalUnitBaseValueTime(unit string) (int64, error) {
	name := strings.ToLower(unit)

	if value, ok := g_unit_base_time[name]; ok {
		return value, nil
	}

	msg := fmt.Sprintf("Time unit '%s' not support", unit)
	return int64(0), errors.New(msg)
}

// evalUnitBaseValueByte() - Returns base value of byte unit.
//
// @unit: Byte unit name.
func evalUnitBaseValueByte(unit string) (int64, error) {
	name := strings.ToLower(unit)

	if value, ok := g_unit_base_byte[name]; ok {
		return value, nil
	}

	msg := fmt.Sprintf("Time unit '%s' not support", unit)
	return int64(0), errors.New(msg)
}

// evalTimeUnitValue() - Returns timestamp of @value.
//
// @value: Value of timestamp.
//         int64:  Return value itself.
//         string: Parse "2017-04-21 13:29:56.000" to timestamp.
func evalTimeValue(v interface{}, fmtStr string) (int64, error) {
	switch v.(type) {
	case int64:
		return v.(int64), nil
	case string:
		timeStr := v.(string) + time.Now().Format("-07:00")
		if fmtStr == "" {
			fmtStr = "2006-01-02 15:04:05.999-07:00"
		}

		tv, err := time.Parse(fmtStr, timeStr)
		if err != nil {
			return int64(0), err
		}

		return int64(tv.UnixNano() / 1000000), nil
	}

	msg := fmt.Sprintf("invalid time value '%s'", v)
	return int64(0), errors.New(msg)
}

// evalTimeNow() - Returns now() timestamp in ms.
//
// @ctx: Script context.
func evalTimeNow(ctx *script.Cntx) (int64, error) {
	now, err := tools.GetTimeNow(ctx)
	if err != nil {
		return 0, err
	}

	var offset int64 = 0
	if ctxValue := ctx.GetX("__time_offset__"); ctxValue != nil {
		if v, ok := ctxValue.(int64); ok {
			offset = v
		}
	}

	return now.Timestamp - offset, nil
}

// evalTimeNow() - Returns now() timestamp in ms.
//
// @ctx: Script context.
func evalTimeNowDay(ctx *script.Cntx) (int64, error) {
	nowValue, err := evalTimeNow(ctx)
	if err != nil {
		return int64(0), err
	}

	tv := time.Unix(int64(nowValue)/1000, (int64(nowValue)%1000)*1000000)
	tv, _ = time.Parse("2006-01-02 -07:00", tv.Format("2006-01-02 -07:00"))

	return int64(tv.UnixNano() / 1000000), nil
}

// evalTokenTimeInterval() - Returns value of time interval tokens.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func evalTokenTimeInterval(dsl *Dsl, token sql.Token, ctx *script.Cntx) (int64, error) {
	t, ok := token.(*sql.TokenNumUnit)
	if !ok {
		msg := fmt.Sprintf("'%s' not NUM_UNIT", token.Type().Name())
		return 0, errors.New(msg)
	}

	base, err := evalUnitBaseValueTime(t.Unit.Str())
	if err != nil {
		return int64(0), err
	}

	v := int64(0)
	switch t.Num.Type() {
	case sql.T_INT:
		v = t.Num.(*sql.TokenInt).Value * base
	case sql.T_FLOAT:
		v = int64(t.Num.(*sql.TokenFloat).Value * float64(base))
	default:
		msg := fmt.Sprintf("invalid time interval '%s'", token.Str())
		return int64(0), errors.New(msg)
	}

	return v, nil
}

// evalTokenBytes() - Returns value of bytes tokens.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func evalTokenBytes(dsl *Dsl, token sql.Token, ctx *script.Cntx) (int64, error) {
	t, ok := token.(*sql.TokenNumUnit)
	if !ok {
		msg := fmt.Sprintf("'%s' not NUM_UNIT", token.Type().Name())
		return int64(0), errors.New(msg)
	}

	base, err := evalUnitBaseValueByte(t.Unit.Str())
	if err != nil {
		return int64(0), err
	}

	v := int64(0)
	switch t.Num.Type() {
	case sql.T_INT:
		v = t.Num.(*sql.TokenInt).Value * base
	case sql.T_FLOAT:
		v = int64(t.Num.(*sql.TokenFloat).Value * float64(base))
	default:
		msg := fmt.Sprintf("invalid bytes '%s'", token.Str())
		return int64(0), errors.New(msg)
	}

	return v, nil
}

// evalTokenUnitNumber() - Returns the value 'NUM UNIT' token.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
//
// This function tries to eval the token as TIME_INTERVAL token then to eval
// the token as BYTES token if any erorrs occurred.
func evalTokenUnitNumber(dsl *Dsl, token sql.Token, ctx *script.Cntx) (interface{}, error) {
	v, err := evalTokenTimeInterval(dsl, token, ctx)
	if err == nil {
		return v, nil
	}

	v, err = evalTokenBytes(dsl, token, ctx)
	if err == nil {
		return v, nil
	}

	msg := fmt.Sprintf("invalid 'NUM UNIT' token '%s'", token.Str())
	return nil, errors.New(msg)
}

func evalIPv4ToUInt32(addr string) (uint32, error) {
	if ip := net.ParseIP(addr); ip != nil {
		if ipv4 := ip.To4(); ipv4 != nil {
			ipValue := uint32(ipv4[0])
			ipValue = (ipValue << 8) + uint32(ipv4[1])
			ipValue = (ipValue << 8) + uint32(ipv4[2])
			ipValue = (ipValue << 8) + uint32(ipv4[3])

			return ipValue, nil
		}
	}

	msg := fmt.Sprintf("'%s' not IPv4 string", addr)
	return 0, errors.New(msg)
}

func evalUnit32ToIPv4(v uint32) string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		(v>>24)&0xFF, (v>>16)&0xFF, (v>>8)&0xFF, v&0xFF,
	)
}

// evalIPv4() - Convert IPv4 string to INTEGER value.
//
// @addr: String value of IP address.
func evalIPv4ToInt64(addr string) (int64, error) {
	ipValue, err := evalIPv4ToUInt32(addr)
	return int64(ipValue), err
}

func evalIPRangeStr(str string) (from, to string, err error) {
	_, ipnet, err := net.ParseCIDR(str)
	if err != nil {
		return str, str, nil
	}

	ones, bits := ipnet.Mask.Size()
	if ones <= 0 || bits != 32 {
		msg := fmt.Sprintf("%s not a IPv4 network", str)
		return "", "", errors.New(msg)
	}

	from, mask := ipnet.IP.String(), (uint32(1)<<uint(bits-ones))-1
	ipValue, _ := evalIPv4ToUInt32(from)

	return from, evalUnit32ToIPv4(ipValue | mask), nil
}

// evalTokenRangeValue() - Returns the range value of SQL token.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
func evalTokenRangeValue(dsl *Dsl, v interface{}, ctx *script.Cntx) (interface{}, error) {
	switch v.(type) {
	case sql.Token:
		switch token := v.(sql.Token); token.Type() {
		case sql.T_STAR:
			return nil, nil
		case sql.T_STR:
			return token.(*sql.TokenStr).Value, nil
		case sql.T_INT:
			return token.(*sql.TokenInt).Value, nil
		case sql.T_FLOAT:
			return token.(*sql.TokenFloat).Value, nil
		default:
			msg := fmt.Sprintf("invalid range value '%s'", token.Str())
			return nil, errors.New(msg)
		}

	case int64, int32, int16, float32, float64:
		return v, nil

	case string:
		if intValue, err := value.ToInt(v); err == nil {
			return intValue, nil
		} else if floatValue, err := value.ToFloat(v); err == nil {
			return floatValue, nil
		}
		return v, nil
	}

	msg := fmt.Sprintf("invalid range value '%s'", value.ToStr(v))
	return nil, errors.New(msg)
}

// evalTokenRange() - Returns the value of RANGE(LIST) token.
//
// @dsl:   ES DSL object.
// @token: SQL token.
// @ctx:   Script context.
//
// A range is a, like {"from": 100, "to": 200 }, dict.
func evalTokenRange(dsl *Dsl, token sql.Token, ctx *script.Cntx) (map[string]interface{}, error) {
	var fromValue, toValue, mask interface{}

	switch token.Type() {
	case sql.T_LIST:
		list := token.(*sql.TokenList).List
		if len(list) != 2 {
			msg := fmt.Sprintf("range '%s' has %d (expected 2) items", token.Str(), len(list))
			return nil, errors.New(msg)
		}
		fromValue, toValue = list[0], list[1]

	case sql.T_STR:
		str := token.(*sql.TokenStr).Value

		list := strings.Fields(str)
		if len(list) != 3 || strings.ToLower(list[1]) != "to" {
			if _, _, err := net.ParseCIDR(str); err == nil {
				mask = str
			} else {
				msg := fmt.Sprintf("invalid range '%s'", token.Str())
				return nil, errors.New(msg)
			}
		} else {
			fromValue, toValue = list[0], list[2]
		}

	default:
		msg := fmt.Sprintf("range '%s' not a STR or LIST", token.Str())
		return nil, errors.New(msg)
	}

	rangeValue := map[string]interface{}{}
	if mask == nil {
		from, err := evalTokenRangeValue(dsl, fromValue, ctx)
		if err != nil {
			msg := fmt.Sprintf("first value in range '%s' %s", token.Str(), err)
			return nil, errors.New(msg)
		}

		to, err := evalTokenRangeValue(dsl, toValue, ctx)
		if err != nil {
			msg := fmt.Sprintf("second value in range '%s' %s", token.Str(), err)
			return nil, errors.New(msg)
		}

		if from == nil && to == nil {
			msg := fmt.Sprintf("invalid range '%s'", token.Str())
			return nil, errors.New(msg)
		}

		if from != nil {
			rangeValue["from"] = from
		}

		if to != nil {
			rangeValue["to"] = to
		}
	} else {
		rangeValue["mask"] = mask
	}

	return rangeValue, nil
}

// evalTokenRangeList() - Returns the value of RANGE(LIST) token.
//
// @dsl:    ES DSL object.
// @tokens: List of SQL token.
// @ctx:    Script context.
//
// A range is a, like {"from": 100, "to": 200 }, dict.
func evalTokenRangeList(dsl *Dsl, tokens []sql.Token, ctx *script.Cntx) ([]interface{}, error) {
	list := []interface{}{}

	for cc, item := range tokens {
		rangeValue, err := evalTokenRange(dsl, item, ctx)
		if err != nil {
			msg := fmt.Sprintf("the %d range %s", cc, err)
			return nil, errors.New(msg)
		}

		list = append(list, rangeValue)
	}

	return list, nil
}

func evalCheckDatetime(now int64, ctx *script.Cntx) bool {
	oldValue := ctx.GetGlobal("__now__")
	ctx.SetGlobal("__now__", now)

	result, err := ctx.Funcs["check_datetime"]([]interface{}{}, ctx)
	if oldValue != nil && oldValue != nil {
		ctx.SetGlobal("__now__", oldValue)
	} else {
		ctx.Clear("__now__")
	}

	if err != nil {
		return true
	}

	return value.IsTrue(result)
}
