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
func GetProjectCampaign(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
		opResult, _ =storage.QueryCampaignsByProjectCode(&(projectCode.ProjectCode))
	}
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
	return result, userCode
}

func GetProjectAssertCampaign(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
		opResult, _ =storage.QueryAssertCampaignsByProjectCode(&(projectCode.ProjectCode))
	}
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
}

func GetProjectUnassertCampaign(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		projectCode, _ := utility.ParseProjectCodeMessage(parameter)
		opResult, _ =storage.QueryUnassertCampaignsByProjectCode(&(projectCode.ProjectCode))
	}
	command := "getAllCampaign"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
}

func AddMission(userCode * string, parameter * string) ([] byte, *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var missionCode * string
	var missionOut * utility.Mission
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(parameter)
		mission.MissionCode = *(utility.GenerateCode(userCode))
		missionCode = &(mission.MissionCode)
		opResult, _ :=storage.InsertIntoMission(mission)
		if opResult == false {
			errorCode = 1
		} else {
			missionOut, _ = storage.QueryMissionByMissionCode(missionCode)
		}
	}
	command := "addMission"
	result := utility.BoolResultToOutMessage(&command, missionOut, errorCode, userCode)
	h.Dispatch(result)
}

func ModifyMission(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var missionCode * string
	var missionOut * utility.Mission
	if (errorCode == 0) {
		mission, _ := utility.ParseMissionMessage(parameter)
		opResult, _ :=storage.ModifyMission(mission)
		missionCode = &(mission.MissionCode)
		if opResult == false {
			errorCode = 1
		} else {
			missionOut, _ = storage.QueryMissionByMissionCode(missionCode)
		}
	}
	command := "modifyMission"
	result := utility.BoolResultToOutMessage(&command, missionOut, errorCode, userCode)
	h.Dispatch(result)
}

func DeleteMission(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ :=storage.DeleteMissionByMissionCode(&(missionCode.MissionCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteMission"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	h.Dispatch(result)
}

func QueryMissionByMissionCode(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult * utility.Mission
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ =storage.QueryMissionByMissionCode(&(missionCode.MissionCode))
	}
	command := "queryMissionByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}


func GetPersonAllWaitingMission(userCode * string, parameter * string, h * connection.HubStruct) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryWaitingMissionByUserCode(userCode)
	}
	command := "getPersonAllWaitingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}

func GetPersonAllUndergoingMission(userCode * string, parameter * string, h * connection.HubStruct) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryUndergoingMissionByUserCode(userCode)
	}
	command := "getPersonAllUndergoingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}

func GetPersonAllReviewingMission(userCode * string, parameter * string, h * connection.HubStruct) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryReviewingMissionByUserCode(userCode)
	}
	command := "getPersonAllReviewingMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}

func GetPersonAllFinishedMission(userCode * string, parameter * string, h * connection.HubStruct) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryFinishedMissionByUserCode(userCode)
	}
	command := "getPersonAllFinishedMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}

func GetAllUndesignatedMission(userCode * string, parameter * string, h * connection.HubStruct) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Mission
	if (errorCode == 0) {
		opResult, _ =storage.QueryAllUndesignatedMission()
	}
	command := "getPersonAllFinishedMission"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}
