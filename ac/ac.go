package ac

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/lib/number_utils"
	"sort"
)

/*
   @Time: 2023/9/28 10:46
   @Author: david
   @File: ac
*/

type AC struct {
	size   int
	base   []int //转移基数
	check  []int //dat 映射父子节点唯一关系性
	fail   []int
	output map[int][]rune
}

func (ins *AC) Init(keywords strKeySlice) {
	ins.fail = make([]int, InitSize)
	ins.output = make(map[int][]rune)

	sort.Sort(keywords)
	trie := new(Trie)
	trie.Build(keywords)

	ins.Build(trie)
}

type Result struct {
	Pattern []rune
	Start   int
}

//MultiPatternSearch 多模式匹配
func (ins *AC) MultiPatternSearch(input []rune) []Result {
	var (
		index int //Node 数组下角标
		state int //Node 转移因子
	)
	patterns := make([]Result, 0)
	state = StateRoot
	for i, r := range input {
		code := int(r)
		for {
			t := state + code
			if ins.exist(t, state) {
				index = t
				state = number_utils.ABS[int](ins.base[index])
				goto M
			}

			if state == StateRoot {
				index = IndexRoot
				break
			}

			index = ins.fail[index]
			state = number_utils.ABS[int](ins.base[index])
		}
	M:
		if info, ok := ins.output[index]; ok {
			for _, l := range info {
				start := i - (int(l) - 1)
				patterns = append(patterns, Result{
					Pattern: input[start : i+1],
					Start:   start,
				})
			}
		}
	}
	return patterns
}

func (ins *AC) exist(index, state int) bool {
	if index >= ins.size {
		return false
	}
	return ins.check[index] == state
}

func (ins *AC) Build(trie *Trie) {
	dat := new(DAT)
	dat.Build()
	trie.BFS(func(father *Node) (stop bool) {
		if len(father.children) == 0 {
			return
		}

		//确定最佳转移基数
		state := number_utils.ABS[int](dat.base[father.index])
		for {
		COMPLETE:
			for i := range father.children {
				node := father.children[i]
				pos := state + int(node.code)
				if dat.check[pos] != 0 {
					state++
					goto COMPLETE
				}
			}
			break
		}

		dat.setState(father.index, state, father.isLeaf)

		for i := range father.children {
			node := father.children[i]
			trie.f(node)

			index := state + int(node.code)
			//记录节点在base 中 index
			node.index = index
			//记录父子节点关系
			dat.check[index] = state
			dat.setState(index, state, node.isLeaf)
			//将 Fail Node 压缩到 fail slice 中
			ins.fail[index] = node.fail.index
			if node.Exist() {
				output := make([]rune, len(node.output))
				copy(output, node.output)
				ins.output[index] = output
			}
		}
		return
	})

	ins.base = dat.base
	ins.check = dat.check
	ins.size = dat.size
}

func (ins *AC) setOutput(node *Node, index int) {
	if node.Exist() {
		output := make([]rune, len(ins.output))
		copy(output, node.output)
		ins.output[index] = output
	}
}

func (ins *AC) Print() {
	fmt.Println("base: ", ins.base)
	fmt.Println("check: ", ins.check)
	fmt.Println("fail: ", ins.fail)
	fmt.Println("output: ", ins.output)
}
