package user

import (
    // "fmt"
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

/**
 * insert user data to database
 * @return mix
 */

// func UserStore(fname interface{}, lname interface{}, born interface{}, isMarreid interface{}) error {
func UserStore(str ...interface{}) error {
    db, err := Conn()
    if err != nil {
        return err
    }

    stmt, err := db.Prepare("insert into users (fname, lname, born, is_married) values (?, ?, ?, ?)")
    if err != nil {
        return err
    }

    _, err = stmt.Exec(str[0], str[1], str[2], str[3])
    if err != nil {
        return err
    }

    return nil

}

