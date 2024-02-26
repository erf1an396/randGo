package appp

import (
	"fmt"
	"run/userr"
)

type App struct {
	Name string
	//StorageFilePath string
	//InMemoryStorage storagee.Memory

	UserStorage UserStore
}

type UserStore interface {
	CreateUser(u userr.User)
	ListUsers() []userr.User
	GetUserByID(id uint) userr.User
}

func (a App) CreateUser(u userr.User) {
	if u.Name == "" {
		fmt.Println("name can't be empty")

		return
	}

	//var fileHandler *os.File
	//if f, err := os.OpenFile(a.StorageFilePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777); err != nil {
	//	fmt.Println("can't open file", err)
	//
	//	return
	//} else {
	//	fileHandler = f
	//}
	//
	//defer fileHandler.Close()
	//
	//data, mErr := json.Marshal(u)
	//if mErr != nil {
	//	fmt.Println("can't marshal user data", mErr)
	//}
	//
	//  ******************  another solution  ****************a.InMemoryStorage.AddUser(u)
	a.UserStorage.CreateUser(u)
	//if _, wErr := fileHandler.Write(data); wErr != nil {
	//	fmt.Println("can't write to the file", wErr)
	//
	//	return
	//}

}

//func (a App) ListUsers() userr.User {
//
//}
//
//func (a App) GetUserByID(id uint) userr.User {
//
//}
