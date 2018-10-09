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

func NewTx(from, to, amount string) *Transaction {
	//xiaohong 给 xiaoqiang 转 1 eth
	inputXiaohong := &TxInput{
		Hash:      "4dccdaf210ecd57a6dc86558d56733e3747d81befee68412904c0b23c88c754e",
		Index:     0,
		ScriptSig: "xiaohong",
	}
	tokenAmount, _ := strconv.Atoi(amount)
	output := &TxOutput{
		Value:        tokenAmount,
		ScriptPubKey: to,
	}
	xiaohongRemainedOutput := &TxOutput{
		Value:        1,
		ScriptPubKey: "xiaohong",
	}

	tx := &Transaction{
		Inputs:  []*TxInput{inputXiaohong},
		Outputs: []*TxOutput{output, xiaohongRemainedOutput},
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
