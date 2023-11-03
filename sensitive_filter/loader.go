package sensitive_filter

import (
	"bufio"
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"io"
	"os"
)

/*
   @Author: orbit-w
   @File: loader
   @2023 11月 周五 11:24
*/

type Loader struct {
	sks aho_corasick.StrKeySlice
}

func NewLoader() *Loader {
	return &Loader{
		sks: make(aho_corasick.StrKeySlice, 0),
	}
}

func (ins *Loader) build(ac *aho_corasick.AC) {
	ac.Build(ins.sks)
}

func (ins *Loader) Merge(words []string) {
	for i := range words {
		word := words[i]
		ins.sks = append(ins.sks, []rune(word))
	}
}

func (ins *Loader) LoadDict(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return ins.Load(file)
}

func (ins *Loader) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		ins.sks = append(ins.sks, []rune(string(line)))
	}
	return nil
}
