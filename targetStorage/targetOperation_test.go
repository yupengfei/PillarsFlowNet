package targetStorage

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

func TestInsertIntoTarget(t * testing.T) {
	target := utility.Target {
		TargetCode: string("ae0b2f1b208f93586e0ad86cb6f16662"),
		MissionCode: string("a0bb2f1b208f93586e0ad86cb6f16668"),
		ProjectCode: string("e0bb2f1b208f93586e0ad86cb6f16668"),
		VersionTag: string("0.1"),
		StoragePosition: string("/home"),
		Picture: string(""),

	}
	//QueryAllProject()
	result, err := InsertIntoTarget(&target)
	if err != nil {
		// panic("insert target error")
		t.Error(err.Error())
	}
	if result == false {
		t.Error("insert Test failed")
	}
}

func TestQueryTargetsByMissionCode(t * testing.T) {
	missionCode := string("a0bb2f1b208f93586e0ad86cb6f16668")
	targets, err := QueryTargetsByMissionCode(&missionCode)

	if err != nil {
		t.Error(err.Error(), "test query target failed")
	}

	fmt.Println(*(utility.ObjectToJsonString(targets)))
}

func TestDeleteTargetByTargetCode(t * testing.T) {
	targetCode := string("ae0b2f1b208f93586e0ad86cb6f16662")
	result, _ := DeleteTargetByTargetCode(&targetCode)
	if result == false {
		t.Error("delete target failed")
	}
}