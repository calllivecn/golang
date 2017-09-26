package main

import (
	"fmt"
	"unsafe"
	)

func max(a,b int)int{
	if a > b {
		return a
	}else{
		return b
	}
}

func main(){
	var result int

	var var_test float32 = 1.0

	result = max(5,7)
	fmt.Println("max 值: ",result)

	a, b := 5,7
	swap(&a,&b)
}



func swap(a,b * int){
	var c int
	c = 100
	fmt.Println("int c :",c)
	*a ,*b = *b ,*a
	fmt.Println("&a &b unsafe.Sizeof(a) unsafe.Sizeof(b)",a,b,unsafe.Sizeof(a),unsafe.Sizeof(b))
}
