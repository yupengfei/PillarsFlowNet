package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	CDialer websocket.Dialer
	CConn   *websocket.Conn
}

func (this *Client) initClient() {
	this.CDialer = websocket.Dialer{}
	this.CDialer.ReadBufferSize = 1024
	this.CDialer.WriteBufferSize = 1024
}

func (this *Client) connectServer(url *string) bool {
	var err error
	this.CConn, _, err = this.CDialer.Dial(*url, nil)
	if err != nil {
		fmt.Println("Content server error: " + err.Error())
		return false
	}
	return true
}

func (this *Client) onReadData() {
	defer this.closeClient()
	for _, recv, err := this.CConn.ReadMessage(); err == nil; _, recv, err = this.CConn.ReadMessage() {
		command, result, err2 := ParseInMessage(recv)
		recv = nil
		fmt.Println("receve message:" + *command)
		if err2.ErrorCode == 1 {
			fmt.Println("parse " + *command + "error:" + err2.ErrorMessage)
			continue
		}
		switch *command {
		case "getAllProject":
			receveForgetAllProject(result)
		case "addProject":
			receveForaddProject(result)
		case "modifyProject":
			receveFormodifyProject(result)
		case "getProjectAssertCampaign":
			receveForgetProjectAssertCampaign(result)
		case "getProjectShotCampaign":
			receveForgetProjectShotCampaign(result)
		case "addMission":
			receveForaddMission(result)
		case "getMissionByMissionCode":
			receveForgetMissionByMissionCode(result)
		case "modifyMission":
			receveFormodifyMission(result)
		case "deleteMission":
			receveFordeleteMission(result)
		case "getCampaignNode":
			receveForgetCampaignNode(result)
		case "addNode":
			receveForaddNode(result)
		case "modifyNode":
			receveFormodifyNode(result)
		case "deleteNode":
			receveFordeleteNode(result)
		case "getCampaignDependency":
			receveForgetCampaignDependency(result)
		case "addDependency":
			receveForaddDependency(result)
		case "modifyDependency":
			receveFormodifyDependency(result)
		case "deleteDependency":
			receveFordeleteDependency(result)
		case "addTarget":
			receveForaddTarget(result)
		case "modifyTarget":
			receveFormodifyTarget(result)
		case "deleteTarget":
			receveFordeleteTarget(result)
		case "getTargetByMissionCode":
			receveForgetTargetByMissionCode(result)
		case "addDaily":
			receveForaddDaily(result)
		case "modifyDaily":
			receveFormodifyDaily(result)
		case "deleteDaily":
			receveFordeleteDaily(result)
		case "getDailyByMissionCode":
			receveForgetDailyByMissionCode(result)
		case "getAllUser":
			receveForgetAllUser(result)
		case "addChart":
			receveForaddChart(result)
		case "receiveChart":
			receveForreceiveChart(result)
		case "getAllUnreceivedChart":
			receveForgetAllUnreceivedChart(result)
		case "addPost":
			receveForaddPost(result)
		case "getAllTargetPost":
			receveForgetAllTargetPost(result)
		case "getPersonAllWaitingMission":
			receveForgetPersonAllWaitingMission(result)
		case "getPersonAllUndergoingMission":
			receveForgetPersonAllUndergoingMission(result)
		case "getPersonAllReviewingMission":
			receveForgetPersonAllReviewingMission(result)
		case "getPersonAllFinishedMission":
			receveForgetPersonAllFinishedMission(result)
		case "getAllUndesignatedMission":
			receveForgetAllUndesignatedMission(result)
		}

	}
}
func (this *Client) onWriteData(data []byte) {
	err := this.CConn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		fmt.Println("onWriteData error: " + err.Error())
	}
}
func (this *Client) closeClient() {
	this.CConn.Close()
	this.CConn = nil
}
