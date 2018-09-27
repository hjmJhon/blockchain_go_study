package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
	"study.com/Day20/db"
	"study.com/Day20/types"
)

func Run() {
	genesisFlag := flag.NewFlagSet("genesis", flag.ExitOnError)
	addBlockFlag := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printBlockchainFlag := flag.NewFlagSet("printBlockchain", flag.ExitOnError)

	genesisFlagValue := genesisFlag.String("data", "send 100 BTC to xiaoming", "genesis block data")
	addBlockFlagValue := addBlockFlag.String("data", "send 100 HPB to xiaoming", "block data")

	args := os.Args

	checkArgsValidate(args)

	switch args[1] {
	case "genesis":
		if err := genesisFlag.Parse(args[2:]); err != nil {
			log.Panic(err)
		}
		break
	case "addBlock":
		if err := addBlockFlag.Parse(args[2:]); err != nil {
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
		types.AddGenesisBlockToBlockchain(*genesisFlagValue)

		defer db.CloseDB()
	}

	//添加区块
	if addBlockFlag.Parsed() {
		if len(*addBlockFlagValue) == 0 {
			fmt.Println("the block data can not be empty!")
			exit()
		}
		blockchain := types.GetBlockchain()
		checkBlockchain(blockchain)

		blockchain.AddBlockToBlockchain(*addBlockFlagValue)

		defer db.CloseDB()
	}

	//打印所有的区块
	if printBlockchainFlag.Parsed() {
		blockchain := types.GetBlockchain()
		checkBlockchain(blockchain)

		types.PrintBlockChain(blockchain)

		defer db.CloseDB()
	}
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
	fmt.Println("genesis -data:create genesis block and add to the blockchain")
	fmt.Println("addBlock -data:create block and add to the blockchain")
	fmt.Println("printBlockchain:print the all block")
}
