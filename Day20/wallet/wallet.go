package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"log"
)

const VERSION = byte(0x00)
const ADDRESS_CHECKSUM_LEN = 4

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
