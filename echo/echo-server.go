/*
# date 2019-08-07 17:24:17
# author calllivecn <calllivecn@outlook.com>
*/

package main


import (
    "fmt"
    "net"
    //"bufio"
    "os"
    //"io"
    //"strings"
    // "bytes"
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
            fmt.Printf("Fail to Accept() %s\n", err)
            os.Exit(1)
        }

        fmt.Printf("recv log: %s --> %s\n", conn.RemoteAddr(), conn.LocalAddr())
        go Echo(conn)
    }
}


func Echo(conn net.Conn) {

    defer conn.Close()
    defer fmt.Println("defer conn.Close() 执行了。")

    for{
        buf := make([]byte, 256)

        data_len, err := conn.Read(buf)
        if err != nil {
            fmt.Printf("disconnect ... data_len: %d\terr: %s\n", data_len, err)
            break
        }

        fmt.Printf("Len: %d, %s\n", data_len, buf[:data_len])
        conn.Write(buf[:data_len])

    }

}
