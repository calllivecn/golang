
package main
/*
// C 标志io头文件，你也可以使用里面提供的函数
#include <stdio.h> 
void pri(){
    printf("hey");
}
int add(int a,int b){
    return a+b;
}
*/
import "C"  // 切勿换行再写这个
 
import "fmt"
 
func main() {
    fmt.Println(C.add(2, 1))
}


/*
上面的代码，直接拷贝运行就能输出结果：3

结论：
1. 但凡要引用与 c/c++ 相关的内容，写到 go 文件的头部注释里面
2. 嵌套的 c/c++ 代码必须符合其语法，不与 go 一样
3. import "C" 这句话要紧随，注释后，不要换行，否则报错
4. go 代码中调用 c/c++ 的格式是: C.xxx()，例如 C.add(2, 1)
*/
