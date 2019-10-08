package test

import (
	"fmt"
)

const default_word = "world"
const helloPrefix = "Hello, "

func Hello(word string) string {
	if word == "" {
		word = "world"
	}
	return helloPrefix + word
}


func main()  {
	fmt.Println(Hello("world"))
}
