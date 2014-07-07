package utility

import "testing"
import "fmt"
import "os"
import "io/ioutil"

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
	//fmt.Print(data)
	//ParserInMessage(data)
	command := ParseInMessage(&data)
	//fmt.Print(command)
	if *command != "login" {
		t.Error("parse wrong")
	}
	// fmt.Print(parameter)
}

func TestObjectToJson(t * testing.T) {
	file, err := os.Open("login.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	data, error := ioutil.ReadAll(file)
	if error != nil {
		panic(error.Error())
	}

	LoginInMessage := ParseLoginInMessage(&data)
	
	bytes := LoginMessageToJson(LoginInMessage)
	//fmt.Println(string(*bytes))

}



