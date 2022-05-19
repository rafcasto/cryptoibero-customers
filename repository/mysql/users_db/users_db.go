package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client   *sql.DB
	username = "cryptoiberodb"
	password = "Crypt01b3r02021"
	host     = "172.105.162.187:3306"
	schema   = "cryptoibero"
)

func init() {
	// cryptoiberodb:Crypt01b3r02021@tcp(172.105.162.187)/cryptoibero
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	println(dataSourceName)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("successful datbase connection...")

}
