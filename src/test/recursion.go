package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

	var k int = 15
	fmt.Printf("\n%d 的阶乘是 %d\n\n", k, Factorial(uint64(k)))
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}

}

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1

}
