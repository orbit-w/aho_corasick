package aho_corasick

import (
	"bufio"
	"io"
	"os"
	"sort"
)

/*
   @Author: orbit-w
   @File: loader
   @2023 10月 周二 18:56
*/

func LoadDict(path string) (IAhoCorasick, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return Load(file)
}

func Load(rd io.Reader) (IAhoCorasick, error) {
	ac := new(AC)
	buf := bufio.NewReader(rd)
	sks := StrKeySlice{}
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		sks = append(sks, []rune(string(line)))
	}

	ac.Build(sks)
	return ac, nil
}

func LoadDictAndBuildDat(path string) (IDat, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
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
				return nil, err
			}
			break
		}
		sks = append(sks, []rune(string(line)))
	}

	dat := new(DAT)
	sort.Sort(sks)
	trie := new(Trie)
	trie.Build(sks)
	dat.Build(trie)
	return dat, nil
}
