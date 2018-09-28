package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"study.com/Day20/utils"
	"time"
)

type Block struct {
	Height    uint64
	Data      []byte
	Nonce     uint64
	Diff      uint64
	PreHash   []byte
	Hash      []byte
	Timestamp int64
}

func (block *Block) setHash() {
	hash := bytes.Join([][]byte{utils.Int64ToByte((int64)(block.Height)), block.Data, utils.Int64ToByte((int64)(block.Nonce)), block.PreHash, utils.Int64ToByte(block.Timestamp)}, []byte{})
	sum := sha256.Sum256(hash)
	block.Hash = sum[:]
}

/*
	创建创世区块
*/
func CreateGenesisBlock(data []byte) *Block {
	return NewBlock(data, 0, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}

/*
	创建区块
*/
func NewBlock(data []byte, height uint64, preHash []byte) *Block {
	block := &Block{
		Height:    height,
		Data:      data,
		PreHash:   preHash,
		Timestamp: time.Now().Unix(),
	}

	//挖矿
	pow := NewPOW(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	block.Diff = pow.Target

	fmt.Println()

	return block
}

/*
	将区块序列化为 []byte
*/
func (block *Block) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

/*
	反序列化
*/
func Deserialize(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

type blockString struct {
	Height    uint64 `json:"height"`
	Data      string `json:"data"`
	Nonce     uint64 `json:"nonce"`
	Diff      uint64 `json:"diff"`
	PreHash   string `json:"preHash"`
	Hash      string `json:"hash"`
	Timestamp string `json:"timestamp"`
}

func (block *Block) ToString() {
	data := string(block.Data)
	preHash := hex.EncodeToString(block.PreHash)
	hash := hex.EncodeToString(block.Hash)
	timeStamp := time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM")

	blockString := blockString{
		Height:    block.Height,
		Data:      data,
		Nonce:     block.Nonce,
		Diff:      block.Diff,
		PreHash:   preHash,
		Hash:      hash,
		Timestamp: timeStamp,
	}

	blockBytes, _ := json.Marshal(blockString)
	fmt.Println(string(blockBytes))
}