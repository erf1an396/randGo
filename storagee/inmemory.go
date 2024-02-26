package storagee

import "run/userr"

type Memory struct {
	users []userr.User
}

func (m *Memory) AddUser(u userr.User) {
	m.users = append(m.users, u)
}

func (m *Memory) CreateUser(u userr.User) {

}

//func (m *Memory) ListUsers() []userr.User {
//
//}
//func (m *Memory) GetUserByID(id uint) userr.User {
//	for _, user := range m.users {
//		if user.ID == id {
//			return user
//		}
//		return userr.User{}
//	}
//
//}
