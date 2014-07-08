package utility

import "encoding/json"
// import "fmt"

func ParseInMessage(message *[] byte) (* string, error) {
	var result InMessage
	
	err := json.Unmarshal(*message, &result)
	// if err != nil {
	// 	panic(err.Error())
	// }
	return &(result.Command), err
}

func ParseLoginInMessage(message * [] byte) (* User, error)  {
	var result User
	
	err := json.Unmarshal(*message, &result)
	// if err != nil {
	// 	panic(err.Error())
	// }
	return &result, err
}


func LoginMessageToJson(object * User) * []byte {
	message, err := json.Marshal(*object)
	if err != nil {
		panic(err.Error())
	}
	return &message
}