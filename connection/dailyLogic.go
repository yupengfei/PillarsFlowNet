package connection

import (
	"PillarsFlowNet/dailyStorage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
)

//向Daily表中增加一条记录
//inputParameters[0]为发起该操作的用户的usercode
//inputParameters[1]为具体的参数
func AddDaily(userCode * string, parameter * string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var dailyCode * string
	var dailyOut * utility.Daily
	if (errorCode == 0) {
		daily, _ := utility.ParseDailyMessage(parameter)
		daily.DailyCode = *(utility.GenerateCode(userCode))
		dailyCode = &(daily.DailyCode)
		opResult, _ :=dailyStorage.InsertIntoDaily(daily)
		if opResult == false {
			errorCode = 1
		} else {
			dailyOut, _ = dailyStorage.QueryDailyByDailyCode(dailyCode)
		}
	}
	var command = "addDaily"
	result := utility.BoolResultToOutMessage(&command, dailyOut, errorCode, userCode)
	Hub.Dispatch(result)
}

//修改Daily表中的某一条数据
//inputParameters[0]为发起该操作的用户的UserCode
//inputParameters[1]为具体的参数
func ModifyDaily(userCode * string, parameter * string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var dailyCode * string
	var dailyOut * utility.Daily
	if (errorCode == 0) {
		daily, _ := utility.ParseDailyMessage(parameter)
		dailyCode = &(daily.DailyCode)
		opResult, _ :=dailyStorage.ModifyDaily(daily)
		if opResult == false {
			errorCode = 1
		} else {
			dailyOut, _ = dailyStorage.QueryDailyByDailyCode(dailyCode)
		}
	}
	var command = "modifyDaily"
	result := utility.BoolResultToOutMessage(&command, dailyOut, errorCode, userCode)
	Hub.Dispatch(result)
}

//删除某条Daily
//inputParameters[0]为发起该操作的用户的UserCOde
//inputParameters[1]为具体的参数
func DeleteDaily(userCode * string, parameter * string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		dailyCode, _ := utility.ParseDailyCodeMessage(parameter)
		opResult, _ :=dailyStorage.DeleteDailyByDailyCode(&(dailyCode.DailyCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteDaily"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	Hub.Dispatch(result)
}

//获取missionCode相关的所有Daily
//inputParameters[0]为发起该操作的用户的code
//inputParameters[1]为具体的参数
func GetDailyByMissionCode(userCode * string, parameter * string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult [] utility.Daily
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ =dailyStorage.QueryDailysByMissionCode(&(missionCode.MissionCode))

	}
	command := "queryDailyByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}