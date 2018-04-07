package main

import (
	"flag"
	"fmt"
	"os"
)

var file, hash string

func main() {
	fmt.Printf("%#v %#v\n", file, hash)
}

func init() {
	flag.StringVar(&file,"f", "", "a file")
	flag.StringVar(&hash,"hash", "md5", "hash al.")
	flag.Parse()

	if fi, err := os.Stat(file); err == nil {
		if fi.IsDir() {
			fmt.Println(file, "不是一个文件")
			os.Exit(1)
		}
	}
	switch {
	case hash == "sha512":
	case hash == "md5":
	default:
		fmt.Println("hash 必须是md5 和sha512中的一种")
		os.Exit(1)
	}
	fmt.Println(file, hash)
}
