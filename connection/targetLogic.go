package connection

import (
	"PillarsFlowNet/targetStorage"
	"PillarsFlowNet/utility"
)

func AddTarget(userCode * string, parameter * string) {
	var errorCode int
	var targetCode * string
	var targetOut * utility.Target
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(parameter)
		target.TargetCode = *(utility.GenerateCode(userCode))
		targetCode = &(target.TargetCode)
		opResult, _ :=targetStorage.InsertIntoTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = targetStorage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "addTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, userCode)
	Hub.Dispatch(result)
}


func ModifyTarget(userCode * string, parameter * string) {
	var errorCode int
	var targetCode * string
	var targetOut * utility.Target
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(parameter)
		targetCode = &(target.TargetCode)
		opResult, _ :=targetStorage.ModifyTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = targetStorage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "modifyTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, userCode)
	Hub.Dispatch(result)
}

func DeleteTarget(userCode * string, parameter * string) {
	var errorCode int
	if (errorCode == 0) {
		targetCode, _ := utility.ParseTargetCodeMessage(parameter)
		opResult, _ :=targetStorage.DeleteTargetByTargetCode(&(targetCode.TargetCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteTarget"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	Hub.Dispatch(result)
}


func GetTargetByMissionCode(userCode * string, parameter * string) {
	var errorCode int
	var opResult [] utility.Target
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ =targetStorage.QueryTargetsByMissionCode(&(missionCode.MissionCode))

	}
	command := "queryTargetByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}