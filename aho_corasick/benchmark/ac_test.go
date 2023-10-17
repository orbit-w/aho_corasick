package ac

import (
	"fmt"
	"testing"
)

/*
   @Time: 2023/8/22 08:20
   @Author: david
   @File: dat_test
*/

// h: 104, e: 101, s: 115, r: 114, i: 105
func TestAC_Fail(t *testing.T) {
	ks := strKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	for i := range ks {
		fmt.Println(ks[i])
	}

	ac := new(AC)
	ac.Build(ks)
	ac.Print()
}

func TestAC_MultiPatternSearch(t *testing.T) {
	ks := strKeySlice{
		[]rune("he"),
		[]rune("she"),
		[]rune("hers"),
		[]rune("his"),
	}

	for i := range ks {
		fmt.Println(ks[i])
	}

	ac := new(AC)
	ac.Build(ks)
	input := []rune("ahishers")
	patterns := ac.MultiPatternSearch(input)
	for _, r := range patterns {
		fmt.Println(string(r.Pattern))
		fmt.Println(r.Start)
	}
}

func TestDAT_Fetch(t *testing.T) {
	keywords := [][]rune{
		strKey("abcf"),
		strKey("abc"),
		strKey("abcd"),
		strKey("abed"),
		strKey("abfdh"),
	}

	ks := strKeySlice{}
	for _, keyword := range keywords {
		var dk strKey = keyword
		ks = append(ks, dk)
	}

	ac := new(AC)
	ac.Build(ks)
}

func toRune(s string) rune {
	switch s {
	case "清":
		return 1
	case "华":
		return 2
	case "大":
		return 3
	case "学":
		return 4
	case "新":
		return 5
	case "中":
		return 6
	case "人":
		return 7
	}
	return 0
}
