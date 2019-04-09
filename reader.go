package main


import (
    "fmt"
    "io"
    "strings"
    )


func main(){

    r := strings.NewReader("我 --在工要最小值 在要械苛在圾要。。。")
    r_len := r.Len()
    //r_cap := cap(r)

    fmt.Println("string的长度：", r_len)
    //fmt.Println("string的cap： ", r_cap)

    b:= make([]byte,3)
    for {
        n, err := r.Read(b)
        //fmt.Printf("n=%v err=%v b=%v\n",n,err,b)
        fmt.Printf("b[:n]=%q\n",b[:n])
        if err == io.EOF{
            break
        }
    }
    fmt.Println("---------------------")

}

