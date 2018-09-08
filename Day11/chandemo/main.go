package main

import "fmt"

//求 100 以内的所有偶数
func main() {
	//开 8 个协程求偶数
	intChan := make(chan int, 100)
	resultChan := make(chan int, 100)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 1; i <= 100; i++ {
			intChan <- i
		}

		close(intChan)
	}()

	//开 8 个协程去计算
	for i := 0; i < 8; i++ {
		go caculate(intChan, resultChan, exitChan)
	}

	//开启一个协程,等待计算的8个协程全部退出
	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(resultChan)
		close(exitChan)
	}()


	for e := range resultChan {
		fmt.Println("偶数 e = ", e)
	}

}

func caculate(intChan chan int, resultChan chan int, exitChan chan bool) {
	for e := range intChan {
		flag := false
		if e%2 == 0 {
			flag = true
		}

		if flag {
			resultChan <- e
		}
	}

	exitChan <- true
}
