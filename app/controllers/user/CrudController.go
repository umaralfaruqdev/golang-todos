package user

import (
	"fmt"
	"net/http"
	userModel "gotodo/app/models/user"
	. "gotodo/system/error"
	"encoding/json"
	"reflect"
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

	var fname interface{} = data.Fname
	var lname string = data.Lname
	var born interface{} = data.Born
	var isMaried interface{} = data.IsMarried

	if fname == "" {
		fname = nil
	}

	if born == "" {
		born = nil
	}

	if isMaried == "" {
		isMaried = nil
	}

	// insert to database
	err = userModel.UserStore(fname, lname, born, isMaried)
	var jsonStringErr = fmt.Sprintf(`{"status": %d, "message": "%v"}`, 403, err)
	var jsonStringErrByte = []byte(jsonStringErr)

	if !Msg(err) {
		w.Write(jsonStringErrByte)
		// http.Error(w, err.Error(), 403)
		return
	}

	w.Write(jsonByte)
}
