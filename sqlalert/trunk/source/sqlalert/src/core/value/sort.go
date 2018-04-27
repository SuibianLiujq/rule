// Sort the list (or slice).
//
// Copyright (C) 2017 Yun Li Lai, Inc. All Rights Reserved.
// Written by: ZHANG Li Dan.
package value

import (
	"sort"
	"strings"
)

// Sorter - Structure of sorter.
//
// @list: List of emlements to sort.
// @key:  Sort key.
type Sorter struct {
	list     []interface{}
	key, cmp string
}

// Init() - Initialize the Sorter instance.
//
// @list: List of emlements to sort.
// @key:  Sort key.
func (this *Sorter) Init(list []interface{}, key, order string) *Sorter {
	this.list, this.key, this.cmp = list, key, "<"
	if strings.ToLower(order) == "desc" {
		this.cmp = ">"
	}
	return this
}

// Len() - Returns number of elements.
func (this *Sorter) Len() int {
	return len(this.list)
}

// Swap() - Swap the elements.
//
// @i, j: The index of two elements to swap.
func (this *Sorter) Swap(i, j int) {
	this.list[i], this.list[j] = this.list[j], this.list[i]
}

// Less() - Return Less() value.
//
// @i, j: The index of two elements to swap.
//
// This function returns true if elem[i] < elem[j], otherwise
// it returns false.
func (this *Sorter) Less(i, j int) bool {
	iItem, jItem := this.list[i], this.list[j]

	if IsDict(iItem) && IsDict(jItem) {
		if iDict := iItem.(map[string]interface{}); this.key != "" {
			if iValue, ok := iDict[this.key]; ok {
				if jValue, ok := jItem.(map[string]interface{})[this.key]; ok {
					return Compare(iValue, this.cmp, jValue)
				}
			}
		}
	}

	return Compare(this.list[i], this.cmp, this.list[j])
}

// Sort() - Copy & Sort list.
//
// @list: List of elements to sort.
// @key:  Sort key.
func CopySort(list []interface{}, key, order string) []interface{} {
	copyList := Copy(list).([]interface{})

	sorter := (&Sorter{}).Init(copyList, key, order)
	sort.Sort(sorter)

	return copyList
}

// Sort() - Sort list.
//
// @list: List of elements to sort.
// @key:  Sort key.
func Sort(list []interface{}, key, order string) {
	sorter := (&Sorter{}).Init(list, key, order)
	sort.Sort(sorter)
}
