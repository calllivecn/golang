package main


import (
	"fmt"
	"bytes"
    "encoding/json"
)

type Movie struct {
	Name string
	Description string
	TimeLen int // seconds
}

func NewMovie() *Movie {
	p := &Movie{}
	return p
}

func main(){
	m1 := &Movie{"张旭", "golang json ", 64}
	fmt.Println("Movie struct --> %#v", m1)

	j := &bytes.Buffer{}

	encoder := json.NewEncoder(j)
	err := encoder.Encode(m1)	

	if err != nil {
		fmt.Println("json encoding Error.")
		return
	}
	jb := j.Bytes()
	fmt.Println("struct --> json: ", string(jb))

	fmt.Println("print :", j)

	// json --> struct 
	// struct_ := &Movie{}
	decoder := json.NewDecoder(j)
	err = decoder.Decode(jb)
	if err != nil {
		fmt.Println("json encoding Error: ", err)
		return
	}
	fmt.Printf("json --> struct : %#v\n", decoder)
}