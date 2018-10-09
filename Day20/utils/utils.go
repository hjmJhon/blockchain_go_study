package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

/*
	将 int64 类型转为 []byte
*/
func Int64ToByte(i int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, i)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func Json2Slice(str string) []string {
	var result []string
	json.Unmarshal([]byte(str), &result)
	return result
}
