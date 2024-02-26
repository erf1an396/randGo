package main

import (
	"run/appp"
	"run/new_in_memory"
	"run/userr"
)

func main() {
	//var a = &storagee.Memory{}
	var newMemoryStorage = &new_in_memory.Store{}
	appplication := appp.App{
		Name:        "sample-apppp",
		UserStorage: newMemoryStorage,
	}

	appplication.CreateUser(userr.User{
		ID:   1,
		Name: "Erfan",
	})
}
