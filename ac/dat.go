package ac

import (
	"github.com/orbit-w/aho_corasick/lib/math"
)

/*
   @Time: 2023/8/22 00:10
   @Author: david
   @File: dat
*/

type DAT struct {
	len   int
	cap   int
	base  []int //转移基数
	check []int //dat 映射父子节点唯一关系性
}

func (ins *DAT) Build() {
	ins.len = InitSize
	ins.base = make([]int, InitSize)
	ins.check = make([]int, InitSize)
	ins.base[0] = StateRoot
	return
}

func (ins *DAT) Find(keyword []rune) bool {
	var index int
	state := StateRoot
	for _, r := range keyword {
		i := state + int(r)
		if ins.check[i] != state {
			return false
		}
		index = i
	}
	return ins.base[index] < 0
}

func (ins *DAT) resize(in int) {
	if ins.cap >= in {
		if ins.len < in {
			ins.len = in
		}
		return
	}

	ins.cap = math.PowerOf2(in)
	ins.len = in

	ins.arrayCopy()
}

func (ins *DAT) arrayCopy() {
	newBase := make([]int, ins.cap)
	copy(newBase, ins.base)
	ins.base = newBase

	newCheck := make([]int, ins.cap)
	copy(newCheck, ins.check)
	ins.check = newCheck
}

func (ins *DAT) Length() int {
	return ins.len
}

func (ins *DAT) Cap() int {
	return ins.cap
}

//setState 更新 base[state]
func (ins *DAT) setState(index, state int, isLeaf bool) {
	if isLeaf {
		ins.base[index] = -state
	} else {
		ins.base[index] = state
	}
}

func (ins *DAT) free(index int) bool {
	if index >= ins.len {
		return false
	}
	return ins.check[index] == 0
}
