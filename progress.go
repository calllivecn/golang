package main


import (
	"fmt"
	"time"
)

func progress(){
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main(){
	go progress()
	time.Sleep(10 * time.Second)
}