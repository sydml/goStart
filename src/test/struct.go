package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	fmt.Println(Book{"go language", "菜鸟", "subject", 123})
	fmt.Println(Book{title: "Go 语言", author: "www.runoob.com"})
}
