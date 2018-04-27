package dsl

import (
	"core/script"
	"core/sql"
	"core/value"
	"errors"
	"fmt"
	"sync"
)

// execList() - Execute the LIST token.
//
// @token: SQL token.
// @ctx:   Script context.
func execList(token sql.Token, ctx *script.Cntx) (interface{}, error) {
	t := token.(*sql.TokenList)

	list := []interface{}{}
	for cc, item := range t.List {
		itemValue, err := execToken(item, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in %d item '%s'", err, cc, item.Str())
			return nil, errors.New(msg)
		}

		list = append(list, itemValue)
	}

	return list, nil
}

// execDict() - Execute the DICT token.
//
// @token: SQL token.
// @ctx:   Script context.
func execDict(token sql.Token, ctx *script.Cntx) (interface{}, error) {
	t := token.(*sql.TokenDict)

	dict := map[string]interface{}{}
	for cc, item := range t.List {
		pair := item.(*sql.TokenPair)

		k, err := execToken(pair.Key, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in '%s' in the %d item", err, pair.Key.Str(), cc)
			return nil, errors.New(msg)
		}

		key, ok := k.(string)
		if !ok {
			msg := fmt.Sprintf("%s (%s) not STR", value.ToStr(k), pair.Key.Str())
			return nil, errors.New(msg)
		}

		v, err := execToken(pair.Value, ctx)
		if err != nil {
			msg := fmt.Sprintf("%s in '%s' in the %d item", err, pair.Value.Str(), cc)
			return nil, errors.New(msg)
		}

		dict[key] = v
	}

	return dict, nil
}

// Execution function for all SQL token.
var g_token_funcs_once sync.Once
var g_token_funcs map[sql.TokenType]func(sql.Token, *script.Cntx) (interface{}, error)

func init_token_funcs() {
	g_token_funcs = map[sql.TokenType]func(sql.Token, *script.Cntx) (interface{}, error){
		sql.T_LIST: execList,
		sql.T_DICT: execDict,
	}
}

func execToken(token sql.Token, ctx *script.Cntx) (interface{}, error) {
	g_token_funcs_once.Do(init_token_funcs)

	switch token.Type() {
	case sql.T_IDENT:
		return token.(*sql.TokenIdent).Value, nil
	case sql.T_VAR:
		return ctx.Get(token.(*sql.TokenVar).Value), nil
	case sql.T_STR:
		return token.(*sql.TokenStr).Value, nil
	case sql.T_INT:
		return token.(*sql.TokenInt).Value, nil
	case sql.T_FLOAT:
		return token.(*sql.TokenFloat).Value, nil
	case sql.T_BOOL:
		return token.(*sql.TokenBool).Value, nil
	case sql.T_NULL:
		return nil, nil
	}

	function, ok := g_token_funcs[token.Type()]
	if !ok {
		msg := fmt.Sprintf("no exec function for '%s(%s)'", token.Type().Name(), token.Str())
		return nil, errors.New(msg)
	}

	return function(token, ctx)
}
