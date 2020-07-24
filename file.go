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
var Qq bool = false

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
	var ch chan string = make(chan string, 10000)
	go GetFileName(f, ch)
	go DeleteFileByName(ch)
	for {
		if Qq {
			time.Sleep(time.Second * 600)
			os.Exit(0)
		} else {
			time.Sleep(time.Second * 60)
			continue
		}
	}

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
		nameList, err := file.Readdirnames(Number)
		if err != nil {
			if err == io.EOF {
				Qq = true
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
	for value := range ch {
		fileName := Path + "/" + value
		err = os.Remove(fileName)
		if err == nil {
			fmt.Println("success to delete:", fileName)
		} else {
			fmt.Println("fail to delete:", fileName)
		}
	}

}
