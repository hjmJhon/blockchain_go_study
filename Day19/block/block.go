package block

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Hash      string
	PrevHash  string
	Data      string
	Index     int
	Timestamp string
	BPM       int
	//取得记账权的节点地址
	Validator string
}

var Blockchain []Block
var TempBlocks []Block

// Block 的通道，任何一个节点在创建一个新区块时都将它发送到这个通道
var CandidateBlocks = make(chan Block)

func GenerateBlock(oldBlock Block, bpm int, data string, address string) Block {
	block := Block{
		PrevHash:  oldBlock.Hash,
		Data:      data,
		Index:     oldBlock.Index + 1,
		Timestamp: time.Now().String(),
		BPM:       bpm,
		Validator: address,
	}
	block.Hash = HashCode(block)
	return block
}

func IsValidBlock(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	return true
}

func HashCode(block Block) string {
	data := string(block.Index) + block.Timestamp + string(block.BPM) + block.Data + block.PrevHash
	return CalculateHash(data)
}

func CalculateHash(data string) string {
	sum := sha256.Sum256([]byte(data))
	return hex.EncodeToString(sum[:])
}
