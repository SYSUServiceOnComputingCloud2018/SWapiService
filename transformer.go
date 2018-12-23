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
	"database/sql"
	"fmt"
	"strings"
	"log"
)

func main(){


	const (
		userName = "root"
		password = "123456"
		ip = "127.0.0.1"
		port = "3306"
		dbName = "wangyx"
	)
	//Db数据库连接池
	var DB *sql.DB
	
	//注意方法名大写，就是public
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil{
		fmt.Println("opon database fail")
		return
	}
	fmt.Println("connnect success")

	/*
	//测试插入数据
	tx, err := DB.Begin()
    if err != nil{
        fmt.Println("tx fail")
    }
    //准备sql语句
    stmt, err := tx.Prepare("INSERT INTO people (`key`, `value`) VALUES (?, ?)")
    if err != nil{
        fmt.Println("Prepare fail")
    }
    //将参数传递到sql语句中并且执行
    res, err := stmt.Exec("王迎旭", "16340226")
    if err != nil{
        fmt.Println("Exec fail")
    }
    //将事务提交
    tx.Commit()
	//获得上一个插入自增的id
	fmt.Println(res.LastInsertId())
	*/
	



	//测试查询数据
	var key = "16340226"
	rows, err := DB.Query(" SELECT value FROM people where `key`= ? ",key)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var value string 
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s  %s\n", key, value)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}


	/*
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/betting?charset=utf8")
	if err != nil{
		fmt.
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
	*/

}
