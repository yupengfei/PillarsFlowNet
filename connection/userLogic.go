package connection

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/userStorage"
)

//result, userName, err
func ValidateUser(parameter * string) (* string,  * string,  error) {
	var err error
	user, err := utility.ParseLoginInMessage(parameter)
	if err != nil {
		//TO DO need to made more rebust
		panic("wrong login message")
	}
	userCode, _ := userStorage.CheckUserNameAndPassword(&((*user).UserName), &((*user).Password))
	var sysError = utility.Error {
			ErrorCode: 0,
			ErrorMessage: "",
	}
	var resultString * string
	if *userCode == "" {
		tmpResultString := ""
		resultString = &tmpResultString
	} else {
		user, _ := userStorage.QueryUserByUserCode(userCode)
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

func GetAllUser(userCode * string, parameter * string) {
	var errorCode int
	var userSlice [] utility.User
	if errorCode == 0 {
		userSlice, _ = userStorage.QueryAllUser()
	}
	command := "getAllUser"
	result := utility.SliceResultToOutMessage(&command, userSlice, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}