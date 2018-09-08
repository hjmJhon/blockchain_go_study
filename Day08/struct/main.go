package main

import (
	"fmt"
	"encoding/json"
)

//定义一个结构体
type Book struct {
	Grade   string
	Subject string
	Id      int
}

//struct 转为json
func toJson(book Book) string {
	bytes, e := json.Marshal(book)
	if e == nil {
		return string(bytes)
	}
	return ""
}

//json 转为 struct
func toStruct(jsonStr string) {
	var book Book
	err := json.Unmarshal([]byte(jsonStr), &book)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("book=", book)

}

//重写 string() 方法
func (p *Book) String() string {
	book := *p
	return fmt.Sprintf("Grade=%s  Subject=%s  Id=%d", book.Grade, book.Subject, book.Id)
}

//给 struct 定义方法
func (p *Book) modifyGrade(grade string) {
	p.Grade = grade
}

//继承
type Student struct {
	Book
	age int
}

func main() {
	//声明方式一:
	var Book1 Book
	Book1.Grade = "一年级"
	Book1.Subject = "数学"
	Book1.Id = 12345

	//声明方式二:值类型
	book2 := Book{
		Grade:   "三年级",
		Subject: "语文",
		Id:      23456,
	}

	//声明方式三:指针类型
	book3 := new(Book)
	fmt.Println("book3", book3)

	printBook(book2)

	printBook2(&Book1)
	printBook(Book1)

	jsonStr := toJson(book2)
	fmt.Println("json:", jsonStr)

	str := `{"Grade":"五年级","Subject":"科学","Id":910987}`

	toStruct(str)

	//String()
	book2Str := book2.String()
	fmt.Println("book2Str:", book2Str)

	//调用结构体的方法
	book2.modifyGrade("初三")
	fmt.Println("book2",book2)

	//继承
	stu := Student{
		age:12,
	}
	stu.Id=89890
	stu.Subject="生物"
	stu.Grade="大一"
	stuStr := stu.String()
	fmt.Println("stu=",stu)
	fmt.Println("stuStr:",stuStr)
}

//结构体值传递
func printBook(Book Book) {
	fmt.Println("Book1=", Book)
}

//结构体引用传递
func printBook2(book *Book) {
	(*book).Grade = "高一"
	fmt.Println("book2=", *book)
}
