package login

import "database/sql"
import "PillarsFlowNet/utility"

func CheckExist(userName * string, password * string, db * sql.DB) bool {
	stmt, err := db.Prepare("SELECT * FROM user WHERE user_name=? AND password=?")
	if err != nil {
		panic(err.Error())
	}
	passwordMd5 := utility.Md5sum(*password)
	result, err := stmt.Query(userName, passwordMd5)
	if err != nil {
		panic(err.Error())
	}
	if result.Next() {
		return true
	} else {
		return false
	}
}

