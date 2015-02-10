package connection

import (
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/projectStorage"
	"PillarsFlowNet/utility"
	"sync"
)

//只查询本公司的项目，这里需要修改ｓｑｌ查询
func GetAllProject(userCode *string, parameter *string) { // ([]byte, *string)
	var errorCode int
	var opResult []utility.Project

	if errorCode == 0 {
		opResult, _ = projectStorage.QueryAllProject()
	}

	command := "getAllProject"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
	//return result, userCode
}

func AddProject(userCode *string, mutex *sync.Mutex, parameter *string) { //([]byte, *string)
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	var projectCode *string
	var projectOut *utility.Project
	if errorCode == 0 {
		project, _ := utility.ParseProjectMessage(parameter)
		project.ProjectCode = *(utility.GenerateCode(userCode))
		projectCode = &(project.ProjectCode)
		opResult, _ := projectStorage.InsertIntoProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = projectStorage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "addProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}

	//return result, userCode
}

func ModifyProject(userCode *string, mutex *sync.Mutex, parameter *string) { //([]byte, *string)
	auth, code := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	} else {
		mutex.Lock()
		defer mutex.Unlock()
	}
	var projectCode *string
	var projectOut *utility.Project
	if errorCode == 0 {
		project, _ := utility.ParseProjectMessage(parameter)
		projectCode = &(project.ProjectCode)
		opResult, _ := projectStorage.ModifyProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = projectStorage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "modifyProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, userCode)
	if auth == true {
		Hub.Dispatch(result, code)
	} else {
		Hub.SendToUserCode(result, userCode)
	}
	//return result, userCode
}
