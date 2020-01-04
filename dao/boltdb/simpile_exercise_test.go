/**
@description 测试boltdb操作方法
@参考[go test test & benchmark](https://studygolang.com/articles/7051)
**/
package main

import (
	"github.com/boltdb/bolt"
	"log"
	"testing"
)

func BenchmarkLoops(b *testing.B) {
	b.StopTimer()
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
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		get(db, key)
	}
}
