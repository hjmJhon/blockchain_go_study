package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"study.com/Day20/db"
	"study.com/Day20/types"
	"study.com/Day20/utils"
)

func Run() {

	nodeId := utils.GetNodeId()
	if len(nodeId) == 0 {
		fmt.Println("nodeId is empty !")
		utils.Exit()
	}

	startNodeFlag := flag.NewFlagSet("startNode", flag.ExitOnError)
	createWalletFlag := flag.NewFlagSet("createWallet", flag.ExitOnError)
	addressListFlag := flag.NewFlagSet("addressList", flag.ExitOnError)
	genesisFlag := flag.NewFlagSet("genesis", flag.ExitOnError)
	sendFlag := flag.NewFlagSet("send", flag.ExitOnError)
	balanceFlag := flag.NewFlagSet("balance", flag.ExitOnError)
	printBlockchainFlag := flag.NewFlagSet("printBlockchain", flag.ExitOnError)
	testFlag := flag.NewFlagSet("test", flag.ExitOnError)

	genesisFlagValue := genesisFlag.String("address", "", "create the genesis block's address")
	sendFlagFromValue := sendFlag.String("from", "", "the address sending asset")
	sendFlagToValue := sendFlag.String("to", "", "the address receiving asset")
	sendFlagAmountValue := sendFlag.String("amount", "", "asset amount")
	sendFlagMine := sendFlag.Bool("mine", false, "mine immediately on this node")
	balanceFlagValue := balanceFlag.String("address", "", "get the balance of the specified address")

	args := os.Args

	checkArgsValidate(args)

	switch args[1] {
	case "createWallet":
		if err := createWalletFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	case "addressList":
		if err := addressListFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	case "genesis":
		if err := genesisFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	case "startNode":
		if err := startNodeFlag.Parse(args[2:]); err != nil {
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
	case "test":
		if err := testFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	default:
		utils.Exit()
	}

	//创建钱包
	if createWalletFlag.Parsed() {
		createWallet()
	}

	//获取所有已经创建的钱包地址
	if addressListFlag.Parsed() {
		addressList()
	}

	//创建创世区块
	if genesisFlag.Parsed() {
		createGenesisBlock(*genesisFlagValue)
	}

	//启动节点
	if startNodeFlag.Parsed() {
		fmt.Println("nodeId:", nodeId)
		startNode()
	}

	if balanceFlag.Parsed() {
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
			send(from, to, amount, *sendFlagMine)
		}
	}

	//打印所有的区块
	if printBlockchainFlag.Parsed() {
		printBlockchain()
	}

	//测试一些 cmd
	if testFlag.Parsed() {
		defer db.CloseDB()
		types.GetAllUTXOs()
		//types.ResetUTXOTable()
	}
}

func checkArgsValidate(args []string) {
	//fmt.Println("args:",args)
	if len(args) == 1 {
		utils.Exit()
	}
}
