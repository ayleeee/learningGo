package main

import (
	"fmt"
	"net/http"

	"github.com/ayleeee/go-course/pkg/handlers"
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

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port : %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
