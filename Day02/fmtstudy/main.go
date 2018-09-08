package main

import "fmt"

func isPrime(a int) bool {
	for i := 2; i < a;  i++{
		if i % a == 0 {
			return false
		}
	}
	return true
}

func main() {
	var a int
	var b int
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	for i := a; i <= b;  i++{
		if isPrime(i) {
			fmt.Println("i = ",i)
		}
	}

}
