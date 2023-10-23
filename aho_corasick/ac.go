package aho_corasick

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/lib/number_utils"
	"sort"
)

/*
   @Author: orbit-w
   @File: ac
   @2023 10月 周二 18:56
*/

type AC struct {
	cap    int
	base   []int //转移基数
	check  []int //dat 映射父子节点唯一关系性
	fail   []int
	output map[int][]rune
}

func New(keywords StrKeySlice) IAhoCorasick {
	ins := new(AC)
	ins.Build(keywords)
	return ins
}

func (ins *AC) Cap() int {
	return ins.cap
}

func (ins *AC) Build(keywords StrKeySlice) {
	ins.output = make(map[int][]rune)

	sort.Sort(keywords)
	trie := new(Trie)
	trie.Build(keywords)

	dat := new(DAT)
	dat.Build(trie)

	length := dat.Length()
	ins.cap = length
	ins.base = make([]int, length)
	copy(ins.base, dat.base)
	ins.check = make([]int, length)
	copy(ins.check, dat.check)

	ins.build(trie)
}

func (ins *AC) Validate(input []rune) bool {
	legal := true
	ins.multiPatternMatching(input, func(res []rune, index int) (stop bool) {
		legal = false
		stop = true
		return
	})
	return legal
}

type Result struct {
	Pattern []rune
	Start   int
}

func (ins *AC) FindAll(input []rune) []Result {
	patterns := make([]Result, 0)
	ins.multiPatternMatching(input, func(res []rune, index int) (stop bool) {
		for i := range res {
			r := res[i]
			start := index - (int(r) - 1)
			patterns = append(patterns, Result{
				Pattern: input[start : index+1],
				Start:   start,
			})
		}
		return
	})
	return patterns
}

func (ins *AC) Replace(input []rune, repl rune) {
	ins.multiPatternMatching(input, func(res []rune, index int) (stop bool) {
		for _, r := range res {
			for i := index - (int(r) - 1); i <= index; i++ {
				input[i] = repl
			}
		}
		return
	})
}

func (ins *AC) multiPatternMatching(input []rune, iter func(res []rune, index int) (stop bool)) {
	var (
		index int
		state int
	)
	state = StateRoot
	for i, r := range input {
		code := int(r)
		for {
			t := state + code
			if ins.Exist(t, index) {
				index = t
				state = ins.State(index)
				goto M
			}

			if state == StateRoot {
				index = IndexRoot
				break
			}

			index = ins.fail[index]
			state = ins.State(index)
		}
		continue
	M:
		if info, ok := ins.output[index]; ok {
			if iter(info, i) {
				return
			}
		}
	}
}

func (ins *AC) State(s int) int {
	return number_utils.ABS[int](ins.base[s])
}

func (ins *AC) Exist(i, state int) bool {
	if i >= ins.cap {
		return false
	}
	return ins.check[i] == state
}

func (ins *AC) build(trie *Trie) {
	ins.fail = make([]int, ins.cap)
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
