package connection

import (
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/graphStorage"
	"PillarsFlowNet/missionStorage"
	"PillarsFlowNet/utility"
	//"fmt"
)

func GetProjectAssertCampaign(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	projectCode, _ := utility.ParseProjectCodeMessage(parameter)
	var opResult *utility.Mission
	var opNode []utility.Graph
	var resultSlice []string
	if errorCode == 0 {
		opNode, _ = graphStorage.QueryAssertNodesByCampaignCode(&(projectCode.ProjectCode))
		opResultLength := len(opNode)
		
		for i := 0; i < opResultLength; i++ {
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(opNode[i]))
			opResult, _ = missionStorage.QueryMissionByMissionCode(&(opNode[i].NodeCode))
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(opResult))
		}
	}
	command := "getProjectAssertCampaign"
	result := utility.SliceResultToOutMessage(&command, resultSlice, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
func GetProjectShotCampaign(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	projectCode, _ := utility.ParseProjectCodeMessage(parameter)
	var opResult *utility.Mission
	var opNode []utility.Graph
	var resultSlice []string
	if errorCode == 0 {
		opNode, _ = graphStorage.QueryShotNodesByCampaignCode(&(projectCode.ProjectCode))
		opResultLength := len(opNode)
		
		for i := 0; i < opResultLength; i++ {
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(opNode[i]))
			opResult, _ = missionStorage.QueryMissionByMissionCode(&(opNode[i].NodeCode))
			resultSlice = append(resultSlice, *utility.ObjectToJsonString(opResult))
		}
	}
	command := "getProjectShotCampaign"
	result := utility.SliceResultToOutMessage(&command, resultSlice, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
// func GetProjectUnassertCampaign(userCode *string, parameter *string) {
// 	auth := authentication.GetAuthInformation(userCode)
// 	var errorCode int
// 	if auth == false {
// 		errorCode = 3
// 	}
// 	var opResult []utility.Mission
// 	if errorCode == 0 {
// 		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
// 		opResult, _ = missionStorage.QueryUnassertCampaignsByProjectCode(&(projectCode.ProjectCode))
// 	}
// 	command := "getAllCampaign"
// 	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
// 	Hub.SendToUserCode(result, userCode)
// }

/*
func GetProjectShotCampaign(userCode *string, parameter *string){

	Hub.SendToUserCode(result, userCode)
}*/
func AddMission(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var missionCode *string
	var missionOut *utility.Mission
	if errorCode == 0 {
		mission, _ := utility.ParseMissionMessage(parameter)
		mission.MissionCode = *(utility.GenerateCode(userCode))
		missionCode = &(mission.MissionCode)
		opResult, _ := missionStorage.InsertIntoMission(mission)
		if opResult == false {
			errorCode = 1
		} else {
			missionOut, _ = missionStorage.QueryMissionByMissionCode(missionCode)
		}
	}
	command := "addMission"
	result := utility.BoolResultToOutMessage(&command, missionOut, errorCode, userCode)
	Hub.Dispatch(result)
}

func ModifyMission(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var missionCode *string
	var missionOut *utility.Mission
	if errorCode == 0 {
		mission, _ := utility.ParseMissionMessage(parameter)
		opResult, _ := missionStorage.ModifyMission(mission)
		missionCode = &(mission.MissionCode)
		if opResult == false {
			errorCode = 1
		} else {
			missionOut, _ = missionStorage.QueryMissionByMissionCode(missionCode)
		}
	}
	command := "modifyMission"
	result := utility.BoolResultToOutMessage(&command, missionOut, errorCode, userCode)
	Hub.Dispatch(result)
}

func DeleteMission(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	if errorCode == 0 {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ := missionStorage.DeleteMissionByMissionCode(&(missionCode.MissionCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteMission"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	Hub.Dispatch(result)
}

func GetMissionByMissionCode(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var opResult *utility.Mission
	if errorCode == 0 {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ = missionStorage.QueryMissionByMissionCode(&(missionCode.MissionCode))
	}
	command := "getMissionByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

func GetPersonAllWaitingMission(userCode *string, parameter *string) {
	var errorCode int
	var opResult []utility.Mission
	if errorCode == 0 {
		opResult, _ = missionStorage.QueryWaitingMissionByUserCode(userCode)
	}
	command := "getPersonAllWaitingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

func GetPersonAllUndergoingMission(userCode *string, parameter *string) {
	var errorCode int
	var opResult []utility.Mission
	if errorCode == 0 {
		opResult, _ = missionStorage.QueryUndergoingMissionByUserCode(userCode)
	}
	command := "getPersonAllUndergoingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

func GetPersonAllReviewingMission(userCode *string, parameter *string) {
	var errorCode int
	var opResult []utility.Mission
	if errorCode == 0 {
		opResult, _ = missionStorage.QueryReviewingMissionByUserCode(userCode)
	}
	command := "getPersonAllReviewingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

func GetPersonAllFinishedMission(userCode *string, parameter *string) {
	var errorCode int
	var opResult []utility.Mission
	if errorCode == 0 {
		opResult, _ = missionStorage.QueryFinishedMissionByUserCode(userCode)
	}
	command := "getPersonAllFinishedMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

func GetAllUndesignatedMission(userCode *string, parameter *string) {
	var errorCode int
	var opResult []utility.Mission
	if errorCode == 0 {
		opResult, _ = missionStorage.QueryAllUndesignatedMission()
	}
	command := "GetAllUndesignatedMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
