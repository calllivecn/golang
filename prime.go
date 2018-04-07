package main

import (
	"fmt"
	"math"
	)

var NUMBER uint64 = 500000

func main(){
	var i uint64
	for i=1;i <= NUMBER;i++{
        prime(i)
	/*if prime(i) {
		fmt.Println(i,"是素数")
		}
    */
		}
fmt.Println("run done Go lang")
}



func prime(n uint64)bool{
	var i uint64
	if n <= 2 {
		return true
		}

	for i=3;i < uint64(math.Sqrt(float64(n)));i+=1 {
		if 0 == n % i {
			return false
			}
		}

	return true
}
