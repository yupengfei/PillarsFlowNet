package postStorage

import (
	"testing"
	"time"
	// "fmt"
	"PillarsFlowNet/utility"
)


func TestStoreToPost(t * testing.T) {
	message := "test"
	replyTo := "123"
	userCode := "456"
	targetCode := "789"

	postCode := utility.GenerateCode(&userCode)
	// PostType：int#0mission自身的消息，1missionDaily的消息，2target的消息
	// Code：post关联到某个daily或者mission自身或者target
	// IsPicture int #0不是图片，1是图片
	post := utility.Post {
		Id: *postCode,
    	Code: targetCode,
    	PostType: 0,
    	IsPicture:0,
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
}

func TestGetAllPostByTargetCode(t * testing.T) {
	targetCode := "789"
	_, err := GetAllPostByTargetCode(&targetCode)
	if err != nil {
		panic(err.Error())
	} else {
		// fmt.Println(*(utility.ObjectToJsonString(result)))
	}
}
