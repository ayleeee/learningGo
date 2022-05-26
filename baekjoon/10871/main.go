package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n, &x)
	defer writer.Flush()

	var arr = make([]int, n)
	for i := range arr {
		fmt.Fscanf(reader, "%d ", &arr[i])
		if arr[i] < x {
			fmt.Fprintf(writer, "%d ", arr[i])
		}
	}
	fmt.Fprint(writer, "\n")
}
