package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
	var numbers1 []int

	printSlice(numbers1)

	if numbers1 == nil {
		fmt.Printf("切片是空的")
	}
	slices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(slices)
	fmt.Println("slices[:3]", slices[:4])
	fmt.Println("slices[4:]", slices[4:])
	slices = append(slices, 9, 10, 11)
	printSlice(slices)
	slices = append(slices[:1], slices[2:]...)
	printSlice(slices)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
