package missionStorage

import (
	"testing"
	"PillarsFlowNet/utility"
	"fmt"
)
// func InsertIntoMission(mission * utility.Mission) bool {
func TestInsertIntoMission(t * testing.T) {
	missionName := string("modify the test")
	missionCode := "a115313c765a01505acd6a5260c7d0ea"

	//projectName := string("very good project")
	projectCode := string("d655313c765a01505acd6a5260c7d1ee")
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
		// Width: 10,
		// Height: 10,
		// XCoordinate: 100,
		// YCoordinate:100,
	}
	result, _ := InsertIntoMission(&mission)
	if (result == false) {
		t.Error("insert mission failed")
	}
}



func TestQueryMissionByMissionCode(t * testing.T) {
	missionCode := string("d655313c765a01505acd6a5260c7d1ea")
	mission, err := QueryMissionByMissionCode(&missionCode)
	if (err != nil) {
		t.Error("query mission failed")
	}
	fmt.Println(*utility.ObjectToJsonString(mission))
}

func TestQueryMissionsByProjectCode(t * testing.T) {
	projectCode := string("d655313c765a01505acd6a5260c7d1ee")
	missions, err := QueryMissionsByProjectCode(&projectCode)
	if (err != nil) {
		t.Error("query missions by projectCode failed")
	}
	fmt.Println(*utility.ObjectToJsonString(missions))
}

func TestDeleteMissionByMissionCode(t * testing.T) {
	missionCode := "a115313c765a01505acd6a5260c7d0ea"
	result, _ := DeleteMissionByMissionCode(&missionCode)
	if result == false {
		t.Error("delete mission failed")
	}
}