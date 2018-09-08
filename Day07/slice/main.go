package main

import (
	"fmt"
	"sort"
)

//切片的定义和基础 api
func slice1() {
	var slice1 = []int{1, 2, 3}
	fmt.Printf("len=%d; cap=%d", len(slice1), cap(slice1))

	fmt.Println()

	var arr = [...]int{1, 2, 3, 4, 5}
	slice2 := arr[0:]
	fmt.Printf("slice2=%v", slice2)

	fmt.Println()

	slice3 := slice2[:len(slice2)-1]
	fmt.Printf("cap=%d len=%d slice3=%v", cap(slice3), len(slice3), slice3)

	fmt.Println()

	slice4 := make([]int, 10, 20)
	slice4[9] = 20
	fmt.Printf("len=%d; cap=%d; slice=%v", len(slice4), cap(slice4), slice4)

	fmt.Println()

	var slice5 []int
	slice5 = append(slice5, 1)
	slice5 = append(slice5, 2)
	slice5 = append(slice5, 3)
	fmt.Printf("cap=%d slice5=%v", cap(slice5), slice5)
}

//切片扩容后引用数组的地址与之前数组的地址
func slice2() {
	arr := []int{1, 2}
	fmt.Println(arr)
	slice1 := arr[:]
	slice1[0] = 100
	fmt.Println(arr)
	fmt.Printf("原数组arr的地址=%p,\n切片引用的数组地址=%p ", &arr, &slice1)

	fmt.Println()

	//扩容
	slice1 = append(slice1, 3, 4, 5)
	//slice2 := append(slice1, 3, 4, 5)
	slice1[0] = 90
	fmt.Printf("原数组arr的地址=%p,\n切片引用的数组地址=%p,\nslice=%v ", &arr, &slice1, slice1)
	fmt.Println(arr)
}

//切片的拷贝
func copySlice() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5}
	i := copy(slice1, slice2)
	fmt.Println(slice1)
	fmt.Println(i)
}

//切片排序,查找
func sortSlice() {
	slice1 := []int{1, 0, 4}
	sort.Ints(slice1)
	fmt.Println(slice1)

	slice2 := []string{"B", "J", "a", "0"}
	sort.Strings(slice2)
	fmt.Println(slice2)

	index := sort.SearchInts(slice1, 1)
	fmt.Println(index)
}

//修改 string 中的字符
func modifyChar() {
	str := "黄江明"
	temp := []rune(str)
	temp[0]='忘'
	str = string(temp)
	fmt.Println(str)
}

func main() {
	slice1()

	fmt.Println()

	slice2()

	fmt.Println()

	copySlice()

	fmt.Println()

	sortSlice()

	fmt.Println()

	modifyChar()

}
