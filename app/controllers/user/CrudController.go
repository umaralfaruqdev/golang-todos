package user

import (
	"fmt"
	"net/http"
	userModel "gotodo/app/models/user"
	. "gotodo/system/error"
	"encoding/json"
)

func init() {
	fmt.Println("Hello from user branch")
}

type UserStuct struct {
	Fname 	  string `json:"fname"`
	Lname 	  string `json:"lname"`
	Born 	  string `json:"born"`
	IsMarried bool   `json:"is_married"`
}

func UserStore(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	
	var data UserStuct // use struct
	var jsonString string = `{"status": 200, "message": "your data successfuly inserted"}`
	var jsonByte []byte = []byte(jsonString)

	var decoder *json.Decoder = json.NewDecoder(r.Body)
	var err error = decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var fname string = data.Fname
	var lname string = data.Lname
	var born string = data.Born
	var isMaried bool = data.IsMarried

	// insert to database
	err = userModel.UserStore(fname, lname, born, isMaried)
	if !Msg(err) {return }

	w.Write(jsonByte)
}
