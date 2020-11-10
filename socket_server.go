package main

import (
	"fmt"
	"net"
)

var connList = make(map[byte]net.Conn)

func handleConn(c net.Conn) {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(err)
		}
		c.Close()
	}()
	//var num int
	for {
		var reByte []byte = make([]byte, 1024)
		_, err = c.Read(reByte)
		if err != nil {
			return
		}
		connList[reByte[0]] = c
		//fmt.Println(reByte[0])
		for _, value := range connList {
			_, err := value.Write(reByte[1:])
			if err != nil {
				return
			}
		}
		//fmt.Println("receive ", num, " byte:", string(reByte[1:]))

	}
}

func main() {
	var err error
	defer func() {
		if err == recover() {
			fmt.Println(err)
		}
	}()
	var l net.Listener
	l, err = net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	for {
		var c net.Conn
		c, err = l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
