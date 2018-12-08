package main

import (
	"os"
	"github.com/hansenbeast/boltdb/service"
	flag "github.com/spf13/pflag"
	// "encoding/json"
	// "fmt"
	// "github.com/hansenbeast/boltdb/dbOperator"
	// "github.com/boltdb/bolt"
)
const (
    PORT string = "8080"
)

func main() {
	// db, err := bolt.Open("my.db", 0600, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()
	// Schema输出
	// jsonData, _ := dbOperator.GetSchemaByBucket(db, "Person")
	// var schema Schema //定义在crawler.go中，需要使用go run service.go crawler.go
	// err = json.Unmarshal(jsonData, &schema)
	// if err == nil {
	// 	fmt.Println(schema)
	// } else {
	// 	fmt.Println(err)
	// }
	port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }

    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }

    server := service.NewServer()
    server.Run(":" + port)
	//查找单个元素
	// v, err := dbOperator.GetElementById(db, "Person", "2")

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	var user swapi.Person
	// 	err = json.Unmarshal(v, &user)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	} else {
	// 		fmt.Println(user.FilmURLs)
	// 	}
	// }
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
