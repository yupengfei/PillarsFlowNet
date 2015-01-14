package post

import (
	"time"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func AddPost(userCode * string, parameter * string, h * connection.HubStruct) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var postOut * utility.Post
	var err error
	//var toUserCode * string
	if (errorCode == 0) {
		post, _ := utility.ParsePostMessage(parameter)
		//toUserCode = &(chart.To)

		post.Id = *(utility.GenerateCode(userCode)
		post.UserCode = inputParameters[0]
		
		post.Deleted = 0
		post.DeletedTime = time.Now().Format("2006-01-02 15:04:05")
		postOut, err = storage.StoreToPost(post)
		if err != nil {
			errorCode = 1
		}
	}
	var command = "addPost"
	result := utility.BoolResultToOutMessage(&command, postOut, errorCode, userCode)
	h.Dispatch(result)
}

func GetAllTargetPost(userCode * string, parameter * string, h * connection.HubStruct) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Post
	var err error
	if (errorCode == 0) {
		targetCode, _ := utility.ParseTargetCodeMessage(parameter)
		opResult, err = storage.GetAllPostByTargetCode(&(targetCode.TargetCode))
		if err != nil {
			errorCode = 1
		}
	}
	var command = "getAllTargetPost"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}