package dbcontroller

import (
	"dbproxy"
	"fmt"
)

// InitializeMYSQL :  Connects the server to the mysql server
// (decouple main from dbproxy)
func InitializeMYSQL() {
	dbproxy.SetupMysqlConnection()
}

// Serve Basic/Entry function to handle input data.
// Creates a new thread for processing and returns control
// to the listener loop
func Serve(data []byte, ip string, n int) {
	// Store data to new memory
	localData := make([]byte, n)
	n = copy(localData, data[0:n])
	// Multithreaded processing
	go asyncServe(localData, ip)
	// Return control
	return
}

func asyncServe(data []byte, ip string) {
	msg := string(data)
	fmt.Println("msg: ", msg, " ip: ", ip)
	dbproxy.StoreMessageToDB(msg, ip)
}
