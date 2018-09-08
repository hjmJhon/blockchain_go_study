package main

import "fmt"

func main() {
	//deferTest1()

	deferTest2()

}

//以下程序, panic 执行的顺序不固定
func deferTest1() {
	defer func() {
		fmt.Println("最后执行")
	}()

	defer func() {
		fmt.Println("第二位执行")
	}()

	defer func() {
		fmt.Println("最先执行")
	}()

	panic("发生异常了")
}

//
func deferTest2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕捉到异常,err:", err)
		}
		fmt.Println("最后执行")
	}()

	defer func() {

		fmt.Println("第二位执行")
	}()

	defer func() {

		fmt.Println("最先执行")
	}()

	panic("发生异常了")
}
