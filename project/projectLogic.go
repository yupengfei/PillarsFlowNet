package project

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"fmt"
)

func GetAllProject() [] byte {
	var sysError = utility.Error {
						ErrorCode: 0,
						ErrorMessage: "",
					}
	projects, _ := storage.QueryAllProject()
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllProject",
						Result: *utility.ObjectToJsonString(projects),
					}
	var result = utility.ObjectToJsonByte(out)
	return result
}

//result, userName, error 
func AddProject(parameter * string) (* string, * string, error) {
	var result string
	var userName string
	var err error

	fmt.Print(*parameter)
	

	result = "success"
	userName = "er.wang"

	return &result, &userName, err
}

//result, userName, error 
func ModifyProject(parameter * string) (* string, * string, error) {
	var result string
	var userName string
	var err error

	fmt.Print(*parameter)

	result = "success"
	userName = "er.wang"

	return &result, &userName, err
}