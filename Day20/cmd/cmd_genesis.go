package cmd

import (
	"study.com/Day20/types"
	"study.com/Day20/utils"
)

func createGenesisBlock(address string) {
	utils.CheckAddress(address)
	txs := types.NewCoinbaseTx(address)
	types.AddGenesisBlockToBlockchain([]*types.Transaction{txs})
}
