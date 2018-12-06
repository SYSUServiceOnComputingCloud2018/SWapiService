package main

import (
	"encoding/json"
	"fmt"

	"github.com/SYSUServiceOnComputingCloud2018/SWapiService/dbOperator"
	"github.com/boltdb/bolt"
	"github.com/peterhellberg/swapi"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	/*
		v, err := dbOperator.GetElementById(db, "Person", "2")
	*/
	v, err := dbOperator.GetElementBySearchFields(db, "Person", "Luke Skywalker")
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
	}

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
