// File: helper.go
//
// This file implements helper functions to use JSON parser.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package json

import (
	"bytes"
	"core/sys"
	"core/value"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Parse() - Parse JSON from given @src bytes.
//
// @src: Source input byte stream.
func Parse(src []byte, name string) (interface{}, error) {
	lexer, err := (&Lexer{}).Init(src, name)
	if err != nil {
		return nil, err
	}

	if yyParse(lexer) != 0 {
		msg := lexer.errMsg
		if lexer.scanner.GotError {
			msg = lexer.scanner.Error
		}

		return nil, errors.New(msg)
	}

	return lexer.value, nil
}

// Parse() - Parse JSON from given @name file.
//
// @name: String of file name.
func ParseFile(name string) (interface{}, error) {
	if content, err := sys.ReadFile(name); err == nil {
		return Parse(content, name)
	} else {
		return nil, err
	}
}

// ToBytes() - Convert @v to JSON string.
//
// @v: Value to convert.
func ToStr(v interface{}) (string, error) {
	if res, err := ToBytes(v); err == nil {
		return string(res), nil
	} else {
		return "", err
	}
}

// ToBytes() - Convert @v to JSON bytes.
//
// @v: Value to convert.
func ToBytes(v interface{}) ([]byte, error) {
	switch v.(type) {
	case []interface{}, map[string]interface{}:
		return json.Marshal(v)
	}
	return nil, errors.New("not LIST or DICT")
}

// DumpBytes() - Convert @v to JSON bytes in pretty format.
//
// @v: Value to convert.
func DumpStr(v interface{}) (string, error) {
	if res, err := DumpBytes(v); err == nil {
		return string(res), nil
	} else {
		return "", err
	}
}

// DumpBytes() - Convert @v to JSON bytes in pretty format.
//
// @v: Value to convert.
func DumpBytes(v interface{}) ([]byte, error) {
	return dump(v, "", true, true)
}

// DumpStrAll() - Convert @v to JSON string.
//
// @v: Value to convert.
//
// This functions calls to DumpStr() to dump JSON object.
// If there is any error this function calls to fmt.Sprintf()
// to format the given value.
func DumpStrAll(v interface{}) string {
	if res, err := DumpBytes(v); err == nil {
		return string(res)
	}

	return fmt.Sprintf("%v", v)
}

// PPrint() - Print @v to STDOUT in pretty format.
//
// @v: Value to print.
func PPrint(v interface{}) {
	if res, err := DumpStr(v); err == nil {
		fmt.Println(res)
	} else {
		fmt.Printf("%v\n", v)
	}
}

// MapTo() - Map JSOn object to structure.
//
// @v: Value to map.
// @o: Structure instance.
func MapTo(v, o interface{}) error {
	res, err := ToBytes(v)
	if err != nil {
		return err
	}

	return json.Unmarshal(res, o)
}

// dump() - Innter implementation of DumpBytes().
//
// @v:         Value to convert.
// @prefix:    Prefix string.
// @hasPrefix: Flag to set prefix.
// @hasSuffix: Flag to set suffix ('\n').
func dump(v interface{}, prefix string, hasPrefix, hasSuffix bool) ([]byte, error) {
	buffer := &bytes.Buffer{}
	if hasPrefix {
		buffer.WriteString(prefix)
	}

	if v == nil {
		buffer.WriteString("null")
	} else {
		switch v.(type) {
		case int, int8, int16, int32, int64, float32, float64, bool:
			buffer.WriteString(value.ToStr(v))
		case string:
			buffer.WriteString(strconv.Quote(v.(string)))
		case []interface{}:
			res, _ := dumpList(v, prefix, hasPrefix, hasSuffix)
			buffer.Write(res)
		case map[string]interface{}:
			res, _ := dumpDict(v, prefix, hasPrefix, hasSuffix)
			buffer.Write(res)
		default:
			msg := fmt.Sprintf("'%v' not JSON object", v)
			return nil, errors.New(msg)
		}
	}

	if hasSuffix {
		buffer.WriteByte('\n')
	}

	return buffer.Bytes(), nil
}

// dumpList() - Innter implementation of dump().
//
// @v:         Value to convert.
// @prefix:    Prefix string.
// @hasPrefix: Flag to set prefix.
// @hasSuffix: Flag to set suffix ('\n').
func dumpList(v interface{}, prefix string, hasPrefix, hasSuffix bool) ([]byte, error) {
	list, buffer, next := value.List(v), &bytes.Buffer{}, prefix+"  "

	if len(list) == 0 {
		buffer.WriteString("[]")
	} else {
		buffer.WriteString("[\n")

		if res, err := dump(list[0], next, true, false); err == nil {
			buffer.Write(res)
		} else {
			return nil, err
		}

		for _, item := range list[1:] {
			buffer.WriteString(",\n")
			if res, err := dump(item, next, true, false); err == nil {
				buffer.Write(res)
			} else {
				return nil, err
			}
		}

		buffer.WriteByte('\n')
		buffer.WriteString(prefix)
		buffer.WriteByte(']')
	}

	return buffer.Bytes(), nil
}

// dumpDict() - Innter implementation of dump().
//
// @v:         Value to convert.
// @prefix:    Prefix string.
// @hasPrefix: Flag to set prefix.
// @hasSuffix: Flag to set suffix ('\n').
func dumpDict(v interface{}, prefix string, hasPrefix, hasSuffix bool) ([]byte, error) {
	dict, buffer, next := value.Dict(v), &bytes.Buffer{}, prefix+"  "

	if len(dict) == 0 {
		buffer.WriteString("{}")
	} else {
		buffer.WriteString("{\n")

		first := true
		for key, item := range dict {
			if !first {
				buffer.WriteString(",\n")
			}

			first = false
			if res, err := dump(key, next, true, false); err == nil {
				buffer.Write(res)
			} else {
				return nil, err
			}

			buffer.WriteString(": ")
			if res, err := dump(item, next, false, false); err == nil {
				buffer.Write(res)
			} else {
				return nil, err
			}
		}

		buffer.WriteByte('\n')
		buffer.WriteString(prefix)
		buffer.WriteByte('}')
	}

	return buffer.Bytes(), nil
}
