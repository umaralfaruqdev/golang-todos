package config

import (
    "os"
    "io/ioutil"
    "log"
    "path/filepath"
    "encoding/json"
    . "gotodo/system/error"
)

type _Configuration struct {
    Server struct {
        Port int `json:"port"`
    }
    Log struct {
        Verbose bool `json:"verbose"`
    }
}

var shared *_Configuration

func Configuration() _Configuration {
    if shared != nil {
        return *shared
    }
    
    return _Configuration{}
}

func init() {

    // get workdir
    var dir, err = os.Getwd()
    if !Msg(err) { return } // chekcing error

    // define file location
    var fileLocation = filepath.Join(dir, "app", "config", "conf.json")

    // checking file
    _, err = os.Stat(fileLocation)
    
    if os.IsNotExist(err) {
        log.Println(err)
        return
    }
    
    fileStr, err := ioutil.ReadFile(fileLocation)
    if !Msg(err) { return }


    shared = new(_Configuration)
    err = json.Unmarshal(fileStr, &shared)

    if !Msg(err) { return }

}




























