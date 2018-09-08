package main

import (
	"fmt"
	"Day01/conststudy"
)

func list(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}

func main() {
	list(5)
	sum := conststudy.Add(1, 9)
	fmt.Println("sum = ", sum)

	name := conststudy.Name
	fmt.Println("cont in conststudy is ", name)
}
