package graphLogic

import (
	"PillarsFlowNet/graphStorage"
	"PillarsFlowNet/missionStorage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
)

//获取特定战役所有的node
//TODO
//将该参数改名为GetCampaignNode
func GetCampaignNode(userCode * string, parameter * string, h * utility.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Graph
	if (errorCode == 0) {
		campaignCode, _ := utility.ParseCampaignCodeMessage(parameter)
		opResult, _ =graphStorage.QueryGraphNodesByCampaignCode(&(campaignCode.CampaignCode))
	}
	var missionSlice []utility.Mission
	opResultLength := len(opResult)
	var i int
	for i = 0; i < opResultLength; i++ {
		var mission * utility.Mission
		mission, err := missionStorage.QueryMissionByMissionCode(&(opResult[i].NodeCode))
		if err != nil {
			mission = new(utility.Mission)
		}
		missionSlice = append(missionSlice, *mission)
	}
	var resultSlice [] string
	for i = 0; i < opResultLength; i++ {
		resultSlice = append(resultSlice, *utility.ObjectToJsonString(opResult[i]))
		resultSlice = append(resultSlice, *utility.ObjectToJsonString(missionSlice[i]))
	}
	command := "getAllNode"
	result := utility.SliceResultToOutMessage(&command, resultSlice, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}

func AddNode(userCode * string, parameter * string, h * utility.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var graphCode * string
	var resultSlice [] string
	if (errorCode == 0) {
		graph, _ := utility.ParseGraphMessage(parameter)
		graph.GraphCode = *(utility.GenerateCode(userCode))
		graphCode = &(graph.GraphCode)
		opResult, _ :=graphStorage.InsertIntoGraph(graph)
		if opResult == false {
			errorCode = 1
		} else {
			graphOut, _ := graphStorage.QueryGraphNodeByGraphCode(graphCode)
		
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(graphOut))
			mission, _ := missionStorage.QueryMissionByMissionCode(&(graphOut.NodeCode))
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(mission))
		}
	}

	var command = "addNode"
	result := utility.BoolResultToOutMessage(&command, resultSlice, errorCode, userCode)
	h.Dispatch(result)
}

func ModifyNode(userCode * string, parameter * string, h * utility.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var graphCode * string
	var graphOut * utility.Graph
	if (errorCode == 0) {
		graph, _ := utility.ParseGraphMessage(parameter)
		graphCode = &(graph.GraphCode)
		opResult, _ :=graphStorage.ModifyGraph(graph)
		if opResult == false {
			errorCode = 1
		} else {
			graphOut, _ = graphStorage.QueryGraphNodeByGraphCode(graphCode)
		}
	}
	var command = "modifyNode"
	result := utility.BoolResultToOutMessage(&command, graphOut, errorCode, userCode)
	h.Dispatch(result)
}


func DeleteNode(userCode * string, parameter * string, h * utility.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}

	if (errorCode == 0) {
		graphCode, _ := utility.ParseGraphCodeMessage(parameter)
		opResult, _ :=graphStorage.DeleteGraphByGraphCode(&(graphCode.GraphCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteNode"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	h.Dispatch(result)
}