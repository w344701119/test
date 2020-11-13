package main

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	var counter int = 0
	var q []rune
	var arr = []rune(s)
	fmt.Println(len(arr))
	for _, value := range arr {
		q = append(q, value)
		if checkStr(q) {
		} else {
			q = q[1:]
		}
		counter = getCounter(q, counter)

	}

	return counter
}

func getCounter(runeList []rune, counter int) int {
	c := len(runeList)
	if c > counter {
		return c
	} else {
		return counter
	}
}

func checkStr(runeList []rune) bool {
	m := make(map[rune]int)
	for _, value := range runeList {
		m[value] = 1
	}
	if len(m) != len(runeList) {
		return false
	} else {
		return true
	}
}

func TestStr(t *testing.T) {

	reuslt := lengthOfLongestSubstring(" ")
	fmt.Println(reuslt)

}
