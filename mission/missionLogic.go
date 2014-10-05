package mission

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)
// 根据projectCode获取所有的compaign
// {
// 	“command”:”getAllCompaign”,
// 	“parameter”:”{
// 		“ProjectCode”:string
// 	}”
// }
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “getAllCompaign”,
// 	“result”:”[{
// 		MissionCode string
//     		MissionName string
//     		ProjectCode string
//     		ProductType string
//     		IsCampaign int
//     		MissionType string
//     		MissionDetail string
//     		PlanBeginDatetime string
//     		PlanEndDatetime string
//     		RealBeginDatetime string
//     		RealEndDatetime string
//     		PersonIncharge string
//     		Status int
//     		Picture string
//     		InsertDatetime string
//     		UpdateDatetime string		
// 	}”
// }
func GetAllCompaign(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		projectCode, _ := utility.ParseProjectCodeMessage(&(inputParameters[1]))
		opResult, _ =storage.QueryCampaignsByProjectCode(&(projectCode.ProjectCode))
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
						Command: "“getAllCompaign”",
						Result:*utility.ObjectToJsonString(opResult),
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// {
// 	“command”:”addMission”,
// 	“parameter”:”{
// 		MissionCode 任意string，不起作用,可以没有
//     		MissionName string
//     		ProjectCode string
//     		ProductType string
//     		IsCampaign int
//     		MissionType string
//     		MissionDetail string
//     		PlanBeginDatetime string
//     		PlanEndDatetime string
//     		RealBeginDatetime string
//     		RealEndDatetime string
//     		PersonIncharge string
//     		Status int
//     		Picture string
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
// 	“commnd”: “addMission”,
// 	“result”:”{
		
// 	}”
// }

func AddMission(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(&(inputParameters[1]))
		mission.MissionCode = *(utility.GenerateCode(&(inputParameters[0])))
		opResult, _ :=storage.InsertIntoMission(mission)
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
						Command: "addMission",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}


// {
// 	“command”:”modifyMission”,
// 	“parameter”:{
// 		MissionCode string
//     		MissionName string
//     		ProjectCode string
//     		ProductType string
//     		IsCampaign int
//     		MissionType string
//     		MissionDetail string
//     		PlanBeginDatetime string
//     		PlanEndDatetime string
//     		RealBeginDatetime string
//     		RealEndDatetime string
//     		PersonIncharge string
//     		Status int
//     		Picture string
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
// 	“commnd”: “modifyMission”,
// 	“result”:”{
		
// 	}”
// }
func ModifyMission(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(&(inputParameters[1]))
		opResult, _ :=storage.ModifyMission(mission)
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
						Command: "modifyMission",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 删除mission
// {
// 	“command”:”deleteMission”,
// 	“parameter”:”{
// 		MissionCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “deleteMission”,
// 	“result”:”{
		
// 	}”
// }

func DeleteMission(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(&(inputParameters[1]))
		opResult, _ :=storage.DeleteMissionByMissionCode(&(missionCode.MissionCode))
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
						Command: "modifyMission",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}