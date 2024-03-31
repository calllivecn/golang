/*
# date 2019-08-08 14:22:50
# author calllivecn <calllivecn@outlook.com>
*/



package main


import (
    "io"
    "os"
    "fmt"
    "bufio"
    //"bytes"
)

import "./zxlib"

import "flag"

var filename *string = flag.String("filename", "-", "请输入一个文件名")

func init(){
    flag.Parse()
}

func main(){

    fp, err := os.Open(*filename)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer fp.Close()
    defer fmt.Println("读取完成")

    buf := make([]byte, 4096)
    buf_fp := bufio.NewReader(fp)

    for {
        _, err := buf_fp.Read(buf)
        if err == io.EOF {
            break
        }

        fmt.Printf("%s", string(buf))
        zxlib.Clearbuf(buf)
    }
}
