package main

import (
	"fmt"
	"net"
	"strings"
)

func establishConn() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", ":8888")
	return
}

//func main() {
//	conn, err := establishConn()
//	if err != nil {
//		panic("fail to connent")
//	}
//	go read(conn)
//	fmt.Println("place input something:")
//	reader := bufio.NewReader(os.Stdin)
//	var str string
//	for {
//
//		tmp, err := reader.ReadString('\n')
//		if err != nil {
//			fmt.Println("with err file", err)
//		}
//		str = strings.TrimSpace(tmp)
//		write(conn, str)
//	}
//
//}

func read(conn net.Conn) {
	for {
		var re = make([]byte, 1024)
		var err error
		_, err = conn.Read(re)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("this is return string:", strings.TrimSpace(string(re)))
	}

}

func write(conn net.Conn, str string) {
	//var num int
	var err error
	_, err = conn.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
	}
}
