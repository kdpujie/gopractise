/**
@description  boltDB练习
@author pujie
@data	2018-01-18
@参考
	[BoltDB中文网](http://boltdb.cn/quickstart.html)
**/

package main

import (
	"github.com/boltdb/bolt"
	"log"
)

var bucketName = []byte("test")

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatalf("bold.Open() err, message=%s \n", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(bucketName)
		return err
	})
	if err != nil {
		log.Fatalf("create bucket failed, err=%s \n", err)
	}
	var key = []byte("key_test_1")
	var value = []byte("value_test_1_11")
	set(db, key, value)
	result, _ := get(db, key)
	log.Printf("search value =%s \n", string(result))
}

//写入数据
func set(db *bolt.DB, key, value []byte) {
	db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucketName).Put(key, value)
	})
}

func get(db *bolt.DB, key []byte) (b []byte, err error) {
	db.View(func(tx *bolt.Tx) error {
		b = tx.Bucket(bucketName).Get(key)
		return nil
	})
	return b, nil
}
