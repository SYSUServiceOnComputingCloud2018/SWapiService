package main

import (
	"encoding/json"
	"fmt"

	"github.com/SYSUServiceOnComputingCloud2018/SWapiService/dbOperator"
	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//Schema输出
	jsonData, _ := dbOperator.GetSchemaByBucket(db, "Person")
	var schema Schema //定义在crawler.go中，需要使用go run service.go crawler.go
	err = json.Unmarshal(jsonData, &schema)
	if err == nil {
		fmt.Println(schema)
	} else {
		fmt.Println(err)
	}

	//查找单个元素
	/*
			v, err := dbOperator.GetElementById(db, "Person", "2")

		if err != nil {
			fmt.Println(err)
		} else {
			var user swapi.Person
			err = json.Unmarshal(v, &user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(user)
			}
		}*/
	//查找多个元素
	/*
		v, err := dbOperator.GetElementsBySearchField(db, "Person", "Skywalker")
		if err == nil {
			for i := 0; i < len(v); i++ {
				var user swapi.Person
				err = json.Unmarshal(v[i], &user)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(user)
				}
			}
		} else{
			fmt.Println(err)
		}
	*/
	/*
		n := negroni.Classic()
		router := mux.NewRouter()
		formatter := render.New(render.Options{
			IndentJSON:true,
			Directory: "file",
			Extensions: []string{".html"},
		})

		router.HandleFUnc("/api/")*/
}
