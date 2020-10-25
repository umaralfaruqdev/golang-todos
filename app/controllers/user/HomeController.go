package user

import (
    "net/http"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello world"))
}  
