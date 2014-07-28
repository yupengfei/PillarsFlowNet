package utility

import "testing"
import "os"
import "io/ioutil"
import "PillarsFlowNet/log"


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
		log.Logger.Panic(err.Error())
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


