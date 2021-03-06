package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

//设定删除的路径
var Path = `/webpage/www/video_live/public/channel_resource/channel/image`

//var Path = `F:\wordpress\wp-admin`
var Qq = make(chan int)

//设定每次读取的数目
var Number int = 100

func main() {
	param := os.Args
	if len(param) < 3 {
		fmt.Println("Invalid parameter")
		return
	}
	var pathName string = param[1]
	var numberStr string = param[2]
	pathName = strings.TrimSpace(pathName)
	numberStr = strings.TrimSpace(numberStr)
	if pathName != "" {
		Path = pathName
	}
	num, _ := strconv.Atoi(numberStr)
	if num != 0 {
		Number = num
	}
	//以只读的方式打开目录
	f, err := os.Open(Path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	} else {
		fmt.Println("打开成功！")
	}
	//延迟关闭目录
	defer f.Close()
	//设定chan容量
	var ch chan string = make(chan string, 3000)
	go GetFileName(f, ch)
	go DeleteFileByName(ch)
	for {
		select {
		case <-time.After(time.Second * 60):
			fmt.Println("task is runing")
			continue
		case <-Qq:
			os.Exit(0)
		}
	}
	//for _ = range time.After(time.Second * 60) {
	//
	//}
	//for {
	//	//遇到\n结束读取
	//	nameList, err := f.Readdirnames(10000)
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//		fmt.Println("err:", err.Error())
	//	}
	//	for _, name := range nameList {
	//		fileName := Path + "/" + name
	//		err := os.Remove(fileName)
	//		if err == nil {
	//			fmt.Println("success to delete:", fileName)
	//		} else {
	//			fmt.Println("fail to delete:", fileName)
	//		}
	//	}
	//}
}

func GetFileName(file *os.File, ch chan string) {
	for {
		fmt.Println("begin to read dir names ")
		nameList, err := file.Readdirnames(Number)
		if err != nil {
			if err == io.EOF {
				Qq <- 1
				break
			}
			fmt.Println("err:", err.Error())
			continue
		}
		for _, name := range nameList {
			ch <- name
		}
	}
}

func DeleteFileByName(ch chan string) {
	var err error
	for {
		select {
		//增加超时机制
		case <-time.After(time.Second * 5):
			//超时5秒
			fmt.Println("time sleep channel length:", len(ch))
			time.Sleep(time.Second)
		case value := <-ch:
			fileName := Path + "/" + value
			err = os.Remove(fileName)
			if err == nil {
				fmt.Println("success to delete:", fileName, " channel length:", len(ch))
			} else {
				fmt.Println("fail to delete:", fileName, " channel length:", len(ch))
			}
		}
	}

}
