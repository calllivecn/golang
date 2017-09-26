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
	sha.Write([]byte("张旭"))
	fmt.Printf("%x\n",sha.Sum(nil))
}
