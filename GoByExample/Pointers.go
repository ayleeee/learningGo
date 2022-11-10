package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	// return 1
	i := 1
	fmt.Println("initial:", i)

	// return i
	// zeroval is pass by value
	// zeroval will get a copy of ival distinct from the one in the calling function
	zeroval(i)
	fmt.Println("zeroval:", i)

	// return 0
	// zeroptr is pass by reference
	// Assigning a value to a dereferenced pointer changes the value at the referenced address.
	// change the value because it has a reference to the memroy address for that variable
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
