package main

/*

함수에 고정된 수의 파라미터를 전달하지 않고 다양한 숫자의 파라미터를 전달하고자 할 때
... (3개의 마침표)를 사용한다.

가변 파라미터를 받아들이는 함수 = Variadic Function

Closure : 함수 바깥에 있는 변수를 참조하는 함수 값

익명 함수 : 함수명을 가지지 않는 함수
일반적으로 함수 전체를 변수에 할당 혹은 다른 함수의 파라미터에 직접 정의되어 사용
일반 함수는 언제든 호출될 수 있기 때문에 항상 메모리를 차지하고 있다.
그러나 익명 함수는 한 번만 사용하는 함수로 만들어졌기 때문에 쓰는 시간 외에 불필요한 시간동안
메모리를 차지하지 않아 메모리를 아낄 수 있다.

*/

func say(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

// Closures
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	say("This", "is", "a", "book")
	say("hello", "world")
	nextInt := nextValue()
	println(nextInt())
	println(nextInt())
	println(nextInt())
	println(nextInt())
	println(nextInt())

	newNext := nextValue()
	println(newNext())
}
