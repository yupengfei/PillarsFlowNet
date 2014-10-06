package post

import (
	"time"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func AddPost(userCodeAndParameter * string) ([] byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var postOut * utility.Post
	var err error
	//var toUserCode * string
	if (errorCode == 0) {
		post, _ := utility.ParsePostMessage(&(inputParameters[1]))
		//toUserCode = &(chart.To)

		post.PostCode = *(utility.GenerateCode(&inputParameters[0]))
		post.UserCode = inputParameters[0]
		
		post.Deleted = 0
		post.DeletedTime = time.Now().Format("2006-01-02 15:04:05")
		postOut, err = storage.StoreToPost(post)
		if err != nil {
			errorCode = 1
		}
	}
	var command = "addPost"
	result := utility.BoolResultToOutMessage(&command, postOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

func GetAllTargetPost(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Post
	var err error
	if (errorCode == 0) {
		targetCode, _ := utility.ParseTargetCodeMessage(&(inputParameters[1]))
		opResult, err = storage.GetAllPostByTargetCode(&(targetCode.TargetCode))
		if err != nil {
			errorCode = 1
		}
	}
	var command = "getAllTargetPost"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}