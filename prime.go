package main

import (
	"fmt"
    "flag"
	"math"
)

var NUMBER uint64 = 99
var P bool = false

func init(){
    flag.Uint64Var(&NUMBER, "number", 99, "求N以内的质数。")
    flag.BoolVar(&P, "print", false, "打印质数(default: false)")
}

func main() {

    if P {
        no_print(NUMBER)
    }else{
        print_(NUMBER)
    }

	fmt.Println("run done Go lang")
}

func prime(n uint64) bool {
	var i uint64

    if n == 1 {
        return false
    }

	for i = 2; i <= uint64(math.Sqrt(float64(n))); i++ {
		if 0 == n%i {
			return false
		}
	}

	return true
}

func no_print( n uint64) {

	var i uint64

	for i = 1; i <= n; i++ {
		prime(i)
	}
}

func print_( n uint64) {

	var i uint64

	for i = 1; i <= n; i++ {
		if prime(i) {
			fmt.Println(i, "是素数")
		}
	}
}
