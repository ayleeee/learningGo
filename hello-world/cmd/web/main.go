package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ayleeee/go-course/pkg/config"
	"github.com/ayleeee/go-course/pkg/handlers"
	"github.com/ayleeee/go-course/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	// path name
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "hello world!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	//Sprintf : 형식을 지정해서 문자열로 만든다
	// 	fmt.Println(fmt.Sprintf("Number of bytes written : %d", n))

	// })

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port : %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
