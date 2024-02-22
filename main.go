package main

import "fmt"

func main() {

	i := 0

start:
	fmt.Println("i", i)
	i++
	if i < 10 {
		goto start
	}
	fmt.Println("Finish")
}
