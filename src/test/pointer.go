package main

import "fmt"

/**
指针
*/
func main() {
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
	if ptr == nil {
		fmt.Println("ptr = nil")
	} else {
		fmt.Println("ptr is ", ptr)

	}
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	var ptrA, ptrB = 100, 200
	swapPtr(&ptrA, &ptrB)
	fmt.Printf("交换后 a 的值 : %d\n", ptrA)
	fmt.Printf("交换后 b 的值 : %d\n", ptrB)
}

func swapPtr(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
