package main

import (
	"fmt"
)

type addFunc func(int, int) int

func add(a int, b int) (result int) {
	sum := a + b
	return sum
}

func sub(a int, b int) int {
	return a - b
}

//函数作为变量传递
func operator(addFunc addFunc, a int, b int) int {
	return addFunc(a, b)
}

//函数多返回值
func methodMulRetnValue(a int, b int) (int, string) {
	return a + b, "result"
}

//可变参数
func mulArgs(a int, arg ... int) int {
	result := a
	for _, v := range arg {
		result += v
	}
	return result
}

//defer 语句及其特征
func deferDemo(a int) int {
	defer fmt.Printf("defer 语句定义在上面,a=%d", a)
	defer fmt.Println("defer 定义在下面")
	a++
	fmt.Printf("defer   --------- a = %d", a)
	fmt.Println()

	return a
}

//rune
func runeDemo(str string) {
	length := len(str)
	fmt.Printf("length=%d", length)
	fmt.Println()
	runes := []rune(str)
	length2 := len(runes)
	fmt.Printf("length2=%d", length2)
	fmt.Println()
	for k, v := range runes {
		fmt.Printf("k=%d v=%c \t", k, v)
	}
}

//内置函数:panic,recover
func buildInMethod() {
	defer fmt.Println("hahah")
	defer func() {
		if recover() != nil {
			fmt.Println("发生异常了")
		}
	}()

	panic("error")

}

func main() {
	c := add
	sum := operator(c, 1, 2)
	fmt.Println(sum)

	d := sub
	result := operator(d, 1, 2)
	fmt.Println(result)

	a, b := methodMulRetnValue(1, 9)
	fmt.Printf("第一个返回值 a = %d, 第二个返回值 result = %s", a, b)

	fmt.Println()

	mulArgsResult := mulArgs(1, 2, 3)
	fmt.Printf("可变参数 返回结果 mulArgsResult = %d", mulArgsResult)

	fmt.Println()

	deferDemo(2)

	fmt.Println()
	runeDemo("黄小明")

	fmt.Println()
	buildInMethod()
}
