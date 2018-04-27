// File: context.go
//
// This file implements the context of execution of SCRIPT.
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by ZHANG Li Dan <lidan.zhang@clearclouds-global.com>.
package script

import (
	stack "container/list"
	"core/value"
	"fmt"
	"strings"
	"sync"
)

// Type of builtin functions.
type Func func([]interface{}, *Cntx) (interface{}, error)

// Cntx - Script execution context.
//
// @Global: Global context.
// @Local:  Local context.
//
// The VARIABLE token seach the context to find the variable value
// by call to Get() method.
//
// The ASSIGN statements set the variable value by call to Set() method.
type Cntx struct {
	Global   map[string]interface{}
	Local    map[string]interface{}
	LocalRef int

	Session *sync.Map
	//	Session map[string]interface{}

	Funcs    map[string]Func
	Defines  map[string]Token
	Includes map[string]Token

	CtxGlobal *Cntx
	Stack     *stack.List
}

// Init() - Initialize Context instance.
//
// This function returns the instance itself for chain operation.
func (this *Cntx) Init() *Cntx {
	this.Global = make(map[string]interface{})
	this.Local = make(map[string]interface{})
	this.LocalRef = 0

	this.Funcs = make(map[string]Func)
	this.Defines = make(map[string]Token)
	this.Includes = make(map[string]Token)

	this.Stack = (&stack.List{}).Init()
	return this
}

// GetLocal()/GetGlobal() - Return value from local/global context.
//
// @k: Variable name.
func (this *Cntx) GetLocal(k string) interface{}  { v, _ := this.Local[k]; return v }
func (this *Cntx) GetGlobal(k string) interface{} { v, _ := this.Global[k]; return v }

// Get() - Get variable value by name.
//
// @k: Variable name.
//
// This function seach the local context first for the given name
// and then seach the global context if @name not found in local
// context.
//
// It return nil if @name neither found in local context nor found
// in global context.
func (this *Cntx) Get(k string) interface{} {
	if v, ok := this.Local[k]; ok {
		return v
	}

	if v, ok := this.Global[k]; ok {
		return v
	}

	return nil
}

// GetXxx() - Get variable value and convert to corresponding type.
//
// @k: Variable name.
func (this *Cntx) GetInt(k string) (int64, bool)     { v, ok := this.Get(k).(int64); return v, ok }
func (this *Cntx) GetFloat(k string) (float64, bool) { v, ok := this.Get(k).(float64); return v, ok }
func (this *Cntx) GetStr(k string) (string, bool)    { v, ok := this.Get(k).(string); return v, ok }
func (this *Cntx) GetBool(k string) (bool, bool)     { v, ok := this.Get(k).(bool); return v, ok }
func (this *Cntx) GetDict(k string) (map[string]interface{}, bool) {
	v, ok := this.Get(k).(map[string]interface{})
	return v, ok
}
func (this *Cntx) GetList(k string) ([]interface{}, bool) {
	v, ok := this.Get(k).([]interface{})
	return v, ok
}

// GetFirst() - Get first variable value by given list.
//
// @list: List of variable name.
//
// This function seach the local context first for the given name
// and then seach the global context if @name not found in local
// context.
//
// This function will returns the value at the first matching and
// it will return NULL if none key matches.
func (this *Cntx) GetFirst(list []string) interface{} {
	for _, item := range list {
		if v, ok := this.Local[item]; ok {
			return v
		}

		if v, ok := this.Global[item]; ok {
			return v
		}
	}

	return nil
}

// GetFirstXxx() - Get first variable value by given list and
//                 convert to corresponding type.
//
// @k: Variable name.
func (this *Cntx) GetFirstInt(list []string) (int64, bool) {
	v, ok := this.GetFirst(list).(int64)
	return v, ok
}
func (this *Cntx) GetFirstFloat(list []string) (float64, bool) {
	v, ok := this.GetFirst(list).(float64)
	return v, ok
}
func (this *Cntx) GetFirstStr(list []string) (string, bool) {
	v, ok := this.GetFirst(list).(string)
	return v, ok
}
func (this *Cntx) GetFirstBool(list []string) (bool, bool) {
	v, ok := this.GetFirst(list).(bool)
	return v, ok
}
func (this *Cntx) GetFirstDict(list []string) (map[string]interface{}, bool) {
	v, ok := this.GetFirst(list).(map[string]interface{})
	return v, ok
}
func (this *Cntx) GetFirstList(list []string) ([]interface{}, bool) {
	v, ok := this.GetFirst(list).([]interface{})
	return v, ok
}

// GetX() - Get variable value by name ignore case.
//
// @k: Variable name.
//
// This function seach the local context first for the given name
// and then seach the global context if @name not found in local
// context.
//
// It return nil if @name neither found in local context nor found
// in global context.
func (this *Cntx) GetX(k string) interface{} {
	k = strings.ToLower(k)

	for key, item := range this.Local {
		if strings.ToLower(key) == k {
			return item
		}
	}

	for key, item := range this.Global {
		if strings.ToLower(key) == k {
			return item
		}
	}

	return nil
}

// GetXXxx() - Get variable value by given list ignore case and convert to corresponding type.
//
// @k: Variable name.
func (this *Cntx) GetXInt(k string) (int64, bool)     { v, ok := this.GetX(k).(int64); return v, ok }
func (this *Cntx) GetXFloat(k string) (float64, bool) { v, ok := this.GetX(k).(float64); return v, ok }
func (this *Cntx) GetXStr(k string) (string, bool)    { v, ok := this.GetX(k).(string); return v, ok }
func (this *Cntx) GetXBool(k string) (bool, bool)     { v, ok := this.GetX(k).(bool); return v, ok }
func (this *Cntx) GetXDict(k string) (map[string]interface{}, bool) {
	v, ok := this.GetX(k).(map[string]interface{})
	return v, ok
}
func (this *Cntx) GetXList(k string) ([]interface{}, bool) {
	v, ok := this.GetX(k).([]interface{})
	return v, ok
}

// GetFirstX() - Get variable value by given list ignore case.
//
// @list: List of variable name.
//
// This function seach the local context first for the given name
// and then seach the global context if @name not found in local
// context.
//
// This function will returns the value at the first matching and
// it will return NULL if none key matches.
func (this *Cntx) GetFirstX(list []string) interface{} {
	for _, str := range list {
		str = strings.ToLower(str)

		for key, item := range this.Local {
			if strings.ToLower(key) == str {
				return item
			}
		}

		for key, item := range this.Global {
			if strings.ToLower(key) == str {
				return item
			}
		}
	}

	return nil
}

// GetFirstXxx() - Get first variable value by given list and
//                 convert to corresponding type.
//
// @k: Variable name.
func (this *Cntx) GetFirstXInt(list []string) (int64, bool) {
	v, ok := this.GetFirstX(list).(int64)
	return v, ok
}
func (this *Cntx) GetFirstXFloat(list []string) (float64, bool) {
	v, ok := this.GetFirstX(list).(float64)
	return v, ok
}
func (this *Cntx) GetFirstXStr(list []string) (string, bool) {
	v, ok := this.GetFirstX(list).(string)
	return v, ok
}
func (this *Cntx) GetFirstXBool(list []string) (bool, bool) {
	v, ok := this.GetFirstX(list).(bool)
	return v, ok
}
func (this *Cntx) GetFirstXDict(list []string) (map[string]interface{}, bool) {
	v, ok := this.GetFirstX(list).(map[string]interface{})
	return v, ok
}
func (this *Cntx) GetFirstXList(list []string) ([]interface{}, bool) {
	v, ok := this.GetFirstX(list).([]interface{})
	return v, ok
}

// Set() - Set variable value.
//
// @k:  Variable name.
// @v: Variable value.
//
// This function set the value to local context if it is found in local
// context. If the variable is found in global but not found in local
// context this function set the value to global context.
//
// It set value to local context if @name is neither found in local
// context nor found in global context.
func (this *Cntx) SetLocal(k string, v interface{})  { this.Local[k] = v }
func (this *Cntx) SetGlobal(k string, v interface{}) { this.Global[k] = v }
func (this *Cntx) Set(k string, v interface{}) {
	if this.LocalRef > 0 {
		this.Local[k] = v
	} else {
		if _, ok := this.Local[k]; ok {
			this.Local[k] = v
		}

		this.Global[k] = v
	}
}

// Push() - Push local context to stack.
func (this *Cntx) PushAndRefer() { this.Push(); this.ReferLocal() }
func (this *Cntx) Push() {
	this.Stack.PushBack(this.Local)
	this.Local = value.CopyDict(this.Local)
	//this.Local = make(map[string]interface{})
}

// Pop() - Pop local context from stack.
func (this *Cntx) PopAndDefer() { this.Pop(); this.DeferLocal() }
func (this *Cntx) Pop() {
	local := this.Stack.Back()
	this.Stack.Remove(local)
	this.Local = local.Value.(map[string]interface{})
}

// Refer & Derefer local context.
func (this *Cntx) ReferLocal() { this.LocalRef++ }
func (this *Cntx) DeferLocal() {
	this.LocalRef--
	if this.LocalRef < 0 {
		this.LocalRef = 0
	}
}

// Clear() - Clear context variable.
//
// @key: Varialbe name.
func (this *Cntx) Clear(k string) {
	if _, ok := this.Global[k]; ok {
		delete(this.Global, k)
	}

	if _, ok := this.Local[k]; ok {
		delete(this.Local, k)
	}
}

// Get/Set value from/to session.
//
// @k: String of key.
// @v: Value to set.
func (this *Cntx) SetSession(k string, v interface{}) { this.Session.Store(k, v) }
func (this *Cntx) GetSession(k string, v interface{}) interface{} {
	res, _ := this.Session.LoadOrStore(k, v)
	return res
}

func (this *Cntx) Copy() *Cntx {
	ctx := &Cntx{}
	ctx.Global = value.CopyDict(this.Global)
	ctx.Local = value.CopyDict(this.Local)

	ctx.Funcs = make(map[string]Func)
	ctx.Defines = make(map[string]Token)
	ctx.Includes = make(map[string]Token)
	ctx.Session = this.Session
	ctx.Stack = (&stack.List{}).Init()
	ctx.CtxGlobal = this.CtxGlobal

	return ctx
}

func (this *Cntx) CopyAll() *Cntx {
	ctx := this.Copy()

	for k, v := range this.Funcs {
		ctx.Funcs[k] = v
	}
	for k, v := range this.Defines {
		ctx.Defines[k] = v
	}
	for k, v := range this.Includes {
		ctx.Includes[k] = v
	}

	return ctx
}

func (this *Cntx) CopyGlobalCtx() *Cntx {
	ctx := this.CtxGlobal.CopyAll()
	ctx.Session = this.Session

	if val := this.GetX("__now__"); val != nil {
		ctx.Set("__now__", val)
	}

	if val := this.GetX("__now_fmt__"); val != nil {
		ctx.Set("__now_fmt__", val)
	}

	return ctx
}

// For debug.
func (this *Cntx) Print() {
	fmt.Println()

	if len(this.Global) == 0 {
		fmt.Println("Context-Global: empty")
	} else {
		fmt.Println("Context-Global:")
		for key, item := range this.Global {
			fmt.Printf("%s = %v\n", key, value.ToStr(item))
		}
	}

	if len(this.Local) == 0 {
		fmt.Println("\nContext-Local: empty")
	} else {
		fmt.Println("\nContext-Local:")
		for key, item := range this.Local {
			fmt.Printf("%s = %v\n", key, value.ToStr(item))
		}
	}
}

// NewContext() - Create script execution context.
//
// This function returns a new empty script execution context.
func NewContext() *Cntx {
	return (&Cntx{}).Init()
}
