package user

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/storage"
	"strings"
	"PillarsFlowNet/authentication"
)

//result, userName, err
func ValidateUser(parameter * string) (* string,  * string,  error) {
	var err error
	user, err := utility.ParseLoginInMessage(parameter)
	if err != nil {
		//TO DO need to made more rebust
		panic("wrong login message")
	}
	userCode, _ := storage.CheckUserNameAndPassword(&((*user).UserName), &((*user).Password))
	var sysError = utility.Error {
			ErrorCode: 0,
			ErrorMessage: "",
	}
	var resultString * string
	if *userCode == "" {
		tmpResultString := ""
		resultString = &tmpResultString
	} else {
		user, _ := storage.QueryUserByUserCode(userCode)
		resultString = utility.ObjectToJsonString(user)
	}

	var out = utility.OutMessage {
			Error: sysError,
			Command: "login",
			UserCode: *userCode,
			Result: *resultString,
		}

	result := utility.ObjectToJsonString(&out)
	return result, userCode, err
}

func GetAllUser(userCode * string, parameter * string) ([] byte, *string) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var userSlice [] utility.User
	if errorCode == 0 {
		userSlice, _ = storage.QueryAllUser()
	}
	command := "getAllUser"
	result := utility.SliceResultToOutMessage(&command, userSlice, errorCode, userCode)
	return result, userCode
}