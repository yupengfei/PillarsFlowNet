package dependency

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

//获取特定战役所有的依赖
//inputParameters[0]为发起该操作的用户的usercode
//inputParameters[1]为具体的参数，即战役的code
//TODO
//将该函数改名为GetCampaignDependency
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
	}
	command := "getAllDependency"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}

//增加依赖
//inputParameters[0]为发起该操作的用户的UserCode
//inputParameters[1]为具体的参数
func AddDependency(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var dependencyCode * string
	var dependencyOut * utility.Dependency
	if (errorCode == 0) {
		dependency, _ := utility.ParseDependencyMessage(&(inputParameters[1]))
		dependency.DependencyCode = *(utility.GenerateCode(&(inputParameters[0])))
		dependencyCode = &(dependency.DependencyCode)
		opResult, _ :=storage.InsertIntoDependency(dependency)
		if opResult == false {
			errorCode = 1
		} else {
			dependencyOut, _ = storage.QueryDependencyByDependencyCode(dependencyCode)
		}
	}
	var command = "addDependency"
	result := utility.BoolResultToOutMessage(&command, dependencyOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

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

	var command = "deleteDependency"
	result := utility.StringResultToOutMessage(&command, &inputParameters[1], errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

func ModifyDependency(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var dependencyCode * string
	var dependencyOut * utility.Dependency
	if (errorCode == 0) {
		dependency, _ := utility.ParseDependencyMessage(&(inputParameters[1]))
		dependencyCode = &(dependency.DependencyCode)
		opResult, _ :=storage.ModifyDependency(dependency)
		if opResult == false {
			errorCode = 1
		} else {
			dependencyOut, _ = storage.QueryDependencyByDependencyCode(dependencyCode)
		}
	}
	var command = "modifyDependency"
	result := utility.BoolResultToOutMessage(&command, dependencyOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

