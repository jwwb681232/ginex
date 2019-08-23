package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"ginex/config"
	"fmt"
)

var DB *sql.DB

func init() {
	c := config.DatabaseConfig()
	DB, _ = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DbUsername, c.DbPassword, c.DbHost, c.DbPort, c.DbDatabase))

}
