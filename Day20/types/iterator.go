package types

import (
	"github.com/boltdb/bolt"
	"study.com/Day20/db"
)

type BlockchainIterator struct {
	NextHash []byte
	DB       *bolt.DB
}

func (blc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{
		NextHash: blc.CurrHash,
		DB:       blc.DB,
	}
}

func (blcIter *BlockchainIterator) Next() *Block {
	blockBytes := db.Query(blcIter.NextHash)
	block := Deserialize(blockBytes)
	blcIter.NextHash = block.PreHash
	return block

}
