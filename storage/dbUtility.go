package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "reflect"
	// "fmt"
)
var db * sql.DB
var err error
func ConnectToDB(userName string, password string, host string, port string, database string) * sql.DB {
	if db != nil {
		//fmt.Println("db connection has enstablished")
		return db
	}
	db, err = sql.Open("mysql", userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database)
	//fmt.Println(reflect.TypeOf(db))
	//fmt.Println(reflect.TypeOf(err))
	if err != nil {
		panic(err.Error())
	}
	return db
}
func CloseDBConnection() {
	defer db.Close()
}