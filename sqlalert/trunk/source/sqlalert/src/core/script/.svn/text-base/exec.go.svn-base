// File: exec.go
//
// This file implements the execution of SCRIPT.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package script

import (
	"core/value"
	"errors"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

// Regular expression of placeholder in string.
var g_regexp_placeholder = regexp.MustCompile(`%\((\w+)\)`)

type (
	Continue struct{}
	Break    struct{}
	Return   struct{ Value interface{} }
	Exit     struct{}
)

// isInterrupt() - Test whether the value is an interruption value.
//
// @v: Value to test.
func isInterrupt(v interface{}) bool {
	switch v.(type) {
	case *Continue, *Break, *Return, *Exit:
		return true
	}

	return false
}

// execTestIndexable() - Execute token and test DICT or LIST.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execTestIndexable(token Token, ctx *Cntx) (interface{}, error) {
	v, err := Exec(token, ctx)
	if err != nil {
		return nil, err
	}

	switch v.(type) {
	case map[string]interface{}:
		return v, nil
	case []interface{}:
		return v, nil
	}

	msg := fmt.Sprintf("%s (%s) not DICT or LIST", value.ToStr(v), token.Label())
	return nil, errors.New(msg)
}

// execTestInt() - Execute token and test INT value.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execTestInt(token Token, ctx *Cntx) (int64, error) {
	v, err := Exec(token, ctx)
	if err != nil {
		return 0, err
	}

	switch v.(type) {
	case int64:
		return v.(int64), nil
	}

	msg := fmt.Sprintf("%s (%s) not INT", value.ToStr(v), token.Label())
	return 0, errors.New(msg)
}

// execTestStr() - Execute token and test INT value.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execTestStr(token Token, ctx *Cntx) (string, error) {
	v, err := Exec(token, ctx)
	if err != nil {
		return "", err
	}

	switch v.(type) {
	case string:
		return v.(string), nil
	}

	msg := fmt.Sprintf("%s (%s) not STR", value.ToStr(v), token.Label())
	return "", errors.New(msg)
}

// execNegative() - Execute token and returns the negative value.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execNegative(token Token, ctx *Cntx) (interface{}, error) {
	v, err := Exec(token, ctx)
	if err != nil {
		return nil, err
	}

	switch v.(type) {
	case int64:
		return -v.(int64), nil
	case float64:
		return -v.(float64), nil
	}

	msg := fmt.Sprintf("%s (%s) not NUM", value.ToStr(v), token.Label())
	return nil, errors.New(msg)
}

// execIdent() - Execute the IDENT token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
//
// This function returns the IDENT value or nil if not found.
func execIdent(token Token, ctx *Cntx) (interface{}, error) {
	return ctx.Get(token.(*TokenIdent).Value), nil
}

// Execution function of base values.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execInt(token Token, ctx *Cntx) (interface{}, error)   { return token.(*TokenInt).Value, nil }
func execFloat(token Token, ctx *Cntx) (interface{}, error) { return token.(*TokenFloat).Value, nil }
func execBool(token Token, ctx *Cntx) (interface{}, error)  { return token.(*TokenBool).Value, nil }
func execNull(token Token, ctx *Cntx) (interface{}, error)  { return nil, nil }

func execStr(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenStr)

	matches := g_regexp_placeholder.FindAllStringSubmatch(t.Value, -1)
	if len(matches) > 0 {
		result := t.Value

		for _, list := range matches {
			if len(list) == 2 {
				result = strings.Replace(result, list[0], value.ToStr(ctx.Get(list[1])), -1)
			}
		}

		return result, nil
	}

	return token.(*TokenStr).Value, nil
}

// execList() - Execute the LIST token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execList(token Token, ctx *Cntx) (interface{}, error) {
	result := []interface{}{}

	for _, item := range token.(*TokenList).List {
		if v, err := Exec(item, ctx); err == nil {
			result = append(result, v)
		} else {
			return nil, err
		}
	}

	return result, nil
}

// execDict() - Execute the DICT token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execDict(token Token, ctx *Cntx) (interface{}, error) {
	t, dict := token.(*TokenDict), make(map[string]interface{})

	for _, item := range t.List {
		pair := item.(*TokenPair)

		key, err := execTestStr(pair.Key, ctx)
		if err != nil {
			key = pair.Key.Str()
		}

		value, err := Exec(pair.Value, ctx)
		if err != nil {
			return nil, err
		}

		dict[key] = value
	}

	return dict, nil
}

// execOperSelf() - Execute the SELF-OPER token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execOperSelf(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenOper)

	v, err := Exec(t.Left, ctx)
	if err != nil {
		return nil, err
	}

	k := t.Left.Str()
	switch v.(type) {
	case int64:
		if t.Operator == T_INC {
			ctx.Set(k, v.(int64)+1)
			return v.(int64) + 1, nil
		} else if t.Operator == T_DEC {
			ctx.Set(k, v.(int64)-1)
			return v.(int64) - 1, nil
		}

	case float64:
		if t.Operator == T_INC {
			ctx.Set(k, v.(float64)+1)
			return v.(float64) + 1, nil
		} else if t.Operator == T_DEC {
			ctx.Set(k, v.(float64)-1)
			return v.(float64) - 1, nil
		}
	}

	msg := fmt.Sprintf("%s (%s) not NUM", value.ToStr(v), t.Left.Label())
	return nil, errors.New(msg)
}

// execOper() - Execute the OPER token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execOper(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenOper)

	if t.Left == nil && t.Operator == T_SUB {
		return execNegative(t.Right, ctx)
	}

	if t.Operator == T_INC || t.Operator == T_DEC {
		return execOperSelf(token, ctx)
	}

	left, err := Exec(t.Left, ctx)
	if err != nil {
		return nil, err
	}

	right, err := Exec(t.Right, ctx)
	if err != nil {
		return nil, err
	}

	v, err := value.Operate(left, t.Operator.Name(), right)
	if err != nil {
		msg := fmt.Sprintf("%s (%s)", err, t.Label())
		return nil, errors.New(msg)
	}

	return v, nil
}

// execComp() - Execute the COMP token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execComp(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenComp)

	left, err := Exec(t.Left, ctx)
	if err != nil {
		return nil, err
	}

	right, err := Exec(t.Right, ctx)
	if err != nil {
		return nil, err
	}

	return value.Compare(left, t.Operator.Name(), right), nil
}

// execLogical() - Execute the LOGICAL token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execLogical(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenLogical)

	if t.Operator == T_NOT {
		right, err := Exec(t.Right, ctx)
		if err != nil {
			return nil, err
		}

		return value.IsFalse(right), nil
	} else {
		left, err := Exec(t.Left, ctx)
		if err != nil {
			return nil, err
		}

		switch t.Operator {
		case T_AND:
			if value.IsFalse(left) {
				return false, nil
			}

			right, err := Exec(t.Right, ctx)
			if err != nil {
				return nil, err
			}

			return value.IsTrue(right), nil

		case T_OR:
			right, err := Exec(t.Right, ctx)
			if err != nil {
				return nil, err
			}

			return value.IsTrue(left) || value.IsTrue(right), nil
		}
	}

	return false, nil
}

// execCond() - Execute the COND token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execCond(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenCond)

	left, err := Exec(t.Left, ctx)
	if err != nil {
		return nil, err
	}

	if value.IsTrue(left) {
		return Exec(t.Middle, ctx)
	}

	return Exec(t.Right, ctx)
}

// ExecDefineFunc() - Execute the DEFINE-FUNC token.
//
// @token: Script token.
// @args:  Function arguments.
// @ctx:   Script execution Cntx.
func ExecDefineFunc(token Token, args []interface{}, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenDefine)

	if len(args) > len(t.Args) {
		msg := fmt.Sprintf("too many args (%s)", t.Label())
		return nil, errors.New(msg)
	}

	for cc, item := range t.Args {
		if cc < len(args) {
			ctx.Set(item, args[cc])
		} else {
			ctx.Clear(item)
		}
	}

	if t.Stmts != nil && len(t.Stmts) != 0 {
		if res, err := execStmtList(t.Stmts, ctx, true); err == nil {
			return res, nil
		} else {
			return nil, err
		}
	}

	return nil, nil
}

// execFunc() - Execute the FUNC token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execFunc(token Token, ctx *Cntx) (interface{}, error) {
	t, args := token.(*TokenFunc), []interface{}{}

	for _, item := range t.List {
		v, err := Exec(item, ctx)
		if err != nil {
			return nil, err
		}
		args = append(args, v)
	}

	if funcToken, ok := ctx.Defines[t.Name]; ok {
		ctx.PushAndRefer()
		res, err := ExecDefineFunc(funcToken, args, ctx)
		ctx.PopAndDefer()

		if err != nil {
			msg := fmt.Sprintf("%s\ncalled by %s", err, t.Label())
			return nil, errors.New(msg)
		}

		return res, nil
	}

	if function, ok := ctx.Funcs[strings.ToLower(t.Name)]; ok {
		res, err := function(args, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s\nin %s", err, t.Label())
			return nil, errors.New(msg)
		}

		return res, nil
	}

	msg := fmt.Sprintf("func not found %s", t.Label())
	return nil, errors.New(msg)
}

// execIndex() - Execute the INDEX token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execIndex(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenIndex)

	object, err := Exec(t.Object, ctx)
	if err != nil {
		return nil, err
	}

	switch object.(type) {
	case map[string]interface{}:
		key, err := execTestStr(t.Key, ctx)
		if err != nil {
			return nil, err
		}

		v, _ := object.(map[string]interface{})[key]
		return v, nil

	case []interface{}:
		key, err := execTestInt(t.Key, ctx)
		if err != nil {
			return nil, err
		}

		list := object.([]interface{})
		if key < int64(len(list)) {
			return list[key], nil
		}

		return nil, nil
	}

	msg := fmt.Sprintf("%s (%s) not indexable", value.ToStr(object), t.Object.Label())
	return nil, errors.New(msg)
}

// execIn() - Execute the IN token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execIn(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenIn)

	key, err := Exec(t.Key, ctx)
	if err != nil {
		return nil, err
	}

	object, err := Exec(t.Object, ctx)
	if err != nil {
		return nil, err
	}

	switch object.(type) {
	case map[string]interface{}:
		if keyStr, ok := key.(string); ok {
			_, ok = object.(map[string]interface{})[keyStr]
			return ok, nil
		}

		return false, nil

	case []interface{}:
		for _, item := range object.([]interface{}) {
			if value.Compare(key, "==", item) {
				return true, nil
			}
		}
		return false, nil
	}

	msg := fmt.Sprintf("%s (%s) not iterable", value.ToStr(object), t.Object.Label())
	return nil, errors.New(msg)
}

// execStmtList() - Execute the statement list.
//
// @list: List of tokens.
// @ctx:  Script execution context.
//
// This function exec the token list one by one. If the statement returns
// Continue it exec the next loop and if the statement returns Break it
// stop the execution.
func execStmtList(list []Token, ctx *Cntx, filter bool) (interface{}, error) {
	var lastValue interface{}

	for _, item := range list {
		result, err := Exec(item, ctx)
		if err != nil {
			return nil, err
		}

		lastValue = result
		if isInterrupt(result) {
			if retValue, ok := result.(*Return); ok && filter {
				lastValue = retValue.Value
			}
			break
		}
	}

	return lastValue, nil
}

// execStmts() - Execute the STMTS token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execStmts(token Token, ctx *Cntx) (interface{}, error) {
	v, err := execStmtList(token.(*TokenStmts).List, ctx, true)
	if isInterrupt(v) {
		if _, ok := v.(*Exit); !ok {
			return nil, err
		}
	}
	return v, err
}

// execAssign() - Execute the ASSIGN token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execAssign(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenAssign)

	right, err := Exec(t.Right, ctx)
	if err != nil {
		return nil, err
	}

	if t.Left.Type() == T_INDEX {
		left := t.Left.(*TokenIndex)

		object, err := execTestIndexable(left.Object, ctx)
		if err != nil {
			return nil, err
		}

		switch object.(type) {
		case map[string]interface{}:
			key, err := execTestStr(left.Key, ctx)
			if err != nil {
				return nil, err
			}
			object.(map[string]interface{})[key] = right

		case []interface{}:
			key, err := execTestInt(left.Key, ctx)
			if err != nil {
				return nil, err
			}

			list := object.([]interface{})
			if key < int64(len(list)) {
				list[key] = right
			}
		}
	} else {
		ctx.Set(t.Left.Str(), right)
	}

	return right, nil
}

// execExpr() - Execute the EXPR token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execExpr(token Token, ctx *Cntx) (interface{}, error) {
	return Exec(token.(*TokenExpr).Token, ctx)
}

// execIfElse() - Execute the IFELSE token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execIfElse(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenIfElse)
	ifToken := t.If.(*TokenIf)

	ifCond, err := Exec(ifToken.Cond, ctx)
	if err != nil {
		return nil, err
	}

	if value.IsTrue(ifCond) {
		return execStmtList(ifToken.List, ctx, false)
	}

	if t.ElseIf != nil {
		for _, item := range t.ElseIf {
			eiToken := item.(*TokenElseIf)

			res, err := Exec(eiToken.Cond, ctx)
			if err != nil {
				return nil, err
			}

			if value.IsTrue(res) {
				return execStmtList(eiToken.List, ctx, false)
			}
		}
	}

	if t.Else != nil {
		return execStmtList(t.Else.(*TokenElse).List, ctx, false)
	}

	return nil, nil
}

// execFor() - Execute the FOR token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execFor(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenFor)
	iter := t.Iter.(*TokenForIter)

	var lastValue interface{}
	var res interface{}
	var err error

	ctx.ReferLocal()
	if iter.Start != nil {
		res, err = Exec(iter.Start, ctx)
		if err != nil {
			lastValue = nil
			goto _EXIT
		}
	}

	for {
		res, err = Exec(iter.Cond, ctx)
		if err != nil || value.IsFalse(res) {
			lastValue = nil
			goto _EXIT
		}

		if res, err = execStmtList(t.List, ctx, false); err != nil {
			lastValue = nil
			goto _EXIT
		} else {
			lastValue = res
			switch res.(type) {
			case *Continue:
				lastValue = nil
				goto _NEXT
			case *Break:
				lastValue = nil
				goto _EXIT
			case *Exit:
				lastValue = nil
				goto _EXIT
			case *Return:
				goto _EXIT
			}
		}

	_NEXT:
		if iter.Next != nil {
			if res, err = Exec(iter.Next, ctx); err != nil {
				lastValue = nil
				goto _EXIT
			}
		}
	}

_EXIT:
	ctx.DeferLocal()
	return lastValue, err
}

// execForin() - Execute the FORIN token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execForin(token Token, ctx *Cntx) (v interface{}, e error) {
	t := token.(*TokenForin)
	iter := t.Iter.(*TokenForinIter)

	object, err := Exec(iter.Object, ctx)
	if err != nil {
		return nil, err
	}

	ctx.ReferLocal()
	switch object.(type) {
	case map[string]interface{}:
		for key, value := range object.(map[string]interface{}) {
			if iter.Key == nil {
				ctx.Set(iter.Value.Str(), key)
			} else {
				ctx.Set(iter.Key.Str(), key)
				ctx.Set(iter.Value.Str(), value)
			}

			if res, err := execStmtList(t.List, ctx, false); err != nil {
				v, e = nil, err
				goto _EXIT
			} else {
				v = res
				switch res.(type) {
				case *Continue:
					v = nil
					goto _NEXT_1
				case *Break:
					v = nil
					goto _EXIT
				case *Exit:
					v = nil
					goto _EXIT
				case *Return:
					v, e = res, nil
					// v, e = res.(*Return).Value, nil
					goto _EXIT
				default:
					v = res
				}
			}

		_NEXT_1:
			continue
		}

	case []interface{}:
		for key, value := range object.([]interface{}) {
			if iter.Key != nil {
				ctx.Set(iter.Key.Str(), int64(key))
			}
			ctx.Set(iter.Value.Str(), value)

			if res, err := execStmtList(t.List, ctx, false); err != nil {
				v, e = nil, err
				goto _EXIT
			} else {
				switch res.(type) {
				case *Continue:
					v = nil
					goto _NEXT_2
				case *Break:
					v = nil
					goto _EXIT
				case *Exit:
					v = nil
					goto _EXIT
				case *Return:
					v, e = res, nil
					//v, e = res.(*Return).Value, nil
					goto _EXIT
				default:
					v = res
				}
			}

		_NEXT_2:
			continue
		}

	default:
		msg := fmt.Sprintf("%s (%s) not iterable", value.ToStr(object), iter.Object.Label())
		v, e = nil, errors.New(msg)
	}

_EXIT:
	ctx.DeferLocal()
	return v, e
}

// execContinue()/execBreak() - Execute the CONTINUE/BREAK token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execContinue(token Token, ctx *Cntx) (interface{}, error) { return &Continue{}, nil }
func execBreak(token Token, ctx *Cntx) (interface{}, error)    { return &Break{}, nil }

// execReturn() - Execute the RETURN token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execReturn(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenReturn)

	if t.Token != nil {
		res, err := Exec(t.Token, ctx)
		if err != nil {
			return nil, err
		}

		return &Return{Value: res}, nil
	}

	return &Return{Value: nil}, nil
}

// execDefine() - Execute the DEFINE token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execDefine(token Token, ctx *Cntx) (interface{}, error) {
	t := token.(*TokenDefine)

	ctx.Defines[t.Name] = token
	return nil, nil
}

// execInclude() - Execute the INCLUDE token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execInclude(token Token, ctx *Cntx) (interface{}, error) {
	t, name := token.(*TokenInclude), ""

	if v, e := execStr(t.Name, ctx); e == nil && v != nil {
		name = value.ToStr(v)
	} else {
		name = t.Name.Str()
	}

	if !filepath.IsAbs(name) {
		etc := ctx.GetX("__etc_scripts__")
		switch etc.(type) {
		case string:
			name = filepath.Join(etc.(string), name)
		}
	}

	if _, ok := ctx.Includes[name]; !ok {
		fileToken, err := ParseFile(name)
		if err != nil {
			msg := fmt.Sprintf("%s\nincluded by %s", err, t.Label())
			return nil, errors.New(msg)
		}

		ctx.Includes[name] = fileToken
		_, err = Exec(fileToken, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s\nincluded by %s", err, t.Label())
			return nil, errors.New(msg)
		}
	}

	return nil, nil
}

// execImport() - Execute the IMPORT token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func execImport(token Token, ctx *Cntx) (interface{}, error) {
	t, name := token.(*TokenImport), ""

	if v, e := execStr(t.Name, ctx); e == nil && v != nil {
		name = value.ToStr(v)
	} else {
		name = t.Name.Str()
	}

	if !filepath.IsAbs(name) {
		etc := ctx.GetX("__etc_scripts__")
		switch etc.(type) {
		case string:
			name = filepath.Join(etc.(string), name)
		}
	}

	fileToken, err := ParseFile(name)
	if err != nil {
		msg := fmt.Sprintf("%s\nimport by %s", err, t.Label())
		return nil, errors.New(msg)
	}

	_, err = Exec(fileToken, ctx)
	if err != nil {
		msg := fmt.Sprintf("%s\nincluded by %s", err, t.Label())
		return nil, errors.New(msg)
	}

	return nil, nil
}

// Execution function for all SCRIPT token.
var g_executions_once sync.Once
var g_executions map[TokenType]func(Token, *Cntx) (interface{}, error)

func init_executions() {
	g_executions = map[TokenType]func(Token, *Cntx) (interface{}, error){
		T_IDENT:   execIdent,
		T_STR:     execStr,
		T_INT:     execInt,
		T_FLOAT:   execFloat,
		T_BOOL:    execBool,
		T_NULL:    execNull,
		T_LIST:    execList,
		T_DICT:    execDict,
		T_OPER:    execOper,
		T_COMP:    execComp,
		T_LOGICAL: execLogical,
		T_COND:    execCond,
		T_FUNC:    execFunc,
		T_INDEX:   execIndex,
		T_IN:      execIn,

		T_STMTS:    execStmts,
		T_ASSIGN:   execAssign,
		T_EXPR:     execExpr,
		T_IF_ELSE:  execIfElse,
		T_FOR:      execFor,
		T_FORIN:    execForin,
		T_CONTINUE: execContinue,
		T_BREAK:    execBreak,
		T_RETURN:   execReturn,
		T_DEFINE:   execDefine,
		T_INCLUDE:  execInclude,
		T_IMPORT:   execImport,
	}
}

// Exec() - Execute the script token.
//
// @token: Script token.
// @ctx:   Script execution Cntx.
func Exec(token Token, ctx *Cntx) (interface{}, error) {
	g_executions_once.Do(init_executions)

	if ctx == nil {
		ctx = NewContext()
	}

	function, ok := g_executions[token.Type()]
	if !ok {
		msg := fmt.Sprintf("no exec function for %s %s", token.Type().Name(), token.Label())
		return nil, errors.New(msg)
	}

	return function(token, ctx)
}
