package aho_corasick_v2

/*
   @Author: orbit-w
   @File: heap
   @2023 10月 周三 19:08
*/

type NodeStack []*Node

func (i NodeStack) Length() int {
	return len(i)
}

func (i *NodeStack) Push(x *Node) {
	*i = append(*i, x)
}

func (i *NodeStack) Pop() *Node {
	old := *i
	n := len(old)
	x := old[n-1]
	*i = old[0 : n-1]
	return x
}
