package types

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"strconv"
	"study.com/Day20/db"
	"study.com/Day20/utils"
)

type Blockchain struct {
	CurrHash []byte
	DB       *bolt.DB
}

/*
	将区块添加进区块链
*/
func (blc *Blockchain) AddBlockToBlockchain(txs []*Transaction) {
	//验签
	if blc.verifyTx(txs) == false {
		log.Panic("error:invalidate tx")
	}
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
	defer db.CloseDB()

	var result int
	utxos := blc.GetUTXOs(address, []*Transaction{})
	for _, utxo := range utxos {
		result += utxo.Output.Value
	}
	return result
}

func (blc *Blockchain) GetUTXOs(address string, txs []*Transaction) []*UTXO {
	var result []*UTXO
	var spentOutputs = make(map[string][]int)

	for _, tx := range txs {
		if tx.Inputs[0].Index != -1 {
			for _, in := range tx.Inputs {
				if in.TxInputCanUnLock(address) {
					spentOutputs[in.Hash] = append(spentOutputs[in.Hash], in.Index)
				}
			}
		}
	}

	for _, tx := range txs {
	B:
		for index, out := range tx.Outputs {
			if out.TxOutputCanUnLock(address) == false {
				continue
			}
			if len(spentOutputs) > 0 {
				var isUTXOSpent bool
				for txhash, indexSlice := range spentOutputs {
					for _, spentIndex := range indexSlice {
						if index == spentIndex && txhash == tx.TxHash {
							isUTXOSpent = true
							continue B
						}
					}
				}

				if isUTXOSpent == false {
					utxo := &UTXO{
						TxHash: tx.TxHash,
						Index:  index,
						Output: out,
					}
					result = append(result, utxo)
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

	iterator := blc.Iterator()
	for {
		block := iterator.Next()
		for i := len(block.Txs) - 1; i >= 0; i-- {
			tx := block.Txs[i]
			//inputs
			//排除掉 coinbase 交易的 input
			if tx.IsCoinbase() == false {
				for _, in := range tx.Inputs {
					if in.TxInputCanUnLock(address) {
						spentOutputs[in.Hash] = append(spentOutputs[in.Hash], in.Index)
					}
				}
			}

			//outputs
			fmt.Println("spentOutputs: ", spentOutputs)
		A:
			for index, out := range tx.Outputs {
				if out.TxOutputCanUnLock(address) == false {
					continue
				}
				if len(spentOutputs) > 0 {
					var isUTXOSpent bool
					for txhash, indexSlice := range spentOutputs {
						for _, spentIndex := range indexSlice {
							if index == spentIndex && txhash == tx.TxHash {
								isUTXOSpent = true
								continue A
							}
						}
					}

					if isUTXOSpent == false {
						utxo := &UTXO{
							TxHash: tx.TxHash,
							Index:  index,
							Output: out,
						}
						result = append(result, utxo)
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

/*
	获取用于交易的 UTXO
*/
func (blc *Blockchain) GetSpendableUTXOs(from, to, amount string, txs []*Transaction) (int, []*UTXO) {
	utxos := blc.GetUTXOs(from, txs)

	var value int
	var utxosResult []*UTXO

	number, err := strconv.Atoi(amount)
	if err != nil {
		log.Panic(err)
	}
	for _, utxo := range utxos {
		value += utxo.Output.Value
		utxosResult = append(utxosResult, utxo)
		if value >= number {
			break
		}
	}
	if value < number {
		fmt.Println("余额不足")
		return -1, utxosResult
	}

	return value, utxosResult
}
func (blc *Blockchain) findTxByTxHash(txHash string) *Transaction {
	iterator := blc.Iterator()
	for {
		block := iterator.Next()
		for _, tx := range block.Txs {
			if tx.TxHash == txHash {
				return tx
			}
		}

		if block.Height == 0 {
			break
		}
	}

	return nil
}

/*
	验签
*/
func (blc *Blockchain) verifyTx(txs []*Transaction) bool {
	for _, tx := range txs {
		prevTxs := make(map[string]*Transaction)
		for _, in := range tx.Inputs {
			transaction := blc.findTxByTxHash(in.Hash)
			if transaction != nil {
				prevTxs[in.Hash] = transaction
			}
		}
		if tx.verify(prevTxs) == false {
			return false
		}
	}

	return true
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

func CheckBlockchain(blockchain *Blockchain) {
	if blockchain == nil {
		fmt.Println("please create the genesis block first!")
		utils.Exit()
	}
}
