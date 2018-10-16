package db

import (
	"github.com/boltdb/bolt"
	"log"
)

var DB *bolt.DB

const DBPATH = "blockchain.db"
const TABLENAME_BLOCK = "block"

//const TABLENAME_UTXO = "utxo"

func openDB() {
	var err error
	if DB == nil {
		DB, err = bolt.Open(DBPATH, 0600, nil)
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
