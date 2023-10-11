package ac

import (
	"github.com/orbit-w/aho_corasick/lib/number_utils"
)

/*
   @Time: 2023/8/22 00:10
   @Author: david
   @File: dat
*/

type DAT struct {
	size  int   //容量
	base  []int //转移基数
	check []int //dat 映射父子节点唯一关系性
}

func (ins *DAT) Build() {
	ins.size = InitSize
	ins.base = make([]int, InitSize)
	ins.check = make([]int, InitSize)
	ins.base[0] = StateRoot
	return
}

func (ins *DAT) Find(keyword []rune) bool {
	var index int
	for _, r := range keyword {
		temp := number_utils.ABS[int](ins.base[index]) + int(r)
		//TODO: 有问题！
		if ins.check[temp] != index {
			return false
		}
		index = temp
	}
	return ins.base[index] < 0
}

//Resize TODO: 调整容量方式？如何小范围调整不合适？
func (ins *DAT) Resize(size int) {
	if size <= ins.size {
		return
	}

	//deep copy
	newBase := make([]int, size)
	copy(newBase, ins.base)
	ins.base = newBase

	newCheck := make([]int, size)
	copy(newCheck, ins.check)
	ins.check = newCheck

	ins.size += size
}

func (ins *DAT) Size() int {
	return ins.size
}

//setState 更新 base[state]
func (ins *DAT) setState(index, state int, isLeaf bool) {
	if isLeaf {
		ins.base[index] = -state
	} else {
		ins.base[index] = state
	}
}
