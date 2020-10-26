package user

import (
    "fmt"
    "net/http"
    "os"
    "html/template"
    . "gotodo/system/error"
    "gotodo/app/models/user"
    "reflect"
    "regexp"
    "strconv"
    "time"
)

type UserStruct struct {
    Id        int
    Fname     string
    Lname     string
    Born      string
    IsMarried bool
}

func (u UserStruct) UserIndexing(num int) int {
    return num + 1
}


var funcMap = template.FuncMap{
    "unescape": func (s string) template.HTML {
        return template.HTML(s)
    },
    "getByIndex": func (s interface{}) UserStruct {
        fmt.Printf("%+v\n", s)
        fmt.Println(s.(UserStruct).Fname)
        fmt.Println(reflect.ValueOf(s).Kind())
        return s.(UserStruct)
    },
    "parseDate": func(s string) string {
        var regex, _ = regexp.Compile(`-`)
        var born = regex.Split(s, -1)

        var year = born[0]
        var month = born[1]
        var day = born[2]
        
        // convert month to integer
        // @type int
        var monthInt, err = strconv.Atoi(month)
        if !Msg(err) { return "" }

        // array month
        var monthArr = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
        var monthArrLen = len(monthArr)
        var monthStr string

        for i := 1; i <= monthArrLen; i++ {
            if i == monthInt {
                monthStr = monthArr[i-1]
                break
            }
        }

        monthStr = fmt.Sprintf("%s %s, %s", day, monthStr, year)

        return monthStr
    },
    "getYearOld": func (s string) string {
        var regex, _ = regexp.Compile(`-`)
        var born = regex.Split(s, -1)

        var year, err = strconv.Atoi(born[0])
        if !Msg(err) { return "" }

        var now = time.Now()
        yearNow := now.Year()

        age := strconv.Itoa(yearNow - year)

        return age
    },
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
        "users": result,
    }

    // get work dir
    var dir, _ = os.Getwd()

    // parse template
    tmpl, err := template.New("index").Funcs(funcMap).ParseFiles(""+dir+"/app/views/user/index.html")
    if !Msg(err) { return }
    
    // execute template
    err = tmpl.Execute(w, data)
    if !Msg(err) { return }

} 













































