package utility

import "encoding/json"
// import "fmt"
//input a slice contains message, return command and parameter
func ParseInMessage(message [] byte) (* string, * string, error) {
	var result InMessage
	err := json.Unmarshal(message, &result)
	return &result.Command, &result.Parameter, err
}

func ParseLoginInMessage(message * string) (* UserLogin, error)  {
	var result UserLogin
	err := json.Unmarshal([]byte(*message), &result)
	return &result, err
}


func ObjectToJson(object interface{}) []byte {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	return message
}

func ObjectToJsonString(object interface{}) * string {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	messageString := string(message)
	return &messageString
}

func ParseChartMessage(message [] byte) (* string, * string, error) {
	var result ChartMessage
	err := json.Unmarshal(message, &result)
	return &result.Message, &result.To, err
}