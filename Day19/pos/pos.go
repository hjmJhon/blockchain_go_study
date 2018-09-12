package pos

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"strconv"
	"study.com/Day19/block"
	"sync"
	"time"
)

//向所有几点广播新区块的通道
var announcements = make(chan string)

var Mutex = &sync.Mutex{}

//节点的map，同时也会保存每个节点持有的 token 数量
var validators = make(map[string]int)

//30s 出一个区块
func PickWinner() {
	time.Sleep(time.Duration(30) * time.Second)

	Mutex.Lock()
	tempt := block.TempBlocks
	Mutex.Unlock()

	lotteryPool := make([]string, 0)

	if len(tempt) > 0 {
	outer:
		for _, b := range tempt {
			for _, node := range lotteryPool {
				if node == b.Validator {
					continue outer
				}
			}

			Mutex.Lock()
			setValidators := validators
			Mutex.Unlock()

			//有多少 token ,就向 lotteryPool 追加多少个 address
			tokenAmount, ok := setValidators[b.Validator]
			if ok {
				for i := 0; i < tokenAmount; i++ {
					lotteryPool = append(lotteryPool, b.Validator)
				}
			}
		}
		//从 lotteryPool 中随机选出一个节点地址作为记账者
		source := rand.NewSource(time.Now().Unix())
		r := rand.New(source)
		winner := lotteryPool[r.Intn(len(lotteryPool))]

		//将记账者写到区块链,并通知所有的节点
		for _, b := range tempt {
			if b.Validator == winner {
				Mutex.Lock()
				block.Blockchain = append(block.Blockchain, b)
				Mutex.Unlock()

				announcements <- "the winner is " + winner

				break
			}
		}

		Mutex.Lock()
		tempt = []block.Block{}
		Mutex.Unlock()
	}
}

//其他节点连接请求,用于验证区块
func HandleConn(conn net.Conn) {
	if conn == nil {
		return
	}
	defer conn.Close()

	go func() {
		message := <-announcements
		conn.Write([]byte(message))
	}()

	var address string

	conn.Write([]byte("enter your token amount:\t"))
	tokenScanner := bufio.NewScanner(conn)
	for tokenScanner.Scan() {
		tokenAmount, err := strconv.Atoi(tokenScanner.Text())
		if err != nil {
			log.Printf("not a number: %v", tokenScanner.Text())
		}

		//生成地址
		t := time.Now().String()
		address = block.CalculateHash(t)

		//存储address和对应的token
		validators[address] = tokenAmount

		break
	}

	conn.Write([]byte("enter the BPM:\t"))
	bpmScanner := bufio.NewScanner(conn)
	go func() {
		for {
			for bpmScanner.Scan() {
				bpm, err := strconv.Atoi(bpmScanner.Text())
				if err != nil {
					log.Printf("%v not a number: %v", tokenScanner.Text(), err)
					//删除此 validator
					delete(validators, address)
					conn.Close()
				}

				//生成区块
				Mutex.Lock()
				oldBlock := block.Blockchain[len(block.Blockchain)-1]
				newBlock := block.GenerateBlock(oldBlock, bpm, "区块数据", address)
				//将生成的区块发送到 CandidateBlocks
				if block.IsValidBlock(newBlock, oldBlock) {
					block.CandidateBlocks <- newBlock
				}
				Mutex.Unlock()

				conn.Write([]byte("enter the new BPM:\t"))
			}
		}
	}()

	for true {
		time.Sleep(time.Duration(10) * time.Second)
		Mutex.Lock()
		encoder := json.NewEncoder(conn)
		encoder.Encode(block.Blockchain)
		Mutex.Unlock()
	}

}
