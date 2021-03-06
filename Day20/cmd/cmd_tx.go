package cmd

import (
	"study.com/Day20/db"
	"study.com/Day20/types"
	"study.com/Day20/utils"
)

//发送交易
func send(from []string, to []string, amount []string, mineNow bool) {
	defer db.CloseDB()

	blockchain := types.GetBlockchain()
	types.CheckBlockchain(blockchain)
	utils.CheckTxArgs(from, to, amount)

	var txs []*types.Transaction
	for index, fromAddr := range from {
		tx := types.NewTx(fromAddr, to[index], amount[index], txs)
		if tx == nil {
			continue
		}
		txs = append(txs, tx)
	}

	if mineNow {
		blockchain.AddBlockToBlockchain(txs)
	} else {
		for _, tx := range txs {
			types.SendTx(types.KnownNodes[0], tx)
		}
	}
}
