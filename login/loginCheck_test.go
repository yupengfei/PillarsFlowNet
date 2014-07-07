package login

import (
	"testing"
	"PillarsFlowNet/storage"
	//"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func TestCheckExist(t * testing.T) {
	db := storage.ConnectToDB()
	userName := "er.wang"
	password := "e10adc3949ba59abbe56e057f20f883e"
	isExist := CheckExist(&userName, &password, db)
	if !isExist {
		t.Error("UserName or password wrong")
	}
	userName = userName + "1"
	isExist = CheckExist(&userName, &password, db)
	if isExist {
		t.Error("UserName or password wrong")
	}
	storage.CloseDBConnection()
}