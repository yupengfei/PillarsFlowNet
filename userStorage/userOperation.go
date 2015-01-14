package userStorage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/mysqlUtility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)

func InsertIntoUser(user * utility.User) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO user(user_code, user_name, password, `group`, display_name, position, picture, email, phone) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		// fmt.Print(err.Error())
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.UserCode, user.UserName, user.Password,
		user.Group, user.DisplayName, user.Position, user.Picture, user.Email,
		user.Phone)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		panic(err.Error())
	}
	return true, err
}

func DeleteUserByUserName(userName * string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("DELETE FROM user WHERE user_name = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(userName)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

func QueryUserByUserName(userName * string) (* utility.User, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT user_code, user_name, `group`, display_name, position, picture, email, phone, insert_datetime, update_datetime FROM user WHERE user_name=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(userName)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user utility.User
	if result.Next() {
		err = result.Scan(&(user.UserCode), &(user.UserName),
		&(user.Group), &(user.DisplayName), &(user.Position), &(user.Picture), &(user.Email),
		&(user.Phone), &(user.InsertDatetime), &(user.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &user, err

}

func QueryUserByUserCode(userCode * string) (* utility.User, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT user_code, user_name, `group`, display_name, position, picture, email, phone, insert_datetime, update_datetime FROM user WHERE user_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(userCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var user utility.User
	if result.Next() {
		err = result.Scan(&(user.UserCode), &(user.UserName),
		&(user.Group), &(user.DisplayName), &(user.Position), &(user.Picture), &(user.Email),
		&(user.Phone), &(user.InsertDatetime), &(user.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &user, err

}

func QueryUserCode(userName * string) (* string, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT user_code FROM user WHERE user_name=?")
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
	return &user_code, err
}

func QueryAllUser() ([] utility.User, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT user_code, user_name, `group`, display_name, position, picture, email, phone, insert_datetime, update_datetime FROM user")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var users [] utility.User
	for result.Next() {
		var user utility.User
		err = result.Scan(&(user.UserCode), &(user.UserName),
		&(user.Group), &(user.DisplayName), &(user.Position), &(user.Picture), &(user.Email),
		&(user.Phone), &(user.InsertDatetime), &(user.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		users = append(users, user)
	}
	return users, err

}

func CheckUserNameAndPassword(userName * string, password * string) (* string, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT user_code FROM user WHERE user_name=? AND password=?")
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
	var userCode string
	if result.Next() {
		result.Scan(&userCode)
		return &userCode, err
		
	} 
	return &userCode, err
}
