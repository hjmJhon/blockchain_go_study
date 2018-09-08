package main

import (
	"fmt"
)

type Student struct {
	Name string
}

//chan的声明和初始化
func chanDemo1() {
	var intChan chan int
	intChan = make(chan int, 10)

	intChan <- 19

	value := <-intChan
	fmt.Println("value", value)
}

//chan 读写 struct
func chanStruct() {
	var stuChan chan *Student
	stuChan = make(chan *Student, 10)
	stu := Student{
		Name: "小明",
	}
	stuChan <- &stu

	value := <-stuChan
	fmt.Println("value=", (*value).Name)
}

//chan 读写任意类型
func chanStruct2(a interface{}) {
	var aChan chan interface{}
	aChan = make(chan interface{}, 10)

	aChan <- a

	value := <-aChan
	fmt.Println("value=", value)
}

//缓冲
func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Println("write i=", i)
	}
}

func read(ch chan int) {
	for {
		b := <-ch
		fmt.Println("read b=", b)
	}
}

//关闭 chan
func closeChan(size int) {
	intChan := make(chan string, size)
	for i := 1; i < 10; i++ {
		intChan <- "hah"
	}
	close(intChan)

	//读取方式一
	//for {
	//	b, ok := <-intChan
	//	if ok == false {
	//		fmt.Println("chan closed")
	//		break
	//	}
	//	fmt.Println("b=", b)
	//}

	//读取方式二:
	for e := range intChan {
		fmt.Println("e=", e)
	}

}

//chan 的退出机制
func exitChan() {
	intChan := make(chan int, 10)
	exitChan := make(chan struct{}, 2)

	go sendChan(intChan, exitChan)
	go receiveChan(intChan, exitChan)

	totalCount := 0
	for _ = range exitChan {
		totalCount++
		fmt.Println("totalCount=", totalCount)
		if totalCount == 2 {
			break
		}
	}
}

//
func receiveChan(intChan chan int, exitChan chan struct{}) {
	for {
		value, ok := <-intChan
		if ok == false {
			break
		}
		fmt.Println("receiveChan value=", value)
	}

	var a struct{}
	exitChan <- a
}

//
func sendChan(intChan chan int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	close(intChan)

	var a struct{}
	exitChan <- a
}

//select 语句
func selectChan() {
	intChan := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			intChan <- i
		}
		close(intChan)
	}()


	for {
		select {
		case a ,ok := <-intChan:
			if ok == false {
				break
			}
			fmt.Println("a=", a)
		default:
			fmt.Println("break")

		}
	}
}

func main() {
	//chanDemo1()
	//
	//chanStruct()
	//
	//stu := Student{
	//	Name: "hahah",
	//}
	//chanStruct2(&stu)

	//intChan := make(chan int, 10)
	//go write(intChan)
	//go read(intChan)
	//time.Sleep(time.Second * 5)

	//closeChan(10)

	//exitChan()

	selectChan()
}
