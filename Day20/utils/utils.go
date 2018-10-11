package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"study.com/Day20/wallet"
)

/*
	将 int64 类型转为 []byte
*/
func Int64ToByte(i int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, i)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func Json2Slice(str string) []string {
	var result []string
	json.Unmarshal([]byte(str), &result)
	return result
}

func Exit() {
	printUseage()
	os.Exit(1)
}

func CheckAddress(address string) {
	if wallet.IsValidaAddress(address) == false {
		fmt.Println("the address is invalidate")
		Exit()
	}
}

func CheckTxArgs(from, to, amount []string) {
	if len(from) != len(to) || len(from) != len(amount) {
		fmt.Println("invalidate arguments")
		Exit()
	}

	for index, from := range from {
		CheckAddress(from)
		CheckAddress(to[index])
	}
}

func printUseage() {
	fmt.Println("createWallet: create the wallet")
	fmt.Println("genesis -address:create genesis block and add to the blockchain")
	fmt.Println("balabce -address:get the balance of the specified address")
	fmt.Println("send -from  -to  -amount:send transaction to the blockchain")
	fmt.Println("printBlockchain: print the all block")
}
