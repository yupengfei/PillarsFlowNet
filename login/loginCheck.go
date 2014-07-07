package login

import "database/sql"

func CheckExist(userName * string, password * string, db * sql.DB) bool {
	stmt, err := db.Prepare("SELECT * FROM user WHERE user_name=? AND password=?")
	if err != nil {
		panic(err.Error())
	}
	result, err := stmt.Query(userName, password)
	if err != nil {
		panic(err.Error())
	}
	if result.Next() {
		return true
	} else {
		return false
	}
}

