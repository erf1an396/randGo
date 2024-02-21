package main

import (
	"fmt"
	_ "run/setup2"

	_ "run/setup"
)

func main() {
	fmt.Println("Bye")
	output := call(func(s string, i int) bool {
		fmt.Println("i", i)
		return s == "hello"

	})
	fmt.Println("out", output)
}

type eri int

type sFunc func(string, int) bool

func pr(s string, _ int) bool {

	return s == "hello"
}

func call(f sFunc) string {
	result := f("hello", 4)

	if result == true {
		return "yse"
	}

	return "no"
}
