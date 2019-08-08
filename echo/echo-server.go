/*
# date 2019-08-07 17:24:17
# author calllivecn <c-all@qq.com>
*/

package main


import (
    "fmt"
    "net"
    //"bufio"
    "os"
    //"io"
    //"strings"
    "bytes"
)


func main(){

    accept, err := net.Listen("tcp", "0.0.0.0:6789")
    if err != nil {
        fmt.Println("start verser Failed, %s\n", err)
        os.Exit(1)
    }

    fmt.Println("server start done")

    for{
        conn, err := accept.Accept()

        if err != nil {
            fmt.Printf("Fail to connect %s\n", err)
        }

        fmt.Printf("recv log: %s --> %s\n", conn.RemoteAddr(), conn.LocalAddr())
        go Echo(conn)
    }
}



func Echo(conn net.Conn) {

    defer conn.Close()
    defer fmt.Println("defer conn.Close() 执行了。")

    //empty := byte(0)
    //got := []byte("Got: ")
    for{
        buf := make([]byte, 128)

        data_len, err := conn.Read(buf)
        if err != nil || data_len == 0 {
            fmt.Printf("disconnect ... data_len: %d\terr: %s\n", data_len, err)
            //conn.Close()
            break
        }

        tmp := mergeByteSlice([]byte("Got: "), buf)
        fmt.Printf("Len: %d, %s", data_len, tmp)
        conn.Write(tmp)

    }

}

func mergeByteSlice(s ... []byte) []byte{
    return bytes.Join(s, []byte(""))
}

