package routes

import (
    "fmt"
    "log"
    "net/http"
    "gotodo/app/config"
    "gotodo/app/controllers/user"
)

type CustomMux struct {
    http.ServeMux
}

var conf = config.Configuration()
var confVerbose = conf.Log.Verbose

func (c CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if confVerbose {
        log.Println("Incoming request from:", r.URL.String(), "method: "+r.Method+"")
    }

    c.ServeMux.ServeHTTP(w, r)
}

func Routes() {
    mux := new(CustomMux)

    mux.HandleFunc("/", user.UserIndex)

    // running message
    fmt.Println("Server running")

    var confPort int = conf.Server.Port

    // init
    var port int = 8080
    if confPort != 0 {
        port = confPort
    }

    fmt.Println(port)

    server := new(http.Server) // create server instance
    server.Addr = fmt.Sprintf(":%d", port) // server port
    server.Handler = mux
    server.ListenAndServe() // launch
}


