package storage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)
//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry 


//insert is a Transaction
func InsertIntoMission(mission * utility.Mission) bool {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("INSERT INTO mission(mission_code, mission_name, project_code, product_type, mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		// fmt.Print(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(mission.MissionCode, mission.MissionName, mission.ProjectCode,
		mission.ProductType, mission.MissionType, mission.MissionDetail,
		mission.PlanBeginDatetime, mission.PlanEndDatetime, mission.RealBeginDatetime, 
		mission.RealEndDatetime, mission.PersonInCharge,
		mission.Status, mission.Picture)
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

func QueryMissionByMissionCode(missionCode * string) * utility.Mission {
	stmt, err := DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM mission WHERE project_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var mission utility.Project
	if result.Next() {
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.MissionType), &(mission.MissionDetail)
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonInCharge), &(mission.Status), &(mission.Picture), &(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
	}
	return &mission
}

func QueryMissionsByProjectCode(project_code * string) [] utility.Mission {
	stmt, err := DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM project where project_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var missionSlice [] utility.Mission
	for result.Next() {
		var mission utility.Mission
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.MissionType), &(mission.MissionDetail)
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonInCharge), &(mission.Status), &(mission.Picture), &(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		missionSlice := append(missionSlice, project)
	}
	return missionSlice
}

