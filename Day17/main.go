package main

import (
	"fmt"
	"study.com/Day17/hashmap"
	"study.com/Day17/linklist"
)

func main() {

	hashMap := hashmap.CreateHashMap()
	hashMap.Put("key0", "value0")
	valu0 := hashMap.Get("key0")
	fmt.Println("key0: ", valu0)

	hashMap.Put("key0", "小明")
	valu00 := hashMap.Get("key0")
	fmt.Println("valu00: ", valu00)

	hashMap.Put("key1","value1")
	hashMap.Put("key2","value2")
	hashMap.Put("key3","value3")
	hashMap.Put("key4","value4")
	hashMap.Put("key5","value5")
	hashMap.Put("key6","value6")
	hashMap.Put("key7","value7")

	fmt.Println("hashMap Size: ",hashMap.Size())
	fmt.Println("hashMap Cap: ",hashMap.Cap())

	fmt.Println("---------------------------------------")

	linkList := linklist.CreateLinkList()
	linkList.Append("xiaoming")
	linkList.Append("hahah")
	linkList.Insert(1,"insert")
	linkList.String()

	node := linkList.IndexOf(4)
	fmt.Println(node)

	linkList.Remove(1)
	linkList.String()

	linkList.Remove(0)
	linkList.String()
	head := linkList.Head()
	fmt.Println("head",*head)
	size := linkList.Size()
	fmt.Println("size:",size)


}
