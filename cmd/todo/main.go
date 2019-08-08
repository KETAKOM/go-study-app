package main

import (
	"fmt"

	hand "github.com/KETAKOM/go-study-app/application/handler"
)

func main() {
	todo := hand.NewTodo()
	fmt.Println(todo.Create("Golang勉強", "環境構築").Show())
}
