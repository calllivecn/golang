package main

import "fmt"

func sum(a []int, c chan int) {

	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	a := []int{12, 23, 3, 5, 4, 6, 45, 7, 567, 456, 7456, 4, 98, 3, 344, 5, 345}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	fmt.Println("-----------------------")

	cc := make(chan int, 2)
	cc <- 1
	cc <- 2
	fmt.Println(<-cc)
	fmt.Println(<-cc)

	fmt.Println("-----------------------")

	ccc := make(chan int, 10)
	go fibonacci(cap(ccc), ccc)
	for i := range ccc {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	defer close(c)

	a, b := 0, 1

	for i := 0; i < n; i++ {
		c <- b
		a, b = b, a+b
	}
}
