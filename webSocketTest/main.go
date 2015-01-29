// webSocketTest project main.go
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var client Client
var usercode string
var missioncode string
var projectcode string = "asdssssssssssssssssssss"
var nodecode string
var node_mission string
var denpendcode string
var targetcode string
var dailycode string
var ticker *time.Ticker

func main() {
	url := "ws://localhost:1234/ws" //////!!!
	client = Client{}
	client.initClient()
	ok := client.connectServer(&url)
	if ok {
		fmt.Println("connect success")
	} else {
		fmt.Println("connect false")
		return
	}
	go client.onReadData()

	TestLogin()
	ticker = time.NewTicker(3 * time.Second)
	/*****************Test add Logic*********************/
	<-ticker.C
	TestAddProject()
	//<-ticker.C
	//TestAddMission(&projectcode)
	<-ticker.C
	TestaddNode(&projectcode)
	<-ticker.C
	TestaddDependency(&missioncode, &missioncode, &projectcode)
	<-ticker.C
	TestaddDaily(&missioncode, &projectcode)
	<-ticker.C
	TestaddTarget(&missioncode, &projectcode)
	<-ticker.C
	TestaddChart()
	<-ticker.C
	TestaddPost(&missioncode, &dailycode)
	<-ticker.C
	/*******************Test query  Logic************************/
	TestgetAllTargetPost(&targetcode)
	<-ticker.C
	TestgetAllUndesignatedMission()
	<-ticker.C
	TestgetAllUnreceivedChart()
	<-ticker.C
	TestgetAllUser()
	<-ticker.C
	TestgetCampaignDependency(&denpendcode)
	<-ticker.C
	TestgetCampaignNode(&nodecode)
	<-ticker.C
	TestgetDailyByMissionCode(&missioncode)
	<-ticker.C
	TestgetPersonAllFinishedMission()
	<-ticker.C
	TestgetPersonAllReviewingMission()
	<-ticker.C
	TestgetPersonAllUndergoingMission()
	<-ticker.C
	TestgetPersonAllWaitingMission()
	<-ticker.C
	TestgetProjectAssertCampaign(&missioncode)
	<-ticker.C
	TestgetProjectShotCampaign(&missioncode)
	<-ticker.C
	TestgetTargetByMissionCode(&missioncode)
	<-ticker.C
	/*****************Test modify Logic***************************/
	TestmodifyProject(&projectcode)
	<-ticker.C
	TestmodifyMission(&projectcode, &missioncode)
	<-ticker.C
	TestmodifyDaily(&dailycode)
	<-ticker.C
	TestmodifyDependency(&missioncode, &missioncode, &projectcode, &denpendcode)
	<-ticker.C
	TestmodifyNode(&missioncode, &projectcode, &missioncode, &nodecode)
	<-ticker.C
	TestmodifyTarget(&missioncode, &projectcode, &targetcode)
	<-ticker.C
	/********************Test   delete Logic***************************/

	TestdeleteDaily(&dailycode)
	<-ticker.C
	TestdeleteDependency(&denpendcode)
	<-ticker.C
	TestdeleteMission(&missioncode)
	<-ticker.C
	TestdeleteNode(&nodecode)
	<-ticker.C
	TestdeleteTarget(&targetcode)

	<-ticker.C
	fmt.Println("Test is over .")
}
func OnTest() {

}
func TestLogin() {
	var logon UserLogin
	logon.Email = "2@163.com"
	logon.Password = "111"
	var message OutMessage
	message.Command = "login"
	msg, _ := json.Marshal(logon)
	message.Parameter = string(msg)
	client.onWriteData(ObjectToJsonByte(message))

}
func TestAddProject() {
	projectName := string("very maim")

	project := Project{
		ProjectCode:       "",
		ProjectName:       projectName,
		ProjectDetail:     "project detail test",
		PlanBeginDatetime: "2014-07-15 00:12:00",
		PlanEndDatetime:   "2014-08-29 00:05:00",
		RealBeginDatetime: "2014-07-15 00:12:00",
		RealEndDatetime:   "2014-08-29 00:05:00",
		PersonInCharge:    "yupengfei",
		Status:            1,
		Picture:           "123",
	}
	var message OutMessage
	message.Command = "addProject"
	msg, _ := json.Marshal(project)
	message.Parameter = string(msg)
	client.onWriteData(ObjectToJsonByte(message))
}
func TestmodifyProject(code *string) {
	projectName := string("very maim")
	//projectCode := string("890b2f1b208f93586e0ad86cb6f16668")
	project := Project{
		ProjectCode:       *code,
		ProjectName:       projectName,
		ProjectDetail:     "project detail Modify",
		PlanBeginDatetime: "2014-07-15 00:12:00",
		PlanEndDatetime:   "2014-08-29 00:05:00",
		RealBeginDatetime: "2014-07-15 00:12:00",
		RealEndDatetime:   "2014-08-29 00:05:00",
		PersonInCharge:    "hanser",
		Status:            1,
		Picture:           "123",
	}
	var message OutMessage
	message.Command = "modifyProject"
	msg, _ := json.Marshal(project)
	message.Parameter = string(msg)
	client.onWriteData(ObjectToJsonByte(message))
}
func TestgetProjectShotCampaign(code *string) {
	var _code ProjectCode
	_code.ProjectCode = *code
	var message OutMessage
	message.Command = "getProjectShotCampaign"
	msg, _ := json.Marshal(_code)
	message.Parameter = string(msg)
	client.onWriteData(ObjectToJsonByte(message))
}
func TestgetProjectAssertCampaign(code *string) {
	var _code ProjectCode
	_code.ProjectCode = *code
	var message OutMessage
	message.Command = "getProjectAssertCampaign"
	msg, _ := json.Marshal(_code)
	message.Parameter = string(msg)
	client.onWriteData(ObjectToJsonByte(message))
}
func TestAddMission(project *string) {
	cmd := "addMission"
	missionName := string("modify the test")
	//missionCode := "a115313c765a01505acd6a5260c7d0ea"

	//projectCode := string("d655313c765a01505acd6a5260c7d1ee")
	mission := Mission{
		MissionCode:       "",
		MissionName:       missionName,
		ProjectCode:       *project,
		ProductType:       "1",
		IsCampaign:        0,
		MissionDetail:     "test mission detail",
		PlanBeginDatetime: "2014-07-15 00:12:00",
		PlanEndDatetime:   "2014-08-29 00:05:00",
		RealBeginDatetime: "2014-07-15 00:12:00",
		RealEndDatetime:   "2014-08-29 00:05:00",
		PersonIncharge:    "yupengfei",
		Status:            1,
		Picture:           "123",
	}
	msg, _ := json.Marshal(mission)
	go Trace(&cmd, msg)
}

func TestmodifyMission(project, code *string) {
	cmd := "modifyMission"
	Object := Mission{
		MissionCode:       *code,
		MissionName:       "modify mission",
		ProjectCode:       *project,
		ProductType:       "0",
		IsCampaign:        0,
		MissionDetail:     "test mission detail",
		PlanBeginDatetime: "2014-07-15 00:12:00",
		PlanEndDatetime:   "2014-08-29 00:05:00",
		RealBeginDatetime: "2014-07-15 00:12:00",
		RealEndDatetime:   "2014-08-29 00:05:00",
		PersonIncharge:    "yupengfei",
		Status:            1,
		Picture:           "456",
	}

	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestdeleteMission(code *string) {
	cmd := "deleteMission"
	var Object MissionCode
	Object.MissionCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestgetMissionByMissionCode(code *string) {
	cmd := "getMissionByMissionCode"
	var Object MissionCode
	Object.MissionCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestgetCampaignNode(code *string) {
	cmd := "getCampaignNode"
	var Object CampaignCode
	Object.CampaignCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestaddNode(project *string) {
	//cmd := "addNode"
	width := 10
	height := 10
	xCoordinate := 100
	yCoordinate := 100
	node := Graph{
		GraphCode:    "",
		CampaignCode: "",
		ProjectCode:  *project,
		NodeCode:     "", //因为它就是下面的missionCode,服务端处理
		Width:        width,
		Height:       height,
		XCoordinate:  xCoordinate,
		YCoordinate:  yCoordinate,
	}
	_mission := Mission{
		MissionCode:       "", ///这个地方有问题　前台和后台该不该处理，应该了解业务
		MissionName:       "add a Node's mission",
		ProjectCode:       *project,
		ProductType:       "0",
		IsCampaign:        1,
		MissionDetail:     "test mission detail",
		PlanBeginDatetime: "2014-07-15 00:12:00",
		PlanEndDatetime:   "2014-08-29 00:05:00",
		RealBeginDatetime: "2014-07-15 00:12:00",
		RealEndDatetime:   "2014-08-29 00:05:00",
		PersonIncharge:    "yupengfei",
		Status:            1,
		Picture:           "123",
	}
	Object := AddNodeMsg{}
	node_str, _ := json.Marshal(node)
	miss_str, _ := json.Marshal(_mission)
	Object.Content[0] = string(node_str)
	Object.Content[1] = string(miss_str)
	msg, _ := json.Marshal(Object)
	//fmt.Println(string(msg))
	go Trace(&cmd, msg)
}
func TestmodifyNode(camCode, project, mission, code *string) {
	cmd := "modifyNode"
	Object := Graph{
		GraphCode:    *code,
		CampaignCode: *camCode,
		ProjectCode:  *project,
		NodeCode:     *mission,
		Width:        100,
		Height:       100,
		XCoordinate:  50,
		YCoordinate:  50,
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestdeleteNode(code *string) {
	cmd := "deleteNode"
	Object := GraphCode{}
	Object.GraphCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestgetCampaignDependency(code *string) {
	cmd := "getCampaignDependency"
	Object := CampaignCode{}
	Object.CampaignCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestaddDependency(star, end, project *string) {
	cmd := "addDependency"
	Object := Dependency{
		DependencyCode:   "",
		ProjectCode:      *project,
		StartMissionCode: *star,
		EndMissionCode:   *end,
		DependencyType:   1,
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestdeleteDependency(code *string) {
	cmd := "deleteDependency"
	Object := DependencyCode{}
	Object.DependencyCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestmodifyDependency(star, end, project, code *string) {
	cmd := "modifyDependency"
	Object := Dependency{
		DependencyCode:   *code,
		ProjectCode:      *project,
		StartMissionCode: *star,
		EndMissionCode:   *end,
		DependencyType:   1,
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestaddTarget(mission, project *string) {
	cmd := "addTarget"
	Object := Target{
		TargetCode:      "",
		MissionCode:     *mission,
		ProjectCode:     *project,
		VersionTag:      string("0.1"),
		StoragePosition: string("/home"),
		Picture:         string(""),
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestmodifyTarget(mission, project, code *string) {
	cmd := "modifyTarget"
	Object := Target{
		TargetCode:      *code,
		MissionCode:     *mission,
		ProjectCode:     *project,
		VersionTag:      string("1.0"),
		StoragePosition: string("/home"),
		Picture:         string("asdfasdf"),
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestdeleteTarget(code *string) {
	cmd := "deleteTarget"
	Object := TargetCode{}
	Object.TargetCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestgetTargetByMissionCode(code *string) {
	cmd := "getTargetByMissionCode"
	Object := MissionCode{}
	Object.MissionCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestaddChart() {
	cmd := "addChart"
	Object := Chart{
		Id:           "",
		IsPicture:    0,
		Message:      "asdfasdf",
		From:         "hanser",
		SendTime:     time.Now().Format("2006-01-02 15:04:05"),
		To:           "123456",
		ReceivedTime: time.Now().Format("2006-01-02 15:04:05"),
		IsReceived:   0,
		Deleted:      0,
		DeletedTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestreceiveChart(code *string) {
	cmd := "receiveChart"
	var Object ChartCode
	Object.ChartCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestgetAllUnreceivedChart() {
	cmd := "getAllUnreceivedChart"
	go Trace(&cmd, []byte{})
}
func TestaddPost(mission, code *string) {
	cmd := "addPost"
	Object := Post{
		Id:          "",
		MissionCode: *mission,
		PostType:    0,
		Code:        *code,
		IsPicture:   0,
		Message:     "111111111",
		ReplyTo:     "hanser",
		UserCode:    "",
		PostTime:    time.Now().Format("2006-01-02 15:04:05"),
		Deleted:     0,
		DeletedTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestgetAllTargetPost(code *string) {
	cmd := "getAllTargetPost"
	var Object TargetCode
	Object.TargetCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestgetAllUser() {
	cmd := "getAllUser"
	go Trace(&cmd, []byte{})
}
func TestgetPersonAllWaitingMission() {
	cmd := "getPersonAllWaitingMission"
	go Trace(&cmd, []byte{})
}
func TestgetPersonAllUndergoingMission() {
	cmd := "getPersonAllUndergoingMission"
	go Trace(&cmd, []byte{})
}
func TestgetPersonAllReviewingMission() {
	cmd := "getPersonAllReviewingMission"
	go Trace(&cmd, []byte{})
}
func TestgetPersonAllFinishedMission() {
	cmd := "getPersonAllFinishedMission"
	go Trace(&cmd, []byte{})
}
func TestgetAllUndesignatedMission() {
	cmd := "getAllUndesignatedMission"
	go Trace(&cmd, []byte{})
}
func TestgetDailyByMissionCode(code *string) {
	cmd := "getDailyByMissionCode"
	var Object MissionCode
	Object.MissionCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestaddDaily(mission, project *string) {
	cmd := "addDaily"
	Object := Daily{
		DailyCode:       "",
		MissionCode:     *mission,
		ProjectCode:     *project,
		VersionTag:      string("0.1"),
		StoragePosition: string("/home"),
		Picture:         string(""),
	}
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestmodifyDaily(code *string) {
	cmd := "modifyDaily"
	var Object DailyCode
	Object.DailyCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func TestdeleteDaily(code *string) {
	cmd := "deleteDaily"
	Object := DailyCode{}
	Object.DailyCode = *code
	msg, _ := json.Marshal(Object)
	go Trace(&cmd, msg)
}
func Trace(cmd *string, msg []byte) {
	var message OutMessage
	message.Command = *cmd
	message.Parameter = string(msg)
	client.onWriteData(ObjectToJsonByte(message))
}
