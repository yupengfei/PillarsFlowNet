package utility

import "encoding/json"
// import "fmt"

type Error struct {
	errorCode string
	errorMessage string
}
type OutMessage struct {
	error Error
	result string
}
type InMessage struct {
	Command string
	Parameter interface{}
}
type LoginIn struct {
	UserName string
	Password string
}
type LoginInMessage struct {
	Command string
	Parameter LoginIn
}


func ParseInMessage(message *[] byte) * string {
	var result InMessage
	
	err := json.Unmarshal(*message, &result)
	if err != nil {
		panic(err.Error())
	}
	return &result.Command
}

func ParseLoginInMessage(message * [] byte) * LoginInMessage {
	var result LoginInMessage
	
	err := json.Unmarshal(*message, &result)
	if err != nil {
		panic(err.Error())
	}
	return &result
}


func LoginMessageToJson(object * LoginInMessage) * []byte {
	message, err := json.Marshal(*object)
	if err != nil {
		panic(err.Error())
	}
	return &message
}