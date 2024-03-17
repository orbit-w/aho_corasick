package aho_corasick

import (
	"bufio"
	"fmt"
	"github.com/orbit-w/aho_corasick/lib/math"
	"github.com/orbit-w/aho_corasick/lib/number_utils"
	"io"
	"os"
	"sort"
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

func (ins *DAT) LoadDict(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	buf := bufio.NewReader(file)
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

	sort.Sort(sks)
	trie := new(Trie)
	trie.Build(sks)
	ins.Build(trie)
	return nil
}

func (ins *DAT) init() {
	ins.cap = InitSize
	ins.base = make([]int, InitSize)
	ins.check = make([]int, InitSize)
	ins.base[IndexRoot] = StateRoot
	return
}

func (ins *DAT) Build(trie *Trie) {
	ins.init()
	builder := NewBuilder()
	trie.BFS(func(father *Node) (stop bool) {
		builder.insert(ins, father)
		return
	})
}

func (ins *DAT) Find(keyword []rune) bool {
	var index int = IndexRoot
	for _, r := range keyword {
		i := ins.getState(index) + int(r)
		if !ins.exist(i, index) {
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

func (ins *DAT) exist(i, state int) bool {
	if i >= ins.cap {
		return false
	}
	return ins.check[i] == state
}

func (ins *DAT) getState(i int) int {
	return number_utils.ABS[int](ins.base[i])
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

// setState 更新 base[state]
func (ins *DAT) setState(index, state int, isLeaf bool) {
	if isLeaf {
		ins.base[index] = -state
	} else {
		ins.base[index] = state
	}
}

func (ins *DAT) Print() {
	fmt.Println("Base: ", ins.base)
	fmt.Println("Check: ", ins.check)
}
