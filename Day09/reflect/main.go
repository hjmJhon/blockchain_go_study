package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string `json:"StudentName"`
	age  int
	score float64
}

func (s *Student) Study() {
	fmt.Println("study method is running")
}

//反射结构体
func reflect1(a interface{}) {
	ty := reflect.TypeOf(a)
	fmt.Println("type:", ty)

	value := reflect.ValueOf(a)
	fmt.Println("value: ", value)

	kind := value.Kind()
	fmt.Println("kind:", kind)

	if kind != reflect.Ptr && kind != reflect.Struct {
		return
	}

	fieldNum := value.Elem().NumField()
	fmt.Println("FieldNum = ", fieldNum)

	value.Elem().Field(0).SetString("huangjiangming")

	methodNum := value.Elem().NumMethod()
	fmt.Println("methodNum=", methodNum)

	for i := 0; i < fieldNum; i++ {
		val := value.Elem().Field(i)
		fmt.Println("val", "[",i,"]=",val)
	}

	numField := ty.Elem().NumField()
	numMethod := ty.Elem().NumMethod()
	fmt.Println("numField=",numField,"numMethod=",numMethod)

	field := ty.Elem().Field(1)
	fmt.Println("field",field)
	tag := field.Tag.Get("json")
	fmt.Println("tag",tag)

}

func main() {

	stu := Student{
		Name: "小明",
		age:  0,
		score:90.9,
	}

	stu.Study()

	reflect1(&stu)

	elemDemo()
}

//反射基本类型
func elemDemo() {
	var a float64 = 1
	value := reflect.ValueOf(&a)
	elem := value.Elem()//Elem()方法的调用者必须是指针!
	value.Elem().SetFloat(9)
	fmt.Println("elemDemo elem=", elem,"value=",value,"a=",a)
}
