package main

import (
	"log"
	"net/http"

	"github.com/KETAKOM/go-study-app/application/handler"
	"github.com/julienschmidt/httprouter"
)

func main() {
	log.Printf("Server Start >> http://localhost:3000")

	router := httprouter.New()
	router.GET("/", handler.BookIndex)
	router.GET("/todo", handler.TodoIndex)
	router.GET("/todo/add", handler.TodoAdd)
	log.Fatal(http.ListenAndServe(":3000", router))
}
