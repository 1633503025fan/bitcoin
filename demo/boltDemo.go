package main

import (
	"github.com/bolt"
	"log"
	"os"
	"fmt"
)

func main()  {
	//打开数据库
	//func Open(path string, mode os.FileMode, options *Options) (*DB, error) {
	db,err:=bolt.Open("test.db",0600,nil)
	if err!=nil{
		log.Panic("bolt.Open failed",err)
		os.Exit(1)
	}
	defer db.Close()
	//func (db *DB) Update(fn func(*Tx) error) error {
	//写数据库
	db.Update(func(tx *bolt.Tx) error {
		//找到抽屉bucket,不存在创建
		//func (tx *Tx) Bucket(name []byte) *Bucket {
		bucket:=tx.Bucket([]byte("firstBucket"))
		if bucket==nil{
			//func (tx *Tx) CreateBucket(name []byte) (*Bucket, error) {
			bucket,err=tx.CreateBucket([]byte("firstBucket"))
			if err!=nil{
				log.Panic("createBucket failed",err)
				os.Exit(1)
			}
		}
		bucket.Put([]byte("aaaaa"),[]byte("hello"))
		bucket.Put([]byte("bbbbb"),[]byte("world"))
		return nil
	})
	//读数据库
	db.View(func(tx *bolt.Tx) error {
		bucket:=tx.Bucket([]byte("firstBucket"))
		if bucket==nil{
			log.Panic("抽屉不能为空")
			os.Exit(1)

		}
		//读书据
		//func (b *Bucket) Get(key []byte) []byte {
		res1:=bucket.Get([]byte("aaaaa"))
		res2:=bucket.Get([]byte("bbbbb"))
		fmt.Printf("res1=%s\n",string(res1))
		fmt.Printf("res2=%s\n",string(res2))
		return nil
	})

}
