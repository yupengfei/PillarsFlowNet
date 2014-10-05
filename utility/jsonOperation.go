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

func ParseChartMessage(message [] byte) (* string, * string, error) {
	var result ChartMessage
	err := json.Unmarshal(message, &result)
	return &result.Message, &result.To, err
}

func ParseProjectMessage(message * string) (* Project, error) {
	var project Project
	err := json.Unmarshal([]byte(*message), &project)
	return &project, err
}

func ParseMissionMessage(message * string) (*Mission, error) {
	var mission Mission
	err := json.Unmarshal([]byte(*message), &mission)
	return &mission, err
}

func ParseTargetMessage(message * string) (*Target, error) {
	var target Target
	err := json.Unmarshal([]byte(*message), &target)
	return &target, err
}

func ParseDependencyMessage(message * string) (* Dependency , error) {
	var dependency Dependency
	err := json.Unmarshal([]byte(*message), &dependency)
	return &dependency, err
}

func ParseCampaignMessage(message * string) (* Campaign , error) {
	var campaign Campaign
	err := json.Unmarshal([]byte(*message), &campaign)
	return &campaign, err
}


func ParseProjectCodeMessage(message * string) (* ProjectCode, error) {
	var projectCode ProjectCode
	err := json.Unmarshal([]byte(*message), &projectCode)
	return &projectCode, err
}

func ParseMissionCodeMessage(message * string) (* MissionCode, error) {
	var missionCode MissionCode
	err := json.Unmarshal([]byte(*message), &missionCode)
	return &missionCode, err
}


func ObjectToJsonByte(object interface{}) []byte {
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

