package main

import (
	"fmt"
	"time"
)

func main() {

	// Variables declared without a corresponding initialization are zero-valued
	var e int
	fmt.Println(e)
	// so e becomes 0

	f := "apple"
	fmt.Println(f)
	// := is shorthand for declaring and initializing a variable

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	//  빈 인터이스 (empty Interface) : interface{}
	//  어떠한 타입도 담을 수 있는 컨테이너라고 볼 수 있다.

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	// 하나의 변수명을 가지고 bool값, int 값, string 값을 모두 사용할 수 있다.
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	/*
		Array (배열) : 고정된 배열크기 안에 동일한 타입의 데이터를 연속적으로 저장. 배열 크기 동적 변경 X, 부분배열 발췌 불가
		Slice (슬라이스) : 내부적으로 배열에 기초하였으나 크기를 동적으로 변경할 수 있고, 부분 배열 발췌 가능
		Maps : 파이썬의 딕셔너리 형태와 비슷함
	*/

	//Array
	arr := [5]int{1, 2, 3, 4, 5}
	var a [5]int
	var twoD [2][4]int

	//Slice
	s := make([]string, 3)
	l := s[2:]

	// Maps
	m := make(map[string]int)
	n := map[string]int{"foo": 1, "bar": 2}

}
