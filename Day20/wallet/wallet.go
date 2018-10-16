package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/gob"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"io/ioutil"
	"log"
	"os"
	"study.com/Day20/constants"
	"study.com/Day20/utils"
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewWallet() *Wallet {
	priKey, pubKey := newKeyPair()
	return &Wallet{
		PrivateKey: priKey,
		PublicKey:  pubKey,
	}
}

/*
	根据钱包地址获取钱包
*/
func GetWallet(address string) Wallet {
	var result Wallet
	wMapList := walletList()
A:
	for _, wMap := range wMapList {
		for key, w := range wMap {
			if address == key {
				result = w
				break A
			}
		}
	}

	return result
}

/*
	获取钱包地址
*/
func (w *Wallet) GetAddress() string {
	return w.generateAddress()
}

/*
	产生钱包地址
*/
func (w *Wallet) generateAddress() string {
	ripemdBytes := utils.Ripemd160Hash(w.PublicKey)

	//拼接 version 和 ripemdBytes
	versionAppendBytes := append([]byte{constants.VERSION}, ripemdBytes...)

	//计算 checksum
	checksum := utils.Checksum(versionAppendBytes)

	//拼接 versionAppendBytes 和  checksum
	checksumAppendBytes := append(versionAppendBytes, checksum...)

	return base58.Encode(checksumAppendBytes)

}

/*
	生成公私钥对
*/
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	//fmt.Println("privateKey:", *privateKey)
	//fmt.Printf("pubKey: %x\n", pubKey)

	return *privateKey, pubKey
}

/*
	将创建的钱包写入文件中
*/
func (w *Wallet) SaveWalletToFile(address string) {
	wMap := make(map[string]Wallet)
	wMap[address] = *w

	if _, err := os.Stat(constants.WALLETDIR); err == nil {
		//do nothing
	} else {
		err := os.Mkdir(constants.WALLETDIR, 0755)
		if err != nil {
			log.Panic(err)
		}
	}

	path := constants.WALLETDIR + address + constants.WALLET_SUFFIX
	wBytes := Serialize(wMap)
	err := ioutil.WriteFile(path, wBytes, 0644)
	if err != nil {
		log.Panic(err)
	}
}

/*
	获取所有已经创建的钱包地址
*/
func AddressList() []string {
	wMapList := walletList()

	addressList := make([]string, 0)
	for _, wMap := range wMapList {
		for k, _ := range wMap {
			addressList = append(addressList, k)
		}
	}

	return addressList
}

/*
	获取所有已创建的钱包
*/
func walletList() []map[string]Wallet {
	fileInfos, e := ioutil.ReadDir(constants.WALLETDIR)
	if e != nil {
		log.Panic(e)
	}
	wMapList := make([]map[string]Wallet, 0)
	for _, fileInfo := range fileInfos {
		fileName := constants.WALLETDIR + fileInfo.Name()
		fmt.Println("wallet fileName:", fileName)
		contentBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
			continue
		}

		wMap := Deserialize(contentBytes)
		wMapList = append(wMapList, wMap)
	}
	return wMapList
}

/*
	将钱包序列化为 []byte
*/
func Serialize(wMap map[string]Wallet) []byte {
	var result bytes.Buffer

	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(wMap)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

/*
	反序列化
*/
func Deserialize(b []byte) map[string]Wallet {
	wMap := make(map[string]Wallet)

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(&wMap)
	if err != nil {
		log.Panic(err)
	}

	return wMap
}
