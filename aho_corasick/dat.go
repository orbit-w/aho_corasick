package aho_corasick

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/lib/math"
	"github.com/orbit-w/aho_corasick/lib/number_utils"
)

/*
   @Author: orbit-w
   @File: dat
   @2023 10月 周二 18:56
*/

type DAT struct {
	len   int
	cap   int   // 底层数据的真实容量
	base  []int // 转移基数
	check []int // dat 映射父子节点唯一关系性
}

func (ins *DAT) init() {
	ins.cap = InitSize
	ins.base = make([]int, InitSize)
	ins.check = make([]int, InitSize)
	ins.base[0] = StateRoot
	return
}

func (ins *DAT) Build(trie *Trie) {
	ins.init()
	trie.BFS(func(father *Node) (stop bool) {
		if len(father.children) == 0 {
			return
		}

		//确定最佳转移基数
		state := ins.State(father.index)
		var max int
		for {
		BEGIN:
			for i := range father.children {
				node := father.children[i]
				index := state + int(node.code)
				max = number_utils.Max[int](max, index)
				if !ins.Empty(index) {
					state++
					goto BEGIN
				}
			}
			break
		}

		ins.resize(max + 1)
		ins.setState(father.index, state, father.isLeaf)

		for i := range father.children {
			node := father.children[i]
			index := state + int(node.code)
			//记录节点在base 中 index
			node.index = index
			//记录父子节点关系
			ins.check[index] = father.index
			ins.setState(index, state, node.isLeaf)
		}
		return
	})
}

func (ins *DAT) Find(keyword []rune) bool {
	var index int
	for _, r := range keyword {
		i := ins.State(index) + int(r)
		if ins.check[i] != index {
			return false
		}
		index = i
	}
	return ins.base[index] < 0
}

func (ins *DAT) Length() int {
	return ins.len
}

func (ins *DAT) Cap() int {
	return ins.cap
}

func (ins *DAT) State(s int) int {
	return number_utils.ABS[int](ins.base[s])
}

func (ins *DAT) Empty(s int) bool {
	if s >= ins.len {
		return true
	}

	return ins.check[s] == 0 && ins.base[s] == 0
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

	ins.malloc()
}

func (ins *DAT) malloc() {
	newBase := make([]int, ins.cap)
	copy(newBase, ins.base)
	ins.base = newBase

	newCheck := make([]int, ins.cap)
	copy(newCheck, ins.check)
	ins.check = newCheck
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

func (ins *DAT) Print() {
	fmt.Println("Base: ", ins.base)
	fmt.Println("Check: ", ins.check)
}
