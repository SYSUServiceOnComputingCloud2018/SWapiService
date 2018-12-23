//用来将bolt.db的数据导入到mySQL中
/*
MySQL数据结构
一个数据库：
	七个表：People,Planet,Starships,Film,Species,Vehicles,Schema
	七个表，表有两个项目：key和value
*/
package main

import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/boltdb/bolt"
	"database/sql"
	"fmt"
)

func main(){
	db,err = sql.Open("mysql","...")
	if err != nil{
		panic(err)
	}
	defer db.Close()

	boltdb, err := bolt.Open("my.db", 0600, nil)
	defer boltdb.Close()
	if err != nil {
		log.Fatal(err)
	}

	name:= bolt
	_,err = db.Exec("CREATE DATABASE "+name)
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("USE "+name)
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE People (key VARCHAR(255),value TEXT)")
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE Planet (key VARCHAR(255),value TEXT)")
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE Starships (key VARCHAR(255),value TEXT)")
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE Film (key VARCHAR(255),value TEXT)")
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE Species (key VARCHAR(255),value TEXT)")
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE Vehicles (key VARCHAR(255),value TEXT)")
	if err != nil{
		panic(err)
	}

	_,err = db.Exec("CREATE TABLE Schema (key VARCHAR(255),value TEXT)")
	if err != nil{
		panic(err)
	}



	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockName))
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			str := "INSERT "+blockName+" VALUES ("+ k +"," + v +")"
		}
		return nil
	})

}
