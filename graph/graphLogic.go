package graph

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func GetAllNode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Graph
	if (errorCode == 0) {
		campaignCode, _ := utility.ParseCampaignCodeMessage(&(inputParameters[1]))
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
	result := utility.SliceResultToOutMessage(&command, resultSlice, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

func AddNode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var graphCode * string
	var resultSlice [] string
	if (errorCode == 0) {
		graph, _ := utility.ParseGraphMessage(&(inputParameters[1]))
		graph.GraphCode = *(utility.GenerateCode(&(inputParameters[0])))
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
	result := utility.BoolResultToOutMessage(&command, resultSlice, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

func ModifyNode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var graphCode * string
	var graphOut * utility.Graph
	if (errorCode == 0) {
		graph, _ := utility.ParseGraphMessage(&(inputParameters[1]))
		graphCode = &(graph.GraphCode)
		opResult, _ :=storage.ModifyGraph(graph)
		if opResult == false {
			errorCode = 1
		} else {
			graphOut, _ = storage.QueryGraphNodeByGraphCode(graphCode)
		}
	}
	var command = "modifyNode"
	result := utility.BoolResultToOutMessage(&command, graphOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}


func DeleteNode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}

	if (errorCode == 0) {
		graphCode, _ := utility.ParseGraphCodeMessage(&(inputParameters[1]))
		opResult, _ :=storage.DeleteGraphByGraphCode(&(graphCode.GraphCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteNode"
	result := utility.StringResultToOutMessage(&command, &inputParameters[1], errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}