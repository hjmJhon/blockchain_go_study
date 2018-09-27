package main

import (
	"study.com/Day20/db"
	"study.com/Day20/types"
)

func main() {

	blockchain := types.AddGenesisBlockToBlockchain("genesis block")
	blockchain.AddBlockToBlockchain("second block")

	types.PrintBlockChain(blockchain)

	defer db.CloseDB()
}
