package main

import (
	"fmt"
	"reflect"
	"run/richerror"
	"time"
)

func main() {
	rErr := richerror.RichError{
		Message: "id is not valid",
		MetaData: map[string]string{
			"id": "0",
		},
		Operation: "main",
		Time:      time.Now(),
	}

	value := reflect.ValueOf(rErr)
	switch value.Kind() {
	case reflect.Struct:
		fmt.Println("number of field", value.NumField())
		for i := 0; i < value.NumField(); i++ {

			fmt.Printf("field index: %d, type: %s , value: %s/n", i,
				value.Type().Field(i).Type,
				value.Type().Field(i).Name,
			)
		}
	}
}
