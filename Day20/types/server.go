package types

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"study.com/Day20/utils"
)

const NETWORK = "tcp"
const NODE_VERSION = 1
const COMMAND_LENGTH = 12

var nodeAddress string
var KnownNodes = []string{"localhost:3000"}
var blocksInTransit = [][]byte{}
var txMemPool = make(map[string]*Transaction)

/*
	区块之间检查是否有新区块时的交互数据
*/
type version struct {
	Version int
	//当前节点区块的最新高度
	BestHeight uint64
	//发送者地址
	AddrFrom string
}

/*
	区块之间获取最新区块时的交互数据
*/
type getblocks struct {
	AddrFrom string
}

type inv struct {
	AddrFrom string
	Type     string
	Items    [][]byte
}

type getdata struct {
	AddrFrom string
	Type     string
	Item     []byte
}

type block struct {
	AddrFrom string
	Block    []byte
}

type tx struct {
	AddFrom     string
	Transaction []byte
}

func StartServer() {
	nodeId := utils.GetNodeId()
	nodeAddress = fmt.Sprintf("localhost:%s", nodeId)
	listener, e := net.Listen(NETWORK, nodeAddress)
	if e != nil {
		log.Panic(e)
	}
	defer listener.Close()

	blc := GetBlockchain()
	CheckBlockchain(blc)

	//向其他节点广播当前节点的区块信息
	if nodeAddress != KnownNodes[0] {
		sendVersion(KnownNodes[0], blc)
	}

	for {
		conn, e := listener.Accept()
		if e != nil {
			continue
		}
		go handleConn(conn, blc)
	}

}

func sendVersion(addr string, blc *Blockchain) {
	v := version{
		Version:    NODE_VERSION,
		BestHeight: blc.GetLatestHeight(),
		AddrFrom:   nodeAddress,
	}
	data := gobEncode(v)
	request := append(commandToBytes("version"), data...)

	sendData(addr, request)
}

func sendData(addr string, request []byte) {
	conn, e := net.Dial(NETWORK, addr)
	if e != nil {
		fmt.Printf("%s is not available \n", addr)
		var tempNodeAddrs []string
		for _, node := range KnownNodes {
			if addr != node {
				tempNodeAddrs = append(tempNodeAddrs, node)
			}
		}
		KnownNodes = tempNodeAddrs

		return
	}
	defer conn.Close()

	_, e = io.Copy(conn, bytes.NewReader(request))
	if e != nil {
		log.Fatal(e)
	}
}

func handleConn(conn net.Conn, blc *Blockchain) {
	data, e := ioutil.ReadAll(conn)
	if e != nil {
		return
	}

	command := bytesToCommand(data[:COMMAND_LENGTH])
	fmt.Println("command:", command)

	switch command {
	case "version":
		handleVersion(data, blc)
		break
	case "getblocks":
		handleGetBlocks(data, blc)
		break
	case "inv":
		handleInv(data, blc)
		break
	case "getdata":
		handleGetData(data, blc)
		break
	case "block":
		handleBlock(data, blc)
		break
	case "tx":
		handleTx(data, blc)
		break
	default:
		fmt.Println("command can not be identified")
		break
	}
}
func handleTx(request []byte, blc *Blockchain) {
	var buff bytes.Buffer
	var payload tx

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	//将交易添加到交易缓存池
	tx := DeserializeTx(payload.Transaction)
	txMemPool[tx.TxHash] = tx

	//中心节点不挖矿,而将交易消息发送给其他节点
	if nodeAddress == KnownNodes[0] {
		for _, node := range KnownNodes {
			if node != nodeAddress && node != payload.AddFrom {
				sendInv(node, "tx", [][]byte{[]byte(tx.TxHash)})
			}
		}
	} else {
		//挖矿
		for len(txMemPool) > 0 {
			var validaTxs []*Transaction
			for _, tx := range txMemPool {
				//验证每一个交易
				txArr := []*Transaction{tx}
				if blc.verifyTx(txArr) {
					validaTxs = append(validaTxs, tx)
				}
			}

			if len(validaTxs) == 0 {
				fmt.Println("no valida tx ")
				return
			}

			blc.AddBlockToBlockchain(validaTxs)

			for _, tx := range validaTxs {
				delete(txMemPool, tx.TxHash)
			}

			//向其他节点广播最新的区块消息
			for _, node := range KnownNodes {
				if nodeAddress != node {
					sendInv(node, "block", [][]byte{blc.CurrHash})
				}
			}
		}
	}

}

func SendTx(addr string, transaction *Transaction) {
	data := tx{
		AddFrom:     nodeAddress,
		Transaction: transaction.Serialize(),
	}

	request := append(commandToBytes("tx"), gobEncode(data)...)

	sendData(addr, request)
}

func handleBlock(request []byte, blc *Blockchain) {
	var buff bytes.Buffer
	var payload block

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := Deserialize(payload.Block)
	fmt.Println("received a block:")
	blockData.ToString()

	blc.AddBlock(blockData)
	fmt.Println("Added block, hash is", blockData.Hash)

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		sendGetData(payload.AddrFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	} else {
		ResetUTXOTable()
	}

}
func handleGetData(request []byte, blc *Blockchain) {
	var buff bytes.Buffer
	var payload getdata

	buff.Write(request[COMMAND_LENGTH:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	if payload.Type == "block" {
		block := GetBlockByHash(payload.Item)
		if block == nil {
			return
		}
		sendBlock(payload.AddrFrom, block)
	}

	if payload.Type == "tx" {
		txHash := hex.EncodeToString(payload.Item)
		tx := txMemPool[txHash]

		SendTx(payload.AddrFrom, tx)
	}
}

func sendBlock(addr string, b *Block) {
	data := block{
		AddrFrom: nodeAddress,
		Block:    b.Serialize(),
	}
	request := append(commandToBytes("block"), gobEncode(data)...)

	sendData(addr, request)
}

func handleInv(request []byte, blc *Blockchain) {
	var buff bytes.Buffer
	var payload inv

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("received inv %d 个,type is %s \n", len(payload.Items), payload.Type)

	if payload.Type == "block" {
		blocksInTransit = payload.Items
		blockHash := blocksInTransit[0]
		sendGetData(payload.AddrFrom, "block", blockHash)

		newItems := [][]byte{}
		for _, item := range blocksInTransit {
			if bytes.Compare(item, blockHash) != 0 {
				newItems = append(newItems, item)
			}
		}
		blocksInTransit = newItems
	}

	if payload.Type == "tx" {
		txHash := payload.Items[0]

		if txMemPool[hex.EncodeToString(txHash)] == nil {
			sendGetData(payload.AddrFrom, "tx", txHash)
		}
	}

}
func sendGetData(addr string, kind string, item []byte) {
	data := getdata{
		AddrFrom: nodeAddress,
		Type:     kind,
		Item:     item,
	}
	request := append(commandToBytes("getdata"), gobEncode(data)...)

	sendData(addr, request)

}

func handleGetBlocks(request []byte, blc *Blockchain) {
	var buff bytes.Buffer
	var payload getblocks

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	hashes := blc.GetBlockHashes()

	sendInv(payload.AddrFrom, "block", hashes)

}

func sendInv(addr string, kind string, items [][]byte) {
	invData := inv{
		AddrFrom: nodeAddress,
		Type:     kind,
		Items:    items,
	}
	dataByte := gobEncode(invData)
	request := append(commandToBytes("inv"), dataByte...)

	sendData(addr, request)
}

func handleVersion(request []byte, blc *Blockchain) {
	var buff bytes.Buffer
	var payload version

	buff.Write(request[COMMAND_LENGTH:])
	decoder := gob.NewDecoder(&buff)
	err := decoder.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	latestHeight := blc.GetLatestHeight()
	if latestHeight > payload.BestHeight {
		sendVersion(payload.AddrFrom, blc)
	} else if latestHeight < payload.BestHeight {
		sendGetBlocks(payload.AddrFrom)
	}

	if !nodeIsKnown(payload.AddrFrom) {
		KnownNodes = append(KnownNodes, payload.AddrFrom)
	}
}

func sendGetBlocks(addr string) {
	data := getblocks{AddrFrom: nodeAddress}
	request := append(commandToBytes("getblocks"), gobEncode(data)...)

	sendData(addr, request)
}

func gobEncode(data interface{}) []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(data)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func commandToBytes(command string) []byte {
	var bytes [COMMAND_LENGTH]byte

	for i, c := range command {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func bytesToCommand(bytes []byte) string {
	var command []byte

	for _, b := range bytes {
		if b != 0x0 {
			command = append(command, b)
		}
	}

	return fmt.Sprintf("%s", command)
}

func nodeIsKnown(addr string) bool {
	for _, node := range KnownNodes {
		if node == addr {
			return true
		}
	}

	return false
}
