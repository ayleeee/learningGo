package main

import (
	"fmt"

	"aylee.test/practice2/mydict"
)

func main() {
	// dictionary := mydict.Dictionary{"first": "first word"}
	// definition, err := dictionary.Search("first")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	// definition := "Greeting"
	// err := dictionary.Add(word, definition)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// hello, _ := dictionary.Search(word)
	// fmt.Println(hello)

	// err2 := dictionary.Add(word, definition)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	dictionary.Add(baseword, "First")
	// err := dictionary.Update(baseword, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseword)
	// fmt.Println(word)
	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)

}
