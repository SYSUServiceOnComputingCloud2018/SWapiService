package dbOperator

import (
	"errors"

	"github.com/boltdb/bolt"
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
