package main

import (
	"html/template"
	"net/http"
)

type TodoItem struct {
	Title string
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/add-todo", addTodoHandler)

	http.ListenAndServe(":3000", router)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	homeTmpl := template.Must(template.ParseFiles("./templates/index.html"))
	homeTmpl.Execute(w, nil)
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	todoName := r.PostFormValue("todo")
	todoItem := template.Must(template.ParseFiles("./templates/todo-item.html"))

	todoItem.Execute(w, TodoItem{
		Title: todoName,
	})
}
