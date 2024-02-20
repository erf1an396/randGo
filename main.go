package main

import (
	"fmt"
	_ "run/setup2"

	_ "run/setup"
)

func init() {
	fmt.Println("step 0 MAIN")
}

func main() {
	fmt.Println("Bye")

}
