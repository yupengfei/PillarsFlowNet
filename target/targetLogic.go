package target

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func AddTarget(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var targetCode * string
	var targetOut * utility.Target
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(&(inputParameters[1]))
		target.TargetCode = *(utility.GenerateCode(&(inputParameters[0])))
		targetCode = &(target.TargetCode)
		opResult, _ :=storage.InsertIntoTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = storage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "addTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}


func ModifyTarget(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var targetCode * string
	var targetOut * utility.Target
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(&(inputParameters[1]))
		targetCode = &(target.TargetCode)
		opResult, _ :=storage.ModifyTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = storage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "modifyTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

func DeleteTarget(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		targetCode, _ := utility.ParseTargetCodeMessage(&(inputParameters[1]))
		opResult, _ :=storage.DeleteTargetByTargetCode(&(targetCode.TargetCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteTarget"
	result := utility.BoolResultToOutMessage(&command, &inputParameters[1], errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}


func QueryTargetByMissionCode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult [] utility.Target
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(&(inputParameters[1]))
		opResult, _ =storage.QueryTargetsByMissionCode(&(missionCode.MissionCode))

	}
	command := "queryTargetByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}