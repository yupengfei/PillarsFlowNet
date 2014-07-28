package utility

import "testing"
import "os"
import "io/ioutil"
import "PillarsFlowNet/pillarsLog"
import "fmt"


func TestParseInMessage(t * testing.T) {
	file, err := os.Open("login.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	data, error := ioutil.ReadAll(file)
	if error != nil {
		panic(error.Error())
	}

	command, parameter, err := ParseInMessage(data)
	if err != nil {
		pillarsLog.Logger.Panic(err.Error())
	}

	if *command != "login" {
		t.Error("parse wrong")
	} else {
		
		user, error := ParseLoginInMessage(parameter)
		if error != nil {
			t.Error("parse login message wrong")
		}
		if (*user).UserName != "yupengfei" {
			t.Error("username wrong")
		}
		if (*user).Password != "123456" {
			t.Error("password wrong")
		}
	}
}

func TestObjectToJson(t * testing.T) {
	var error = Error {
		ErrorCode: 0,
		ErrorMessage: "",
	}
	var login = LoginInMessage {
		Auth: "success",
		AuthMessage : "",
	}
	loginStr := string(LoginMessageToJson(login))
	var out = OutMessage {
		Error: error,
		Command: "login",
		Result: loginStr,
	}
	fmt.Println(string(LoginMessageToJson(out)))
}


