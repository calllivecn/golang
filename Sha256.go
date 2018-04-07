package main

import (
	"fmt"
	"crypto/sha256"
	)

//SHA256 := 32

func main(){

	//fmt.Println()
	Sha256()

}



func Sha256(){
	sha := sha256.New()
    text := "你以为你是谁啊？"
	sha.Write([]byte(text))
	fmt.Printf("%x  %s\n",sha.Sum(nil),text)
}
