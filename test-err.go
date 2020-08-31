package main

import (
	"fmt"
	"errors"
)


func main(){
	e := errors.New("这是一个err")
	if e != nil {
		fmt.Println("Err: ", e)
	}else{
		fmt.Println("没有Err")
	}
}