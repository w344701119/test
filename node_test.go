package main

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateNode(s []int) *ListNode {
	var head = new(ListNode)

	nums := len(s)
	var tail, curNode *ListNode
	if nums > 0 {

		head = &ListNode{s[0], curNode}
		tail = head
		for i := 1; i < nums; i++ {
			fmt.Println(i)
			tail.Next = &ListNode{s[i], curNode}
			tail = tail.Next
		}
	} else {
		head = &ListNode{0, curNode}
	}

	return head
}

func CheckArr(arr []int) []int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 0 {
			break
		} else {
			arr = arr[:i]
		}
	}
	return arr
}

func FetchArr(s1, s2 []int) []int {
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
		m, n := GetSum(number1, number2)
		var m1 int
		m1, re[i] = GetSum(re[i], n)
		re = append(re, m+m1)
	}
	return re
}

func GetSum(i, j int) (int, int) {
	t := i + j
	if t > 9 {
		return t / 10, t % 10
	} else {
		return 0, t
	}
}

func TestNodes(t *testing.T) {
	var s1 []int = []int{0}
	var s2 []int = []int{0}
	reArr := FetchArr(s1, s2)
	reArr = CheckArr(reArr)

	re := CreateNode(reArr)

	for re != nil {
		fmt.Println(re.Val)
		re = re.Next
	}

}
