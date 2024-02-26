package new_in_memory

import "run/userr"

type Store struct {
	users map[uint]userr.User
}

func (s *Store) GetUserByID(id uint) userr.User {
	return s.users[id]

}

func (s *Store) CreateUser(u userr.User) {

}
func (s *Store) ListUsers() []userr.User {

	return []userr.User{}
}
