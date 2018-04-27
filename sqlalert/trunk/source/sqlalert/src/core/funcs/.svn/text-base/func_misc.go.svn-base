// MISC functions.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package funcs

import (
	"bytes"
	"core/json"
	"core/script"
	"core/tools"
	"core/value"
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

// fnPrint() - Print (to stdout) all arguments.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscPrint(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	v := []interface{}{}

	for _, item := range args {
		str, err := json.ToStr(item)
		if err != nil {
			str = value.ToStr(item)
		}

		v = append(v, str)
	}

	fmt.Println(v...)
	return nil, nil
}

func miscPrintContext(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	ctx.Print()
	return nil, nil
}

// miscPrintList() - Print a list of values.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscPrintList(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) >= 1 {
		if list, ok := value.AsList(args[0]); ok {
			for _, item := range list {
				if str, err := json.ToStr(item); err == nil {
					fmt.Println(str)
				} else {
					fmt.Println(value.ToStr(item))
				}
			}
		}
	}

	return nil, nil
}

// miscPPrint() - Print (to stdout) in pretty format.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscPPrint(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	v := []interface{}{}

	for _, item := range args {
		v = append(v, json.DumpStrAll(item))
	}

	fmt.Println(v...)
	return nil, nil
}

// miscExit() - Returns EXIT value to stop the scirpt execution.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscExit(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) > 0 {
		miscPrint(args, ctx)
	}

	return &script.Exit{}, nil
}

// miscError() - Returns an error and stop the script execution.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscError(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	msg := ""
	if len(args) > 0 {
		res, _ := miscJoin([]interface{}{args, " "}, ctx)
		msg = value.ToStr(res)
	}

	return nil, errors.New(msg)
}

// miscCopy() - Returns a new copy of given argument.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscCopy(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) <= 0 {
		return nil, nil
	}

	return value.Copy(args[0]), nil
}

// miscNow() - Return timestamp of now.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscNow(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if args != nil && len(args) != 0 {
		msg := fmt.Sprintf("argument mismatch %d (expected 0)", len(args))
		return nil, errors.New(msg)
	}

	now, err := tools.GetTimeNow(ctx)
	if err != nil {
		return nil, err
	}

	return now.Timestamp, nil
}

// miscSysNow() - Return timestamp of now (system).
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscSysNow(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if args != nil && len(args) != 0 {
		msg := fmt.Sprintf("argument mismatch %d (expected 0)", len(args))
		return nil, errors.New(msg)
	}

	now, err := tools.GetTimeNow(nil)
	if err != nil {
		return nil, err
	}

	return now.Timestamp, nil
}

// miscTime() - Return time string of now.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscTime(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	spliter := " "

	if args != nil {
		switch len(args) {
		case 0:
		case 1:
			if spliter = value.ToStr(args[0]); spliter == "" {
				spliter = " "
			}
		default:
			msg := fmt.Sprintf("argument mismatch %d(expected <= 1)", len(args))
			return nil, errors.New(msg)
		}
	}

	now, err := tools.GetTimeNow(ctx)
	if err != nil {
		return nil, err
	}

	fmtStr := fmt.Sprintf("2006-01-02%s15:04:05.999-07:00", spliter)
	return now.ToStr(fmtStr), nil
}

// miscTime() - Return time string of now.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscSysTime(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	spliter := " "

	if args != nil {
		switch len(args) {
		case 0:
		case 1:
			if spliter = value.ToStr(args[0]); spliter == "" {
				spliter = " "
			}
		default:
			msg := fmt.Sprintf("argument mismatch %d(expected <= 1)", len(args))
			return nil, errors.New(msg)
		}
	}

	now, err := tools.GetTimeNow(nil)
	if err != nil {
		return nil, err
	}

	fmtStr := fmt.Sprintf("2006-01-02%s15:04:05.999-07:00", spliter)
	return now.ToStr(fmtStr), nil
}

// miscKeys() - Returns all keys of dict.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscKeys(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	dict, ok := args[0].(map[string]interface{})
	if !ok {
		msg := fmt.Sprintf("%s of args[0] not a DICT", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	for key, _ := range dict {
		list = append(list, key)
	}

	return list, nil
}

// miscValues() - Returns all values of dict.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscValues(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list := []interface{}{}
	dict, ok := args[0].(map[string]interface{})
	if !ok {
		msg := fmt.Sprintf("%s of args[0] not a DICT", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	for _, value := range dict {
		list = append(list, value)
	}

	return list, nil
}

// miscAppend() - Append value to a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscAppend(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected > 2)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := args[0].([]interface{})
	if !ok {
		msg := fmt.Sprintf("%s of args[0] not a DICT", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	for _, item := range args[1:] {
		list = append(list, item)
	}

	return list, nil
}

// miscAppendFirst() - Append value to a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscAppendFirst(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected > 2)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := args[0].([]interface{})
	if !ok {
		msg := fmt.Sprintf("%s of args[0] not a DICT", value.ToStr(args[0]))
		return nil, errors.New(msg)
	}

	newlist := make([]interface{}, 0, len(list)+len(args[1:]))
	for _, item := range args[1:] {
		newlist = append(newlist, item)
	}

	for _, item := range list {
		newlist = append(newlist, item)
	}

	return newlist, nil
}

// miscAppendList() - Append a list to anothor.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscAppendList(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected > 2)", len(args))
		return nil, errors.New(msg)
	}

	listDst, ok := value.AsList(args[0])
	if !ok {
		return nil, errors.New(fmt.Sprintf("args[0] not a list"))
	}

	listSrc, ok := value.AsList(args[1])
	if !ok {
		return nil, errors.New(fmt.Sprintf("args[1] not a list"))
	}

	for _, item := range listSrc {
		listDst = append(listDst, item)
	}

	return listDst, nil
}

// miscRemoveFirst() - Remove the first element of a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscRemoveFirst(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	if list, ok := value.AsList(args[0]); ok {
		if len(list) != 0 {
			return list[1:], nil
		} else {
			return list, nil
		}
	}

	msg := fmt.Sprintf("'%s' not a list", value.ToStr(args[0]))
	return nil, errors.New(msg)
}

// miscRemoveLast() - Remove the last element of a list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscRemoveLast(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	if list, ok := value.AsList(args[0]); ok {
		if len(list) != 0 {
			return list[0 : len(list)-1], nil
		} else {
			return list, nil
		}
	}

	msg := fmt.Sprintf("'%s' not a list", value.ToStr(args[0]))
	return nil, errors.New(msg)
}

// miscDelete() - Delete the key from a dict.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscDelete(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	if dict, ok := value.AsDict(args[0]); ok {
		key := value.ToStr(args[1])
		if _, ok := dict[key]; ok {
			delete(dict, key)
		}
	}

	return nil, nil
}

// miscIsEmpty() - Test whether a DICT or LIST is empty.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsEmpty(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.IsFalse(args[0]), nil
}

// miscIsNull() - Test whether the value is null
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsNull(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return args[0] == nil, nil
}

// miscIsInt() - Test whether the value is a integer.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsInt(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.IsInt(args[0]), nil
}

// miscIsFloat() - Test whether the value is a float number.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsFloat(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.IsFloat(args[0]), nil
}

// miscIsNum() - Test whether the value is a number.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsNum(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.IsNum(args[0]), nil
}

// miscIsStr() - Test whether the value is a string.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsStr(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.IsStr(args[0]), nil
}

// miscIsStrList() - Test whether the value is a list of strings.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsStrList(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		return false, nil
	}

	for _, item := range list {
		if !value.IsStr(item) {
			return false, nil
		}
	}

	return true, nil
}

// miscIsList() - Test whether the value is a LIST.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsList(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.IsList(args[0]), nil
}

// miscIsDict() - Test whether the value is a DICT.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsDict(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.IsDict(args[0]), nil
}

// miscIsFunc() - Test whether the value is a FUNC.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIsFunc(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	name := value.ToStr(args[0])
	if _, ok := ctx.Funcs[name]; ok {
		return true, nil
	}

	if _, ok := ctx.Defines[name]; ok {
		return true, nil
	}

	return false, nil
}

// miscLoadJson() - Load JSON file.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscLoadJson(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	etcPath, ok := ctx.GetStr("__etc_scripts__")
	if !ok {
		msg := fmt.Sprintf("Contex '__etc_scripts__' not found")
		return nil, errors.New(msg)
	}

	argPath := value.ToStr(args[0])
	if filepath.IsAbs(argPath) {
		return json.ParseFile(argPath)
	}

	return json.ParseFile(filepath.Join(etcPath, argPath))
}

// miscExecFile() - Load & execute script dynamically.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscExecFile(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	var ctxNew *script.Cntx = nil

	switch len(args) {
	case 2:
		if dict, ok := value.AsDict(args[1]); ok {
			ctxNew = ctx.CopyGlobalCtx()
			for k, v := range dict {
				ctxNew.Local[k] = v
			}
		}
		fallthrough

	case 1:
		if ctxNew == nil {
			ctxNew = ctx
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	etcPath, ok := ctx.GetStr("__etc_scripts__")
	if !ok {
		msg := fmt.Sprintf("Contex '__etc_scripts__' not found")
		return nil, errors.New(msg)
	}

	return script.ExecFile(filepath.Join(etcPath, value.ToStr(args[0])), ctxNew)
}

// miscExecAsync() - Load & execute script dynamically.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscExecAsync(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	var ctxNew *script.Cntx = nil

	switch len(args) {
	case 2:
		if dict, ok := value.AsDict(args[1]); ok {
			ctxNew = ctx.CopyGlobalCtx()
			for k, v := range dict {
				ctxNew.Local[k] = v
			}
		}
		fallthrough

	case 1:
		if ctxNew == nil {
			ctxNew = ctx.CopyAll()
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	etcPath, ok := ctx.GetStr("__etc_scripts__")
	if !ok {
		msg := fmt.Sprintf("Contex '__etc_scripts__' not found")
		return nil, errors.New(msg)
	}

	go script.ExecFile(filepath.Join(etcPath, value.ToStr(args[0])), ctxNew)
	return nil, nil
}

// miscSetGlobal() - Set value to Global context.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscSetGlobal(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	ctx.SetGlobal(value.ToStr(args[0]), args[1])
	return args[1], nil
}

// miscSetLocal() - Set value to Local context.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscSetLocal(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	ctx.SetLocal(value.ToStr(args[0]), args[1])
	return args[1], nil
}

// miscGetGlobal() - Get value from Global context.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscGetGlobal(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return ctx.GetGlobal(value.ToStr(args[0])), nil
}

// miscGetSession() - Get value from session.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscGetSession(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return ctx.GetSession(value.ToStr(args[0]), nil), nil
}

// miscSetSession() - Get value from session.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscSetSession(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	ctx.SetSession(value.ToStr(args[0]), args[1])
	return args[1], nil
}

// miscList() - Create an empty list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscList(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) > 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected <=1)", len(args))
		return nil, errors.New(msg)
	}

	size := int64(0)
	if len(args) == 1 {
		if argValue, ok := args[0].(int64); ok && argValue >= 0 {
			size = argValue
		}
	}

	list := make([]interface{}, size)
	for cc := int64(0); cc < size; cc++ {
		list[cc] = nil
	}

	return list, nil
}

func miscListToDict(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	list, keys, spliter, ok := []interface{}{}, []interface{}{}, "", false

	switch len(args) {
	case 3:
		if args[2] != nil {
			spliter = value.ToStr(args[2])
		}
		fallthrough

	case 2:
		if keys, ok = value.AsList(args[1]); !ok {
			keys = append(keys, value.ToStr(args[1]))
		}

		if list, ok = value.AsList(args[0]); !ok {
			msg := fmt.Sprintf("args[0] not a list")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 3)", len(args))
		return nil, errors.New(msg)
	}

	dictRes := map[string]interface{}{}
	for _, item := range list {
		if res, err := miscJoinValues([]interface{}{item, keys, spliter}, ctx); err == nil {
			key := value.ToStr(res)
			if _, ok = dictRes[key]; ok {
				dictRes[key] = append(value.List(dictRes[key]), item)
			} else {
				dictRes[key] = []interface{}{item}
			}
		}
	}

	return dictRes, nil
}

// miscListSlice() - Return the slice of given list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscListSlice(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	argList, argFrom, argTo, ok := []interface{}(nil), int64(-1), int64(-1), false

	switch len(args) {
	case 3:
		if argTo, ok = value.AsInt(args[2]); !ok {
			return nil, errors.New(fmt.Sprintf("args[2] not a integer"))
		}

		fallthrough

	case 2:
		if argFrom, ok = value.AsInt(args[1]); !ok {
			return nil, errors.New(fmt.Sprintf("args[1] not a integer"))
		}

		fallthrough

	case 1:
		if argList, ok = value.AsList(args[0]); !ok {
			return nil, errors.New(fmt.Sprintf("args[0] not a list"))
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 3)", len(args))
		return nil, errors.New(msg)
	}

	if argFrom < 0 {
		argFrom = 0
	}

	if argFrom > int64(len(argList))-1 {
		return []interface{}{}, nil
	}

	if argTo < 0 || argTo > int64(len(argList)) {
		argTo = int64(len(argList))
	}

	return argList[argFrom:argTo], nil
}

// miscStr() - Convert value to STRING.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscStr(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.ToStr(args[0]), nil
}

// miscLen() - Returns the length of the object.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscLen(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return value.Len(args[0]), nil
}

// miscJoin() - Join the values together.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscJoin(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	list, spliter, ok := []interface{}{}, "", false
	switch len(args) {
	case 2:
		spliter = value.ToStr(args[1])
		fallthrough
	case 1:
		if list, ok = value.AsList(args[0]); !ok {
			return nil, errors.New(fmt.Sprintf("arg[0] not a list"))
		}

	default:
		return nil, errors.New(fmt.Sprintf("argument mismatch %d (expected 1 or 2)", len(args)))
	}

	buffer := &bytes.Buffer{}
	if len(list) != 0 {
		buffer.WriteString(value.ToStr(list[0]))

		for _, item := range list[1:] {
			buffer.WriteString(spliter)
			buffer.WriteString(value.ToStr(item))
		}
	}

	return string(buffer.String()), nil
}

// miscJoinValues() - Join the values of a dict together.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscJoinValues(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	dict, keys, spliter, ok := map[string]interface{}{}, []interface{}{}, "", false

	switch len(args) {
	case 3:
		if args[2] != nil {
			spliter = value.ToStr(args[2])
		}
		fallthrough

	case 2:
		if keys, ok = value.AsList(args[1]); !ok {
			keys = append(keys, args[1])
		}

		if dict, ok = value.AsDict(args[0]); !ok {
			msg := fmt.Sprintf("arg[0] not a dict")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 3)", len(args))
		return nil, errors.New(msg)
	}

	listRes := []interface{}{}
	for _, item := range keys {
		if v, ok := dict[value.ToStr(item)]; ok {
			listRes = append(listRes, v)
		}
	}

	return miscJoin([]interface{}{listRes, spliter}, ctx)
}

// miscGetContex() - Return the value of the context item.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscGetContex(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return ctx.GetX(value.ToStr(args[0])), nil
}

// miscCheckDateTime() - Test whether now is 'on' depends on the __workdays__.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscCheckDateTime(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	var cfg interface{}

	switch len(args) {
	case 0:
	case 1:
		cfg = args[0]
	default:
		msg := fmt.Sprintf("argument mismatch %d (expected <=1 )", len(args))
		return nil, errors.New(msg)
	}

	return tools.DTCheck(nil, cfg, ctx)
}

// miscCallBuiltin() - Call a function and retruns the return value.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscCall(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected >1)", len(args))
		return nil, errors.New(msg)
	}

	name := value.ToStr(args[0])
	if function, ok := ctx.Funcs[name]; ok {
		return function(args[1:], ctx)
	}

	if token, ok := ctx.Defines[name]; ok {
		ctx.PushAndRefer()
		result, err := script.ExecDefineFunc(token, args[1:], ctx)
		ctx.PopAndDefer()

		return result, err
	}

	msg := fmt.Sprintf("function '%v' not found", args[0])
	return nil, errors.New(msg)
}

type __MResult struct {
	Error  error
	Result interface{}
}

func __runRoutine(name string, args interface{}, ctx *script.Cntx, ch chan __MResult) {
	callArgs := []interface{}{name}

	if args != nil {
		switch value.Type(args) {
		case value.LIST:
			for _, item := range value.List(args) {
				callArgs = append(callArgs, item)
			}
		default:
			callArgs = append(callArgs, args)
		}
	}

	result, err := miscCall(callArgs, ctx)
	ch <- __MResult{Error: err, Result: result}
}

// miscMultiCall() - Call a function in multi-routines and retruns the return value.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscMultiCall(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	funcName, funcArgList, nCall := "", []interface{}{}, 0

	switch len(args) {
	case 3:
		if list, ok := value.AsList(args[2]); ok {
			funcArgList = list
		} else {
			funcArgList = append(funcArgList, args[2])
		}
		fallthrough

	case 2:
		if vInt, err := value.ToInt(args[1]); err == nil {
			nCall = int(vInt)
		} else {
			return nil, err
		}
		fallthrough

	case 1:
		if str, ok := value.AsStr(args[0]); ok {
			funcName = str
		} else {
			msg := fmt.Sprintf("arg[0] not a string value: %v", args[0])
			return nil, errors.New(msg)
		}
		break

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	if nCall == 0 {
		nCall = len(funcArgList)
	}

	if nCall == 0 {
		nCall = 1
	}

	chanList := []chan __MResult{}
	for cc := 0; cc < nCall; cc++ {
		ch, funcCtx, funcArgs := make(chan __MResult, 1), ctx.CopyAll(), interface{}(nil)
		chanList = append(chanList, ch)

		if cc < len(funcArgList) {
			funcArgs = funcArgList[cc]
		}

		go __runRoutine(funcName, funcArgs, funcCtx, ch)
	}

	result, err := []interface{}{}, error(nil)
	for cc := 0; cc < nCall; cc++ {
		res := <-chanList[cc]
		if res.Error != nil {
			err = res.Error
		}

		result = append(result, res.Result)
	}

	return result, err
}

func __runFile(name string, args map[string]interface{}, ctx *script.Cntx, ch chan error) {
	for k, v := range args {
		ctx.Local[k] = v
	}

	_, err := miscExecFile([]interface{}{name}, ctx)
	ch <- err
}

// miscMultiExec() - Run scripts in multi-routines.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscMultiExec(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	listFile, dictArgs, ok := []interface{}{}, map[string]interface{}{}, false

	switch len(args) {
	case 2:
		if dictArgs, ok = value.AsDict(args[1]); !ok {
			msg := fmt.Sprintf("arg[1] not a dict")
			return nil, errors.New(msg)
		}
		fallthrough

	case 1:
		if listFile, ok = value.AsList(args[0]); !ok {
			msg := fmt.Sprintf("arg[0] not a list")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 1 or 2)", len(args))
		return nil, errors.New(msg)
	}

	chanList := []chan error{}
	for _, name := range listFile {
		ctxNew, ch := ctx.CopyGlobalCtx(), make(chan error, 1)
		chanList = append(chanList, ch)

		go __runFile(value.ToStr(name), dictArgs, ctxNew, ch)
	}

	errRes := error(nil)
	for _, ch := range chanList {
		if err := <-ch; err != nil {
			errRes = err
		}
	}

	return nil, errRes
}

func miscMultiQuery(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	result, err := miscMultiCall(args, ctx)
	if err != nil {
		return nil, err
	}

	listRes := []interface{}{}
	for _, listItem := range value.List(result) {
		for _, item := range value.List(listItem) {
			listRes = append(listRes, item)
		}
	}

	return listRes, nil
}

// miscCallBuiltin() - Call builtin function and retruns the return value.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscCallBuiltin(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) < 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected >1)", len(args))
		return nil, errors.New(msg)
	}

	function, ok := ctx.Funcs[value.ToStr(args[0])]
	if !ok {
		msg := fmt.Sprintf("builtin '%v' not found", args[0])
		return nil, errors.New(msg)
	}

	return function(args[1:], ctx)
}

// miscCallList() - Call functions one-by-one.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscCallList(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		msg := fmt.Sprintf("args[0] not a list")
		return nil, errors.New(msg)
	}

	result, err := interface{}(nil), error(nil)
	for cc, item := range list {
		subArgs := []interface{}{}

		if value.IsStr(item) {
			subArgs = append(subArgs, item)
			subArgs = append(subArgs, result)
		} else if dict, ok := value.AsDict(item); ok {
			if _, ok := dict["name"]; !ok {
				err = errors.New(fmt.Sprintf("option 'name' not found in item[%d]", cc))
				break
			}

			subArgs = append(subArgs, dict["name"])
			subArgs = append(subArgs, result)
			if option, ok := dict["args"]; ok {
				subArgs = append(subArgs, option)
			}
		} else {
			err = errors.New(fmt.Sprintf("item[%d] not an valid function list option"))
			break
		}

		if result, err = miscCall(subArgs, ctx); err != nil {
			return nil, err
		}
	}

	return result, err
}

// miscIpIn() - Check whether IP address in the given list.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIpIn(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 2 {
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	if ipaddr, ok := value.AsStr(args[0]); ok {
		iplist := []string{}
		switch args[1].(type) {
		case string:
			iplist = append(iplist, args[1].(string))
		case []interface{}:
			for _, item := range args[1].([]interface{}) {
				if str, ok := value.AsStr(item); ok {
					iplist = append(iplist, str)
				}
			}
		}

		return tools.IPInList(ipaddr, iplist)
	}

	return false, nil
}

// miscHasFunc() - Check whether given function is defined.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscHasFunc(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	name := value.ToStr(args[0])
	if _, ok := ctx.Defines[name]; ok {
		return true, nil
	}

	if _, ok := ctx.Funcs[name]; ok {
		return true, nil
	}

	return false, nil
}

// miscHasPrefix() - Check whether given string has the prefix string.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscHasPrefix(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	str, prefix := "", ""

	switch len(args) {
	case 2:
		str, prefix = value.ToStr(args[0]), value.ToStr(args[1])

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	if args[1] == nil || prefix == "" {
		return true, nil
	}

	if args[0] == nil || str == "" {
		return false, nil
	}

	return strings.HasPrefix(str, prefix), nil
}

// miscHasSuffix() - Check whether given string has the suffix string.
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscHasSuffix(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	str, suffix := "", ""

	switch len(args) {
	case 2:
		str, suffix = value.ToStr(args[0]), value.ToStr(args[1])

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2)", len(args))
		return nil, errors.New(msg)
	}

	if args[1] == nil || suffix == "" {
		return true, nil
	}

	if args[0] == nil || str == "" {
		return false, nil
	}

	return strings.HasSuffix(str, suffix), nil
}

// miscIPMap() - Map IP values from to another (anything).
//
// @args: Arguments for function.
// @ctx:  Script execution context.
func miscIPMap(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	list, key, dict, newkey, ok := []interface{}{}, "", map[string]interface{}{}, "", false

	switch len(args) {
	case 4:
		newkey = value.ToStr(args[3])
		fallthrough

	case 3:
		if list, ok = value.AsList(args[0]); !ok {
			msg := fmt.Sprintf("arg[0] not a list")
			return nil, errors.New(msg)
		}

		key = value.ToStr(args[1])
		if dict, ok = value.AsDict(args[2]); !ok {
			msg := fmt.Sprintf("arg[2] not a dict")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 3 or 4)", len(args))
		return nil, errors.New(msg)
	}

	if newkey == "" {
		newkey = key
	}

	for _, item := range list {
		if dictItem, ok := value.AsDict(item); ok {
			if itemValue, ok := dictItem[key]; ok {
				if mapValue, ok := dict[value.ToStr(itemValue)]; ok {
					dictItem[newkey] = mapValue
				}
			}
		}
	}

	return list, nil
}

func miscListEmpty(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	if list, ok := value.AsList(args[0]); ok && len(list) > 0 {
		return false, nil
	}

	return true, nil
}

func miscListNotEmpty(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	if list, ok := value.AsList(args[0]); ok && len(list) > 0 {
		return true, nil
	}

	return false, nil
}

func miscSort(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	key, list, ok := "", []interface{}{}, false

	switch len(args) {
	case 2:
		key = value.ToStr(args[1])
		fallthrough

	case 1:
		if list, ok = value.AsList(args[0]); !ok {
			return args[0], nil
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 1)", len(args))
		return nil, errors.New(msg)
	}

	return tools.Sort(list, key, "<"), nil
}

func miscSortReverse(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	key, list, ok := "", []interface{}{}, false

	switch len(args) {
	case 2:
		key = value.ToStr(args[1])
		fallthrough

	case 1:
		if list, ok = value.AsList(args[0]); !ok {
			return args[0], nil
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 1)", len(args))
		return nil, errors.New(msg)
	}

	return tools.Sort(list, key, ">"), nil
}

func miscSplit(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	str, spliter, list, ok := "", "", []interface{}{}, false

	switch len(args) {
	case 2:
		spliter = value.ToStr(args[1])
		fallthrough

	case 1:
		str, ok = value.AsStr(args[0])
		if !ok {
			msg := fmt.Sprintf("args[0] not a string")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 2 or 1)", len(args))
		return nil, errors.New(msg)
	}

	if spliter == "" {
		for _, item := range strings.Fields(str) {
			list = append(list, item)
		}
	} else {
		for _, item := range strings.Split(str, spliter) {
			list = append(list, item)
		}
	}

	return list, nil
}

func miscTrim(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	return strings.TrimSpace(value.ToStr(args[0])), nil
}

func miscReverse(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	if len(args) != 1 {
		msg := fmt.Sprintf("argument mismatch %d (expected 1)", len(args))
		return nil, errors.New(msg)
	}

	list, ok := value.AsList(args[0])
	if !ok {
		msg := fmt.Sprintf("args[0] not a list")
		return nil, errors.New(msg)
	}

	newlist, length := make([]interface{}, 0, len(list)), len(list)
	for cc := 0; cc < length; cc++ {
		newlist = append(newlist, list[(length-1)-cc])
	}

	return newlist, nil
}

func miscRun(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	result, err := miscCall(args, ctx)
	if err != nil {
		return nil, err
	}

	if result != nil {
		if listRules, ok := ctx.GetXList("__sub_rules__"); ok {
			args := map[string]interface{}{"__alert_result__": result}
			return miscMultiExec([]interface{}{listRules, args}, ctx)
		}
	}

	return result, err
}

func miscDictGet(args []interface{}, ctx *script.Cntx) (interface{}, error) {
	dict, key, defValue, ok := map[string]interface{}{}, "", interface{}(nil), false

	switch len(args) {
	case 3:
		defValue = args[2]
		fallthrough

	case 2:
		key = value.ToStr(args[1])
		dict, ok = value.AsDict(args[0])
		if !ok {
			msg := fmt.Sprintf("args[0] not a dict")
			return nil, errors.New(msg)
		}

	default:
		msg := fmt.Sprintf("argument mismatch %d (expected 3 or 2)", len(args))
		return nil, errors.New(msg)
	}

	if _, ok := dict[key]; ok {
		return dict[key], nil
	}

	return defValue, nil
}
