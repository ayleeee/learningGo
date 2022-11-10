package main

/*

구조체 = 필드들의 집합체
인터페이스 = 메서드들의 집합체

인터페이스는 타입이 구현해야 하는 메서드 원형을 정의
*/

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	r := Rect{10., 20.}
	showArea(r)

}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}
