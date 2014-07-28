package login

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/storage"
)

func QueryUserCode(userName * string) * string {
	stmt, err := storage.DBConn.Prepare("SELECT user_code FROM user WHERE user_name=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(userName)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user_code string

	if result.Next() {
		result.Scan(&user_code)
		
	} 
	return &user_code
}

func CheckUserNameAndPassword(userName * string, password * string) bool {
	stmt, err := storage.DBConn.Prepare("SELECT user_code FROM user WHERE user_name=? AND password=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	passwordMd5 := utility.Md5sum(password)
	result, err := stmt.Query(userName, passwordMd5)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	if result.Next() {
		return true
		
	} 
	return false
}

