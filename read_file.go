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
var FilePath = `/webpage/www/video_live/public/channel_resource/channel/image`

var SaveFile = "output.txt"

//var Path = `F:\wordpress\wp-admin`
var Qqs = make(chan int)

//设定每次读取的数目
var Numbers int = 100

var SaveFileInfo *os.File

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
		FilePath = pathName
	}
	num, _ := strconv.Atoi(numberStr)
	if num != 0 {
		Numbers = num
	}
	//以只读的方式打开目录
	f, err := os.Open(FilePath)
	defer f.Close()
	defer SaveFileInfo.Close()
	SaveFileInfo, err = os.OpenFile(SaveFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0766)
	if err != nil {
		fmt.Println("open the file err:", err)
		os.Exit(0)
	}
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	} else {
		fmt.Println("打开成功！")
	}
	//延迟关闭目录
	defer f.Close()
	for {
		fmt.Println("begin to read dir names ")
		nameList, err := f.Readdirnames(Numbers)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			fmt.Println("err:", err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		fileList := make([]string, Numbers)
		for key, name := range nameList {
			fileList[key] = FilePath + "/" + name
		}
		_, err = SaveFileInfo.WriteString(strings.Join(fileList, "\n") + "\n")
		if err != nil {
			fmt.Println("save the file err", err)
			continue
		}
	}

}
