package main

import (
	"fmt"
	"unsafe"
)

import "stringutil"

func main() {

	var hello string = "hello world."
	var cn string = "中文呢？"
	fmt.Println(hello, len(hello))
	fmt.Println(cn, len(cn))

	print_char(hello)
	print_char(cn)
	fmt.Println(stringutil.Reverse("这是一个返转的字符串。"))
}

func print_char(s string) {
	for _, c := range s {
		fmt.Println(string(c), "rune size:", unsafe.Sizeof(c))
	}
}
