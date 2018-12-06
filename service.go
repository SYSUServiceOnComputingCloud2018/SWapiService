package main

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
//爬虫相关
	if peopleSchema, err := CrawlSchema("people/schema"); err == nil {
		jsonCode, _ := json.Marshal(peopleSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Person"), jsonCode)
			return nil
		})
	}
	if filmSchema, err := CrawlSchema("films/schema"); err == nil {
		jsonCode, _ := json.Marshal(filmSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Film"), jsonCode)
			return nil
		})
	}
	if starshipSchema, err := CrawlSchema("starships/schema"); err == nil {
		jsonCode, _ := json.Marshal(starshipSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Starship"), jsonCode)
			return nil
		})
	}
	if vehicleSchema, err := CrawlSchema("vehicles/schema"); err == nil {
		jsonCode, _ := json.Marshal(vehicleSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Vehicle"), jsonCode)
			return nil
		})
	}
	if planetSchema, err := CrawlSchema("planets/schema"); err == nil {
		jsonCode, _ := json.Marshal(planetSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Planet"), jsonCode)
			return nil
		})
	}
	if speciesSchema, err := CrawlSchema("species/schema"); err == nil {
		jsonCode, _ := json.Marshal(speciesSchema)
		db.Update(func(tx *bolt.Tx) error {
			schemaBucket, err := tx.CreateBucketIfNotExists([]byte("Schema"))
			if err != nil {
				return err
			}
			schemaBucket.Put([]byte("Species"), jsonCode)
			return nil
		})
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
