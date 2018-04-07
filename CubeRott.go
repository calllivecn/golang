package main

import "fmt"

func main() {

	fmt.Printf("%T,%v\n", nil, nil)

	//fmt.Println(CubeRoot(9.0))
}

func CubeRoot(x float64) float64 {

	z := x / 3
	for i := 0; i < 1e6; i++ {
		//prevz := z
		z -= (z*z*z - x) / (3 * z * z)
		if z < 0.00000001 {
			return z
		}
	}
	panic(fmt.Sprintf("CubeRoot(%g) did not caonverge", x))
}
