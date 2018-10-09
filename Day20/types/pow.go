package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"study.com/Day20/utils"
)

const Target = 4

type POW struct {
	Block *Block
	//目标值,即hash中至少有多少个前导0
	Target uint64
}

func NewPOW(block *Block) *POW {
	return &POW{
		Block:  block,
		Target: Target,
	}
}

/*
	挖矿
*/
func (pow *POW) Run() (nonce uint64, hash []byte) {
	block := pow.Block
	preFix := strings.Repeat("0", int(pow.Target))
	n := uint64(0)
	h := []byte{}
	sum := [32]byte{}
	var hashStr string
	for {
		h = bytes.Join([][]byte{utils.Int64ToByte((int64)(block.Height)), BuildTxHashBytes(block.Txs), utils.Int64ToByte(int64(n)), block.PreHash, utils.Int64ToByte(block.Timestamp)}, []byte{})
		sum = sha256.Sum256(h)
		fmt.Printf("\r%x", sum)
		hashStr = hex.EncodeToString(sum[:])
		if strings.HasPrefix(hashStr, preFix) {
			break
		}

		n += 1
	}

	return n, sum[:]
}
