package main

import "fmt"

//map基本使用
func map1() {
	strMap := map[string]string{}
	fmt.Println(strMap)

	strMap["name"] = "小明"
	strMap["age"] = "19"
	strMap["sex"] = "man"
	fmt.Println(strMap)

	delete(strMap, "age")
	fmt.Println(strMap)

	//遍历
	for k, v := range strMap {
		fmt.Println(k, "=", v)
	}

	for e := range strMap {
		fmt.Println(e)
	}

	//反转map
	reversedMap := map[string]string{}
	var slice = make([]string, len(reversedMap))

	for k := range strMap {
		slice = append(slice, k)
	}
	fmt.Println("slice=",slice)

	for _, v := range slice {
		reversedMap[strMap[v]] = v
	}

	fmt.Println("reversedMap=",reversedMap)

}

//嵌套 map
func nestedMap() {
	nestedMap1 := map[string]map[string]int{}
	nestedMap1["key0"] = map[string]int{}
	nestedMap1["key0"]["key001"] = 0001
	nestedMap1["key0"]["key002"] = 0002
	nestedMap1["key1"] = map[string]int{}
	nestedMap1["key1"]["key1001"] = 1001
	nestedMap1["key1"]["key1002"] = 1002

	for k, v := range nestedMap1 {
		for k2, v2 := range v {
			fmt.Println(k,"[",k2,"]=",v2)
		}
	}

	fmt.Println("nestedMap1=",nestedMap1)

}

func main(){
	map1()

	nestedMap()
}