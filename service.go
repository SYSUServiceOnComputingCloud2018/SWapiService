package main

import (
	"fmt"
	"encoding/json"
	"github.com/SYSUServiceOnComputingCloud2018/SWapiService/dbOperator"
	"github.com/peterhellberg/swapi"
	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

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
