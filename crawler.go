package main

//负责从API将数据转移到数据库中
import (
	"log"
	"reflect"

	"github.com/boltdb/bolt"
	"github.com/peterhellberg/swapi"
)

func personDownLoad(db *bolt.DB, c swapi.DefaultClient, id int) error {
	//获得数据Person
	person, err := c.Person(id)
	if err != nil {
		return err
	}
	//开始事物
	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	//拿到存储区buckets
	personBuckets, err := tx.CreateBucketIfNotExists([]byte("Person"))
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			if value, ok := v.Field(i).Interface().(string); ok {
				key := t.Field(i).Name
				personBuckets.Put([]byte(key), []byte(value))
			}
		}
	}
	filmURLsBucket, err := personBuckets.CreateBucketIfNotExists([]byte("FilmURLs"))
	if err != nil {
		return err
	}
	for i:=0 ; i < len(person.FilmURLs);i++{
		key
	}
	speciesURLsBucket, err := personBuckets.CreateBucketIfNotExists([]byte("SpeciesURLs"))
	if err != nil {
		return err
	}
	vehicleURLsBucket, err := personBuckets.CreateBucketIfNotExists([]byte("VehicleURLs"))
	if err != nil {
		return err
	}
	starshipURLsBucket, err := personBuckets.CreateBucketIfNotExists([]byte("StarshipURLs"))
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func main() {
	c := swapi.DefaultClient
	db, err := bolt.Open("my.db", 0600, nil)
	defer db.close()
	if err != nil {
		log.Fatal(err)
	}
	if err := personDownLoad(1); err != nil {
		log.Fatal(err)
	}
}
