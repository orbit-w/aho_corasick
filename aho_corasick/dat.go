package aho_corasick

import (
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
	cap   int
	base  []int //转移基数
	check []int //dat 映射父子节点唯一关系性
}

func (ins *DAT) init() {
	ins.len = InitSize
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
		state := number_utils.ABS[int](ins.base[father.index])
		var max int
		for {
		COMPLETE:
			for i := range father.children {
				node := father.children[i]
				pos := state + int(node.code)
				max = number_utils.Max[int](max, pos)
				if pos < ins.len && ins.check[pos] != 0 {
					state++
					goto COMPLETE
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
			ins.check[index] = state
			ins.setState(index, state, node.isLeaf)
		}
		return
	})
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
