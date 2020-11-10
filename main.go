package main

import (
	"encoding/xml"
	"fmt"
)

type SConfig struct {
	Name       xml.Name    `xml:"tv"` // 指定最外层的标签为config
	Programmes []Programme `xml:"programme"`
}

type Programme struct {
	Start  string `xml:"start,attr"`
	Stop   string `xml:"stop,attr"`
	Remark string `xml:"channel,attr"`
	Title  string `xml:"title"`
	Desc   string `xml:"desc"`
}

func main() {

	var s []byte = []byte{'a', 'b', 'c', 'd'}

	fmt.Println(s)

	modifySlice(s)

	fmt.Println(s)

	//file, err := os.Open("IN.xml") // For read access.
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//v := SConfig{}
	//err = xml.Unmarshal(data, &v)
	//if err != nil {
	//	fmt.Printf("error: %v", err)
	//	return
	//}
	//fmt.Println(v.Programmes)
}

func modifySlice(s []byte) {
	s[0] = 'e'
}

//func main() {
//	param := os.Args
//	if len(param) < 2 {
//		fmt.Println("Invalid parameter")
//		return
//	}
//	var chartFile string = param[1]
//	var saveFile string = param[2]
//	chartFile = strings.TrimSpace(chartFile)
//	saveFile = strings.TrimSpace(saveFile)
//	if chartFile == "" || saveFile == "" {
//		fmt.Println("chart file or save file is empty")
//		os.Exit(0)
//	}
//	_, e := os.Stat(chartFile)
//	if e != nil {
//		fmt.Println("FILE：", chartFile, "does not exist")
//		os.Exit(0)
//	}
//	chat, err := os.Open(chartFile)
//	if err != nil {
//		fmt.Println(err.Error())
//		os.Exit(0)
//	}
//	//var saveFile = "./chat.txt"
//	var sf *os.File
//	sf, err = os.OpenFile(saveFile, os.O_APPEND|os.O_RDWR, 0766)
//
//	if sf == nil {
//		fmt.Println("open file err", err)
//		os.Exit(0)
//	}
//	//建立缓冲区，把文件内容放到缓冲区中
//	buf := bufio.NewReader(chat)
//	reg := regexp.MustCompile("(?U)\"source\":\\s*\"(.*)\"")
//	for {
//		//遇到\n结束读取
//		b, errR := buf.ReadBytes('\n')
//		if errR != nil {
//			if errR == io.EOF {
//				break
//			}
//			fmt.Println(errR.Error())
//		}
//		s := reg.FindAllString(string(b), 1)
//		if len(s) > 0 {
//			str := s[0]
//			str = strings.Replace(str, "\"source\":", "", 1)
//			str = strings.TrimSpace(str)
//			str = strings.Trim(str, "\"")
//			str = strings.TrimLeft(str, "\"")
//			_, err = sf.WriteString(str + "\n")
//		}
//	}
//
//}
