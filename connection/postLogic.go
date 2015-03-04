package connection

import (
	"PillarsFlowNet/postStorage"
	"PillarsFlowNet/utility"
	"time"
)

func AddPost(userCode *string, parameter *string) {
	var errorCode int
	var postOut *utility.Post
	var err error
	//var toUserCode * string
	if errorCode == 0 {
		post, _ := utility.ParsePostMessage(parameter)
		//toUserCode = &(chart.To)

		post.Id = *(utility.GenerateCode(userCode))
		post.UserCode = *userCode

		post.Deleted = 0
		post.DeletedTime = time.Now().Format("2006-01-02 15:04:05")
		postOut, err = postStorage.StoreToPost(post)
		if err != nil {
			errorCode = 1
		}
	}
	var command = "addPost"
	result := utility.BoolResultToOutMessage(&command, postOut, errorCode, userCode)
	//Hub.Dispatch(result)
	Hub.SendToUserCode(result, userCode)
}

func GetAllTargetPost(userCode *string, parameter *string) {
	var errorCode int
	var opResult []utility.Post
	var err error
	if errorCode == 0 {
		targetCode, _ := utility.ParseTargetCodeMessage(parameter)
		opResult, err = postStorage.GetAllPostByTargetCode(&(targetCode.TargetCode))
		if err != nil {
			errorCode = 1
		}
	}
	var command = "getAllTargetPost"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
