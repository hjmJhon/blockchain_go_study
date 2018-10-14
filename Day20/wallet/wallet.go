package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const VERSION = byte(0x00)
const ADDRESS_CHECKSUM_LEN = 4

const walletDir = "." + string(filepath.Separator) + "wallets" + string(filepath.Separator)
const walletSuffix = ".dat"

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
	判断钱包地址是否合法
*/
func IsValidaAddress(address string) bool {
	addrBytes := base58.Decode(address)
	//取出 addrBytes 的后 4 个 byte
	version := addrBytes[len(addrBytes)-ADDRESS_CHECKSUM_LEN:]

	//取出剩下的 byte
	versionAppendBytes := addrBytes[:len(addrBytes)-ADDRESS_CHECKSUM_LEN]

	//计算 checksum
	checksum := Checksum(versionAppendBytes)

	return bytes.Compare(version, checksum) == 0
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
	//sha256
	sum256 := sha256.Sum256(w.PublicKey)

	//ripemd160
	hash := ripemd160.New()
	hash.Write(sum256[:])
	ripemdBytes := hash.Sum(nil)

	//拼接 version 和 ripemdBytes
	versionAppendBytes := append([]byte{VERSION}, ripemdBytes...)

	//计算 checksum
	checksum := Checksum(versionAppendBytes)

	//拼接 versionAppendBytes 和  checksum
	checksumAppendBytes := append(versionAppendBytes, checksum...)

	return base58.Encode(checksumAppendBytes)

}

func Checksum(bytes []byte) []byte {
	sum256 := sha256.Sum256(bytes)
	i := sha256.Sum256(sum256[:])
	return i[:ADDRESS_CHECKSUM_LEN]
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

	if _, err := os.Stat(walletDir); err == nil {
		//do nothing
	} else {
		err := os.Mkdir(walletDir, 0755)
		if err != nil {
			log.Panic(err)
		}
	}

	path := walletDir + address + walletSuffix
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
	fileInfos, e := ioutil.ReadDir(walletDir)
	if e != nil {
		log.Panic(e)
	}
	wMapList := make([]map[string]Wallet, 0)
	for _, fileInfo := range fileInfos {
		fileName := walletDir + fileInfo.Name()
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
