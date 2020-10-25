package user

import (
    "net/http"
    "os"
    "html/template"
    . "gotodo/system/error"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
    var data = struct {
        Name string
    } {
        Name: "Umar",
    }

    var dir, _ = os.Getwd()

    var tmpl, err = template.ParseFiles(""+dir+"/app/views/user/index.html")
    if !Msg(err) { return }

    err = tmpl.Execute(w, data)
    if !Msg(err) { return }

}  
