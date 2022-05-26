package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &x)
	defer writer.Flush()

	var arr = make([]int, x)
	num := -1000000
	num2 := 1000000
	for i := 0; i < x; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
		if arr[i] > num {
			num = arr[i]
		}
		if arr[i] < num2 {
			num2 = arr[i]
		}
	}

	fmt.Fprintf(writer, "%d %d", num2, num)
}
