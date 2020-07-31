package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

//设定删除的路径
var FileName = `/webpage/www/output.txt`

//var Path = `F:\wordpress\wp-admin`
var Q = make(chan int)

func main() {
	param := os.Args
	if len(param) < 2 {
		fmt.Println("Invalid parameter")
		return
	}
	var pathName string = param[1]
	var numberStr string = param[2]
	pathName = strings.TrimSpace(pathName)
	numberStr = strings.TrimSpace(numberStr)
	if pathName != "" {
		FileName = pathName
	}
	//以只读的方式打开目录
	f, err := os.Open(FileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	} else {
		fmt.Println("打开成功！")
	}
	//延迟关闭目录
	defer f.Close()
	//设定chan容量
	var ch chan string = make(chan string, 5000)
	go ReadFile(f, ch)
	go DeleteFileByPath(ch)
	for {
		select {
		case <-time.After(time.Second * 60):
			fmt.Println("task is runing")
			continue
		case <-Q:
			os.Exit(0)
		}
	}
}

func ReadFile(file *os.File, ch chan string) {
	for {
		fmt.Println("begin to read dir names ")
		buf := bufio.NewReader(file)
		for {
			//遇到\n结束读取
			b, ierr := buf.ReadBytes('\n')
			if ierr != nil {
				if ierr == io.EOF {
					Q <- 1
				}
				fmt.Println("ierr:", ierr.Error())
			}
			name := string(b)
			if strings.Index(name, "channel_resource") != -1 {
				ch <- name
			}
		}
	}
}

func DeleteFileByPath(ch chan string) {
	var err error
	for {
		select {
		//增加超时机制
		case <-time.After(time.Second * 5):
			//超时5秒
			fmt.Println("time sleep channel length:", len(ch))
			time.Sleep(time.Second)
		case value := <-ch:
			err = os.Remove(value)
			if err == nil {
				fmt.Println("success to delete:", value, " channel length:", len(ch))
			} else {
				fmt.Println("fail to delete:", value, " channel length:", len(ch))
			}
		}
	}

}
