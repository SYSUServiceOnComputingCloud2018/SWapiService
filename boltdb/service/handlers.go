package service

import (
    "net/http"
    "github.com/unrolled/render"
    "encoding/json"
	"fmt"
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

func homeHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        formatter.HTML(w, http.StatusOK, "index", struct {
            ID      string `json:"id"`
            Content string `json:"content"`
        }{ID: "8675309", Content: "Hello from Go!"})
    }
}

func perosonHandler(formatter *render.Render) http.HandlerFunc{
  
//   j, err := json.MarshalIndent(person, "", "  ")
//   if err == nil{
//     fmt.Printf("%s\n", j)
//   }
	// jsonData, _ := dbOperator.GetSchemaByBucket(db, "Person")
	// var schema Schema //定义在crawler.go中，需要使用go run service.go crawler.go
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
		
		// decoder := json.NewDecoder(r.Body)
		// var param GetUserParam
		// err := decoder.Decode(&param)
		// if err != nil {
		// 	WriteResponse(w, ErrorResponseCode, "request param is invalid, please check!", nil)
		// 	return
		// }

		v, err := dbOperator.GetElementById(db, "Person", "2")
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


// func dump(data interface{}, err error)([]byte) {
// 	if err != nil {
// 		return
// 	}

// 	if j, err := json.MarshalIndent(data, "", "  "); err == nil {
//     return j
// 	}
// }