package dependency

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)
// 获取Campaign的全部dependency
// {
// 	“command”:”getAllDependency”,
// 	“parameter”:”{
// 		CampaignCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “getAllDependency”,
// 	“result”:”[{
// 		DependencyCode string
//     		CampaignCode string
//     		ProjectCode string
//     		StartMissionCode string
//     		EndMissionCode string
//     		DependencyType int
//     		InsertDatetime string
//     		UpdateDatetime string
// 	}]”
// }

func GetAllDependency(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Dependency
	if (errorCode == 0) {
		campaignCode, _ := utility.ParseCampaignCodeMessage(&(inputParameters[1]))
		opResult, _ =storage.QueryDependenciesByCampaignCode(&(campaignCode.CampaignCode))
		// if opResult == false {
		// 	errorCode = 1
		// }
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllDependency",
						Result:*utility.ObjectToJsonString(opResult),
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 添加dependency
// {
// 	“command”:”addDependency”,
// 	“parameter”:{
// 		DependencyCode 任意string，不起作用,可以没有
//     		CampaignCode string
//     		ProjectCode string
//     		StartMissionCode string
//     		EndMissionCode string
//     		DependencyType int
//     		InsertDatetime 任意string，不起作用,可以没有
//     		UpdateDatetime 任意string，不起作用,可以没有
// 	}
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “addDependency”,
// 	“result”:”{
		
// 	}”
// }
func AddDependency(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		dependency, _ := utility.ParseDependencyMessage(&(inputParameters[1]))
		dependency.DependencyCode = *(utility.GenerateCode(&(inputParameters[0])))
		opResult, _ :=storage.InsertIntoDependency(dependency)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "addDependency",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 删除dependency
// {
// 	“command”:”deleteDependency”,
// 	“parameter”:”{
// 		DependencyCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “deleteDependency”,
// 	“result”:”{
		
// 	}”
// }

func DeleteDependency(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		dependencyCode, _ := utility.ParseDependencyCodeMessage(&(inputParameters[1]))
		opResult, _ :=storage.DeleteDependencyByDependencyCode(&(dependencyCode.DependencyCode))
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "deleteDependency",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// 修改dependency
// {
// 	“command”:”modifyDependency”,
// 	“parameter”:{
// 		DependencyCode string
//     		CampaignCode string
//     		ProjectCode string
//     		StartMissionCode string
//     		EndMissionCode string
//     		DependencyType int
//     		InsertDatetime 任意string，不起作用,可以没有
//     		UpdateDatetime 任意string，不起作用,可以没有
// 	}
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “addDependency”,
// 	“result”:”{
		
// 	}”
// }

func ModifyDependency(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		dependency, _ := utility.ParseDependencyMessage(&(inputParameters[1]))
		opResult, _ :=storage.ModifyDependency(dependency)
		if opResult == false {
			errorCode = 1
		}
	}

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "addDependency",
						Result: "{}",
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

