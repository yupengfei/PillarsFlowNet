package project

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func GetAllProject() [] byte {
	var sysError = utility.Error {
						ErrorCode: 0,
						ErrorMessage: "",
					}
	projects, _:= storage.QueryAllProject()
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllProject",
						Result: *utility.ObjectToJsonString(projects),
					}
	var result = utility.ObjectToJsonByte(&out)
	return result
}
// {
// 	“command”:”addProject”,
// 	“parameter”:”{
// 		“ProjectCode”: 任意string，不起作用,可以没有
// 		“ProjectName”: string,
//     		“ProjectDetail": string,
//     		“PlanBeginDatetime”: string,
//     		“PlanEndDatetime”: string,
//     		“RealBeginDatetime”: string,
//     		“RealEndDatetime”: string,
//     		“PersonInCharge”: string,
//     		“Status”: int,
//     		“Picture”: string,
// 		“InsertDatetime”: 任意string，不起作用,可以没有
//     		“UpdateDatetime”: 任意string，不起作用,可以没有
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “addProject”,
// 	“result”:”{
		// “ProjectCode”: string
		// “ProjectName”: string,
  //   		“ProjectDetail": string,
  //   		“PlanBeginDatetime”: string,
  //   		“PlanEndDatetime”: string,
  //   		“RealBeginDatetime”: string,
  //   		“RealEndDatetime”: string,
  //   		“PersonInCharge”: string,
  //   		“Status”: int,
  //   		“Picture”: string,
		// “InsertDatetime”: string
  //   		“UpdateDatetime”: string
// 	}”
// }
//userCode + @ + parameter
func AddProject(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	var projectCode * string
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(&(inputParameters[1]))
		project.ProjectCode = *(utility.GenerateCode(&(inputParameters[0])))
		projectCode = &(project.ProjectCode)
		opResult, _ :=storage.InsertIntoProject(project)
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
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addProject",
						UserCode: inputParameters[0],
						Result: "{}",
					}
		out = & tempout
	} else {
		projectOut, _ := storage.QueryProjectByProjectCode(projectCode)
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addProject",
						UserCode: inputParameters[0],
						Result: *utility.ObjectToJsonString(projectOut) ,
					}
		out = & tempout
	}

	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// {
// 	“command”:”modifyProject”,
// 	“parameter”:”{
// 		“ProjectCode”: 任意string，不起作用,可以没有
// 		“ProjectName”: string,
//     		“ProjectDetail": string,
//     		“PlanBeginDatetime”: string,
//     		“PlanEndDatetime”: string,
//     		“RealBeginDatetime”: string,
//     		“RealEndDatetime”: string,
//     		“PersonInCharge”: string,
//     		“Status”: int,
//     		“Picture”: string
// 		“InsertDatetime”: 任意string，不起作用,可以没有
//     		“UpdateDatetime”: 任意string，不起作用,可以没有
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “modifyProject”,
// 	“result”:”{
		// “ProjectCode”: string
		// “ProjectName”: string,
  //   		“ProjectDetail": string,
  //   		“PlanBeginDatetime”: string,
  //   		“PlanEndDatetime”: string,
  //   		“RealBeginDatetime”: string,
  //   		“RealEndDatetime”: string,
  //   		“PersonInCharge”: string,
  //   		“Status”: int,
  //   		“Picture”: string,
		// “InsertDatetime”: string
  //   		“UpdateDatetime”: string
// 	}”
// }
func ModifyProject(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var projectCode * string
	if (errorCode == 0) {
		project, _ := utility.ParseProjectMessage(&(inputParameters[1]))
		projectCode = &(project.ProjectCode)
		opResult, _ :=storage.ModifyProject(project)
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
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "modifyProject",
						UserCode: inputParameters[0],
						Result: "{}",
					}
		out = & tempout
	} else {
		projectOut, _ := storage.QueryProjectByProjectCode(projectCode)
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "modifyProject",
						UserCode: inputParameters[0],
						Result: *utility.ObjectToJsonString(projectOut) ,
					}
		out = & tempout
	}

	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

