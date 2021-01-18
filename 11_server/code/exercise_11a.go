package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

// Todo is a todo with a title and content
type Todo struct {
	Title string
	Content string
}

//PageVariables are variables sent to html template
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Get Todos",
		PageTodos:  todos,
	}

	t, err := template.ParseFiles("todos.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, pageVariables)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Request parsing error: ", err)
	}

	todo := Todo{
		Title: r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print(todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)
	// even wrapped in log fatal, the code will still execute
	// will fallback to log fatal if there is an error
	log.Fatal(http.ListenAndServe(":8080", nil))
}
