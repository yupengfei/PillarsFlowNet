package main

import (
	"encoding/json"
	"fmt"
)

func ParseInMessage(message []byte) (*string, *string, Error) {
	var result InMessage
	err := json.Unmarshal(message, &result)
	if err != nil {
		print("utility line 11 ParseInMessage has error: " + err.Error())
		return nil, nil, Error{}
	}
	return &result.Command, &result.Result, result.Error
}
func receveForgetAllProject(receve *string) {
	var result []Project
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForaddProject(receve *string) {
	var result Project
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	projectcode = result.ProjectCode
}
func receveFormodifyProject(receve *string) {
	var result Project
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetProjectAssertCampaign(receve *string) {
	var result []Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetProjectShotCampaign(receve *string) {
	var result []Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForaddMission(receve *string) {
	var result Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	missioncode = result.MissionCode
}
func receveForgetMissionByMissionCode(receve *string) {
	var result Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveFormodifyMission(receve *string) {
	var result Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveFordeleteMission(receve *string) {
	var result MissionCode
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetCampaignNode(receve *string) {
	fmt.Println(*receve)
	var result []string
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	if len(result) == 0 {
		return
	}
	if len(result)%2 != 0 {
		fmt.Println("this Campaign's lenght !=2!!!!!!!!!!!")
		return
	}
	for i := 0; i < len(result); {
		fmt.Println(result[i], result[i+1])
		i += 2
	}
}
func receveForaddNode(receve *string) {
	var result []string
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	var node Graph
	var mision Mission
	json.Unmarshal([]byte(result[0]), &node)
	json.Unmarshal([]byte(result[1]), &mision)
	nodecode = node.GraphCode
	missioncode = mision.MissionCode
	fmt.Println("node: ", node)
	fmt.Println("mission: ", mision)
}
func receveFormodifyNode(receve *string) {
	fmt.Println("STRING: " + *receve)
	var result Graph
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveFordeleteNode(receve *string) {
	var result GraphCode
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetCampaignDependency(receve *string) {
	var result []Dependency
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForaddDependency(receve *string) {
	var result Dependency
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	denpendcode = result.DependencyCode
}
func receveFormodifyDependency(receve *string) {
	var result Dependency
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveFordeleteDependency(receve *string) {
	var result DependencyCode
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForaddTarget(receve *string) {
	var result Target
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	targetcode = result.TargetCode
}
func receveFormodifyTarget(receve *string) {
	var result Target
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveFordeleteTarget(receve *string) {
	var result TargetCode
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetTargetByMissionCode(receve *string) {
	var result []Target
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForaddDaily(receve *string) {
	var result Daily
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	dailycode = result.DailyCode
}
func receveFormodifyDaily(receve *string) {
	var result Daily
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveFordeleteDaily(receve *string) {
	var result DailyCode
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetDailyByMissionCode(receve *string) {
	var result []Daily
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

////////////////////!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func receveForgetAllUser(receve *string) {
	var result []User
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	s := len(result)
	for i := 0; i < s; i++ {
		fmt.Println(result[i].Email, result[i].DisplayName)
	}
}
func receveForaddChart(receve *string) {
	var result Chart
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForreceiveChart(receve *string) {
	var result ChartCode
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetAllUnreceivedChart(receve *string) {
	var result []Chart
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForaddPost(receve *string) {
	var result Post
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetAllTargetPost(receve *string) {
	var result []Post
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetPersonAllWaitingMission(receve *string) {
	var result []Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetPersonAllUndergoingMission(receve *string) {
	var result []Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetPersonAllReviewingMission(receve *string) {
	var result []Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetPersonAllFinishedMission(receve *string) {
	var result []Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func receveForgetAllUndesignatedMission(receve *string) {
	var result []Mission
	err := json.Unmarshal([]byte(*receve), &result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
func ObjectToJsonByte(object interface{}) []byte {
	message, err := json.Marshal(object)
	if err != nil {
		panic(err.Error())
	}
	return message
}
