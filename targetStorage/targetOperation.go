package targetStorage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/mysqlUtility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)
//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry 


//insert is a Transaction
func InsertIntoTarget(target * utility.Target) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO target(target_code, mission_code, project_code, version_tag, storage_position, picture) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(target.TargetCode, target.MissionCode, target.ProjectCode, target.VersionTag, target.StoragePosition, target.Picture)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		panic(err.Error())
	}
	return true, err
}

func ModifyTarget(target * utility.Target) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE target SET mission_code=?, project_code=?, version_tag=?, storage_position=?, picture=? WHERE target_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(target.MissionCode, target.ProjectCode, target.VersionTag, target.StoragePosition, target.Picture, target.TargetCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		panic(err.Error())
	}
	return true, err
}

func DeleteTargetByTargetCode(targetCode * string) (bool, error){
	stmt, err := mysqlUtility.DBConn.Prepare("DELETE FROM target WHERE target_code = ?")
	defer stmt.Close()
	_, err = stmt.Exec(targetCode)
	if err != nil {
		pillarsLog.PillarsLogger.Print(err.Error())
		panic(err.Error())
	}
	return true, err
}

func QueryTargetsByMissionCode(missionCode * string) ([] utility.Target, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT target_code, mission_code, project_code, version_tag, storage_position, picture, insert_datetime, update_datetime FROM target WHERE mission_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(missionCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var targetSlice [] utility.Target
	for result.Next() {
		var target utility.Target
		err = result.Scan(&(target.TargetCode), &(target.MissionCode), &(target.ProjectCode), &(target.VersionTag), &(target.StoragePosition),
		&(target.Picture), &(target.InsertDatetime), &(target.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		targetSlice = append(targetSlice, target)
	}
	return targetSlice, err
}

func QueryTargetByTargetCode(targetCode * string) (* utility.Target, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT target_code, mission_code, project_code, version_tag, storage_position, picture, insert_datetime, update_datetime FROM target WHERE target_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(targetCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var target utility.Target
	if result.Next() {
		err = result.Scan(&(target.TargetCode), &(target.MissionCode), &(target.ProjectCode), &(target.VersionTag), &(target.StoragePosition),
		&(target.Picture), &(target.InsertDatetime), &(target.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &target, err
}




