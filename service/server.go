package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/codegangsta/negroni"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	// 默认html文件在assets中
	// formatter := render.New(render.Options{
	//     IndentJSON: true,
	// })

	// 指定了html模板的目录，模板文件的扩展名
	formatter := render.New(render.Options{
		// Directory:  "src/github.com/hansenbeast/boltdb/templates",
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			// root为$GOPATH
			// webRoot = root + "/src/github.com/hansenbeast/boltdb"
			webRoot = root
			fmt.Println(webRoot)
		}
	}

	// 打开boltdb
	db, err := sql.Open("mysql", "root:123456@tcp(0.0.0.0:3306)/test")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		fmt.Println("opon database fail")
		panic(err)
	}

	mx.HandleFunc("/api/", rootHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/people/", peopleHandler(formatter, db)).Methods("GET")
	mx.HandleFunc("/api/people/schema", peopleSchemaHandler(formatter, db)).Methods("GET")
	mx.HandleFunc("/api/people/{id:[0-9]+}/", peopleIdHandler(formatter, db)).Methods("GET")
	mx.HandleFunc("/api/planets/{id:[0-9]+}/", planetsHandler(formatter, db)).Methods("GET")
	mx.HandleFunc("/api/films/{id:[0-9]+}/", filmsHandler(formatter, db)).Methods("GET")
	mx.HandleFunc("/api/species/{id:[0-9]+}/", speciesHandler(formatter, db)).Methods("GET")
	mx.HandleFunc("/api/vehicles/{id:[0-9]+}/", vehiclesHandler(formatter, db)).Methods("GET")
	mx.HandleFunc("/api/starships/{id:[0-9]+}/", starshipsHandler(formatter, db)).Methods("GET")

}
