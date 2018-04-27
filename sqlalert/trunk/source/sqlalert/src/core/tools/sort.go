package tools

import (
	"core/value"
	"sort"
)

type __KeySort struct {
	list     []interface{}
	key, cmp string
}

func (this *__KeySort) Len() int {
	return len(this.list)
}

func (this *__KeySort) Less(i, j int) bool {
	left, ok := value.AsDict(this.list[i])
	if !ok {
		return value.Compare(this.list[i], this.cmp, this.list[j])
	}

	right, ok := value.AsDict(this.list[j])
	if !ok {
		return value.Compare(this.list[i], this.cmp, this.list[j])
	}

	return value.Compare(left[this.key], this.cmp, right[this.key])
}

func (this *__KeySort) Swap(i, j int) {
	this.list[i], this.list[j] = this.list[j], this.list[i]
}

func Sort(list []interface{}, key, cmp string) []interface{} {
	data := &__KeySort{list: list, key: key, cmp: cmp}
	sort.Sort(data)
	return data.list
}
