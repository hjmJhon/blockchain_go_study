package main

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	"net"
	"study.com/Day19/block"
	"study.com/Day19/pos"
	"sync"
	"time"
)

var Mutex = &sync.Mutex{}

func main() {

	//创建创世区块
	firstBlock := block.Block{
		PrevHash:  "",
		Data:      "创世区块",
		Index:     1,
		Timestamp: time.Now().String(),
		BPM:       1,
		Validator: "xiaoming",
	}
	firstBlock.Hash = block.HashCode(firstBlock)
	block.Blockchain = append(block.Blockchain, firstBlock)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("err: ", err)
		return
	}
	fmt.Println("listener is running on port 8080")
	defer listener.Close()

	go func() {
		for cb := range block.CandidateBlocks {
			Mutex.Lock()
			block.TempBlocks = append(block.TempBlocks, cb)
			Mutex.Unlock()
		}
	}()

	go func() {
		for {
			pos.PickWinner()
		}
	}()

	// 接收验证者节点的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("err: ", err)
		}
		go pos.HandleConn(conn)
	}

}
