/*
# date 2020-01-08 14:32:08
# author calllivecn <c-all@qq.com>
*/

package learn

import "fmt"

var PI float32 = 3.14159

var AUTHOR string = "ZhangXu"

var DATE string = "2020-01-01"



func PrintPI() {
    fmt.Println(PI)
}


func Pkginfo() {
    fmt.Println(AUTHOR, " -- ", DATE)
}
