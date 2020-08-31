package main


import (
	"os"
	"fmt"
)


func main(){
	stdout := os.Stdout
	stdout.Write([]byte("这是一个标准输出。"))
	fmt.Println("执行成功。")
}