package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    . "gotodo/system/error"
)

func Conn() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:godev@tcp(localhost:3306)/godb")
    if !Msg(err) { return nil, err }
    return db, nil
}
