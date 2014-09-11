package project

import (
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"fmt"
)

func GetAllMissionsOfProject() (parameter * string) (* string, * string, error) {
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

// `mission_code` char(32) not null unique,
// 	`mission_x` int,
// 	`mission_y` int,
// 	`mission_width` int,
// 	`mission_height` int,
// 	`mission_name` char(50) NOT NULL,
// 	`project_code` char(32) NOT NULL,#任务所属项目
// 	`product_type` char(50) NOT NULL,#标识是资产还是镜头
// 	`mission_type` char(50) NOT NULL,#标识更具体的
// 	`mission_detail` varchar(200) ,
// 	`plan_begin_datetime` datetime,
// 	`plan_end_datetime` datetime,
// 	`real_begin_datetime` datetime,
// 	`real_end_datetime` datetime,
// 	`person_in_charge` char(32),#存储`user_code`
// 	`status` int default 0, #0未开始，1已经完成,2已经通过
// 	`picture` text,#照片的base64编码
// {
//	"Command": "AddMission"
//	"Parameter": "{
//		"UserName": "er.wang"
//		"ProjectCode": "xjlkfdjlkjlk"
//		
//	}"
// }
//result, userName, error 
func AddMission(parameter * string) (* string, * string, error) {
	var result string
	var userName string
	var err error

	fmt.Print(*parameter)

	result = "success"
	userName = "er.wang"

	return &result, &userName, err
}

//result, userName, error 
func ModifyMission(parameter * string) (* string, * string, error) {
	var result string
	var userName string
	var err error

	fmt.Print(*parameter)

	result = "success"
	userName = "er.wang"

	return &result, &userName, err
}