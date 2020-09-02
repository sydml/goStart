package main

import "fmt"

/* 定义结构体 */
type Circle struct {
	radius float64
}

/* 方法的定义方式
func (variable_name variable_data_type) function_name() [return_type]{
	实现
}*/
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	var c Circle
	c.radius = 10.00
	fmt.Println("圆面积 = ", c.getArea())
}
