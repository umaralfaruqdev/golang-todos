package user

import (
    "database/sql"
    . "gotodo/app/config"
    . "gotodo/system/error"
)

func UserGetAll() (*sql.Rows, error) {
    db, err := Conn()
    if !Msg(err) { return nil, err }

    rows, err := db.Query("select id, fname, lname, born, is_married from users")
    if !Msg(err) { return nil, err }

    return rows, nil
}

