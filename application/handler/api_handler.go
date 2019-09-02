package handler

import (
	"encoding/json"
	"net/http"

	"github.com/KETAKOM/go-study-app/application/domain/model"
	"github.com/KETAKOM/go-study-app/application/usecase"
	"github.com/julienschmidt/httprouter"
)

func BookIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var books []*model.Book
	var err error

	books, err = usecase.BookUsecase{}.GetAll()
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func TodoIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var todoList []*model.Todo
	var err error

	todoList, err = usecase.TodoUsecase{}.GetList()
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(todoList); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func TodoAdd(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var todo model.Todo
	var result bool
	var err error

	// TODO: validate
	params := r.URL.Query()
	todo.Title = params["title"][0]
	todo.Detail = params["detail"][0]
	todo.Auther = params["auther"][0]

	result, err = usecase.TodoUsecase{}.AddTodo(&todo)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func TodoEdit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// var todo model.Todo
	var result bool
	var err error

	// // TODO: validate
	// params := r.URL.Query()
	// todo.Id, err = strconv.Atoi(params["id"][0])
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	// todo.Title = params["title"][0]
	// todo.Detail = params["detail"][0]
	// todo.Auther = params["auther"][0]

	// result, err = usecase.TodoUsecase{}.EditTodo(&todo)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	result = true
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
