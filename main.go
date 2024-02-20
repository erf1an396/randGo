package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var scores = []int{10, 25, 32, 54, 95}

	for i := 0; i < len(scores); i++ {
		fmt.Printf("i : %d , value : %d\n", i, scores[i])
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	expath := filepath.Dir(ex)
	fmt.Println("executable binary path", expath)

	time.Sleep(100 * time.Second)

}
