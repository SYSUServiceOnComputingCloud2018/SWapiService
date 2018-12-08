package service

import (
    "net/http"
    "github.com/unrolled/render"
    "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hansenbeast/boltdb/dbOperator"
	"github.com/peterhellberg/swapi"
	"github.com/boltdb/bolt"

)
const (
	ErrorResponseCode   = 1000 // 错误响应code
	SuccessResponseCode = 0    // 正确响应code
)

type ResponseMessage struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func rootHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
		// 
    }
}


func peopleHandler(formatter *render.Render) http.HandlerFunc{
	
	// 输出schema
	// jsonData, _ := dbOperator.GetSchemaByBucket(db, "Person")
	// var schema Schema //定义在crawler.go中
	// err = json.Unmarshal(jsonData, &schema)
	// if err == nil {
	// 	fmt.Println(schema)
	// } else {
	// 	fmt.Println(err)
	// }
 
  	return func(w http.ResponseWriter, req *http.Request) {

		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		// fmt.Println("URL", req.URL, "HOST", req.Host, "Method", req.Method, "RequestURL", req.RequestURI, "RawQuery", req.URL.RawQuery)

		// 获取id
		vars := mux.Vars(req)
		id := vars["id"]

		// 从db中获得Person Struct
		v, err := dbOperator.GetElementById(db, "Person", id)
		if err != nil {
			fmt.Println(err)
			WriteResponse(w, ErrorResponseCode, "failed", nil)
		} else {
			var user swapi.Person
			err = json.Unmarshal(v, &user)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(user.Name)
				WriteResponse(w, SuccessResponseCode, "success", user)
			}
		}
	}
}

func planetsHandler(formatter *render.Render) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {
		//
	}
}

func filmsHandler(formatter *render.Render) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {
		//
	}
}

func speciesHandler(formatter *render.Render) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {
		//
	}
}

func vehiclesHandler(formatter *render.Render) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {
		//
	}
}

func starshipsHandler(formatter *render.Render) http.HandlerFunc{
  
	return func(w http.ResponseWriter, req *http.Request) {
		//
	}
}

func WriteResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// resp := ResponseMessage{Code: code, Message: message, Data: data}
	b, err := json.Marshal(data)
	if err != nil {
		// logrus.Warnf("error when marshal response message, error:%v\n", err)
		fmt.Println(err)
	}
	w.Write(b)
}
