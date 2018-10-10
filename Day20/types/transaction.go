package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
	"strconv"
)

type Transaction struct {
	TxHash  string
	Inputs  []*TxInput
	Outputs []*TxOutput
}

func NewCoinbaseTx(address string) *Transaction {
	input := &TxInput{
		Hash:      "",
		Index:     -1,
		ScriptSig: "genesis block",
	}
	output := &TxOutput{
		Value:        10,
		ScriptPubKey: address,
	}
	tx := &Transaction{
		Inputs:  []*TxInput{input},
		Outputs: []*TxOutput{output},
	}

	tx.TxHash = Hash(tx)
	return tx
}

/*
	创建交易
*/
func NewTx(from, to, amount string, blc *Blockchain, txs []*Transaction) *Transaction {
	value, utxos := blc.GetSpendableUTXOs(from, to, amount, txs)
	if value == -1 {
		return nil
	}

	var inputs []*TxInput
	for _, utxo := range utxos {
		input := &TxInput{
			Hash:      utxo.TxHash,
			Index:     utxo.Index,
			ScriptSig: from,
		}
		inputs = append(inputs, input)
	}

	tokenAmount, _ := strconv.Atoi(amount)
	output := &TxOutput{
		Value:        tokenAmount,
		ScriptPubKey: to,
	}
	remainedOutput := &TxOutput{
		Value:        value - tokenAmount,
		ScriptPubKey: from,
	}

	tx := &Transaction{
		Inputs:  inputs,
		Outputs: []*TxOutput{output, remainedOutput},
	}

	tx.TxHash = Hash(tx)

	return tx
}

func Hash(tx *Transaction) string {

	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	return hex.EncodeToString(hash[:])
}
