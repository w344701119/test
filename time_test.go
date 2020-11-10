package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestDateToInt(t *testing.T) {
	var dataToInt = func(d string) int64 {
		t, _ := time.Parse("2006-01-02", d)
		return t.Unix()
	}
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	fmt.Printf("Alloc:%d HeapIdle:%d HeapReleased:%d", ms.Alloc, ms.HeapIdle, ms.HeapReleased)
	openChat()
	var ms2 runtime.MemStats
	runtime.ReadMemStats(&ms2)
	fmt.Printf("--Alloc:%d HeapIdle:%d HeapReleased:%d", ms2.Alloc, ms2.HeapIdle, ms2.HeapReleased)
	var tt = ms.HeapIdle - ms2.HeapIdle
	fmt.Println(tt)
	fmt.Println(dataToInt(""))
	//time.Sleep(10 * time.Second)
}

func openChat() {
	var f = "tb_chat_201807.json"
	fileObj, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
	}
	buf := bufio.NewReader(fileObj)
	_, bErr := buf.ReadBytes('\n')
	if bErr != nil {
		fmt.Println(bErr)
	}
}
