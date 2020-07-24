package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

type NormalNumber struct {
	Number int `json:"number"`
	Num    int `json:"num"`
}

var Times = 9999999

type NormalSlice []NormalNumber

func (p NormalSlice) Len() int {
	return len(p)
}
func (p NormalSlice) Less(i, j int) bool {
	return p[i].Num < p[j].Num
}

func (p NormalSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func TestRand(t *testing.T) {
	for i := 0; i < 2; i++ {
		result := make([]int, Times, Times)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < Times; i++ {
			num := r.Intn(34)
			if num == 0 {
				i--
				continue
			}
			result[i] = num
		}
		list := CountList(result)
		//sort.Sort(sort.Reverse(list))
		sort.Sort(list)
		var s []int
		for j := 0; j < 6; j++ {
			s = append(s, list[j].Number)
		}
		sort.Ints(s)
		s = append(s, SelectOne())
		t.Log(s)
	}
	//jsonData, _ := json.Marshal(list)
	//t.Log(string(jsonData))
}

func SelectOne() int {
	result := make([]int, Times, Times)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < Times; i++ {
		num := r.Intn(17)
		if num == 0 {
			i--
			continue
		}
		result[i] = num
	}
	list := CountList(result)
	return list[0].Number
}

func CountList(s []int) (countSlice NormalSlice) {
	tMap := make(map[int]*NormalNumber)
	if len(s) > 0 {
		for _, value := range s {
			if _, ok := tMap[value]; ok {
				tMap[value].Num++
			} else {
				tMap[value] = &NormalNumber{Number: value, Num: 1}
			}
		}
	}
	if len(tMap) > 0 {
		tSlice := make([]NormalNumber, 0, len(tMap)*2)
		for _, value := range tMap {
			tSlice = append(tSlice, *value)
		}
		countSlice = NormalSlice(tSlice)
	}
	return
}
