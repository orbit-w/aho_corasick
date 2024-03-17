package aho_corasick

import (
	"github.com/orbit-w/aho_corasick/lib/number_utils"
)

/*
   @Author: orbit-w
   @File: aho_corasick
   @2023 10月 周三 12:03
*/

type DATBuilder struct {
	cursor int
	depth  int
	max    int
}

func NewBuilder() *DATBuilder {
	return &DATBuilder{
		cursor: 0,
	}
}

func (ins *DATBuilder) insert(dat *DAT, father *Node) {
	if len(father.children) == 0 {
		return
	}

	var (
		maxV int
	)

	state := ins.heuristicState(father)
	//启发函数
	heuristic := ins.heuristicFunc()
	for {
	BEGIN:
		for i := range father.children {
			node := father.children[i]
			index := state + node.Code()
			maxV = number_utils.Max[int](maxV, index)
			if !dat.Empty(index) {
				state++
				goto BEGIN
			}
			heuristic(index)
		}
		break
	}

	dat.resize(maxV + 1)
	dat.setState(father.index, state, father.isLeaf)

	for i := range father.children {
		node := father.children[i]
		index := state + node.Code()
		//记录节点在base 中 index
		node.index = index
		//记录父子节点关系
		dat.check[index] = father.index
		dat.setState(index, state, node.isLeaf)
	}
	return
}

func (ins *DATBuilder) heuristicState(father *Node) (state int) {
	head := father.children[0]
	pos := head.Code() + StateBase
	pos = number_utils.Max[int](pos, ins.cursor)
	state = pos - head.Code()
	return
}

func (ins *DATBuilder) heuristicFunc() func(pos int) {
	var next int
	return func(pos int) {
		if next == 0 {
			next = pos
			ins.cursor = next
		}
	}
}
