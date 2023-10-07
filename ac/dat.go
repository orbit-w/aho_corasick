package ac

import (
	"fmt"
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
	ins.base = make([]int, InitSize)
	ins.check = make([]int, InitSize)
	ins.base[0] = RootState
	return
}

func (ins *DAT) Find(keyword []rune) bool {
	var index int
	for _, r := range keyword {
		temp := ins.state(index) + int(r)
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

func (ins *DAT) state(index int) int {
	return number_utils.ABS[int](ins.base[index])
}

//setState 更新 base[state]
func (ins *DAT) setState(index, state int, isLeaf bool) {
	if isLeaf {
		ins.base[index] = -state
	} else {
		ins.base[index] = state
	}
}

//setCheck 更新check[index]
func (ins *DAT) setCheck(father *Node, i int) {
	if father.Root() {
		ins.check[i] = RootIndex
	} else {
		ins.check[i] = father.index
	}
}

func Print(dat *DAT) {
	fmt.Println("base: ", dat.base)
	fmt.Println("check: ", dat.check)
}
