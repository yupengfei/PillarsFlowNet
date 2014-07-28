package login

import (
	"testing"
	"PillarsFlowNet/storage"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
)

func TestCheckExist(t * testing.T) {
	// db := storage.ConnectToDB()
	userName := "er.wang"
	password := "123456"
	user_code := QueryUserCode(&userName, &password, storage.DBConn)
	// if !isExist {
	// 	t.Error("UserName or password wrong")
	// }
	if *user_code == "" {
		t.Error("UserName or password wrong")
	}
	// fmt.Println(*user_code)
	userName = userName + "1"
	user_code = QueryUserCode(&userName, &password, storage.DBConn)
	if *user_code != "" {
		t.Error("UserName or password wrong")
	}
	storage.CloseDBConnection()
}