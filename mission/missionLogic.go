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
func GetAllCampaign(userCodeAndParameter * string) ([] byte, *string) {
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
						Command: "getAllCampaign",
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
		// MissionCode string
  //   		MissionName string
  //   		ProjectCode string
  //   		ProductType string
  //   		IsCampaign int
  //   		MissionType string
  //   		MissionDetail string
  //   		PlanBeginDatetime string
  //   		PlanEndDatetime string
  //   		RealBeginDatetime string
  //   		RealEndDatetime string
  //   		PersonIncharge string
  //   		Status int
  //   		Picture string
  //   		InsertDatetime string
  //   		UpdateDatetime string
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
	var missionCode * string
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(&(inputParameters[1]))
		mission.MissionCode = *(utility.GenerateCode(&(inputParameters[0])))
		missionCode = &(mission.MissionCode)
		opResult, _ :=storage.InsertIntoMission(mission)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * utility.OutMessage
	if errorCode != 0 {
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addMission",
						Result: "{}",
					}
		out = & tempout
	} else {
		missionOut, _ := storage.QueryMissionByMissionCode(missionCode)
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addMission",
						Result: *utility.ObjectToJsonString(missionOut) ,
					}
		out = & tempout
	}

	var result = utility.ObjectToJsonByte(*out)

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
// MissionCode string
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
func ModifyMission(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var missionCode * string
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(&(inputParameters[1]))
		opResult, _ :=storage.ModifyMission(mission)
		missionCode = &(mission.MissionCode)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * utility.OutMessage
	if errorCode != 0 {
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "modifyMission",
						Result: "{}",
					}
		out = & tempout
	} else {
		missionOut, _ := storage.QueryMissionByMissionCode(missionCode)
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "modifyMission",
						Result: *utility.ObjectToJsonString(missionOut) ,
					}
		out = & tempout
	}

	var result = utility.ObjectToJsonByte(*out)

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
		// MissionCode string
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
	var out * utility.OutMessage
	if errorCode !=0 {
		tempout := utility.OutMessage {
						Error: sysError,
						Command: "deleteMission",
						Result: "{}",
					}
		out = & tempout
	} else {
		tempout := utility.OutMessage {
						Error: sysError,
						Command: "deleteMission",
						Result: inputParameters[1],
					}
		out = & tempout
	}
	
	var result = utility.ObjectToJsonByte(*out)

	return result, &(inputParameters[0])
}

func QueryMissionByMissionCode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult * utility.Mission
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(&(inputParameters[1]))
		opResult, _ =storage.QueryMissionByMissionCode(&(missionCode.MissionCode))
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
						Command: "queryMissionByMissionCode",
						Result:*utility.ObjectToJsonString(opResult),
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}