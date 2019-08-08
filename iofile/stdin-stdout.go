/*
# date 2019-08-08 10:00:47
# author calllivecn <c-all@qq.com>
加上点中文注
*/


package main

import (
    "fmt"
    "bufio"
    //"strings"
    "os"
    //"io"
)

func main() {

    //fmt.Printf("os.Stdin(): %v\t%t\t%x\n", os.Stdin, os.Stdin, *os.Stdin)
    stdin := bufio.NewReader(os.Stdin)
    content := make([]byte, 8)
    for {

        //str, ok := stdin.ReadString(' ')
        n, ok := stdin.Read(content)
        if ok != nil {
            fmt.Println("stdin.Read() error")
            os.Exit(1)
        }

        fmt.Printf("%d: %s\n", n, string(content))

        clearbuf(content)
    }
}


func clearbuf(buf []byte) (int, error) {
    //block := make([]byte, 1024)

    empty := byte(0)

    l := len(buf)

    var i int = 0

    for ;i<l; i++ {

        if buf[i] == empty {
            break
        }

        buf[i] = empty
    }

    return i, nil

}
