package storage

import (
	"testing"
	"time"
	// "fmt"
	"PillarsFlowNet/utility"
)


func TestStoreToPost(t * testing.T) {
	ConnectToMgo()
	message := "test"
	replyTo := "123"
	userCode := "456"
	targetCode := "789"

	postCode := utility.GenerateCode(&userCode)
	post := utility.Post {
		PostCode: *postCode,
    	TargetCode: targetCode,
    	Message: message,
    	ReplyTo: replyTo,
    	UserCode: userCode,
    	PostTime: time.Now().Format("2006-01-02 15:04:05"),
    	Deleted: 0,
    	DeletedTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err := StoreToPost(&post)
	if err != nil {
		panic(err.Error())
	} else {
		// fmt.Println(*(utility.ObjectToJsonString(result)))
	}
	CloseMgoConnection()
	
}

func TestGetAllPostByTargetCode(t * testing.T) {
	ConnectToMgo()
	targetCode := "789"
	_, err := GetAllPostByTargetCode(&targetCode)
	if err != nil {
		panic(err.Error())
	} else {
		// fmt.Println(*(utility.ObjectToJsonString(result)))
	}
	CloseMgoConnection()
}
