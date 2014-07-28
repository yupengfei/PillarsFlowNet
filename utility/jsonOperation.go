package utility

import "encoding/json"
// import "fmt"
//input a slice contains message, return command and parameter
func ParseInMessage(message [] byte) (* string, * string, error) {
	var result InMessage
	err := json.Unmarshal(message, &result)
	return &result.Command, &result.Parameter, err
}

func ParseLoginInMessage(message * string) (* User, error)  {
	var result User
	err := json.Unmarshal([]byte(*message), &result)
	return &result, err
}


func LoginMessageToJson(object interface{}) []byte {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	return message
}