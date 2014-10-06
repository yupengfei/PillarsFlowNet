package user

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/pillarsLog"
	"strings"
	"PillarsFlowNet/authentication"
)

//result, userName, err
func ValidateUser(parameter * string) (* string,  * string,  error) {
	var result string
	var userName string
	var err error
	user, err := utility.ParseLoginInMessage(parameter)
	if err != nil {
		pillarsLog.Logger.Println("parse login message error")
		return &result, &userName, err
	}
	userCode, _ := storage.CheckUserNameAndPassword(&((*user).UserName), &((*user).Password))
	var sysError = utility.Error {
			ErrorCode: 0,
			ErrorMessage: "",
	}
	var loginMessage utility.LoginInMessage
	if *userCode != "" {
		loginMessage = utility.LoginInMessage {
			Auth: "success",
			AuthMessage : "",
		}
	} else {
		loginMessage = utility.LoginInMessage {
			Auth: "failed",
			AuthMessage : "userName or Password wrong",
		}
	}
	loginStr := string(utility.ObjectToJsonByte(loginMessage))	

	var out = utility.OutMessage {
			Error: sysError,
			Command: "login",
			Result: loginStr,
		}

	result = string(utility.ObjectToJsonByte(out))
	return &result, userCode, err
}

// getAllUser 获取所有的用户列表
// {
// 	“command”:”getAllUser”,
// 	”parameter“：”{
// 	}“
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “getAllUser”,
// 	“result”:”[{
// 		UserCode string
//     		UserName string
//     		Password string，返回只里面该字段为空
//     		Group string
//     		DisplayName string
//     		Position string
//     		Picture string
//     		Email string
//     		Phone string
//     		InsertDatetime string
//     		UpdateDatetime string
// 	}]”
// }

func GetAllUser(userCodeAndParameter * string) ([] byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * utility.OutMessage
	if errorCode != 0 {
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "getAllUser",
						Result: "{}",
					}
		out = & tempout
	} else {
		userSlice, _ := storage.QueryAllUser()
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "getAllUser",
						Result: *utility.ObjectToJsonString(userSlice) ,
					}
		out = & tempout
	}

	var result = utility.ObjectToJsonByte(*out)

	return result, &(inputParameters[0])
}