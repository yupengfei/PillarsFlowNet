package login

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"PillarsFlowNet/storage"
	// "fmt"
)

func TestQueryUserCode(t * testing.T) {
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
		t.Error("UserName or password wrong")
	}
	storage.CloseDBConnection()
}

