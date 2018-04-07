package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { //这样相当于引用省内存？？？---> yes
	v.X++
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	//v := &Vertex{3, 4}  //???
	v := Vertex{3, 4}
	fmt.Println("向量v的值:", v.Abs(), "向量v的原始值:", v)
	v.Scale(2)
	fmt.Println("v.Scale(2)", v)
	//-----------------------

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs, "-math.Sqrt2 : ", -math.Sqrt2)
}
