package main

import (
	"fmt"

	"github.com/KETAKOM/go-study-app/registory"
)

func main() {
	s, _ := registory.InitializeBaz()
	fmt.Println(s)
}
