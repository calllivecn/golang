/*
# date 2020-09-05 02:15:22
# author calllivecn <c-all@qq.com>
*/


package main

import (
    "fmt"

)


func main(){
    buf := make([]byte, 16)

    headByte := []byte(`head`)

    playload := []byte(` end`)


    // buf[:4] = headByte
    // buf[4:8] = playload

    copy(buf[:4], headByte)
    copy(buf[4:8], playload)

    fmt.Printf("buf: `%s`\n", buf)
}
