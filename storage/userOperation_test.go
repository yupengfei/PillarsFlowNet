package storage

import (
	"testing"
	// "PillarsFlowNet/utility"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func TestQueryUserCode(t * testing.T) {
	DBConn = ConnectToDB()
	userName := "er.wang"
	user_code := QueryUserCode(&userName)
	// if !isExist {
	// 	t.Error("UserName or password wrong")
	// }
	if *user_code == "" {
		t.Error("UserName or password wrong")
	}
	// fmt.Println(*user_code)
	userName = userName + "1"
	user_code = QueryUserCode(&userName)
	if *user_code != "" {
		t.Error("UserName or password wrong 2")
	}
	CloseDBConnection()
}

// func TestInsertIntoUser1(t * testing.T) {
// 	DBConn = ConnectToDB()
// 	var user utility.User

// 	user.UserName = "yupengfei"
// 	user.UserCode = *(utility.GenerateCode(&(user.UserName)))
// 	pass := "123456"
// 	user.Password =  *(utility.GenerateCode(&(pass)))
// 	user.Group = "fjkdjflk"
// 	user.DisplayName = "fdafae"

// 	result := InsertIntoUser(&user)
// 	if result != true {
// 		t.Error("insert failed")
// 	}
// 	CloseDBConnection()
// }

func TestQueryByUserName(t * testing.T) {
	DBConn = ConnectToDB()

	userName := "yupengfei"

	user := QueryUserByUserName(&userName)
	fmt.Println("testing query user by name " + user.UserName)
	CloseDBConnection()
}


