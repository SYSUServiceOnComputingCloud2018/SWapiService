package service

import (
    "os"
    "fmt"
    "github.com/codegangsta/negroni"
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
        Directory: "templates",
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

    mx.HandleFunc("/api/",rootHandler(formatter)).Methods("GET")
    mx.HandleFunc("/api/people/{id:[0-9]+}",peopleHandler(formatter)).Methods("GET")
    mx.HandleFunc("/api/planets/{id:[0-9]+}",planetsHandler(formatter)).Methods("GET")
    mx.HandleFunc("/api/films/{id:[0-9]+}",filmsHandler(formatter)).Methods("GET")
    mx.HandleFunc("/api/species/{id:[0-9]+}",speciesHandler(formatter)).Methods("GET")
    mx.HandleFunc("/api/vehicles/{id:[0-9]+}",vehiclesHandler(formatter)).Methods("GET")
    mx.HandleFunc("/api/starships/{id:[0-9]+}",starshipsHandler(formatter)).Methods("GET")
    
}