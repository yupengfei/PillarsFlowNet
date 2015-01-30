package connection

import (
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/graphStorage"
	"PillarsFlowNet/missionStorage"
	"PillarsFlowNet/utility"
	"fmt"
)

//获取特定战役所有的node
//TODO
//将该参数改名为GetCampaignNode
func GetCampaignNode(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var opResult []utility.Graph
	if errorCode == 0 {
		campaignCode, _ := utility.ParseCampaignCodeMessage(parameter)
		opResult, _ = graphStorage.QueryGraphNodesByCampaignCode(&(campaignCode.CampaignCode))
	}
	var missionSlice []utility.Mission
	opResultLength := len(opResult)
	var i int
	for i = 0; i < opResultLength; i++ {
		var mission *utility.Mission
		mission, err := missionStorage.QueryMissionByMissionCode(&(opResult[i].NodeCode))
		if err != nil {
			mission = new(utility.Mission)
		}
		missionSlice = append(missionSlice, *mission)
	}
	var resultSlice []string
	for i = 0; i < opResultLength; i++ {
		resultSlice = append(resultSlice, *utility.ObjectToJsonString(opResult[i]))
		resultSlice = append(resultSlice, *utility.ObjectToJsonString(missionSlice[i]))
	}
	command := "getCampaignNode"
	result := utility.SliceResultToOutMessage(&command, resultSlice, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

func AddNode(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	fmt.Println(*parameter)
	var graphCode *string
	var resultSlice utility.AddNodeMsg
	//var graphOut *utility.Graph
	if errorCode == 0 {
		//这里好像没有处理发来的mission结构，添加节点的时候应该也要插入mission结构
		nodeMsg, _ := utility.ParseStringSlice(parameter)
		///////////////添加新节点
		graph, _ := utility.ParseGraphMessage(&nodeMsg.Content[0])
		graph.GraphCode = *(utility.GenerateCode(userCode))
		graphCode = &(graph.GraphCode)
		///////////////添加mission
		fmt.Println(nodeMsg.Content[1])
		mission, _ := utility.ParseMissionMessage(&nodeMsg.Content[1])
		fmt.Println(mission)
		mission.MissionCode = *(utility.GenerateCode(userCode))
		graph.NodeCode = mission.MissionCode //graph的NodeCode等于关联的MissionCode
		opResult, _ := graphStorage.InsertIntoGraph(graph)
		opResult1, _ := missionStorage.InsertIntoMission(mission)
		if opResult == false || opResult1 == false {
			errorCode = 1
		} else {
			graphOut, _ := graphStorage.QueryGraphNodeByGraphCode(graphCode)
			resultSlice.Content[0] = *utility.ObjectToJsonString(graphOut)
			//resultSlice = append(resultSlice.Content, *utility.ObjectToJsonString(graphOut))
			missionOut, _ := missionStorage.QueryMissionByMissionCode(&(mission.MissionCode))
			resultSlice.Content[1] = *utility.ObjectToJsonString(missionOut)
			//resultSlice = append(resultSlice, *utility.ObjectToJsonString(missionOut))
		}
	}

	var command = "addNode"
	result := utility.BoolResultToOutMessage(&command, resultSlice, errorCode, userCode)
	Hub.Dispatch(result)
}

func ModifyNode(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var graphCode *string
	var graphOut *utility.Graph
	if errorCode == 0 {
		graph, _ := utility.ParseGraphMessage(parameter)
		graphCode = &(graph.GraphCode)
		opResult, _ := graphStorage.ModifyGraph(graph)
		if opResult == false {
			errorCode = 1
		} else {
			graphOut, _ = graphStorage.QueryGraphNodeByGraphCode(graphCode)
		}
	}
	var command = "modifyNode"
	result := utility.BoolResultToOutMessage(&command, graphOut, errorCode, userCode)
	Hub.Dispatch(result)
}

func DeleteNode(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}

	if errorCode == 0 {
		graphCode, _ := utility.ParseGraphCodeMessage(parameter)
		opResult, _ := graphStorage.DeleteGraphByGraphCode(&(graphCode.GraphCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteNode"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	Hub.Dispatch(result)
}
