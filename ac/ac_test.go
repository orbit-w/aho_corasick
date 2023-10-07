package ac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
   @Time: 2023/8/22 08:20
   @Author: david
   @File: dat_test
*/

func TestAC_AutomationInit(t *testing.T) {
	keywords := [][]string{
		{"清", "华"},
		{"清", "华", "大", "学"},
		{"清", "新"},
		{"中", "华"},
		{"华", "人"},
	}
	ks := strKeySlice{}

	for _, keyword := range keywords {
		dk := make(strKey, 0)
		for _, k := range keyword {
			dk = append(dk, toRune(k))
		}
		ks = append(ks, dk)
	}
	ac := new(AC)
	ac.Init(ks)
	Print(ac.dat)
	ac.Print()
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
	ac.Init(ks)
}

func TestDAT_Find(t *testing.T) {
	keywords := [][]string{
		{"清", "华"},
		{"清", "华", "大", "学"},
		{"清", "新"},
		{"中", "华"},
		{"华", "人"},
	}
	ks := strKeySlice{}
	for _, keyword := range keywords {
		dk := make(strKey, 0)
		for _, k := range keyword {
			dk = append(dk, toRune(k))
		}
		ks = append(ks, dk)
	}
	ac := new(AC)
	ac.Init(ks)

	assert.True(t, ac.Find([]rune{1, 2}))       //"清", "华"
	assert.False(t, ac.Find([]rune{1, 6}))      //"清", "中"
	assert.True(t, ac.Find([]rune{1, 5}))       //"清", "新"
	assert.True(t, ac.Find([]rune{6, 2}))       //“中”, "华"
	assert.True(t, ac.Find([]rune{2, 7}))       //"华", "人"
	assert.True(t, ac.Find([]rune{1, 2, 3, 4})) //"清", "华", "大", "学"
}

func toString(b rune) string {
	switch b {
	case 1:
		return "清"
	case 2:
		return "华"
	case 3:
		return "大"
	case 4:
		return "学"
	case 5:
		return "新"
	case 6:
		return "中"
	case 7:
		return "人"
	}
	return ""
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
