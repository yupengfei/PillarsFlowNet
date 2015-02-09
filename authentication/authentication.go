package authentication

import (
	"CopyPillarsFlowNet/mysqlUtility"
	//"fmt"
)

//init when start
//read data from mysql and authentication.properties
var AuthMap map[*string]string

func init() {

}

//todo
//this function has not been complished
func GetAuthInformation(userCode *string) bool {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT group FROM user WHERE user_code=? `)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	//stmt.QueryRow(*userCode)
	return true
}
