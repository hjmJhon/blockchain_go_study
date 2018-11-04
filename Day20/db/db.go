package db

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"study.com/Day20/utils"
)

var DB *bolt.DB

const DBPATH = "blockchain_%s.db"
const TABLENAME_BLOCK = "block"
const TABLENAME_UTXO = "utxo"

func openDB() {
	nodeId := utils.GetNodeId()
	fmt.Println("nodeId:", nodeId)
	dbPath := fmt.Sprintf(DBPATH, nodeId)
	var err error
	if DB == nil {
		DB, err = bolt.Open(dbPath, 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

/*
	一次添加多个 数据
*/
func AddArray(keys, values [][]byte, tableName string) {
	if len(keys) != len(values) {
		log.Panic("error: len(keyArray) != len(valueArray)")
	}
	openDB()
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx, tableName)
		if bucket != nil {
			for index, _ := range keys {
				e := bucket.Put(keys[index], values[index])
				if e != nil {
					log.Panic(e)
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func Add(key, value []byte, tableName string) {
	openDB()
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx, tableName)
		if bucket != nil {
			e := bucket.Put(key, value)
			if e != nil {
				log.Panic(e)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

/*
	更新单个数据
*/
func Update(deleteKey, addKey, addValue []byte, tableName string) {
	if addKey == nil || len(addKey) == 0 {
		return
	}
	openDB()
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx, tableName)
		if bucket != nil {
			e := bucket.Delete(deleteKey)
			if e != nil {
				log.Panic(e)
			}
			e = bucket.Put(addKey, addValue)
			if e != nil {
				log.Panic(e)
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

/*
	一次更新多个数据
*/
func UpdateArray(deleteKeys, addKeys, addValues [][]byte, tableName string) {
	if deleteKeys == nil {
		return
	}
	if len(addKeys) != len(addValues) {
		return
	}
	openDB()
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx, tableName)
		if bucket != nil {
			for _, key := range deleteKeys {
				e := bucket.Delete(key)
				if e != nil {
					log.Panic(e)
				}
			}

			for index, addKey := range addKeys {
				e := bucket.Put(addKey, addValues[index])
				if e != nil {
					log.Panic(e)
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

func Query(key []byte, tableName string) []byte {
	openDB()
	var result []byte
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx, tableName)
		if bucket != nil {
			result = bucket.Get(key)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return result
}

/*
	查询所有数据
*/
func QueryAll(tableName string) map[string][]byte {
	openDB()
	var result = make(map[string][]byte)
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx, tableName)
		bucket.ForEach(func(k, v []byte) error {
			result[string(k)] = v
			return nil
		})
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return result
}

func DeleteTable(tableName string) {
	openDB()
	err := DB.View(func(tx *bolt.Tx) error {
		ifExist, _ := bucketIfExist(tx, tableName)
		if ifExist {
			tx.DeleteBucket([]byte(tableName))
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}

/*
	创建 bucket:存在则直接返回对应bucket,不存在则创建
*/
func createBucketIfNotExist(tx *bolt.Tx, tableName string) *bolt.Bucket {
	ifExist, bucket := bucketIfExist(tx, tableName)
	if ifExist == true {
		return bucket
	} else {
		bucket, _ = tx.CreateBucket([]byte(tableName))
		return bucket
	}
}

func bucketIfExist(tx *bolt.Tx, tableName string) (bool, *bolt.Bucket) {
	bucket := tx.Bucket([]byte(tableName))
	if bucket == nil {
		return false, nil
	} else {
		return true, bucket
	}
}
