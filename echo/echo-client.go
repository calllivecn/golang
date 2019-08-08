/*
# date 2019-08-08 09:49:15
# author calllivecn <c-all@qq.com>
*/

package main

import (
	"fmt"
	"net"
	"os"
	//"strings"
	"flag"
)

//import "./zxlib"

var addr_port *string = flag.String("address", "127.0.0.1:2222", "IP:PORT")

func init() {
	flag.Parse()
}

func main() {

	conn, err := net.Dial("tcp", *addr_port)
	if err != nil {
        fmt.Println("connect fail:", addr_port, "Error: ", err)
		os.Exit(1)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 256)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		if string(buf) == "quit" {
			fmt.Println("Bye")
			break
		}
		//fmt.Fprint(os.Stdin, buf)
		//fmt.Println(buf)
        conn.Write(buf[:n])

        r_buf := make([]byte, 256)
        r_n, _ := conn.Read(r_buf)
        fmt.Println(string(buf[:r_n]))
	}
}
