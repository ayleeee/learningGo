package main

import (
	"fmt"
	"time"
)

func main() {
	//go는 main이 작동하는 동안에만 살아있다.
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dodo", "bibi", "nako"}
	for _, person := range people {
		go isSexy(person, c)
	} // receiving a message = block opeartaion
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- person + "is sexy"
}
