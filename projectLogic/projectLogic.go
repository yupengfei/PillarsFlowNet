package project

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/connection"
	// "fmt"
	"strings"
)

func GetAllProject(userCode * string, parameter * string, h * connection.HubStruct) ([] byte, *string) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult [] utility.Project

	if (errorCode == 0) {
		opResult, _ = storage.QueryAllProject()
	}

	command := "getAllProject"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	h.SendToUserCode(result, userCode)
	return result, userCode
}

func AddProject(userCode * string, parameter * string, h * connection.HubStruct) ([] byte, *string) {
	var errorCode int
	
	if (auth == false) {
		errorCode = 3
	}
	var projectCode * string
	var projectOut * utility.Project
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(parameter)
		project.ProjectCode = *(utility.GenerateCode(userCode))
		projectCode = &(project.ProjectCode)
		opResult, _ :=storage.InsertIntoProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = storage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "addProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, userCode)
	h.Dispatch(result)
	return result, userCode
}


func ModifyProject(userCode * string, parameter * string, h * connection.HubStruct) ([] byte, *string) {
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var projectCode * string
	var projectOut * utility.Project
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(parameter)
		projectCode = &(project.ProjectCode)
		opResult, _ :=storage.ModifyProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = storage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "modifyProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, userCode)
	h.Dispatch(result)
	return result, userCode
}

