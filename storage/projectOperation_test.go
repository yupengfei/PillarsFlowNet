package storage

import (
	"testing"	
	"fmt"
	"PillarsFlowNet/utility"
)

func TestQueryAllProject(t * testing.T) {
	DBConn = ConnectToDB()
	QueryAllProject()
	// projects, _ := QueryAllProject()
	//message, _ := json.Marshal(projects)
	//fmt.Println(*utility.ObjectToJsonString(projects))
	CloseDBConnection()
}

func TestInsertIntoProject(t * testing.T) {
	DBConn = ConnectToDB()
	projectName := string("anthor test")
	projectCode := string("890b2f1b208f93586e0ad86cb6f16668")
	project := utility.Project{
		ProjectCode: projectCode,
		ProjectName: projectName,
		ProjectDetail: "project detail test",
		PlanBeginDatetime: "2014-07-15 00:12:00",
		PlanEndDatetime: "2014-08-29 00:05:00",
		RealBeginDatetime: "2014-07-15 00:12:00",
		RealEndDatetime: "2014-08-29 00:05:00",
		PersonInCharge: "yupengfei",
		Status: 1,
		Picture: "123",
	}
	result, _ := InsertIntoProject(&project)
	if (result == false) {
		t.Error("insert project error")
	}
	CloseDBConnection()
}
// func QueryProjectByProjectCode(projectCode * string) * utility.Project
func TestQueryProjectByProjectCode(t * testing.T) {
	DBConn = ConnectToDB()
	projectCode := string("890b2f1b208f93586e0ad86cb6f16669")
	project, err := QueryProjectByProjectCode(&projectCode)
	if (err != nil) {
		t.Error(err.Error())
	}
	fmt.Println(*utility.ObjectToJsonString(project))
	CloseDBConnection()
}

func TestDeleteProjectByProjectCode(t * testing.T) {
	DBConn = ConnectToDB()
	projectCode := string("890b2f1b208f93586e0ad86cb6f16668")
	result, _ := DeleteProjectByProjectCode(&projectCode)
	if result == false {
		t.Error("delete project failed")
	}
	CloseDBConnection()
}