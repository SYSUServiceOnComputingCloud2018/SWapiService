package dbOperator

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/peterhellberg/swapi"
)

func GetElementById(db *bolt.DB, blockName string, id string) ([]byte, error) {
	var codedata []byte
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockName))
		codedata = bucket.Get([]byte(id))
		return nil
	})
	if err != nil {
		return []byte(""), err
	} else if len(codedata) == 0 {
		return []byte(""), errors.New("Empty data")
	}

	return codedata, nil
}

func GetElementsBySearchField(db *bolt.DB, blockName string, value string) ([][]byte, error) {
	storeData := make([][]byte, 0)
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockName))
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {

			switch blockName {
			case "Person":
				{
					var data swapi.Person
					err := json.Unmarshal(v, &data)
					if err != nil {
						return err
					}
					if strings.Contains(data.Name, value) {
						storeData = append(storeData, v)
					}
				}
			}
		}
		return nil
	})

	if err == nil {
		return storeData, nil
	} else if err != nil {
		return storeData, err
	} else {
		return storeData, errors.New("Not Found.")
	}
}
