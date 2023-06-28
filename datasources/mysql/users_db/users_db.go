package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	MySqlUserName = "mysql_username"
	MySqlPassword = "mysql_password"
	MySqlHost     = "mysql_host"
	MySqlSchema   = "mysql_schema"
)

var (
	Client   *sql.DB
	username = os.Getenv(MySqlUserName)
	password = os.Getenv(MySqlPassword)
	host     = os.Getenv(MySqlHost)
	schema   = os.Getenv(MySqlSchema)
)

func init() {
	datasource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
