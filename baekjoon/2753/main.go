package main

import "fmt"

func main() {
	var yr int
	fmt.Scan(&yr)

	if yr%4 == 0 && yr%100 != 0 || yr%400 == 0 {
		fmt.Print(1)
	} else {
		fmt.Print(0)
	}
}
