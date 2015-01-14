package storage

import (
	"testing"
	"PillarsFlowNet/utility"
	"fmt"
)

// type Target struct {
//     TargetCode string
//     MissionCode string
//     ProjectCode string
//     VersionTag string
//     StoragePosition string
//     Picture string
//     InsertDatetime string
//     UpdateDatetime string
// }

func TestInsertIntoDaily(t * testing.T) {
	DBConn = ConnectToDB()

	daily := utility.Daily {
		DailyCode: string("ae0b2f1b208f93586e0ad86cb6f16662"),
		MissionCode: string("a0bb2f1b208f93586e0ad86cb6f16668"),
		ProjectCode: string("e0bb2f1b208f93586e0ad86cb6f16668"),
		VersionTag: string("0.1"),
		StoragePosition: string("/home"),
		Picture: string(""),

	}
	//QueryAllProject()
	result, err := InsertIntoDaily(&daily)
	if err != nil {
		// panic("insert daily error")
		t.Error(err.Error())
	}
	if result == false {
		t.Error("insert Test failed")
	}
	CloseDBConnection()

}

func TestQueryDailysByMissionCode(t * testing.T) {
	DBConn = ConnectToDB()

	missionCode := string("a0bb2f1b208f93586e0ad86cb6f16668")
	dailys, err := QueryDailysByMissionCode(&missionCode)

	if err != nil {
		t.Error(err.Error(), "test query daily failed")
	}

	fmt.Println(*(utility.ObjectToJsonString(dailys)))
	CloseDBConnection()
}

func TestDeleteDailyByDailyCode(t * testing.T) {
	DBConn = ConnectToDB()

	dailyCode := string("ae0b2f1b208f93586e0ad86cb6f16662")
	result, _ := DeleteDailyByDailyCode(&dailyCode)
	if result == false {
		t.Error("delete daily failed")
	}
	CloseDBConnection()
}