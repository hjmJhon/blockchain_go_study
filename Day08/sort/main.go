package main

import "fmt"

//冒泡排序
func bubbleSort() {
	arr := [...]int{2, 1, 0, 4, 3}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}
		}
	}
	fmt.Println(arr)
}

func main() {
	bubbleSort()
}
