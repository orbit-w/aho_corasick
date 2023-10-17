package ac

/*
   @Time: 2023/8/22 00:12
   @Author: david
   @File: key_slice
*/

import (
	"github.com/orbit-w/aho_corasick/lib/number_utils"
)

type strKey []rune
type strKeySlice []strKey

func (ins strKeySlice) Max() (max int) {
	for i := range ins {
		l := len(ins[i])
		if l > 0 && l > max {
			max = l
		}
	}
	return
}

func (ins strKeySlice) Len() int {
	return len(ins)
}

func (ins strKeySlice) Less(i, j int) bool {
	min := number_utils.Min[int](len(ins[i]), len(ins[j]))
	for k := 0; k < min; k++ {
		if ins[i][k] != ins[j][k] {
			return ins[i][k] < ins[j][k]
		}
	}
	return len(ins[i]) < len(ins[j])
}

func (ins strKeySlice) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
}
