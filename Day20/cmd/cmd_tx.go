package cmd

import "study.com/Day20/types"

//发送交易
func send(from []string, to []string, amount []string) {
	blockchain := types.GetBlockchain()
	types.CheckBlockchain(blockchain)
	checkTxArgs(from, to, amount)

	var txs []*types.Transaction
	for index, fromAddr := range from {
		tx := types.NewTx(fromAddr, to[index], amount[index], blockchain, txs)
		txs = append(txs, tx)
	}
	blockchain.AddBlockToBlockchain(txs)
}
