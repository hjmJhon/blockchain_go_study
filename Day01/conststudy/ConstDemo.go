package conststudy

import "fmt"


//常量定义的几种方式
const Name = 100
const (
	A = 10
	B = 10
)

func init() {
	fmt.Println("init method is running ,and the const value is ", Name)
}

func Add(a int, b int) int {
	fmt.Println("a + b method is running")
	return a + b
}
