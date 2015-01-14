package graph

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

//获取特定战役所有的node
//TODO
//将该参数改名为GetCampaignNode
func GetCampaignNode(userCode * string, parameter * string) ([] byte, *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Graph
	if (errorCode == 0) {
		campaignCode, _ := utility.ParseCampaignCodeMessage(parameter)
		opResult, _ =storage.QueryGraphNodesByCampaignCode(&(campaignCode.CampaignCode))
	}
	var missionSlice []utility.Mission
	opResultLength := len(opResult)
	var i int
	for i = 0; i < opResultLength; i++ {
		var mission * utility.Mission
		mission, err := storage.QueryMissionByMissionCode(&(opResult[i].NodeCode))
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
	return result, userCode
}

func AddNode(userCode * string, parameter * string) ([] byte, *string) {
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
		opResult, _ :=storage.InsertIntoGraph(graph)
		if opResult == false {
			errorCode = 1
		} else {
			graphOut, _ := storage.QueryGraphNodeByGraphCode(graphCode)
		
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(graphOut))
			mission, _ := storage.QueryMissionByMissionCode(&(graphOut.NodeCode))
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(mission))
		}
	}

	var command = "addNode"
	result := utility.BoolResultToOutMessage(&command, resultSlice, errorCode, userCode)
	return result, userCode
}

func ModifyNode(userCode * string, parameter * string) ([] byte, *string) {
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
		opResult, _ :=storage.ModifyGraph(graph)
		if opResult == false {
			errorCode = 1
		} else {
			graphOut, _ = storage.QueryGraphNodeByGraphCode(graphCode)
		}
	}
	var command = "modifyNode"
	result := utility.BoolResultToOutMessage(&command, graphOut, errorCode, userCode)
	return result, userCode
}


func DeleteNode(userCode * string, parameter * string) ([] byte, *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}

	if (errorCode == 0) {
		graphCode, _ := utility.ParseGraphCodeMessage(parameter)
		opResult, _ :=storage.DeleteGraphByGraphCode(&(graphCode.GraphCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteNode"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	return result, userCode
}