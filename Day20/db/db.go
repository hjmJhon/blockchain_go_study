package db

import (
	"github.com/boltdb/bolt"
	"log"
)

var DB *bolt.DB

const DBPATH = "blockchain.db"
const TABLENAME_BLOCK = "block"

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

func Add(key, value []byte) {
	openDB()
	err := DB.Update(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx)
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

func Query(key []byte) []byte {
	openDB()
	var result []byte
	err := DB.View(func(tx *bolt.Tx) error {
		bucket := createBucketIfNotExist(tx)
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
	创建 bucket:存在则直接返回对应bucket,不存在则创建
*/
func createBucketIfNotExist(tx *bolt.Tx) *bolt.Bucket {
	bucket := tx.Bucket([]byte(TABLENAME_BLOCK))
	if bucket == nil {
		bucket, _ = tx.CreateBucket([]byte(TABLENAME_BLOCK))
	}
	return bucket
}
