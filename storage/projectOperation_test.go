package storage

import (
	"testing"	
	"fmt"
	"PillarsFlowNet/utility"
)

func TestQueryAllProject(t * testing.T) {
	DBConn = ConnectToDB()
	projects := QueryAllProject()
	//message, _ := json.Marshal(projects)
	fmt.Println(*utility.ObjectToJsonString(projects))
	CloseDBConnection()
}