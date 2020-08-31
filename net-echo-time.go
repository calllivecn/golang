package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:6789")
	if err != nil {
		fmt.Println("net.Listen() Error")
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept() Error: ", err)
			continue
		}
		fmt.Printf("client: %T, %#v %s\n", conn.RemoteAddr, conn.RemoteAddr, conn.RemoteAddr())
		go conHandler(conn)
	}

}

func conHandler(c net.Conn) {
	defer c.Close()
	for {
		t := time.Now()
		// t.Format() 这完意有毒。。。。只能是这个 指定的格式串， 记忆方法， 6  1 2 3 4 5
		_, err := io.WriteString(c, t.Format("2006-01-02 15:04:05.00000") + "\n")
		if err != nil {
			fmt.Println("client Error: ", c.RemoteAddr(), err)
			break
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("client 退出。", c.RemoteAddr())
}
