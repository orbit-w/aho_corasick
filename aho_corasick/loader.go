package aho_corasick

import (
	"bufio"
	"io"
	"os"
)

/*
   @Author: orbit-w
   @File: loader
   @2023 10月 周二 18:56
*/

func LoadDict(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	ac := new(AC)
	return Load(ac, file)
}

func Load(ac *AC, rd io.Reader) error {
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

	ac.Build(sks)
	return nil
}
