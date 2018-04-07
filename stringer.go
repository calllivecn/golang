package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {

	a := Person{"zhangxu", 24}
	b := Person{"张旭", 26}
	fmt.Println(a, b)
	//-----------------------
	fmt.Println("----------------")

	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, a := range addrs {

		fmt.Printf("%v:%v\n", n, a)
		fmt.Println(n, a)
	}
}
