package main

import (
	"fmt"
	"time"
)

func say(s string) {
	s_len := len(s)
	for i := 0; i < s_len; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}

}

func main() {
	go say("world")
	say("hello")
}
