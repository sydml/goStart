package main

import "fmt"

type cb func(int) int

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i

	}

}

func main() {
	var b, c = 1, 3

	var ret = max(b, c)
	fmt.Println(ret)
	var testA, testB = "hello", "world"
	//类型推断
	result1, result2 := swap(testA, testB)
	fmt.Println(result1, result2)

	// 函数作为参数传入，先定义函数，回调函数
	testCallBack(1, callBack)
	// 回调函数，直接在入参中实现，类似于java的匿名内部类
	testCallBack(2, func(x int) int {
		fmt.Printf("我是回调，x：%d\n", x)
		return x
	})

	nextNumber := getSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

// 入参 多个 类型，出参多个用括号，类型定义
func swap(x, y string) (string, string) {
	return y, x
}

func testCallBack(x int, f cb) {
	f(x)
}

func callBack(x int) int {
	fmt.Printf("我是回调，x：%d\n", x)
	return x
}

/**
函数类型写法 入参 类型，出参类型
*/
func max(num1, num2 int) int {
	var result int
	//可以简写括号
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result

}
