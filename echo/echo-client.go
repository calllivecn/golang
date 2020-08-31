/*
# date 2019-08-08 09:49:15
# author calllivecn <c-all@qq.com>
*/

package main

import (
	"fmt"
	"net"
	"os"
	// "bytes"
	"bufio"
	//"strings"
	"flag"
)

//import "./zxlib"

var addr_port *string = flag.String("address", "127.0.0.1:6789", "IP:PORT")

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

	input := bufio.NewScanner(os.Stdin)

	// rbuf := make([]byte, 256)

	end:
	for {

		fmt.Printf("请输入：")

		if !input.Scan() {
			fmt.Println("退出。")
			break
		}

		text := input.Text()
		switch text {
		case "quit":
			fmt.Println("Bye")
			break end
		case "":
			continue
		default:
			// fmt.Println(text)
			fmt.Println("你的输入:", text, "byte:", input.Bytes())
		}

        conn.Write(input.Bytes())

		// _ = reset(rbuf)
		rbuf := make([]byte, 256)

		rn, _ := conn.Read(rbuf)

        fmt.Println("echo from server: ", string(rbuf[:rn]))
	}
}


func reset(b []byte) error {
	l := len(b)
	var null byte = byte('0')
	for i := 0; i < l; i++ {
		b[i] = null
	}

	return nil
}