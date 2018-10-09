package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"study.com/Day20/types"
	"study.com/Day20/utils"
)

func Run() {
	genesisFlag := flag.NewFlagSet("genesis", flag.ExitOnError)
	sendFlag := flag.NewFlagSet("send", flag.ExitOnError)
	balanceFlag := flag.NewFlagSet("balance", flag.ExitOnError)
	printBlockchainFlag := flag.NewFlagSet("printBlockchain", flag.ExitOnError)

	genesisFlagValue := genesisFlag.String("address", "xiaoming", "create the genesis block's address")
	sendFlagFromValue := sendFlag.String("from", "", "the address sending asset")
	sendFlagToValue := sendFlag.String("to", "", "the address receiving asset")
	sendFlagAmountValue := sendFlag.String("amount", "", "asset amount")
	balanceFlagValue := balanceFlag.String("address", "", "get the balance of the specified address")

	args := os.Args

	checkArgsValidate(args)

	switch args[1] {
	case "genesis":
		if err := genesisFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	case "send":
		if err := sendFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	case "balance":
		if err := balanceFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	case "printBlockchain":
		if err := printBlockchainFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	default:
		exit()
	}

	//创建创世区块
	if genesisFlag.Parsed() {
		if len(*genesisFlagValue) == 0 {
			fmt.Println("the genesis block data can not be empty!")
			exit()
		}

		txs := types.NewCoinbaseTx(*genesisFlagValue)
		types.AddGenesisBlockToBlockchain([]*types.Transaction{txs})
	}

	if balanceFlag.Parsed() {
		if len(*balanceFlagValue) == 0 {
			fmt.Println("the address can not be empty!")
			exit()
		}
		getBalance(*balanceFlagValue)
	}

	//发送交易
	if sendFlag.Parsed() {
		if len(*sendFlagFromValue) == 0 || len(*sendFlagToValue) == 0 || len(*sendFlagAmountValue) == 0 {
			exit()
		} else {
			from := utils.Json2Slice(*sendFlagFromValue)
			to := utils.Json2Slice(*sendFlagToValue)
			amount := utils.Json2Slice(*sendFlagAmountValue)
			send(from, to, amount)
		}
	}

	//打印所有的区块
	if printBlockchainFlag.Parsed() {
		blockchain := types.GetBlockchain()
		checkBlockchain(blockchain)

		types.PrintBlockChain(blockchain)
	}
}

//获取余额
func getBalance(address string) {
	blockchain := types.GetBlockchain()
	checkBlockchain(blockchain)
	balance := blockchain.GetBalance(address)
	fmt.Println("balance:", balance)
}

//发送交易
func send(from []string, to []string, amount []string) {
	blockchain := types.GetBlockchain()
	checkBlockchain(blockchain)

	var txs []*types.Transaction
	txs = append(txs, types.NewTx(from[0], to[0], amount[0]))
	blockchain.AddBlockToBlockchain(txs)
}

func checkBlockchain(blockchain *types.Blockchain) {
	if blockchain == nil {
		fmt.Println("please create the genesis block first!")
		exit()
	}
}

func checkArgsValidate(args []string) {
	//fmt.Println("args:",args)
	if len(args) == 1 {
		exit()
	}
}

func exit() {
	printUseage()
	os.Exit(1)
}

func printUseage() {
	fmt.Println("genesis -address:create genesis block and add to the blockchain")
	fmt.Println("balabce -address:get the balance of the specified address")
	fmt.Println("send -from  -to  -amount:send transaction to the blockchain")
	fmt.Println("printBlockchain:print the all block")
}
