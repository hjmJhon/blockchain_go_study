package types

import (
	"bytes"
	"study.com/Day20/utils"
)

type TxInput struct {
	//tx 的 TxHash
	TxHash string
	Index  int
	//签名
	Signature []byte
	//公钥, 即钱包中的 PublicKey
	PublicKey []byte
}

func (txInput *TxInput) TxInputCanUnLock(address string) bool {
	addressHash := utils.GetAddressHash(address)
	ripemd160Hash := utils.Ripemd160Hash(txInput.PublicKey)
	return bytes.Compare(addressHash, ripemd160Hash) == 0
}
