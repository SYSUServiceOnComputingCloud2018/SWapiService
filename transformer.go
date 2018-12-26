//用来将bolt.db的数据导入到mySQL中
/*
MySQL数据结构
一个数据库：
	七个表：People,Planet,Starships,Film,Species,Vehicles,Schema
	七个表，表有两个项目：key和value
docker run -p 3306:3306 --name mymysql -v $PWD/conf:/etc/mysql/conf.d -v $PWD/logs:/logs -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
*/
package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//设置个人库地址
	const (
		userName = "root"
		password = "123456"
		ip       = "127.0.0.1"
		port     = "3306"
		dbName   = "test"
	)

	//Db数据库连接池

	//注意方法名大写，就是public
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	//path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	//DB, err := sql.Open("mysql", path)
	db, err := sql.Open("mysql", "root:123456@tcp(0.0.0.0:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//fmt.Println(reflect.TypeOf(db).String())

	if err := db.Ping(); err != nil {
		fmt.Println("opon database fail")
		panic(err)
	}
	fmt.Println("connnect success")
	//打开原数据库
	boltdb, err := bolt.Open("my.db", 0600, nil)
	defer boltdb.Close()
	if err != nil {
		log.Fatal(err)
	}

	//创建并使用sql数据库

	_, err = db.Exec("CREATE DATABASE " + dbName)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	blockNameSet := []string{"Person", "Starship", "Planet", "Film", "Species", "Vehicle"}

	for _, blockName := range blockNameSet {
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS `" + blockName + "` (`key` VARCHAR(255),`value` TEXT)")
		if err != nil {
			panic(err)
		}
		err := boltdb.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte(blockName))

			c := bucket.Cursor()
			tx2, err := db.Begin()
			if err != nil {
				panic(err)
			}
			for k, v := c.First(); k != nil; k, v = c.Next() {
				stmt, err := tx2.Prepare("INSERT INTO " + blockName + " (`key`, `value`) VALUES (?, ?)")
				if err != nil {
					fmt.Println("Prepare fail")
					panic(err)
				}
				//将参数传递到sql语句中并且执行
				_, err = stmt.Exec(k, v)
				if err != nil {
					fmt.Println("Exec fail")
					panic(err)
				}
			}
			tx2.Commit()
			return nil
		})
		if err != nil {
			panic(err)
		}
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `Schema` (`key` VARCHAR(255),`value` TEXT)")
	if err != nil {
		panic(err)
	}
	err = boltdb.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("Schema"))

		c := bucket.Cursor()
		tx2, err := db.Begin()
		if err != nil {
			panic(err)
		}
		for k, v := c.First(); k != nil; k, v = c.Next() {
			stmt, err := tx2.Prepare("INSERT INTO `Schema` (`key`, `value`) VALUES (?, ?)")
			if err != nil {
				fmt.Println("Prepare fail")
				panic(err)
			}
			//将参数传递到sql语句中并且执行
			_, err = stmt.Exec(k, v)
			if err != nil {
				fmt.Println("Exec fail")
				panic(err)
			}
		}
		tx2.Commit()
		return nil
	})
	if err != nil {
		panic(err)
	}

}
