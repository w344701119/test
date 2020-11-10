package main

import (
	"fmt"
	"testing"
)

func TestAddNumber(t *testing.T) {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{3, 4, 5, 6, 9}
	re := fetchArr(s1, s2)
	fmt.Println(re)
}

func fetchArr(s1, s2 []int) []int {
	num1 := len(s1)
	num2 := len(s2)
	loop := num1
	if num2 > num1 {
		loop = num2
	}
	var re []int = []int{0}
	for i := 0; i < loop; i++ {
		var number1 = 0
		var number2 = 0
		if num1 > i {
			number1 = s1[i]
		}
		if num2 > i {
			number2 = s2[i]
		}
		m, n := getSum(number1, number2)
		var m1 int
		m1, re[i] = getSum(re[i], n)
		re = append(re, m+m1)
	}
	return re
}

func getSum(i, j int) (int, int) {
	t := i + j
	if t > 9 {
		return t / 10, t % 10
	} else {
		return 0, t
	}
}
