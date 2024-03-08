package main

import (
	"encoding/json"
	"fmt"
)

func main() {

}

func String(err error) string {

	return fmt.Sprintln(err.Error())

}

func StringTwo(err error) string {

	return fmt.Sprintln(err)

}

type simpleData struct {
	ID    uint
	Name  string
	Email string
}

func (s simpleData) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"id": %d, "name": %s, "email": %s}`, s.ID, s.Name, s.Email)), nil
}

type simpleDataTwo struct {
	ID    uint
	Name  string
	Email string
}

func Json(data simpleData) ([]byte, error) {
	return json.Marshal(data)
}

func JsonTwo(data simpleDataTwo) ([]byte, error) {
	return json.Marshal(data)
}
