package dbOperator

import (
	"encoding/json"
	"errors"
	"log"
	"strings"

	"database/sql"

	"github.com/boltdb/bolt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/peterhellberg/swapi"
)

func GetElementById(db *sql.DB, tablename string, key string) ([]byte, error) {
	rows, err := db.Query(" SELECT value FROM "+tablename+" where `key`= ? ", key)
	if err != nil {
		return []byte(""), err
	}
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		return []byte(value), nil
	}

	return []byte(""), err
}

func GetElementsBySearchField(db *sql.DB, tablename string, value string) ([][]byte, error) {
	storeData := make([][]byte, 0)

	rows, err := db.Query(" SELECT value FROM " + tablename)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var value string
		if err := rows.Scan(&v); err != nil {
			return nil, err
		}
		v := []byte(value)

		switch blockName {
		case "Person":
			{
				var data swapi.Person
				err := json.Unmarshal(v, &data)
				if err != nil {
					return nil, err
				}
				if strings.Contains(data.Name, value) {
					storeData = append(storeData, v)
				}
			}
		case "Film":
			{
				var data swapi.Film
				err := json.Unmarshal(v, &data)
				if err != nil {
					return nil, err
				}
				if strings.Contains(data.Title, value) {
					storeData = append(storeData, v)
				}
			}
		case "Starship":
			{
				var data swapi.Starship
				err := json.Unmarshal(v, &data)
				if err != nil {
					return nil, err
				}
				if strings.Contains(data.Name, value) {
					storeData = append(storeData, v)
				} else if strings.Contains(data.Model, value) {
					storeData = append(storeData, v)
				}
			}
		case "Vehicle":
			{
				var data swapi.Vehicle
				err := json.Unmarshal(v, &data)
				if err != nil {
					return nil, err
				}
				if strings.Contains(data.Name, value) {
					storeData = append(storeData, v)
				} else if strings.Contains(data.Model, value) {
					storeData = append(storeData, v)
				}
			}
		case "Planet":
			{
				var data swapi.Planet
				err := json.Unmarshal(v, &data)
				if err != nil {
					return nil, err
				}
				if strings.Contains(data.Name, value) {
					storeData = append(storeData, v)
				}
			}
		case "Species":
			{
				var data swapi.Species
				err := json.Unmarshal(v, &data)
				if err != nil {
					return nil, err
				}
				if strings.Contains(data.Name, value) {
					storeData = append(storeData, v)
				}
			}
		}
	}

	if err == nil {
		return storeData, nil
	} else if err != nil {
		return storeData, err
	} else {
		return storeData, errors.New("Not Found.")
	}
}

func GetAllResources(db *sql.DB, tablename string) ([][]byte, error) {
	storeData := make([][]byte, 0)

	rows, err := db.Query(" SELECT value FROM "+tablename+" where `key`= ? ", key)
	if err != nil {
		return []byte(""), err
	}
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		v := []byte(value)
		storeData = append(storeData, v)
	}

	if err == nil {
		return storeData, nil
	} else if err != nil {
		return storeData, err
	} else {
		return storeData, errors.New("Not Found.")
	}
}

func GetSchemaByBucket(db *bolt.DB, blockName string) ([]byte, error) {
	var codedata []byte

	rows, err := db.Query(" SELECT value FROM Schema where `key`= ? ", blockName)
	if err != nil {
		return []byte(""), err
	}
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			log.Fatal(err)
		}
		v := []byte(value)
		return v, nil
	}

	return []byte(""), err
}
