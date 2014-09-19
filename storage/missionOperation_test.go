package storage

import (
	"testing"
	"PillarsFlowNet/utility"
	"fmt"
)
// func InsertIntoMission(mission * utility.Mission) bool {
func TestInsertIntoMission(t * testing.T) {
	DBConn = ConnectToDB()
	missionName := string("modify the test")
	missionCode := "a115313c765a01505acd6a5260c7d1ef"

	//projectName := string("very good project")
	projectCode := string("d655313c765a01505acd6a5260c7d1ef")
	mission := utility.Mission{
		MissionCode: missionCode,
		MissionName: missionName,
		ProjectCode: projectCode,
		ProductType: "test ProductType",
		MissionType: "test Mission type",
		MissionDetail: "test mission detail",
		PlanBeginDatetime: "2014-07-15 00:12:00",
		PlanEndDatetime: "2014-08-29 00:05:00",
		RealBeginDatetime: "2014-07-15 00:12:00",
		RealEndDatetime: "2014-08-29 00:05:00",
		PersonIncharge: "yupengfei",
		Status: 1,
		Picture: "123",
		Width: 10,
		Height: 10,
		XCoordinate: 100,
		YCoordinate:100,
	}
	result, _ := InsertIntoMission(&mission)
	if (result == false) {
		t.Error("insert mission failed")
	}
	CloseDBConnection()
}



func TestQueryMissionByMissionCode(t * testing.T) {
	DBConn = ConnectToDB()
	missionCode := string("d655313c765a01505acd6a5260c7d1ea")
	mission, err := QueryMissionByMissionCode(&missionCode)
	if (err != nil) {
		t.Error("query mission failed")
	}
	fmt.Println(*utility.ObjectToJsonString(mission))
	CloseDBConnection()
}

func TestQueryMissionsByProjectCode(t * testing.T) {
	DBConn = ConnectToDB()
	projectCode := string("d655313c765a01505acd6a5260c7d1ef")
	missions, err := QueryMissionsByProjectCode(&projectCode)
	if (err != nil) {
		t.Error("query missions by projectCode failed")
	}
	fmt.Println(*utility.ObjectToJsonString(missions))
	CloseDBConnection()
}

func TestDeleteMissionByMissionCode(t * testing.T) {
	DBConn = ConnectToDB()
	missionCode := "a115313c765a01505acd6a5260c7d1ef"
	result, _ := DeleteMissionByMissionCode(&missionCode)
	if result == false {
		t.Error("delete mission failed")
	}
	CloseDBConnection()
}