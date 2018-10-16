package types

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
	"math/big"
	"strconv"
	"study.com/Day20/wallet"
)

type Transaction struct {
	TxHash  string
	Inputs  []*TxInput
	Outputs []*TxOutput
}

func (tx *Transaction) IsCoinbase() bool {
	return tx.Inputs[0].Index == -1 && len(tx.Inputs[0].Hash) == 0
}

func NewCoinbaseTx(address string) *Transaction {
	w := wallet.GetWallet(address)
	input := &TxInput{
		Hash:      "",
		Index:     -1,
		Signature: []byte("genesis block"),
		PublicKey: w.PublicKey,
	}
	output := NewTxOutput(10, address)
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
	w := wallet.GetWallet(from)
	for _, utxo := range utxos {
		input := &TxInput{
			Hash:      utxo.TxHash,
			Index:     utxo.Index,
			Signature: nil,
			PublicKey: w.PublicKey,
		}
		inputs = append(inputs, input)
	}

	tokenAmount, _ := strconv.Atoi(amount)
	output := NewTxOutput(tokenAmount, to)
	remainedOutput := NewTxOutput(value-tokenAmount, from)

	tx := &Transaction{
		Inputs:  inputs,
		Outputs: []*TxOutput{output, remainedOutput},
	}

	tx.TxHash = Hash(tx)

	tx.signTx(utxos, w.PrivateKey)

	return tx
}

/*
	交易签名
*/
func (tx *Transaction) signTx(utxos []*UTXO, privateKey ecdsa.PrivateKey) {
	if tx.IsCoinbase() == true {
		return
	}

	tx.sign(utxos, privateKey)
}

/*
	签名
*/
func (tx *Transaction) sign(utxos []*UTXO, privateKey ecdsa.PrivateKey) {
	if len(utxos) == 0 {
		log.Panic("error: no previous tx")
	}

	txCopy := tx.trimmedCopy()

	for index, input := range txCopy.Inputs {
		var utxo *UTXO
		for _, u := range utxos {
			if input.Hash == u.TxHash {
				utxo = u
				break
			}
		}
		if utxo == nil {
			log.Panic("error: no previous tx")
		}

		txCopy.Inputs[index].Signature = nil
		txCopy.Inputs[index].PublicKey = utxo.Output.Ripemd160Hash
		txCopy.TxHash = txCopy.trimmedTxHash()
		txCopy.Inputs[index].PublicKey = nil

		hash, _ := hex.DecodeString(txCopy.TxHash)
		r, s, err := ecdsa.Sign(rand.Reader, &privateKey, hash)
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)
		tx.Inputs[index].Signature = signature
	}
}

/*
	验签
*/
func (tx *Transaction) verify(prevTxs map[string]*Transaction) bool {
	if tx.IsCoinbase() == true {
		return true
	}

	if len(prevTxs) == 0 {
		log.Panic("error: no previous tx")
	}

	txCopy := tx.trimmedCopy()
	curve := elliptic.P256()
	for index, in := range tx.Inputs {
		prevTx := prevTxs[in.Hash]
		txCopy.Inputs[index].Signature = nil
		txCopy.Inputs[index].PublicKey = prevTx.Outputs[in.Index].Ripemd160Hash
		txCopy.TxHash = txCopy.trimmedTxHash()
		txCopy.Inputs[index].PublicKey = nil

		r := big.Int{}
		s := big.Int{}
		signLen := len(in.Signature)
		r.SetBytes(in.Signature[:signLen/2])
		s.SetBytes(in.Signature[signLen/2:])

		x := big.Int{}
		y := big.Int{}
		pubKeyLen := len(in.PublicKey)
		x.SetBytes(in.PublicKey[:pubKeyLen/2])
		y.SetBytes(in.PublicKey[pubKeyLen/2:])

		pubKey := &ecdsa.PublicKey{
			Curve: curve,
			X:     &x,
			Y:     &y,
		}

		txHash := txCopy.TxHash
		hash, _ := hex.DecodeString(txHash)
		if ecdsa.Verify(pubKey, hash, &r, &s) == false {
			return false
		}
	}

	return true
}

func (tx *Transaction) trimmedCopy() *Transaction {
	var inputs []*TxInput
	var outputs []*TxOutput
	for _, in := range tx.Inputs {
		input := &TxInput{
			Hash:      in.Hash,
			Index:     in.Index,
			Signature: nil,
			PublicKey: nil,
		}
		inputs = append(inputs, input)
	}
	for _, out := range tx.Outputs {
		output := &TxOutput{
			Value:         out.Value,
			Ripemd160Hash: out.Ripemd160Hash,
		}
		outputs = append(outputs, output)
	}

	return &Transaction{
		TxHash:  tx.TxHash,
		Inputs:  inputs,
		Outputs: outputs,
	}

}

func Hash(tx *Transaction) string {
	ser := tx.Serialize()

	hash := sha256.Sum256(ser)

	return hex.EncodeToString(hash[:])
}

func (tx *Transaction) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

func (tx *Transaction) trimmedTxHash() string {
	txCopy := *tx
	txCopy.TxHash = ""
	ser := (&txCopy).Serialize()
	sum256 := sha256.Sum256(ser)

	return hex.EncodeToString(sum256[:])
}
