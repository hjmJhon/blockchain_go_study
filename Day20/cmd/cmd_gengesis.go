package cmd

import (
	"fmt"
	"study.com/Day20/types"
	"study.com/Day20/utils"
)

func createGenesisBlock(data string) {
	if len(data) == 0 {
		fmt.Println("the genesis block data can not be empty!")
		utils.Exit()
	}
	txs := types.NewCoinbaseTx(data)
	types.AddGenesisBlockToBlockchain([]*types.Transaction{txs})
}
