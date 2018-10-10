package cmd

import (
	"fmt"
	"study.com/Day20/types"
)

//获取余额
func getBalance(address string) {
	blockchain := types.GetBlockchain()
	types.CheckBlockchain(blockchain)
	balance := blockchain.GetBalance(address)
	fmt.Println("balance:", balance)
}
