package logg

import (
	"encoding/json"
	"fmt"
	"os"
	"run/richerror"
)

type Log struct {
	Errros []richerror.RichError
}

func (l Log) Append(r richerror.RichError) {
	l.Errros = append(l.Errros, r)
}

func (l Log) Print() {
	for i, e := range l.Errros {
		fmt.Println("i", i, "error", e)
	}
}

func (l Log) Save() {
	f, _ := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	defer func() {
		cErr := f.Close()
		if cErr != nil {
			fmt.Println("can't close the file", cErr)
		}
	}()

	data, _ := json.Marshal(l.Errros)

	_, wErr := f.Write(data)
	if wErr != nil {
		fmt.Println("can't write to file", wErr)
	}
}
