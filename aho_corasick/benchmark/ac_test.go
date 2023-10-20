package benchmark

import (
	"fmt"
	"github.com/orbit-w/aho_corasick/aho_corasick"
	"testing"
)

/*
   @Time: 2023/8/22 08:20
   @Author: david
   @File: dat_test
*/

// h: 104, e: 101, s: 115, r: 114, i: 105
func TestAC_Fail(t *testing.T) {
	ks := aho_corasick.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	for i := range ks {
		fmt.Println(ks[i])
	}

	ac := new(aho_corasick.AC)
	ac.Build(ks)
	ac.Print()
}

func TestAC_MultiPatternSearch(t *testing.T) {
	ks := aho_corasick.StrKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	for i := range ks {
		fmt.Println(ks[i])
	}

	ac := new(aho_corasick.AC)
	ac.Build(ks)
	input := []rune("ahishers")
	patterns := ac.MultiPatternSearch(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}

	fmt.Println("===================================================")
	//中文模式串匹配
	input = []rune("听说独立日这部电影是美国人拍的")
	ks = aho_corasick.StrKeySlice{
		[]rune("独立"),
		[]rune("独立日"),
	}
	ac = new(aho_corasick.AC)
	ac.Build(ks)
	patterns = ac.MultiPatternSearch(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}
}

func Test_ACLoadAndSearch(t *testing.T) {
	aho_corasick.LoadDict("./../../data/filter_dict.txt")
}
