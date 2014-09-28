package user

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/pillarsLog"
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
	validLogin, _ := storage.CheckUserNameAndPassword(&((*user).UserName), &((*user).Password))
	var sysError = utility.Error {
			ErrorCode: 0,
			ErrorMessage: "",
	}
	var loginMessage utility.LoginInMessage
	if validLogin {
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
	return &result, &(user.UserName), err
}