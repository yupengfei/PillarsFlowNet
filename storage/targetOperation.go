package storage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)
//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry 


//insert is a Transaction
func InsertIntoTarget(target * utility.Target) bool {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("INSERT INTO target(target_code, mission_code, version_tag, storage_position, picture) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		// fmt.Print(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(target.TargetCode, target.MissionCode, target.VersionTag, target.StoragePosition, target.Picture)
	if err != nil {
		pillarsLog.Logger.Print(err.Error())
		return false
	}
	//insert return Result, it does not have interface Close
	//query return Rows ,which must be closed
	err = tx.Commit()
	if err != nil {
		pillarsLog.Logger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			pillarsLog.Logger.Panic(err.Error())
		}
		return false
	}
	return true
}

func QueryTargetsByMissionCode(missionCode * string) []] utility.Target{
	

	stmt, err := DBConn.Prepare("SELECT target_code, mission_code, version_tag, storage_position, picture, insert_datetime, update_datetime FROM target WHERE mission_code = ")
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
	if result.Next() {
		var target utility.Target
		err = result.Scan(&(target.TargetCode), &(target.MissionCode), &(target.VersionTag), &(target.StoragePosition),
		&(target.Picture), &(target.InsertDatetime), &(target.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		targetSlice := append(targetSlice, target)
	}
	return targetSlice

}
