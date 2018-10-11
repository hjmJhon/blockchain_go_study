package cmd

import (
	"fmt"
	"study.com/Day20/wallet"
)

func createWallet() {
	w := wallet.NewWallet()
	address := w.GetAddress()
	fmt.Println("wallet address is:", address)
}
