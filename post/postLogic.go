package post

import (
	"time"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)
// addPost
// {
// 	“command”:”addPost”,
// 	“parameter”:”{
// 		PostCode: 任意string，可以没有
// 		TargetCode string
// 		Message string
// 		ReplyTo string
// 		UserCode 任意string，可以没有
// 		PostTime: 任意string，可以没有
// 		Deleted: 任意int，可以没有
// 		DeletedTime:任意string，可以没有
// 	}”
// }
// 返回给所有人
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “addPost”,
// 	“result”:”{
// 		PostCode: string
// 		TargetCode string
// 		Message string
// 		ReplyTo string
// 		UserCode string
// 		PostTime: string
// 		Deleted: 任意int，可以没有
// 		DeletedTime:任意string，可以没有

// 	}”
// }
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
	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * utility.OutMessage
	if errorCode != 0 {
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addPost",
						Result: "{}",
					}
		out = & tempout
	} else {
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addPost",
						Result: *utility.ObjectToJsonString(postOut) ,
					}
		out = & tempout
	}

	var result = utility.ObjectToJsonByte(*out)
	return result, &(inputParameters[0])
}

// getAllTargetPost 根据targetCode获取所有post
// {
// 	“command”：“getAllTargetPost”，
// 	“parameter”：“{
// 		TargetCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “getAllTargetPost”,
// 	“result”:”[{
// 		PostCode: string
// 		TargetCode string
// 		Message string
// 		ReplyTo string
// 		UserCode string
// 		PostTime: string
// 		Deleted: 任意int，可以没有
// 		DeletedTime:任意string，可以没有

// 	}]”
// }

func GetAllUnreceivedChart(userCodeAndParameter * string) ([] byte, *string) {
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

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllTargetPost",
						Result:*utility.ObjectToJsonString(opResult),
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}