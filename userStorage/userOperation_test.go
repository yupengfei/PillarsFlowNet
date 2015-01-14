package userStorage

import (
	"testing"
	"PillarsFlowNet/utility"
	// _ "github.com/go-sql-driver/mysql"
	"fmt"
)

// func TestQueryUserCode(t * testing.T) {
// 	DBConn = ConnectToDB()
// 	userName := "kun.wang"
// 	user_code, _ := QueryUserCode(&userName)
// 	// if !isExist {
// 	// 	t.Error("UserName or password wrong")
// 	// }
// 	if *user_code == "" {
// 		t.Error("UserName or password wrong")
// 	}
// 	// fmt.Println(*user_code)
// 	userName = userName + "1"
// 	user_code, _ = QueryUserCode(&userName)
// 	if *user_code != "" {
// 		t.Error("UserName or password wrong 2")
// 	}
// 	CloseDBConnection()
// }

func TestInsertIntoUser1(t * testing.T) {
	var user utility.User

	user.UserName = "er.wang"
	user.UserCode = *(utility.GenerateCode(&(user.UserName)))
	pass := "aaa"
	user.Password =  string(utility.Md5sum(&pass))
	user.Group = "fjkdjflk"
	user.DisplayName = "fdafae"

	result, _ := InsertIntoUser(&user)
	if result != true {
		t.Error("insert failed")
	}
}

func TestQueryByUserName(t * testing.T) {
	userName := "er.wang"
	user, _ := QueryUserByUserName(&userName)
	fmt.Println("testing query user by name " + user.UserName)
}

func TestDeleteUserByUserName(t * testing.T) {
	userName := string("er.wang")
	result, err := DeleteUserByUserName(&userName)
	if result == false || err != nil {
		t.Error("test delete user failed")
	}
}

func TestQueryAllUser(t * testing.T) {
	result, err := QueryAllUser()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(*(utility.ObjectToJsonString(result)))
	}
}


