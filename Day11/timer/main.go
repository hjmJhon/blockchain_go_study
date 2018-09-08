package main

import (
	"time"
	"fmt"
)

func queryDb(statusChan chan int) {

	//处理数据库连接等业务逻辑
	statusChan <- 200
}

//超时处理
func main() {
	statusChan := make(chan int, 1)

	tick := time.Tick(time.Second * 3)

	go func() {
		time.Sleep(time.Second * 3)
		queryDb(statusChan)
	}()

	select {
	case <-statusChan:
		fmt.Println("连接成功")
	case <-tick:
		fmt.Println("timeout")
	}
}
