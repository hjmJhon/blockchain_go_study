package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
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
		utils.Exit()
	}

	//创建创世区块
	if genesisFlag.Parsed() {
		createGenesisBlock(*genesisFlagValue)
	}

	if balanceFlag.Parsed() {
		if len(*balanceFlagValue) == 0 {
			fmt.Println("the address can not be empty!")
			utils.Exit()
		}
		getBalance(*balanceFlagValue)
	}

	//发送交易
	if sendFlag.Parsed() {
		if len(*sendFlagFromValue) == 0 || len(*sendFlagToValue) == 0 || len(*sendFlagAmountValue) == 0 {
			utils.Exit()
		} else {
			from := utils.Json2Slice(*sendFlagFromValue)
			to := utils.Json2Slice(*sendFlagToValue)
			amount := utils.Json2Slice(*sendFlagAmountValue)
			send(from, to, amount)
		}
	}

	//打印所有的区块
	if printBlockchainFlag.Parsed() {
		printBlockchain()
	}
}

func checkTxArgs(from, to, amount []string) {
	if len(from) != len(to) || len(from) != len(amount) {
		fmt.Println("invalidate arguments")
		utils.Exit()
	}
}

func checkArgsValidate(args []string) {
	//fmt.Println("args:",args)
	if len(args) == 1 {
		utils.Exit()
	}
}
