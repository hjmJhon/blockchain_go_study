package main

import "fmt"

type People struct {
}

type Teacher struct {
	People
}

func (p *People) showA() {
	fmt.Println("showPeopleA")
	p.showB()
}

func (p *People) showB() {
	fmt.Println("showPeopleB")

}

func (t *Teacher) showB() {
	fmt.Println("showTeacherB")
}

func main() {
	t := &Teacher{}
	t.showA()//执行结果:showPeopleA showPeopleB,而不是 showPeopleA,showTeacherB

}
