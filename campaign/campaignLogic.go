package campaign

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
	var opResult []utility.Campaign
	if (errorCode == 0) {
		projectCode, _ := utility.ParseProjectCodeMessage(&(inputParameters[1]))
		opResult, _ =storage.QueryCampaignNodesByCampaignCode(&(projectCode.ProjectCode))
		// if opResult == false {
		// 	errorCode = 1
		// }
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllNode",
						Result:*utility.ObjectToJsonString(opResult),
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}


// {
// 	“command”:”addNode”,
// 	“parameter”:”{
// 		CampaignCode 任意string，不起作用,可以没有
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
		
// 	}”
// }

func AddNode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		campaign, _ := utility.ParseCampaignMessage(&(inputParameters[1]))
		campaign.CampaignCode = *(utility.GenerateCode(&(inputParameters[0])))
		opResult, _ :=storage.InsertIntoCampaign(campaign)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "addNode",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 修改node
// {
// 	“command”:”modifyNode”,
// 	“parameter”:”{
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
	if (errorCode == 0) {
		campaign, _ := utility.ParseCampaignMessage(&(inputParameters[1]))
		opResult, _ :=storage.ModifyCampaign(campaign)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "modifyNode",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 删除node
// {
// 	“command”:”deleteNode”,
// 	“parameter”:”{
// 		NodeCode string
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
		campaignCode, _ := utility.ParseCampaignCodeMessage(&(inputParameters[1]))
		opResult, _ :=storage.DeleteCampaignByCampaignCode(&(campaignCode.CampaignCode))
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "deleteNode",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}