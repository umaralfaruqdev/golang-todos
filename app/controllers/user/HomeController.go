package user

import (
    "net/http"
    "os"
    "html/template"
    . "gotodo/system/error"
    "gotodo/app/models/user"
)

type UserStruct struct {
    Id        int
    Fname     string
    Lname     string
    Born      string
    IsMarried bool
}

func UserIndex(w http.ResponseWriter, r *http.Request) {

    rows, err := user.UserGetAll() // get all users
    if !Msg(err) { return }

    var result []UserStruct // init struct array

    for rows.Next() {
        var each = UserStruct{} // init struct
        var err = rows.Scan(&each.Id, &each.Fname, &each.Lname, &each.Born, &each.IsMarried) // Scanning
        if !Msg(err) { return }
        
        result = append(result, each) // append to slice sturct "UserStruct"
    }

    // map data for passing to view
    // passing result struct to map
    var data = map[string]interface{}{
        "name": result,
    }

    // get work dir
    var dir, _ = os.Getwd()

    // parse template
    tmpl, err := template.ParseFiles(""+dir+"/app/views/user/index.html")
    if !Msg(err) { return }
    
    // execute template
    err = tmpl.ExecuteTemplate(w, "index", data)
    if !Msg(err) { return }

} 













































