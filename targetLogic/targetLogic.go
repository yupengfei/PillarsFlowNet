package target

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func AddTarget(userCode * string, parameter * string) ([] byte, *string) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var targetCode * string
	var targetOut * utility.Target
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(parameter)
		target.TargetCode = *(utility.GenerateCode(userCode))
		targetCode = &(target.TargetCode)
		opResult, _ :=storage.InsertIntoTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = storage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "addTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, userCode)
	return result, userCode
}


func ModifyTarget(userCode * string, parameter * string) ([] byte, *string) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var targetCode * string
	var targetOut * utility.Target
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(parameter)
		targetCode = &(target.TargetCode)
		opResult, _ :=storage.ModifyTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = storage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "modifyTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, userCode)
	return result, userCode
}

func DeleteTarget(userCode * string, parameter * string) ([] byte, *string) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		targetCode, _ := utility.ParseTargetCodeMessage(parameter)
		opResult, _ :=storage.DeleteTargetByTargetCode(&(targetCode.TargetCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteTarget"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	return result, userCode
}


func QueryTargetByMissionCode(userCode * string, parameter * string) ([] byte, *string) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult [] utility.Target
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ =storage.QueryTargetsByMissionCode(&(missionCode.MissionCode))

	}
	command := "queryTargetByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	return result, userCode
}