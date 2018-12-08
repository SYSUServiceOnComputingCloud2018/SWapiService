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
    //使用模版
    mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
    // //处理简单js访问
    // mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
    // //处理静态路径前缀
    // mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
    // //显示静态文件
    // // mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
    // //提交表单，并输出一个表格
    // mx.HandleFunc("/login", loginHandler(formatter))
    // //对 /unknown 给出开发中的提示，返回码 5xx
    // mx.HandleFunc("/unknown", NotImplementedHandler(formatter)).Methods("GET")

    mx.HandleFunc("/api/person",perosonHandler(formatter)).Methods("GET")
    
}