package mission

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

//获取某个Project所有的Campaign
//TODO
//将该函数改名为GetProjectCampaign
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


	// var sysError = utility.Error {
	// 					ErrorCode: errorCode,
	// 					ErrorMessage: "",
	// 				}
	// var out *  utility.OutMessage
	// if opResult == nil {
	// 	out = &(utility.OutMessage {
	// 					Error: sysError,
	// 					Command: "getAllCampaign",
	// 					UserCode: inputParameters[0],
	// 					Result:"{}",
	// 				})
	// } else {
	// 	out = &(utility.OutMessage {
	// 					Error: sysError,
	// 					Command: "getAllCampaign",
	// 					UserCode: inputParameters[0],
	// 					Result:*utility.ObjectToJsonString(opResult),
	// 				})
	// }
	
	// var result = utility.ObjectToJsonByte(out)
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

func AddMission(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var missionCode * string
	var missionOut * utility.Mission
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(&(inputParameters[1]))
		mission.MissionCode = *(utility.GenerateCode(&(inputParameters[0])))
		missionCode = &(mission.MissionCode)
		opResult, _ :=storage.InsertIntoMission(mission)
		if opResult == false {
			errorCode = 1
		} else {
			missionOut, _ = storage.QueryMissionByMissionCode(missionCode)
		}
	}
	command := "addMission"
	result := utility.BoolResultToOutMessage(&command, missionOut, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

func ModifyMission(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var missionCode * string
	var missionOut * utility.Mission
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(&(inputParameters[1]))
		opResult, _ :=storage.ModifyMission(mission)
		missionCode = &(mission.MissionCode)
		if opResult == false {
			errorCode = 1
		} else {
			missionOut, _ = storage.QueryMissionByMissionCode(missionCode)
		}
	}
	command := "modifyMission"
	result := utility.BoolResultToOutMessage(&command, missionOut, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

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
	var command = "deleteMission"
	result := utility.StringResultToOutMessage(&command, &inputParameters[1], errorCode, &inputParameters[0])
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
	}
	command := "queryMissionByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}


func GetPersonAllWaitingMission(userCodeAndParameter * string) ([] byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryWaitingMissionByUserCode(&(inputParameters[0]))
	}
	command := "getPersonAllWaitingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

func GetPersonAllUndergoingMission(userCodeAndParameter * string) ([] byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryUndergoingMissionByUserCode(&(inputParameters[0]))
	}
	command := "getPersonAllUndergoingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

func GetPersonAllReviewingMission(userCodeAndParameter * string) ([]byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryReviewingMissionByUserCode(&(inputParameters[0]))
	}
	command := "getPersonAllReviewingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

func GetPersonAllFinishedMission(userCodeAndParameter * string) ([]byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryFinishedMissionByUserCode(&(inputParameters[0]))
	}
	command := "getPersonAllFinishedMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

func GetAllUndesignatedMission(userCodeAndParameter * string) ([]byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryAllUndesignatedMission()
	}
	command := "getPersonAllFinishedMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}
