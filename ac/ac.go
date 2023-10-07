package ac

import (
	"fmt"
	"sort"
)

/*
   @Time: 2023/9/28 10:46
   @Author: david
   @File: ac
*/

type AC struct {
	dat    *DAT
	fail   []int
	output map[int][]rune
}

func (ins *AC) Init(keywords strKeySlice) {
	ins.fail = make([]int, InitSize)
	ins.output = make(map[int][]rune)

	sort.Sort(keywords)
	trie := new(Trie)
	trie.Build(keywords)

	ins.dat = new(DAT)
	ins.dat.Build()

	ins.Build(trie)
}

func (ins *AC) Find(keyword []rune) bool {
	return ins.dat.Find(keyword)
}

func (ins *AC) MultiPatternSearch() {

}

func (ins *AC) Build(trie *Trie) {
	trie.BFS(func(father *Node) (stop bool) {
		if len(father.children) == 0 {
			return
		}
		//确定最佳转移基数
		dat := ins.dat
		state := dat.state(father.index)
		for {
		COMPLETE:
			for i := range father.children {
				node := father.children[i]
				pos := state + int(node.code)
				if dat.base[pos] != 0 {
					state++
					goto COMPLETE
				}
			}
			break
		}

		dat.setState(father.index, state, father.isLeaf)
		for i := range father.children {
			node := father.children[i]
			index := state + int(node.code)
			//记录节点在base 中 index
			node.index = index
			//记录父子节点关系
			dat.setCheck(father, index)
			dat.setState(index, state, node.isLeaf)
			//将Fail Node 压缩到fail slice 中
			ins.fail[index] = node.fail.index
			if node.Exist() {
				output := make([]rune, len(node.output))
				copy(output, node.output)
				ins.output[index] = output
			}
		}
		return
	})
}

func (ins *AC) setOutput(node *Node, index int) {
	if node.Exist() {
		output := make([]rune, len(ins.output))
		copy(output, node.output)
		ins.output[index] = output
	}
}

func (ins *AC) Print() {
	fmt.Println("fail: ", ins.fail)
	fmt.Println("output: ", ins.output)
}
