package utils

import (
	"bytes"
	"crypto/sha256"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
	"study.com/Day20/constants"
)

func GetAddressHash(address string) []byte {
	addressBytes := base58.Decode(address)
	addressHash := addressBytes[1 : len(addressBytes)-constants.ADDRESS_CHECKSUM_LEN]

	return addressHash
}

/*
	先sha256,再 ripemd160 得到 hash
*/
func Ripemd160Hash(b []byte) []byte {
	//sha256
	sum256 := sha256.Sum256(b)

	//ripemd160
	hash := ripemd160.New()
	hash.Write(sum256[:])
	ripemdBytes := hash.Sum(nil)

	return ripemdBytes
}

/*
	判断钱包地址是否合法
*/
func IsValidaAddress(address string) bool {
	if len(address) == 0 {
		return false
	}
	addrBytes := base58.Decode(address)
	//取出 addrBytes 的后 4 个 byte
	version := addrBytes[len(addrBytes)-constants.ADDRESS_CHECKSUM_LEN:]

	//取出剩下的 byte
	versionAppendBytes := addrBytes[:len(addrBytes)-constants.ADDRESS_CHECKSUM_LEN]

	//计算 checksum
	checksum := Checksum(versionAppendBytes)

	return bytes.Compare(version, checksum) == 0
}

func Checksum(bytes []byte) []byte {
	sum256 := sha256.Sum256(bytes)
	i := sha256.Sum256(sum256[:])
	return i[:constants.ADDRESS_CHECKSUM_LEN]
}
