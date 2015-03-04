package authentication

import (
	"CopyPillarsFlowNet/mysqlUtility"
	"PillarsFlowNet/utility"
)

//init when start
//read data from mysql and authentication.properties
var AuthMap map[string]string

func init() {
	AuthMap = utility.ReadProperty("./auth.properties")
}

//todo
//this function has not been complished
//通过查询该用户的用户组，然后通过对比配置文件中该用户组的权限all or no
//是不是太简单了，还是对用户权限这一块理解不对？？
func GetAuthInformation(userCode *string) (bool, string) {
	//return true, ""
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT `group`,company_code  FROM user WHERE user_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	var auth, code string
	result := stmt.QueryRow(*userCode)
	result.Scan(&auth, &code)
	if AuthMap[auth] == "all" {
		return true, code
	} else {
		return false, ""
	}

}
