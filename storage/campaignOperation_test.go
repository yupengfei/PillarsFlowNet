package storage

import (
	"PillarsFlowNet/utility"
	// "PillarsFlowNet/pillarsLog"
	"testing"
	// "fmt"
)

// DBConn = ConnectToDB()
// 	dependencyCode := string("a11e99a9e0c8b37a2622e6752117cf93")
// 	startMissionCode := string("b00e99a9e0c8b37a2622e6752117c453")
// 	endMissionCode := string("d9e84adb0d94c409d1359177a88fef74")
// 	projectCode := string("d43e421cb4f0d2cd6a91f309facf944b")
// 	dependency := utility.Dependency{
// 		DependencyCode: dependencyCode,
// 		ProjectCode: projectCode,
// 		StartMissionCode: startMissionCode,
// 		EndMissionCode: endMissionCode,
// 		DependencyType: "test dep",
// 	}
// 	result, _ := InsertIntoDependency(&dependency)
// 	if result == false {
// 		t.Error("insert dependency")
// 	}
// 	CloseDBConnection()

// type Campaign struct {
//     CampaignCode string
//     ProjectCode string
//     NodeCode string
//     Width int
//     Height int
//     XCoordinate int
//     YCoordinate int
//     InsertDatetime string
//     UpdateDatetime string
// }
func TestInsertIntoCampaign(t * testing.T) {
	DBConn = ConnectToDB()
	campaignCode := string("a11e99a9e0c8b37a2622e6752117cf92")
	projectCode := string("d43e421cb4f0d2cd6a91f309facf944b")
	nodeCode := string("b00e99a9e0c8b37a2622e6752117c453")
	width := 10
	height := 10
	xCoordinate := 100
	yCoordinate := 100
	campaign := utility.Campaign {
		CampaignCode: campaignCode,
		ProjectCode: projectCode,
		NodeCode: nodeCode,
		Width: width,
		Height: height,
		XCoordinate: xCoordinate,
		YCoordinate: yCoordinate,
	}
	result, _ := InsertIntoCampaign(&campaign)
	if result == false {
		t.Error("insert campaign failed")
	}
	CloseDBConnection()
}

func TestInsertIntoCampaign2(t * testing.T) {
	DBConn = ConnectToDB()
	campaignCode := string("a11e99a9e0c8b37a2622e6752117cf93")
	projectCode := string("d43e421cb4f0d2cd6a91f309facf944b")
	nodeCode := string("b00e99a9e0c8b37a2622e6752117c454")
	width := 10
	height := 10
	xCoordinate := 200
	yCoordinate := 200
	campaign := utility.Campaign {
		CampaignCode: campaignCode,
		ProjectCode: projectCode,
		NodeCode: nodeCode,
		Width: width,
		Height: height,
		XCoordinate: xCoordinate,
		YCoordinate: yCoordinate,
	}
	result, _ := InsertIntoCampaign(&campaign)
	if result == false {
		t.Error("insert campaign failed")
	}
	CloseDBConnection()
}

func TestDeleteCampaignByCampaignCode(t * testing.T) {
	DBConn = ConnectToDB()
	campaignCode := string("a11e99a9e0c8b37a2622e6752117cf92")
	result, _ := DeleteCampaignByCampaignCode(&campaignCode)
	if result == false {
		t.Error("DeleteCampaignByCampaignCode failed")
	}
	CloseDBConnection()
}

func TestDeleteNodeByNodeCode(t * testing.T) {
	DBConn = ConnectToDB()
	nodeCode := string("a11e99a9e0c8b37a2622e6752117cf92")
	result, _ := DeleteNodeByNodeCode(&nodeCode)
	if result == false {
		t.Error("DeleteNodeByNodeCode failed")
	}
	CloseDBConnection()
}