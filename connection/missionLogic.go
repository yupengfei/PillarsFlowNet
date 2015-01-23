package connection

import (
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/missionStorage"
	"PillarsFlowNet/utility"
)

//获取某个Project所有的Campaign
//TODO
//将该函数改名为GetProjectCampaign
func GetProjectCampaign(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var opResult []utility.Mission
	if errorCode == 0 {
		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
		opResult, _ = missionStorage.QueryCampaignsByProjectCode(&(projectCode.ProjectCode))
	}
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

func GetProjectAssertCampaign(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var opResult []utility.Mission
	if errorCode == 0 {
		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
		opResult, _ = missionStorage.QueryAssertCampaignsByProjectCode(&(projectCode.ProjectCode))
	}
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
func GetProjectShotCampaign(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var opResult []utility.Mission
	if errorCode == 0 {
		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
		opResult, _ = missionStorage.QueryShotCampaignsByProjectCode(&(projectCode.ProjectCode))
	}
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
func GetProjectUnassertCampaign(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var opResult []utility.Mission
	if errorCode == 0 {
		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
		opResult, _ = missionStorage.QueryUnassertCampaignsByProjectCode(&(projectCode.ProjectCode))
	}
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

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
	command := "getPersonAllFinishedMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
