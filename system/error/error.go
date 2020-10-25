package error

import (
    "log"
)

func Msg(err error) bool {
    // checking error
    if err != nil {
        log.Println(err)
        return false
    }

    // if not err
    return true
}

