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

func (ins *AC) Build(keywords strKeySlice) {
	ins.output = make(map[int][]rune)

	sort.Sort(keywords)
	trie := new(Trie)
	trie.Build(keywords)

	dat := new(DAT)
	dat.Build(trie)

	length := dat.Length()
	ins.size = length
	ins.base = make([]int, length)
	copy(ins.base, dat.base)
	ins.check = make([]int, length)
	copy(ins.check, dat.check)

	ins.build(trie)
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

func (ins *AC) build(trie *Trie) {
	ins.fail = make([]int, ins.size)
	trie.BFS(func(node *Node) (stop bool) {
		if node.Root() {
			return
		}
		//构建失配指针
		trie.f(node)
		//压缩fail
		ins.fail[node.index] = node.fail.index

		if node.Exist() {
			output := make([]rune, len(node.output))
			copy(output, node.output)
			ins.output[node.index] = output
		}
		return
	})
}

func (ins *AC) Print() {
	fmt.Println("base: ", ins.base)
	fmt.Println("check: ", ins.check)
	fmt.Println("fail: ", ins.fail)
	fmt.Println("output: ", ins.output)
}
