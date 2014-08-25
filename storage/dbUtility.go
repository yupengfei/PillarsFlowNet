package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"PillarsFlowNet/pillarsLog"
	"PillarsFlowNet/utility"
)
//begin with capitial word so it can be accessed by outer
var DBConn * sql.DB

func init() {
	DBConn = ConnectToDB()
}
func ConnectToDB() * sql.DB {
	//connection already exist
	if DBConn != nil {
		return DBConn
	}
	//connection not exist
	propertyMap := utility.ReadProperty("./DB.properties")
	var userName, password, host, port, database string
	userName =  propertyMap["DBUserName"]
	password = propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	port = propertyMap["DBPort"]
	database = propertyMap["DBDatabase"]

	sqlString := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	DBConn, err := sql.Open("mysql", sqlString)
	if err != nil {
		pillarsLog.Logger.Panic("can not connect to mysql server")
	}
	return DBConn
}
func CloseDBConnection() {

	DBConn.Close()
	DBConn = nil
}