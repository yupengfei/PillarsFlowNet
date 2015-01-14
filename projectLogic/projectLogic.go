package projectLogic

import (
	"PillarsFlowNet/projectStorage"
	"PillarsFlowNet/utility"
)

func GetAllProject(userCode * string, parameter * string, h * utility.HubStruct) ([] byte, *string) {
	var errorCode int
	var opResult [] utility.Project

	if (errorCode == 0) {
		opResult, _ = projectStorage.QueryAllProject()
	}

	command := "getAllProject"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
	return result, userCode
}

func AddProject(userCode * string, parameter * string, h * utility.HubStruct) ([] byte, *string) {
	var errorCode int
	var projectCode * string
	var projectOut * utility.Project
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(parameter)
		project.ProjectCode = *(utility.GenerateCode(userCode))
		projectCode = &(project.ProjectCode)
		opResult, _ :=projectStorage.InsertIntoProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = projectStorage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "addProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, userCode)
	h.Dispatch(result)
	return result, userCode
}


func ModifyProject(userCode * string, parameter * string, h * utility.HubStruct) ([] byte, *string) {
	var errorCode int
	var projectCode * string
	var projectOut * utility.Project
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(parameter)
		projectCode = &(project.ProjectCode)
		opResult, _ :=projectStorage.ModifyProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = projectStorage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "modifyProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, userCode)
	h.Dispatch(result)
	return result, userCode
}

