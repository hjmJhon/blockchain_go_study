package main

import (
	"strings"
	"fmt"
)

//闭包
func checkHasSuffix(suffix string) func(string) string {
	return func(url string) string {
		if !strings.HasSuffix(url, suffix) {
			return url + suffix
		} else {
			return url
		}
	}
}

//闭包
func checkHasSuffix2(url string) func(string) string {
	return func(suffix string) string {
		if !strings.HasSuffix(url, suffix) {
			return url + suffix
		} else {
			return url
		}
	}
}

//递归
func recursion(a int) int {
	if a == 1 {
		return a
	}
	return a * recursion(a-1)
}

//数组
func arrayTest() {
	var arr = [...]int{}
	fmt.Println(len(arr))
}

//遍历数组
func iteratorArray(){
	arr:=[...]int{1,2,3,4,5}
	for e := range arr {
		fmt.Println(arr[e])
	}

	for k, v := range arr {
		fmt.Printf("%d=%d \t",k,v)
	}
}

//二维数组
func twoDimensionArray(){
	var arr [2][3] string= [2][3]string{{"1","2","3"},{"-1","-2","-3"}}
	for k, v := range arr {
		for k2, v2 := range v {
			fmt.Printf("\n (%d,%d)=%s \t",k,k2,v2)
		}
	}
}

//数组作为参数时是值传递
func arrayTest2(arr [3]int,index int)  {
	arr[index]=10
	fmt.Println(arr[index])
}


func main() {
	hasSuffix := checkHasSuffix(".com")
	suffix := hasSuffix("https://www.baidu")
	fmt.Println(suffix)

	hasSuffix2 := checkHasSuffix2("https://www.baidu")
	suffix2 := hasSuffix2(".com")
	fmt.Println(suffix2)

	recursionResult := recursion(3)
	fmt.Println(recursionResult)

	arrayTest()

	iteratorArray()

	twoDimensionArray()

	var arr = [...]int{1, 2, 3}
	arrayTest2(arr,0)

	arr[2] = 4
	for _, v := range arr {
		fmt.Println(v)
	}
}
