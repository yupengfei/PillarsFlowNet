package graph

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

// 获取Campaign的全部node
// {
// 	“command”:”getAllNode”,
// 	“parameter”:”{
// 		CampaignCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “getAllNode”,
// 	“result”:”[{
// 		GraphCode string
// 		CampaignCode string
//     		ProjectCode string
//     		NodeCode string
//     		Width int
//     		Height int
//     		XCoordinate int
//     		YCoordinate int
//     		InsertDatetime string
//     		UpdateDatetime string

// 	}]”
// }

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
		// if opResult == false {
		// 	errorCode = 1
		// }
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
	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllNode",
						UserCode: inputParameters[0],
						Result:*utility.ObjectToJsonString(resultSlice),
					}
	var result = utility.ObjectToJsonByte(&out)

	return result, &(inputParameters[0])
}


// 新建node
// {
// 	“command”:”addNode”,
// 	“parameter”:”{
// 		GraphCode string 任意string，不起作用,可以没有
// 		CampaignCode string
//     		ProjectCode string
//     		NodeCode string
//     		Width int
//     		Height int
//     		XCoordinate int
//     		YCoordinate int
//     		InsertDatetime 任意string，不起作用,可以没有
//     		UpdateDatetime 任意string，不起作用,可以没有
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “addNode”,
// 	“result”:”{
// 		GraphCode string
// 		CampaignCode string
//     		ProjectCode string
//     		NodeCode string
//     		Width int
//     		Height int
//     		XCoordinate int
//     		YCoordinate int
//     		InsertDatetime string
//     		UpdateDatetime string
// 	}”

func AddNode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var graphCode * string
	if (errorCode == 0) {
		graph, _ := utility.ParseGraphMessage(&(inputParameters[1]))
		graph.GraphCode = *(utility.GenerateCode(&(inputParameters[0])))
		graphCode = &(graph.GraphCode)
		opResult, _ :=storage.InsertIntoGraph(graph)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out *  utility.OutMessage
	if errorCode != 0 {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "addNode",
						UserCode: inputParameters[0],
						Result: "{}",
					}
		out = & tempout
	} else {
		graphOut, _ := storage.QueryGraphNodeByGraphCode(graphCode)
		var resultSlice [] string
		resultSlice = append(resultSlice, *utility.ObjectToJsonString(graphOut))
		mission, err := storage.QueryMissionByMissionCode(&(graphOut.NodeCode))
		if err != nil {
			mission = new(utility.Mission)
		}
		resultSlice = append(resultSlice, *utility.ObjectToJsonString(mission))
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "addNode",
						UserCode: inputParameters[0],
						Result: *utility.ObjectToJsonString(resultSlice),
					}
		out = & tempout
	}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 修改node
// {
// 	“command”:”modifyNode”,
// 	“parameter”:”{
// 		GraphCode string
// 		CampaignCode string
//     		ProjectCode string
//     		NodeCode string
//     		Width int
//     		Height int
//     		XCoordinate int
//     		YCoordinate int
//     		InsertDatetime 任意string，不起作用,可以没有
//     		UpdateDatetime 任意string，不起作用,可以没有
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “modifyNode”,
// 	“result”:”{
// 		GraphCode string
// 		CampaignCode string
//     		ProjectCode string
//     		NodeCode string
//     		Width int
//     		Height int
//     		XCoordinate int
//     		YCoordinate int
//     		InsertDatetime string
//     		UpdateDatetime string
// 	}”
// }

func ModifyNode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var graphCode * string
	if (errorCode == 0) {
		graph, _ := utility.ParseGraphMessage(&(inputParameters[1]))
		graphCode = &(graph.GraphCode)
		opResult, _ :=storage.ModifyGraph(graph)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out *  utility.OutMessage
	if errorCode != 0 {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "modifyNode",
						UserCode: inputParameters[0],
						Result: "{}",
					}
		out = & tempout
	} else {
		graphOut, _ := storage.QueryGraphNodeByGraphCode(graphCode)
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "modifyNode",
						UserCode: inputParameters[0],
						Result:*utility.ObjectToJsonString(graphOut),
					}
		out = & tempout
	}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 删除node
// {
// 	“command”:”deleteNode”,
// 	“parameter”:”{
// 		GraphCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “deleteNode”,
// 	“result”:”{
// 	GraphCode string	
// 	}”
// }

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

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out *  utility.OutMessage				
	if errorCode !=0 {
		tempout := utility.OutMessage {
						Error: sysError,
						Command: "deleteNode",
						UserCode: inputParameters[0],
						Result: "{}",
					}
		out = & tempout
	} else {
		tempout := utility.OutMessage {
						Error: sysError,
						Command: "deleteNode",
						UserCode: inputParameters[0],
						Result: inputParameters[1],
					}
		out = & tempout
	}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}