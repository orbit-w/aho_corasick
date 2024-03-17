package aho_corasick

/*
   @Author: orbit-w
   @File: key_slice
   @2023 10月 周二 18:56
*/

import (
	"github.com/orbit-w/aho_corasick/lib/number_utils"
)

type StrKey []rune
type StrKeySlice []StrKey

func (ins StrKeySlice) Len() int {
	return len(ins)
}

func (ins StrKeySlice) Less(i, j int) bool {
	minV := number_utils.Min[int](len(ins[i]), len(ins[j]))
	for k := 0; k < minV; k++ {
		if ins[i][k] != ins[j][k] {
			return ins[i][k] < ins[j][k]
		}
	}
	return len(ins[i]) < len(ins[j])
}

func (ins StrKeySlice) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
}
