package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var num int
	// 빠른 처리
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fscanln(reader, &num)

	for i := 1; i <= num; i++ {
		fmt.Fprintln(writer, i)
	}
}
