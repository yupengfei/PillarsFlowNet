package connection

import (
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/targetStorage"
	"PillarsFlowNet/utility"
	"sync"
)

func AddTarget(userCode *string, mutex *sync.Mutex, parameter *string) {
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	var targetCode *string
	var targetOut *utility.Target
	if errorCode == 0 {
		target, _ := utility.ParseTargetMessage(parameter)
		target.TargetCode = *(utility.GenerateCode(userCode))
		targetCode = &(target.TargetCode)
		opResult, _ := targetStorage.InsertIntoTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = targetStorage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "addTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}
}

func ModifyTarget(userCode *string, mutex *sync.Mutex, parameter *string) {
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	var targetCode *string
	var targetOut *utility.Target
	if errorCode == 0 {
		target, _ := utility.ParseTargetMessage(parameter)
		targetCode = &(target.TargetCode)
		opResult, _ := targetStorage.ModifyTarget(target)
		if opResult == false {
			errorCode = 1
		} else {
			targetOut, _ = targetStorage.QueryTargetByTargetCode(targetCode)
		}
	}
	var command = "modifyTarget"
	result := utility.BoolResultToOutMessage(&command, targetOut, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}
}

func DeleteTarget(userCode *string, mutex *sync.Mutex, parameter *string) {
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	if errorCode == 0 {
		targetCode, _ := utility.ParseTargetCodeMessage(parameter)
		opResult, _ := targetStorage.DeleteTargetByTargetCode(&(targetCode.TargetCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteTarget"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}
}

func GetTargetByMissionCode(userCode *string, parameter *string) {
	var errorCode int
	var opResult []utility.Target
	if errorCode == 0 {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ = targetStorage.QueryTargetsByMissionCode(&(missionCode.MissionCode))

	}
	command := "getTargetByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
