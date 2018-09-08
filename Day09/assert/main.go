package main

import "fmt"

type Student struct {
	name string
}

type Car interface {
	run()
}

func (stu *Student) run() {

}

func assert(a interface{}) {
	b, ok := a.(Car)
	if ok {
		fmt.Println("转换成功", b)
	} else {
		fmt.Println("转换失败", b)
	}
}
func main() {
	//var b  = 2
	//b := "2"
	b := new(Student)

	assert(b)
}
