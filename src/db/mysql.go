package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	DB, _ = sqlx.Open("mysql", "root:root123@tcp(114.115.250.201:3306)/xxx")
	err := DB.Ping()
	if err != nil {
		println(err)
	}
}
