
package main

/*
 
#cgo CFLAGS: -I./
 
#cgo LDFLAGS: -L./ -lvideo
 
#include "video.h"
 
*/
import "C"

import "fmt"

func main() {
    cmd := C.CString("ffmpeg -i ./xxx/*.png ./xxx/yyy.mp4")
    C.exeFFmpegCmd(&cmd)
    fmt.Println("done")
}



/*
先回答为什么说这种是最安全的和最不爽的？原因如下：
1. 动态库破解十分困难，如果你的 go 代码泄露，核心动态库没那么容易被攻破
2. 动态库会在被使用的时候被加载，影响速度
操作难度比方式一麻烦不少
结论:
CFLAGS: -I路径 这句话指明头文件所在路径，-Iinclude 指明 当前项目根目录的 include 文件夹
LDFLAGS: -L路径 -l名字 指明动态库的所在路径，-Llib -llibvideo，指明在 lib 下面以及它的名字 video
如果动态库不存在，将会爆找不到定义之类的错误信息
*/
