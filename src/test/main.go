package main

import (
	"fmt"
	"time"
)

func main() {
	// 定义变量方式，变量名 类型，也可以不写类型，有类型推断
	var a string = "hello world"
	fmt.Println(a)
	fmt.Println("hello world")
	var b, c int = 1, 2
	fmt.Println(b, c)
	const d, e, f = 1, false, "str"
	fmt.Println(d, e, f)
	//:= 方式赋值
	sum := 0
	//循环去掉括号 := 赋值
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	time.Sleep(50 * time.Second)
}
