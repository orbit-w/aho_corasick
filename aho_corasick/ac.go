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
	trie.Free()
}

func (ins *AC) Cap() int {
	return ins.cap
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

// ReplaceAll 将匹配到的所有字符全部替换成 repl
func (ins *AC) ReplaceAll(input []rune, repl rune) {
	ins.multiPatternMatching(input, func(res []rune, index int) (stop bool) {
		for _, r := range res {
			for i := index - (int(r) - 1); i <= index; i++ {
				input[i] = repl
			}
		}
		return
	})
}

func (ins *AC) Replace(input []rune, repl rune) {
	var (
		index = IndexRoot
		state = StateRoot
		left  = 0
	)

	for i := 0; i < len(input); i++ {
		code := int(input[i]) + 1
		t := state + code
		if !ins.Exist(t, index) {
			state = StateRoot
			index = IndexRoot
			i = left
			left++
			continue
		}

		index = t
		state = ins.State(index)
		if info, ok := ins.output[index]; ok {
			for _, r := range info {
				for j := i - (int(r) - 1); j <= i; j++ {
					input[j] = repl
				}
			}
		}
	}
}

func (ins *AC) multiPatternMatching(input []rune, iter func(res []rune, index int) (stop bool)) {
	var (
		index = IndexRoot
		state = StateRoot
	)

	for i, r := range input {
		code := int(r) + 1
		out := ins._FSM(S{
			Index: index,
			State: state,
			Code:  code,
		})

		switch out.State {
		case StateFail:
			state = StateRoot
			index = out.Index
		default:
			state = out.State
			index = out.Index
			if info, ok := ins.output[index]; ok {
				if iter(info, i) {
					return
				}
			}
		}
	}
}

func (ins *AC) _FSM(in S) (out S) {
	state := in.State
	index := in.Index
	for {
		t := state + in.Code
		if ins.Exist(t, index) {
			out.Index = t
			out.State = ins.State(out.Index)
			return
		}

		if state == StateRoot {
			out.State = StateFail
			out.Index = IndexRoot
			return
		}

		index = ins.fail[index]
		state = ins.State(index)
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
	for i := range ins.base {
		if i != 0 && ins.base[i] != 0 {
			fmt.Println(i)
		}
	}
}
