package cmd

import (
	"fmt"
	"study.com/Day20/wallet"
)

func createWallet() {
	w := wallet.NewWallet()
	address := w.GetAddress()
	fmt.Println("wallet address is:", address)

	w.SaveWalletToFile(address)
}

func addressList() {
	addrList := wallet.AddressList()
	fmt.Println("addrList", addrList)
}
