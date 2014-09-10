package project

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
)

func GetAllProject() [] byte {
	var sysError = utility.Error {
						ErrorCode: 0,
						ErrorMessage: "",
					}
	projects := storage.QueryAllProject()
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllProject",
						Result: *utility.ObjectToJsonString(projects),
					}
	var result = utility.ObjectToJson(out)
	return result

}