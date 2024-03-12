package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	s := Student{
		ID:       3,
		Name:     "",
		Email:    "",
		IsActive: false,
		JoinDate: time.Time{},
	}

	var t1 = Teacher{
		ID:       2,
		IsActive: true,
	}

	DeactivateUser(&t1, nil)

	DeactivateUSer2(&t1)

	err := DeactivateUSer2(&s)
	if err != nil {
		fmt.Println("can't implant the deactivate")
	}

}

type Teacher struct {
	ID       uint
	Name     string
	Email    string
	IsActive bool
	Grade    string
	Salary   uint
}

func (t *Teacher) Deactivate() error {
	if !t.IsActive {
		return errors.New("the Teacher is deactivated already")
	}
	if t.IsActive {
		t.IsActive = false
	}
	return nil
}

type Student struct {
	ID       uint
	Name     string
	Email    string
	IsActive bool
	JoinDate time.Time
}

func (s *Student) Deactivate() error {
	if !s.IsActive {
		return errors.New("the Teacher is deactivated already")
	}
	if s.IsActive {
		s.IsActive = false
	}
	return nil
}

func DeactivateUser(t *Teacher, s *Student) error {
	if t != nil {
		err := t.Deactivate()
		if err != nil {
			return fmt.Errorf("can't deactivate user : %w", err)
		}
	}
	return nil
}

type User interface {
	Deactivate() error
}

func DeactivateUSer2(u User) error {
	return u.Deactivate()
}
