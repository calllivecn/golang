package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type V2 struct {
	X2, Y2 int
}

var (
	v1 = V2{1, 2}
	v2 = V2{X2: 1}
	v3 = V2{}
	p  = &V2{1, 2}
)

func main() {
	v := Vertex{34, 35}
	fmt.Println(v)
	fmt.Println("Vertex.X and Vertex.Y", v.X, v.Y)
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p, v2, v3)
}
