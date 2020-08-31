package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("请输入 (Ctrl + D 退出)： ")

		if !input.Scan() {
			fmt.Println("退出。")
			break
		}

		text := input.Text()
		if text == "quit" {
			fmt.Println("退出。")
			break
		} else {
			fmt.Println("你的输入: ", text, "byte :", input.Bytes())
		}
	}

	if err := input.Err(); err != nil {
		fmt.Println("退出. err: ", err)
		os.Exit(0)
	}

}
