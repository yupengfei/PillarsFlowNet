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
func AddDaily(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var dailyCode * string
	var dailyOut * utility.Daily
	if (errorCode == 0) {
		daily, _ := utility.ParseDailyMessage(&(inputParameters[1]))
		daily.DailyCode = *(utility.GenerateCode(&(inputParameters[0])))
		dailyCode = &(daily.DailyCode)
		opResult, _ :=storage.InsertIntoDaily(daily)
		if opResult == false {
			errorCode = 1
		} else {
			dailyOut, _ = storage.QueryDailyByDailyCode(dailyCode)
		}
	}
	var command = "addDaily"
	result := utility.BoolResultToOutMessage(&command, dailyOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

//修改Daily表中的某一条数据
//inputParameters[0]为发起该操作的用户的UserCode
//inputParameters[1]为具体的参数
func ModifyDaily(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var dailyCode * string
	var dailyOut * utility.Daily
	if (errorCode == 0) {
		daily, _ := utility.ParseDailyMessage(&(inputParameters[1]))
		dailyCode = &(daily.DailyCode)
		opResult, _ :=storage.ModifyDaily(daily)
		if opResult == false {
			errorCode = 1
		} else {
			dailyOut, _ = storage.QueryDailyByDailyCode(dailyCode)
		}
	}
	var command = "modifyDaily"
	result := utility.BoolResultToOutMessage(&command, dailyOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

//删除某条Daily
//inputParameters[0]为发起该操作的用户的UserCOde
//inputParameters[1]为具体的参数
func DeleteDaily(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		dailyCode, _ := utility.ParseDailyCodeMessage(&(inputParameters[1]))
		opResult, _ :=storage.DeleteDailyByDailyCode(&(dailyCode.DailyCode))
		if opResult == false {
			errorCode = 1
		}
	}
	var command = "deleteDaily"
	result := utility.StringResultToOutMessage(&command, &inputParameters[1], errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

//获取missionCode相关的所有Daily
//inputParameters[0]为发起该操作的用户的code
//inputParameters[1]为具体的参数
func QueryDailyByMissionCode(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult [] utility.Daily
	if (errorCode == 0) {
		missionCode, _ := utility.ParseMissionCodeMessage(&(inputParameters[1]))
		opResult, _ =storage.QueryDailysByMissionCode(&(missionCode.MissionCode))

	}
	command := "queryDailyByMissionCode"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}