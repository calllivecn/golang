package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println(os, ".")
	}
	fmt.Println("time.Saturday", time.Saturday)
	fmt.Println("----------------------")

	t := time.Now()
	switch tt := t.Hour(); {
	case tt < 12:
		fmt.Println("Good morning!",tt)
	case tt < 17:
		fmt.Println("Good afternoon.",tt)
	default:
		fmt.Println("Good evening.",tt)
	}
}
