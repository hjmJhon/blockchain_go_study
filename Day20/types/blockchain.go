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
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {
	blockBytes := db.Query(blc.CurrHash)
	currBlock := Deserialize(blockBytes)
	block := NewBlock(txs, currBlock.Height+1, currBlock.Hash)
	db.Add(block.Hash, block.Serialize())
	db.Add([]byte("hash"), block.Hash)

	blc.CurrHash = block.Hash

	defer db.CloseDB()
}

/*
	将创世区块添加进区块链
*/
func AddGenesisBlockToBlockchain(txs []*Transaction) {
	currHash := db.Query([]byte("hash"))
	if currHash != nil {
		fmt.Println("创世区块已存在")
		return
	}
	block := CreateGenesisBlock(txs)
	db.Add(block.Hash, block.Serialize())
	db.Add([]byte("hash"), block.Hash)

	defer db.CloseDB()
}

/*
	获取余额
*/
func (blc *Blockchain) GetBalance(address string) int {
	var result int
	utxos := blc.GetUTXOs(address)
	for _, utxo := range utxos {
		result += utxo.Output.Value
	}
	return result
}

func (blc *Blockchain) GetUTXOs(address string) []*UTXO {
	var result []*UTXO
	var spentOutputs = make(map[string][]int)
	iterator := blc.Iterator()
	defer db.CloseDB()

	for {
		block := iterator.Next()
		for _, tx := range block.Txs {
			//inputs
			//排除掉 coinbase 交易的 input
			if tx.Inputs[0].Index != -1 {
				for _, in := range tx.Inputs {
					if in.ScriptSig == address {
						spentOutputs[in.Hash] = append(spentOutputs[in.Hash], in.Index)
					}
				}
			}

			//outputs
			fmt.Println("spentOutputs: ", spentOutputs)
			for index, out := range tx.Outputs {
				if out.ScriptPubKey != address {
					continue
				}
				if len(spentOutputs) > 0 {
					for txhash, indexSlice := range spentOutputs {
						for _, spentIndex := range indexSlice {
							if index == spentIndex && txhash == tx.TxHash {
								continue
							} else {
								utxo := &UTXO{
									TxHash: tx.TxHash,
									Index:  index,
									Output: out,
								}
								result = append(result, utxo)
							}
						}
					}
				} else {
					utxo := &UTXO{
						TxHash: tx.TxHash,
						Index:  index,
						Output: out,
					}
					result = append(result, utxo)
				}
			}
		}

		if block.Height == 0 {
			break
		}
	}

	return result
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

	defer db.CloseDB()
}
