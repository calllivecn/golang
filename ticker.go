/*
# date 2020-01-06 10:15:11
# author calllivecn <calllivecn@outlook.com>
*/

package main

import (
    "fmt"
    "time"
)

func main(){

    ticker := time.NewTicker(time.Millisecond * 500)

    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at ", t)
        }
    }()

    time.Sleep(time.Second * 5)
    ticker.Stop()
    fmt.Println("ticker stop")
}
