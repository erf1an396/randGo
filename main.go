package main

import "fmt"

type User struct {
	ID   uint
	Name string
}

func main() {
	i := new(int)
	fmt.Printf("type of i : %T\n", i)

	u := new(User)
	fmt.Printf("type of u : %T\n", u)

	u.ID = 2
	u.Name = "erfan"

	fmt.Printf("u.id , u.name : %+v, %+v\n", u.ID, u.Name)
	uu := &User{
		ID:   1,
		Name: "erfan",
	}
	fmt.Printf("type of uu :%T\n", uu)

	var uuu *User
	fmt.Printf("type of uuu:%T\n", uuu)

	fmt.Println("u.ID", u.ID)
	fmt.Println("uu.ID", uu.ID)
	//fmt.Println("uuu.ID", uuu.ID)
}
