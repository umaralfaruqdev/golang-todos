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
	IsMarried string `json:"is_married"`
}

func UserStore(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// initial varaiables
	var errConv string

	// the function for write response
	// @parameter status_code, message
	var resFunc = func(str ...string) (resByte []byte) {
		var resStr string = `{"status": `+str[0]+`, "message": "`+str[1]+`"}`
		resByte = []byte(resStr)
		
		return
	}
	
	var data UserStuct // use struct

	var decoder *json.Decoder = json.NewDecoder(r.Body)
	var err error = decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(500)
		errConv = fmt.Sprintf("%s", err)
		w.Write(resFunc("500", errConv))
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
	} else if isMaried == "false" {
		isMaried = false
	} else if isMaried == "true" {
		isMaried = true
	}

	// insert to database
	err = userModel.UserStore(fname, lname, born, isMaried)

	if !Msg(err) {
		w.WriteHeader(http.StatusForbidden)
		errConv = fmt.Sprintf("%s", err)
		w.Write(resFunc("403", errConv))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resFunc("200", "Success"))
}
