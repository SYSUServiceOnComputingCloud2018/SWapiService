package main

//负责从API将数据转移到数据库中
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/peterhellberg/swapi"
)

func personDownLoad(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Person
	person, err := c.Person(id)
	if err != nil {
		return err
	} else if person.Name == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	personBuckets, err := tx.CreateBucketIfNotExists([]byte("Person"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(person)
	if err != nil {
		return err
	}

	err = personBuckets.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func filmDownload(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Person
	films, err := c.Film(id)
	if err != nil {
		return err
	} else if films.Title == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	filesBucket, err := tx.CreateBucketIfNotExists([]byte("Film"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(films)
	if err != nil {
		return err
	}

	err = filesBucket.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func starshipsDownload(db *bolt.DB, id int) error {
	c := swapi.DefaultClient
	//获得数据Person
	starships, err := c.Starship(id)
	if err != nil {
		return err
	} else if starships.Name == "" {
		return errors.New("404")
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	starshipsBucket, err := tx.CreateBucketIfNotExists([]byte("Starship"))
	if err != nil {
		return err
	}

	encoded, err := json.Marshal(starships)
	if err != nil {
		return err
	}

	err = starshipsBucket.Put([]byte(strconv.Itoa(id)), encoded)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func main() {

	db, err := bolt.Open("my.db", 0600, nil)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if err := personDownLoad(db, i); err != nil {
			log.Fatal(err)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

}
