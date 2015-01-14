package storage

import (
	"PillarsFlowNet/utility"
	// "PillarsFlowNet/pillarsLog"
	"testing"
	"fmt"
)


func TestInsertIntoDependency(t * testing.T) {
	DBConn = ConnectToDB()
	dependencyCode := string("a11e99a9e0c8b37a2622e6752117cf93")
	startMissionCode := string("b00e99a9e0c8b37a2622e6752117c455")
	endMissionCode := string("d9e84adb0d94c409d1359177a88fef74")
	projectCode := string("d43e421cb4f0d2cd6a91f309facf944b")
	dependency := utility.Dependency{
		DependencyCode: dependencyCode,
		ProjectCode: projectCode,
		StartMissionCode: startMissionCode,
		EndMissionCode: endMissionCode,
		DependencyType: 1,
	}
	result, _ := InsertIntoDependency(&dependency)
	if result == false {
		t.Error("insert dependency")
	}
	CloseDBConnection()
}

func TestQueryDependenciesByProjectCode(t * testing.T) {
	DBConn = ConnectToDB()
	dependencyCode := string("d43e421cb4f0d2cd6a91f309facf944b")
	dependencies, err := QueryDependenciesByProjectCode(&dependencyCode)
	if err != nil {
		t.Error("query dependency failed")
	} else {
		fmt.Println(*utility.ObjectToJsonString(dependencies))
	}

	CloseDBConnection()
}

func TestDeleteDependencyByDependencyCode(t * testing.T) {
	DBConn = ConnectToDB()
	dependencyCode := string("a11e99a9e0c8b37a2622e6752117cf93")
	result, _ := DeleteDependencyByDependencyCode(&dependencyCode)
	if result == false {
		t.Error("delete dependency failed")
	}
	CloseDBConnection()
}