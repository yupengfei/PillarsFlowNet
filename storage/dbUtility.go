package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "reflect"
	// "fmt"
	"PillarsFlowNet/utility"
)
var db * sql.DB
var err error
func ConnectToDB() * sql.DB {
	propertyMap := utility.ReadProperty("../DB.properties")
	var userName, password, host, port, database string
	userName =  propertyMap["DBUserName"]
	password = propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	port = propertyMap["DBPort"]
	database = propertyMap["DBDatabase"]

	if db != nil {
		//fmt.Println("db connection has enstablished")
		return db
	}
	sqlString := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	//fmt.Println(sqlString)
	db, err = sql.Open("mysql", sqlString)
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