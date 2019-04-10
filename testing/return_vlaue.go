
package main


import "fmt"


func main(){

    v, e := return_value()
    fmt.Println(v,e)

}



func return_value()(v string, err error){
    v = "值"
    err = nil
    return v, err
}
