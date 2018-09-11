package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Block struct {
	//上一个区块的 hashcode
	PreHash string
	//当前区块的 hashcode
	HashCode string
	//时间戳
	TimeStamp string
	//难度值
	Diffi int
	//区块交易数据
	Data string
	//区块索引
	Index int
	//随机值
	Nonce int
}

//创建区块请求时传递的数据
type Message struct {
	Data string
}

//互斥锁
var mutex = &sync.Mutex{}

const DIFFI = 1

var Blockchain = make([]Block, 0)

/*
	创建创世区块
*/
func GenerateFirstBlock() Block {
	block := Block{
		PreHash:   "",
		TimeStamp: time.Now().String(),
		Diffi:     0,
		Data:      "创世区块",
		Index:     1,
		Nonce:     0,
	}
	hashCode := hashCode(block)
	block.HashCode = hashCode
	Blockchain = append(Blockchain, block)
	return block
}

/*
	生成下一个区块
*/
func generateNextBlock(data string, oldBlock Block) Block {
	newBlock := Block{
		PreHash:   oldBlock.HashCode,
		TimeStamp: time.Now().String(),
		Diffi:     DIFFI,
		Data:      data,
		Index:     oldBlock.Index + 1,
		Nonce:     0,
	}
	for {
		hashCode := hashCode(newBlock)
		fmt.Println("正在挖矿,hashCode=", hashCode, "Nonce=", newBlock.Nonce)
		repeat := strings.Repeat("0", DIFFI)
		if strings.HasPrefix(hashCode, repeat) {
			newBlock.HashCode = hashCode
			fmt.Println("挖矿成功,Nonce=", newBlock.Nonce)
			break
		}
		newBlock.Nonce++
	}

	return newBlock
}

func Run() error {
	muxRouter := makeMuxRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        muxRouter,
		ReadTimeout:    time.Duration(5) * time.Second,
		WriteTimeout:   time.Duration(5) * time.Second,
		MaxHeaderBytes: 512,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error: ", err)
		return err
	} else {
		fmt.Println("success, blockchain is running... ")
	}

	return nil
}

//路由处理
func makeMuxRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/get", handleReadBlock).Methods("GET")
	router.HandleFunc("/post", handleWriteBlock).Methods("POST")
	return router
}

//生成区块的请求
func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var message Message
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&message)

	mutex.Lock()
	nextBlock := generateNextBlock(message.Data, Blockchain[len(Blockchain)-1])
	if isValidBlock(nextBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, nextBlock)
	}
	mutex.Unlock()

	responseWithJson(w, r, http.StatusOK, nextBlock)
}

//读取区块的请求
func handleReadBlock(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(Blockchain)
}

func responseWithJson(w http.ResponseWriter, r *http.Request, code int, block Block) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(block)
	w.WriteHeader(code)
}

//校验生成的区块是否合法
func isValidBlock(newBlock Block, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}

	if newBlock.PreHash != oldBlock.HashCode {
		return false
	}

	return true
}

func hashCode(block Block) string {
	data := block.PreHash + strconv.Itoa(block.Diffi) + strconv.Itoa(block.Index) + strconv.Itoa(block.Nonce) + block.Data + block.TimeStamp
	cryptedData := sha256.Sum256([]byte(data))
	return hex.EncodeToString(cryptedData[:])

}

func (block *Block) String() string {
	bytes, _ := json.Marshal(block)
	return string(bytes)
}
