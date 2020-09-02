package main

import "fmt"

func main() {
	var mi MI
	result := mi.message("send msg")
	fmt.Println(result)
}

type Phone interface {
	call()
	message() string
}

type MI struct{}

func (mi MI) message(msg string) string {
	fmt.Println("I am mi, send message to you")
	return "I am mi message back"
}
