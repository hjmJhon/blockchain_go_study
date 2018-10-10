package cmd

import "study.com/Day20/types"

func printBlockchain() {
	blockchain := types.GetBlockchain()
	types.CheckBlockchain(blockchain)
	types.PrintBlockChain(blockchain)
}
