package storage

import (
	"PillarsFlowNet/utility"
	// "PillarsFlowNet/pillarsLog"
	"testing"
	// "fmt"
)

func TestInsertIntoGraph(t * testing.T) {
	DBConn = ConnectToDB()
	graphCode := string("a11e99a9e0c8b37a2622e6752117cf96")
	campaignCode := string("a11e99a9e0c8b37a2622e6752117cf92")
	projectCode := string("d43e421cb4f0d2cd6a91f309facf944b")
	nodeCode := string("b00e99a9e0c8b37a2622e6752117c453")
	width := 10
	height := 10
	xCoordinate := 100
	yCoordinate := 100
	graph := utility.Graph {
		GraphCode: graphCode,
		CampaignCode: campaignCode,
		ProjectCode: projectCode,
		NodeCode: nodeCode,
		Width: width,
		Height: height,
		XCoordinate: xCoordinate,
		YCoordinate: yCoordinate,
	}
	result, _ := InsertIntoGraph(&graph)
	if result == false {
		t.Error("insert graph failed")
	}
	CloseDBConnection()
}

// func TestInsertIntoGraph2(t * testing.T) {
// 	DBConn = ConnectToDB()
// 	graphCode := string("a11e99a9e0c8b37a2622e6752117cf94")
// 	campaignCode := string("a11e99a9e0c8b37a2622e6752117cf93")
// 	projectCode := string("d43e421cb4f0d2cd6a91f309facf944b")
// 	nodeCode := string("b00e99a9e0c8b37a2622e6752117c454")
// 	width := 10
// 	height := 10
// 	xCoordinate := 200
// 	yCoordinate := 200
// 	graph := utility.Graph {
// 		GraphCode: graphCode,
// 		CampaignCode: campaignCode,
// 		ProjectCode: projectCode,
// 		NodeCode: nodeCode,
// 		Width: width,
// 		Height: height,
// 		XCoordinate: xCoordinate,
// 		YCoordinate: yCoordinate,
// 	}
// 	result, _ := InsertIntoGraph(&graph)
// 	if result == false {
// 		t.Error("insert graph failed")
// 	}
// 	CloseDBConnection()
// }

func TestDeleteGraphByGraphCode(t * testing.T) {
	DBConn = ConnectToDB()
	graphCode := string("a11e99a9e0c8b37a2622e6752117cf96")
	result, _ := DeleteGraphByGraphCode(&graphCode)
	if result == false {
		t.Error("DeleteGraphByGraphCode failed")
	}
	CloseDBConnection()
}

func TestDeleteNodeByNodeCode(t * testing.T) {
	DBConn = ConnectToDB()
	nodeCode := string("a11e99a9e0c8b37a2622e6752117cf94")
	result, _ := DeleteNodeByNodeCode(&nodeCode)
	if result == false {
		t.Error("DeleteNodeByNodeCode failed")
	}
	CloseDBConnection()
}