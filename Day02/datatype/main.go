package main

import "fmt"

func swap(a *int, b *int) {
	temp := a
	a = b
	b = temp
}
func main() {
	a := 5
	b := 10

	fmt.Println("before swap , a = ",a)
	fmt.Println("before swap ,b = ",b)

	swap(&a, &b)

	fmt.Println("after swap , a = ",a)
	fmt.Println("after swap ,b = ",b)

}
