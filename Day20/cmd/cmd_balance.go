package cmd

import (
	"fmt"
	"study.com/Day20/db"
	"study.com/Day20/types"
)

//获取余额
func getBalance(address string) {
	//utils.CheckAddress(address)
	//blockchain := types.GetBlockchain()
	//types.CheckBlockchain(blockchain)
	//balance := blockchain.GetBalance(address)
	defer db.CloseDB()
	balance := types.GetBalance(address)
	fmt.Println("balance:", balance)
}
