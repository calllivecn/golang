package main



import (
	"fmt"
)


func main(){
	const i = 500
	j := 40

	for m := 0; j > 0; j-- {
		fmt.Printf("m:%d j:%d i:%d\n", m, j, i)
	}

	m := 1
	n := 2
	if m == 0 {
		n = 1
	}else if m != 0 {
		n = 3
	}
}