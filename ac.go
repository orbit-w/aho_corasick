package aho_corasick

import (
	"bufio"
	"fmt"
	"github.com/orbit-w/aho_corasick/lib/number_utils"
	"io"
	"os"
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
	output []Output
}

type Output struct {
	words []Word
}

func (ins *Output) Exist() bool {
	return len(ins.words) > 0
}

func (ins *Output) Rep(i int, iter func(index int)) {
	for _, w := range ins.words {
		for j := i - (int(w.Len) - 1); j <= i; j++ {
			iter(j)
		}
	}
}

type Word struct {
	Len rune
}

func New(keywords StrKeySlice) IAhoCorasick {
	ins := new(AC)
	ins.Build(keywords)
	return ins
}

func (ins *AC) LoadDict(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return ins.loadFile(file)
}

func (ins *AC) loadFile(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	sks := StrKeySlice{}
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		sks = append(sks, []rune(string(line)))
	}

	ins.Build(sks)
	return nil
}

func (ins *AC) Build(keywords StrKeySlice) {
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

	ins.output = make([]Output, length)

	ins.build(trie)
	trie.Free()
}

func (ins *AC) Cap() int {
	return ins.cap
}

// Validate 检测字符是否合法
func (ins *AC) Validate(input []rune) bool {
	legal := true
	ins.multiPatternMatching(input, func(res Output, index int) (stop bool) {
		if res.Exist() {
			legal = false
			stop = true
		}
		return
	})
	return legal
}

func (ins *AC) FindAll(input []rune) []Result {
	patterns := make([]Result, 0)
	ins.multiPatternMatching(input, func(res Output, index int) (stop bool) {
		for i := range res.words {
			word := res.words[i]
			start := index - (int(word.Len) - 1)
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
func (ins *AC) ReplaceAll(pattern string, repl rune) string {
	in := []rune(pattern)
	ins.multiPatternMatching(in, func(res Output, index int) (stop bool) {
		for _, word := range res.words {
			for i := index - (int(word.Len) - 1); i <= index; i++ {
				in[i] = repl
			}
		}
		return
	})
	return string(in)
}

// Replace 按照贪心匹配原则从左向右匹配
// Example: 字典中 {outlearning, learnings}, 模式串 outlearnings 会匹配到 outlearning 而不会匹配到 learnings
func (ins *AC) Replace(pattern string, repl rune) string {
	index := IndexRoot
	state := StateRoot
	in := []rune(pattern)
	m := len(in)
	for i := 0; i < m; i++ {
		for j := i; j < m; j++ {
			code := int(in[j])
			t := state + code
			if !ins.Exist(t, index) {
				state = StateRoot
				index = IndexRoot
				break
			}

			index = t
			state = ins.State(index)
			op := ins.output[index]

			for _, w := range op.words {
				for cur := j - (int(w.Len) - 1); cur <= j; cur++ {
					in[cur] = repl
				}
			}
		}
	}
	return string(in)
}

func (ins *AC) multiPatternMatching(input []rune, iter func(res Output, index int) (stop bool)) {
	var (
		index = IndexRoot
		state = StateRoot
	)

	for i, r := range input {
		code := int(r)
		out := ins._FSM(S{
			Index: index,
			State: state,
			Code:  code,
		})

		switch out.State {
		case StateFail:
			state, index = StateRoot, out.Index
		default:
			state, index = out.State, out.Index
			op := ins.output[index]
			if iter(op, i) {
				return
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
			output := Output{
				words: make([]Word, len(node.output)),
			}
			for i := range node.output {
				r := node.output[i]
				output.words[i] = Word{
					Len: r,
				}
			}
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
