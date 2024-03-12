package main

import "fmt"

func main() {

	var users = map[string]int{}

	users["Hossein Nazari"] = 2

	fmt.Println(users["Hossein Nazari"])

	var u = map[int]int{}

	value, ok := u[3]
	if !ok {
		fmt.Println("the key doesn't exist")
	} else {
		fmt.Println("the value", value)
	}

	type user struct {
		ID    int
		Name  string
		Email string
	}

	var userL = make(map[user]int)

	userL[user{
		ID:    1,
		Name:  "erfan",
		Email: "eric@",
	}] = 100

}
