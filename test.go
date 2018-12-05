package main

import (
	"fmt"

	"github.com/peterhellberg/swapi"
)

func main() {
	c := swapi.DefaultClient
	if atst, err := c.Person(1); err != nil {
		fmt.Println(atst)
		fmt.Println(atst.Name == "")
		fmt.Println("1", atst.Name, "1")
	} else {
		fmt.Println(err)
	}
	/*
		db, _ := bolt.Open("my.db", 0600, nil)
		defer db.Close()
		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Person"))
			v := b.Get([]byte("88"))
			var user swapi.Person
			err := json.Unmarshal(v, &user)
			if err != nil {
				return err
			}
			fmt.Println(user)
			fmt.Println(user.Name)
			return nil
		})*/
}
