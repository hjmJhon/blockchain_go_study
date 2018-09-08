package main

import "fmt"

func main() {

	students := ptrTest()

	for _, v := range students {
		fmt.Println(*v)
	}
}

type student struct {
	Name string
	Age  int
}

//指针
func ptrTest() map[string]*student {

	stus := []student{
		{"xiaoming001", 20},
		{"xiaoming002", 30},
	}

	stuMap := make(map[string]*student)

	for i, _ := range stus {
		//stuMap[v.Name] = &v//这样处理的结果就是, map 里面存的值是一样的(最后存进去的值)
		stuMap[stus[i].Name] = &(stus[i])//这样才能得到正确的结果
	}

	return stuMap
}
