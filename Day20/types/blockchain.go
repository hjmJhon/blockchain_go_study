package types

import (
	"fmt"
	"github.com/boltdb/bolt"
	"math/big"
	"study.com/Day20/db"
)

type Blockchain struct {
	CurrHash []byte
	DB       *bolt.DB
}

/*
	将区块添加进区块链
*/
func (blc *Blockchain) AddBlockToBlockchain(data string) {
	blockBytes := db.Query(blc.CurrHash)
	currBlock := Deserialize(blockBytes)
	block := NewBlock([]byte(data), currBlock.Height+1, currBlock.Hash)
	db.Add(block.Hash, block.Serialize())
	db.Add([]byte("hash"), block.Hash)

	blc.CurrHash = block.Hash
	blc.DB = db.DB
}

/*
	将创世区块添加进区块链
*/
func AddGenesisBlockToBlockchain(data string) {
	block := CreateGenesisBlock([]byte(data))
	db.Add(block.Hash, block.Serialize())
	db.Add([]byte("hash"), block.Hash)
}

func GetBlockchain() *Blockchain {
	currHash := db.Query([]byte("hash"))
	if currHash == nil {
		return nil
	}
	return &Blockchain{
		CurrHash: currHash,
		DB:       db.DB,
	}
}

func PrintBlockChain(blc *Blockchain) {
	iterator := blc.Iterator()
	var preHash big.Int

	fmt.Println()
	for {
		next := iterator.Next()
		next.ToString()
		preHash.SetBytes(next.PreHash)
		if big.NewInt(0).Cmp(&preHash) == 0 {
			break
		}
	}
}
