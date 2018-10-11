package cmd

import (
	"fmt"
	"study.com/Day20/types"
	"study.com/Day20/utils"
)

//获取余额
func getBalance(address string) {
	utils.CheckAddress(address)
	blockchain := types.GetBlockchain()
	types.CheckBlockchain(blockchain)
	balance := blockchain.GetBalance(address)
	fmt.Println("balance:", balance)
}
