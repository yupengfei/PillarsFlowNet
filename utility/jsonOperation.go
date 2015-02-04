package utility

import "encoding/json"

// import "fmt"
//input a slice contains message, return command and parameter
func ParseInMessage(message []byte) (*string, *string, error) {
	var result InMessage
	err := json.Unmarshal(message, &result)
	return &result.Command, &result.Parameter, err
}

func ParseLoginInMessage(message *string) (*UserLogin, error) {
	var result UserLogin
	err := json.Unmarshal([]byte(*message), &result)
	return &result, err
}

func ParseChartMessage(message *string) (*Chart, error) {
	var chart Chart
	err := json.Unmarshal([]byte(*message), &chart)
	return &chart, err
}

func ParsePostMessage(message *string) (*Post, error) {
	var post Post
	err := json.Unmarshal([]byte(*message), &post)
	return &post, err
}

func ParseProjectMessage(message *string) (*Project, error) {
	var project Project
	err := json.Unmarshal([]byte(*message), &project)
	return &project, err
}

func ParseMissionMessage(message *string) (*Mission, error) {
	var mission Mission
	err := json.Unmarshal([]byte(*message), &mission)
	return &mission, err
}

func ParseTargetMessage(message *string) (*Target, error) {
	var target Target
	err := json.Unmarshal([]byte(*message), &target)
	return &target, err
}

func ParseDailyMessage(message *string) (*Daily, error) {
	var daily Daily
	err := json.Unmarshal([]byte(*message), &daily)
	return &daily, err
}

func ParseDependencyMessage(message *string) (*Dependency, error) {
	var dependency Dependency
	err := json.Unmarshal([]byte(*message), &dependency)
	return &dependency, err
}

func ParseGraphMessage(message *string) (*Graph, error) {
	var graph Graph
	err := json.Unmarshal([]byte(*message), &graph)
	return &graph, err
}

func ParseProjectCodeMessage(message *string) (*ProjectCode, error) {
	var projectCode ProjectCode
	err := json.Unmarshal([]byte(*message), &projectCode)
	return &projectCode, err
}

func ParseMissionCodeMessage(message *string) (*MissionCode, error) {
	var missionCode MissionCode
	err := json.Unmarshal([]byte(*message), &missionCode)
	return &missionCode, err
}

func ParseCampaignCodeMessage(message *string) (*CampaignCode, error) {
	var campaignCode CampaignCode
	err := json.Unmarshal([]byte(*message), &campaignCode)
	return &campaignCode, err
}

func ParseGraphCodeMessage(message *string) (*GraphCode, error) {
	var graphCode GraphCode
	err := json.Unmarshal([]byte(*message), &graphCode)
	return &graphCode, err
}

func ParseDependencyCodeMessage(message *string) (*DependencyCode, error) {
	var dependencyCode DependencyCode
	err := json.Unmarshal([]byte(*message), &dependencyCode)
	return &dependencyCode, err
}

func ParseTargetCodeMessage(message *string) (*TargetCode, error) {
	var targetCode TargetCode
	err := json.Unmarshal([]byte(*message), &targetCode)
	return &targetCode, err
}

func ParseDailyCodeMessage(message *string) (*DailyCode, error) {
	var dailyCode DailyCode
	err := json.Unmarshal([]byte(*message), &dailyCode)
	return &dailyCode, err
}

func ParseChartCodeMessage(message *string) (*ChartCode, error) {
	var chartCode ChartCode
	err := json.Unmarshal([]byte(*message), &chartCode)
	return &chartCode, err
}

func ObjectToJsonByte(object interface{}) []byte {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	return message
}

func ObjectToJsonString(object interface{}) *string {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	messageString := string(message)
	return &messageString
}

func ParseStringSlice(message *string) ([]string, error) {
	var msg []string
	err := json.Unmarshal([]byte(*message), &msg)
	return msg, err
}
