package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"os"
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

func printUseage() {
	fmt.Println("genesis -address:create genesis block and add to the blockchain")
	fmt.Println("balabce -address:get the balance of the specified address")
	fmt.Println("send -from  -to  -amount:send transaction to the blockchain")
	fmt.Println("printBlockchain:print the all block")
}
