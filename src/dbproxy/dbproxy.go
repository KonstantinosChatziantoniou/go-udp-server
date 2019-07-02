// Package dbproxy :
// Responsible for the interactions with the mysql database
package dbproxy

import (
	"database/sql"
	"fmt"
	"log"

	// imported driver for mysql queries https://github.com/go-sql-driver/mysql
	// go get -u github.com/go-sql-driver/mysql
	_ "github.com/go-sql-driver/mysql"
)

var (
	// global variable for the database connection reference
	db *sql.DB
)

// SetupMysqlConnection :  Connects the server to the mysql server
func SetupMysqlConnection() {

	fmt.Println("Connecting to MySQL")
	var err error
	db, err = sql.Open("mysql", "gotest:9666@/embiots")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Succesfuly connected to MySQL")
	}

}

// CloseDBConnection : Terminates the connection with the database
func CloseDBConnection() {
	db.Close()
}

// StoreMessageToDB Stores the ip and msg to mysql. ip and msg have to be strings
func StoreMessageToDB(msg string, ip string) {
	query := fmt.Sprintf(`insert into test (ip, msg) values("%s","%s")`, msg, ip)
	insert, err := db.Query(query)
	defer insert.Close()

	if err != nil {
		log.Println(err.Error())
	}

}
