package main

import "fmt"
import "unsafe"

func main(){

const h ,w = 10 ,5

var area int

area = h * w

fmt.Println("面积：",area)

fmt.Println("h : w: area: ",h,w,area)


const (
	c1 = iota
	c2
	c3
	c4
	)

fmt.Println(c1,c2,c3,c4,"sizeof(c1):sizeof(c2)",unsafe.Sizeof(c1),unsafe.Sizeof(c2))


var a = 4
b := 9.9
var ptr * int

ptr = &a

fmt.Println("var a: b: ptr *:",a,b,ptr)
}

