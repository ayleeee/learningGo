package main

import (
	"fmt"

	"aylee.test/practice1/studentInfo"
)

func main() {
	student := studentInfo.NewStudent("aylee")
	student.IdNumber("2018")
	student.Status(3)
	fmt.Println(student)
}
