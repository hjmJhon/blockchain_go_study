package types

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"strconv"
	"study.com/Day20/db"
)

type UTXOSet struct {
	UTXOs []*UTXO
}

func GetBalance(address string) int {
	var result int
	utxos := GetUTXOsByAddress(address)
	for _, utxo := range utxos {
		result += utxo.Output.Value
	}
	return result
}

func GetUTXOsByAddress(address string) []*UTXO {
	var result []*UTXO
	allUTXOs := db.QueryAll(db.TABLENAME_UTXO)
	for _, utxoSetByte := range allUTXOs {
		utxoSet := DeserializeUTXOSet(utxoSetByte)
		for _, utxo := range utxoSet.UTXOs {
			if utxo.Output.TxOutputCanUnLock(address) {
				result = append(result, utxo)
			}
		}
	}

	return result
}

func GetAllUTXOs() {
	allUTXOs := db.QueryAll(db.TABLENAME_UTXO)
	for txHash, utxoSetByte := range allUTXOs {
		fmt.Println("txhash:", txHash)
		utxoSet := DeserializeUTXOSet(utxoSetByte)
		for _, utxo := range utxoSet.UTXOs {
			fmt.Println("utxo:", *utxo)
		}
	}
}

/*
	更新 utxo 表
*/
func UpdateUTXO(txs []*Transaction) {

	var deleteKeys [][]byte
	var addKeys [][]byte
	var addValues [][]byte

	utxoMap := GetUTXOsByTxs(txs)
	for txHash, utxos := range utxoMap {
		addKeys = append(addKeys, []byte(txHash))
		utxoSet := &UTXOSet{UTXOs: utxos}
		addValues = append(addValues, utxoSet.Serialize())
	}

	allUTXOs := db.QueryAll(db.TABLENAME_UTXO)
	for _, tx := range txs {
		for _, input := range tx.Inputs {
			for txHash, utxoSetByte := range allUTXOs {
				utxoSet := DeserializeUTXOSet(utxoSetByte)
				if txHash == input.TxHash {
					var addUTXOs []*UTXO
					for _, utxo := range utxoSet.UTXOs {
						if utxo.Index == input.Index {
							//保存要删除的 key
							deleteKeys = append(deleteKeys, []byte(txHash))
							continue
						}

						//保存于删除的 key 相同的 key 中不该删除的数据
						addUTXOs = append(addUTXOs, utxo)
					}
					addKeys = append(addKeys, []byte(txHash))
					utxoSet := &UTXOSet{UTXOs: addUTXOs}
					addValues = append(addValues, utxoSet.Serialize())
				}

			}
		}
	}

	db.UpdateArray(deleteKeys, addKeys, addValues, db.TABLENAME_UTXO)
}

func GetUTXOsByTxs(txs []*Transaction) map[string][]*UTXO {
	var result = make(map[string][]*UTXO)
	var spentOutputs = make(map[string][]int)
	for _, tx := range txs {
		if tx.Inputs[0].Index != -1 {
			for _, in := range tx.Inputs {
				spentOutputs[in.TxHash] = append(spentOutputs[in.TxHash], in.Index)
			}
		}
	}
	for _, tx := range txs {
	A:
		for index, out := range tx.Outputs {
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

					result[tx.TxHash] = append(result[tx.TxHash], utxo)
				}
			} else {
				utxo := &UTXO{
					TxHash: tx.TxHash,
					Index:  index,
					Output: out,
				}
				result[tx.TxHash] = append(result[tx.TxHash], utxo)
			}
		}
	}

	return result
}

/*
	重置 utxo 表,即删除原有的 utxo 表,遍历链上所有的 utxo,将遍历结果写入到 utxo 表.
*/
func ResetUTXOTable() {
	blc := GetBlockchain()
	CheckBlockchain(blc)

	allUTXOs := blc.FindAllUTXOs()
	if len(allUTXOs) > 0 {
		db.DeleteTable(db.TABLENAME_UTXO)
	}

	var txHashArray [][]byte
	var utxoSetArray [][]byte
	for txHash, utxoSet := range allUTXOs {
		txHashArray = append(txHashArray, []byte(txHash))
		utxoSetArray = append(utxoSetArray, utxoSet.Serialize())
	}

	db.AddArray(txHashArray, utxoSetArray, db.TABLENAME_UTXO)
}

/*
	从 utxo 表中获取交易够用的 utxo
*/
func GetEnoughUTXO(from, amount string, txs []*Transaction) (int, []*UTXO) {
	unPackedUTXOs := GetUnPackedUTXOsByAddress(txs, from)
	tokenAmount, err := strconv.Atoi(amount)
	if err != nil {
		log.Panic(err)
	}

	var money int
	var utxosArr []*UTXO
	for _, utxo := range unPackedUTXOs {
		money += utxo.Output.Value
		utxosArr = append(utxosArr, utxo)

		if money >= tokenAmount {
			return money, utxosArr
		}
	}

	utxos := GetUTXOsByAddress(from)
	for _, utxo := range utxos {
		money += utxo.Output.Value
		utxosArr = append(utxosArr, utxo)

		if money >= tokenAmount {
			break
		}
	}

	if money < tokenAmount {
		log.Panic("error: no enough token to transaction")
		return -1, utxosArr
	}

	return money, utxosArr

}

/*
	区块序列化
*/
func (utxoSet *UTXOSet) Serialize() []byte {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(utxoSet)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

/*
	反序列化
*/
func DeserializeUTXOSet(utxoSetBytes []byte) *UTXOSet {
	var utxoSet UTXOSet

	decoder := gob.NewDecoder(bytes.NewReader(utxoSetBytes))
	err := decoder.Decode(&utxoSet)
	if err != nil {
		log.Panic(err)
	}

	return &utxoSet
}
