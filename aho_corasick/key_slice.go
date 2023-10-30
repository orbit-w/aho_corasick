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

func (ins StrKeySlice) Max() (max int) {
	for i := range ins {
		l := len(ins[i])
		if l > 0 && l > max {
			max = l
		}
	}
	return
}

func (ins StrKeySlice) Len() int {
	return len(ins)
}

func (ins StrKeySlice) Less(i, j int) bool {
	min := number_utils.Min[int](len(ins[i]), len(ins[j]))
	for k := 0; k < min; k++ {
		if ins[i][k] != ins[j][k] {
			return ins[i][k] < ins[j][k]
		}
	}
	return len(ins[i]) < len(ins[j])
}

func (ins StrKeySlice) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
}
