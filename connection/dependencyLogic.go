package connection

import (
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/dependencyStorage"
	"PillarsFlowNet/utility"
	"fmt"
	"sync"
)

//获取特定战役所有的依赖
//inputParameters[0]为发起该操作的用户的usercode
//inputParameters[1]为具体的参数，即战役的code
//TODO
//将该函数改名为GetCampaignDependency
func GetCampaignDependency(userCode *string, parameter *string) {
	//auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	//if auth == false {
	//	errorCode = 3
	//}
	var opResult []utility.Dependency
	if errorCode == 0 {
		campaignCode, _ := utility.ParseCampaignCodeMessage(parameter)
		opResult, _ = dependencyStorage.QueryDependenciesByCampaignCode(&(campaignCode.CampaignCode))
	}
	command := "getCampaignDependency"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

//增加依赖
//inputParameters[0]为发起该操作的用户的UserCode
//inputParameters[1]为具体的参数
func AddDependency(userCode *string, mutex *sync.Mutex, parameter *string) {
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	var dependencyCode *string
	var dependencyOut *utility.Dependency
	if errorCode == 0 {
		dependency, _ := utility.ParseDependencyMessage(parameter)
		dependency.DependencyCode = *(utility.GenerateCode(userCode))
		dependencyCode = &(dependency.DependencyCode)
		opResult, err := dependencyStorage.InsertIntoDependency(dependency)
		if opResult == false {
			errorCode = 1
			fmt.Println(err.Error())
		} else {
			dependencyOut, _ = dependencyStorage.QueryDependencyByDependencyCode(dependencyCode)
		}
	}
	var command = "addDependency"
	result := utility.BoolResultToOutMessage(&command, dependencyOut, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}
}

func DeleteDependency(userCode *string, mutex *sync.Mutex, parameter *string) {
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}

	if errorCode == 0 {
		dependencyCode, _ := utility.ParseDependencyCodeMessage(parameter)
		opResult, _ := dependencyStorage.DeleteDependencyByDependencyCode(&(dependencyCode.DependencyCode))
		if opResult == false {
			errorCode = 1
		}
	}

	var command = "deleteDependency"
	result := utility.StringResultToOutMessage(&command, parameter, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}
}

func ModifyDependency(userCode *string, mutex *sync.Mutex, parameter *string) {
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	var dependencyCode *string
	var dependencyOut *utility.Dependency
	if errorCode == 0 {
		dependency, _ := utility.ParseDependencyMessage(parameter)
		dependencyCode = &(dependency.DependencyCode)
		opResult, _ := dependencyStorage.ModifyDependency(dependency)
		if opResult == false {
			errorCode = 1
		} else {
			dependencyOut, _ = dependencyStorage.QueryDependencyByDependencyCode(dependencyCode)
		}
	}
	var command = "modifyDependency"
	result := utility.BoolResultToOutMessage(&command, dependencyOut, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}
}
