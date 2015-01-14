package daily

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

//向Daily表中增加一条记录
//inputParameters[0]为发起该操作的用户的usercode
//inputParameters[1]为具体的参数
func AddDaily(userCode * string, parameter * string, h * connection.HubStruct) {
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
		opResult, _ :=storage.InsertIntoDaily(daily)
		if opResult == false {
			errorCode = 1
		} else {
			dailyOut, _ = storage.QueryDailyByDailyCode(dailyCode)
		}
	}
	var command = "addDaily"
	result := utility.BoolResultToOutMessage(&command, dailyOut, errorCode, userCode)
	h.Dispatch(result)
}

//修改Daily表中的某一条数据
//inputParameters[0]为发起该操作的用户的UserCode
//inputParameters[1]为具体的参数
func ModifyDaily(userCode * string, parameter * string, h * connection.HubStruct) {
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
		opResult, _ :=storage.ModifyDaily(daily)
		if opResult == false {
			errorCode = 1
		} else {
			dailyOut, _ = storage.QueryDailyByDailyCode(dailyCode)
		}
	}
	var command = "modifyDaily"
	result := utility.BoolResultToOutMessage(&command, dailyOut, errorCode, userCode)
	h.Dispatch(result)
}

//删除某条Daily
//inputParameters[0]为发起该操作的用户的UserCOde
//inputParameters[1]为具体的参数
func DeleteDaily(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		dailyCode, _ := utility.ParseDailyCodeMessage(parameter)
		opResult, _ :=storage.DeleteDailyByDailyCode(&(dailyCode.DailyCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteDaily"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	h.Dispatch(result)
}

//获取missionCode相关的所有Daily
//inputParameters[0]为发起该操作的用户的code
//inputParameters[1]为具体的参数
func QueryDailyByMissionCode(userCode * string, parameter * string, h * connection.HubStruct) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult [] utility.Daily
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(parameter)
		opResult, _ =storage.QueryDailysByMissionCode(&(missionCode.MissionCode))

	}
	command := "queryDailyByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
}