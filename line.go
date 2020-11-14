package main

import (
	"fmt"
	"strconv"
)

func main() {
	re := compressString("aabcccccaaa")
	fmt.Println(re)

}

func compressString(S string) string {
	m := make(map[rune]int)
	for _, value := range []rune(S) {
		if _, ok := m[value]; ok {
			m[value]++
		} else {
			m[value] = 1
		}
	}
	var re = ""
	for k, v := range m {
		re += string(k) + strconv.Itoa(v)
	}
	return re
}
