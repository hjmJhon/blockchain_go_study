package types

import (
	"bytes"
	"study.com/Day20/utils"
)

type TxOutput struct {
	Value int
	//将钱包地址先sha256,再 ripemd160 得到的 hash
	Ripemd160Hash []byte
}

func (txOutput *TxOutput) TxOutputCanUnLock(address string) bool {
	addressHash := utils.GetAddressHash(address)
	return bytes.Compare(addressHash, txOutput.Ripemd160Hash) == 0
}

func (txOutput *TxOutput) Lock(address string) {
	addressHash := utils.GetAddressHash(address)
	txOutput.Ripemd160Hash = addressHash
}

func NewTxOutput(value int, address string) *TxOutput {
	txOutput := &TxOutput{
		Value:         value,
		Ripemd160Hash: nil,
	}
	txOutput.Lock(address)
	return txOutput
}
