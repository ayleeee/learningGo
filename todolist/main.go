package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/unrolled/render"
)

type ToDoItem struct {
	id    int    `json: "todoID"`
	item  string `json :"todoITEM"`
	check bool   `json: "todoCHECK"`
}

var todos = []ToDoItem{}
var idCounter = 0
var rd *render.Render

func addHandler(w http.ResponseWriter, r *http.Request) {
	todo := r.FormValue("todo")
	idCounter += 1
	newtodo := ToDoItem{id: idCounter, item: string(todo), check: false}
	todos = append(todos, newtodo)
	rd.HTML(w, http.StatusOK, "index", string(newtodo.item))
}

func main() {
	rd = render.New(render.Options{
		Directory:  "static",
		Extensions: []string{".html", ".tmpl"},
	})
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/add", addHandler)
	fmt.Println("Starting server at port: 8080 \n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
