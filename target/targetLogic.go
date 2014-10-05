package target

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

// 添加target
// {
// 	“command”:”addTarget”,
// 	“parameter”:”{
// 		TargetCode 任意string，不起作用,可以没有
//     		MissionCode string
//     		ProjectCode string
//     		VersionTag string
//     		StoragePosition string
//     		Picture string
//     		InsertDatetime 任意string，不起作用,可以没有
//     		UpdateDatetime 任意string，不起作用,可以没有
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “addTarget”,
// 	“result”:”{
// 		TargetCode string
//     		MissionCode string
//     		ProjectCode string
//     		VersionTag string
//     		StoragePosition string
//     		Picture string
//     		InsertDatetime string
//     		UpdateDatetime string
// 	}”
// }

func AddTarget(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var targetCode * string
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(&(inputParameters[1]))
		target.TargetCode = *(utility.GenerateCode(&(inputParameters[0])))
		targetCode = &(target.TargetCode)
		opResult, _ :=storage.InsertIntoTarget(target)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out *  utility.OutMessage
	if errorCode != 0 {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "addTarget",
						Result: "{}",
					}
		out = & tempout
	} else {
		targetOut, _ := storage.QueryTargetByTargetCode(targetCode)
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "addTarget",
						Result:*utility.ObjectToJsonString(targetOut),
					}
		out = & tempout
	}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 修改target
// {
// 	“command”:”modifyTarget”,
// 	“parameter”:”{
// 		TargetCode string
//     		MissionCode string
//     		ProjectCode string
//     		VersionTag string
//     		StoragePosition string
//     		Picture string
//     		InsertDatetime 任意string，不起作用,可以没有
//     		UpdateDatetime 任意string，不起作用,可以没有
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “modifyTarget”,
// 	“result”:”{
// 		TargetCode string
//     		MissionCode string
//     		ProjectCode string
//     		VersionTag string
//     		StoragePosition string
//     		Picture string
//     		InsertDatetime string
//     		UpdateDatetime string

// 	}”
// }

func ModifyTarget(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var targetCode * string
	if (errorCode == 0) {
		target, _ := utility.ParseTargetMessage(&(inputParameters[1]))
		targetCode = &(target.TargetCode)
		opResult, _ :=storage.ModifyTarget(target)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * utility.OutMessage
	if errorCode != 0 {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "modifyTarget",
						Result: "{}",
					}
		out = & tempout
	} else {
		targetOut, _ := storage.QueryTargetByTargetCode(targetCode)
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "modifyTarget",
						Result:*utility.ObjectToJsonString(targetOut),
					}
		out = & tempout
	}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 删除target
// {
// 	”command“:”deleteTarget”,
// 	“parameter”:”{
// 		TargetCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “deleteTarget”,
// 	“result”:”{
// 		TargetCode string
// 	}”
// }

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

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out *  utility.OutMessage
	if errorCode != 0 {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "deleteTarget",
						Result: "{}",
					}
		out = & tempout
	} else {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "deleteTarget",
						Result:inputParameters[1],
					}
		out = & tempout
	}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 查询指定mission的target
// {
// 	”command“:”searchTargetByMissionCode”,
// 	“parameter”:”{
// 		MissionCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “searchTargetByMissionCode”,
// 	“result”:”[{
// 		TargetCode string
//     		MissionCode string
//     		ProjectCode string
//     		VersionTag string
//     		StoragePosition string
//     		Picture string
//     		InsertDatetime string
//     		UpdateDatetime string
// 	}]”
// }

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

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "addTarget",
						Result: *utility.ObjectToJsonString(opResult),
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}