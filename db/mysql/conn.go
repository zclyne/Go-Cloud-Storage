package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db,sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/fileserver?charset=utf8")
}