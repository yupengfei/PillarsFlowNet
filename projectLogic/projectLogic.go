package project

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func GetAllProject(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult [] utility.Project

	if (errorCode == 0) {
		opResult, _ = storage.QueryAllProject()
	}

	command := "getAllProject"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))

	return result, &(inputParameters[0])
}

//userCode + @ + parameter
func AddProject(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	
	if (auth == false) {
		errorCode = 3
	}
	var projectCode * string
	var projectOut * utility.Project
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(&(inputParameters[1]))
		project.ProjectCode = *(utility.GenerateCode(&(inputParameters[0])))
		projectCode = &(project.ProjectCode)
		opResult, _ :=storage.InsertIntoProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = storage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "addProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}


func ModifyProject(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var projectCode * string
	var projectOut * utility.Project
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(&(inputParameters[1]))
		projectCode = &(project.ProjectCode)
		opResult, _ :=storage.ModifyProject(project)
		if opResult == false {
			errorCode = 1
		} else {
			projectOut, _ = storage.QueryProjectByProjectCode(projectCode)
		}
	}
	var command = "modifyProject"
	result := utility.BoolResultToOutMessage(&command, projectOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

