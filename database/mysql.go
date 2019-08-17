package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Init() *sql.DB {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gin-blog?parseTime=true")
	return db
}
