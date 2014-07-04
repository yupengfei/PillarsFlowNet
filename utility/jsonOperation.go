package utility

import "encoding/json"

type Error struct {
	errorCode string
	errorMessage string
}
type OutMessage struct {
	error Error
	result string
}
type InMessage struct {
	command string
	parameter string
}
type LoginIn struct {
	userName string
	password string
}


func ParserInMessage(message string) (string, string) {
	var result InMessage
	err := json.Unmarshal(message, result)
	if err != nil {
		panic(err.Error())
	}
	return result.command, result.parameter
}

func ObjectToJson(object interface) string {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	return message
}