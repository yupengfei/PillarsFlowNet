package storage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/pillarsLog"
	"fmt"
)
//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry 


//insert is a Transaction
func InsertIntoMission(mission * utility.Mission) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`INSERT INTO mission(mission_code, mission_name, project_code, product_type, is_campaign,
		mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, 
		real_begin_datetime, real_end_datetime, person_in_charge, status, 
		picture) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(mission.MissionCode, mission.MissionName, mission.ProjectCode,
		mission.ProductType, mission.IsCampaign, mission.MissionType, mission.MissionDetail,
		mission.PlanBeginDatetime, mission.PlanEndDatetime, mission.RealBeginDatetime, 
		mission.RealEndDatetime, mission.PersonIncharge,
		mission.Status, mission.Picture)
	if err != nil {
		panic(err.Error())
	}
	//insert return Result, it does not have interface Close
	//query return Rows ,which must be closed
	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		pillarsLog.Logger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			pillarsLog.Logger.Panic(err.Error())
		}
		return false, err
	}
	return true, err
}

//is_campaign is not checked
func ModifyMission(mission * utility.Mission) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`UPDATE mission SET mission_name=?, project_code=?, product_type=?, is_campaign=?,
		mission_type=?, mission_detail=?, plan_begin_datetime=?, plan_end_datetime=?, 
		real_begin_datetime=?, real_end_datetime=?, person_in_charge=?, status=?, 
		picture=? WHERE mission_code=?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(mission.MissionName, mission.ProjectCode,
		mission.ProductType, mission.IsCampaign, mission.MissionType, mission.MissionDetail,
		mission.PlanBeginDatetime, mission.PlanEndDatetime, mission.RealBeginDatetime, 
		mission.RealEndDatetime, mission.PersonIncharge,
		mission.Status, mission.Picture, mission.MissionCode)
	if err != nil {
		panic(err.Error())
	}
	//insert return Result, it does not have interface Close
	//query return Rows ,which must be closed
	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		pillarsLog.Logger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			pillarsLog.Logger.Panic(err.Error())
		}
		return false, err
	}
	return true, err
}

//delete dependencies and targets related to mission
func DeleteMissionByMissionCode(missionCode * string) (bool, error) {
	tx, err := DBConn.Begin()
	//delete mission
	stmt, err := tx.Prepare("DELETE FROM mission WHERE mission_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(missionCode)
	if err != nil {
		panic(err.Error())
	}
	//delete dependencies 
	stmtDependencyFrom, err := tx.Prepare("DELETE FROM dependency WHERE start_mission_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDependencyFrom.Close()
	_, err = stmtDependencyFrom.Exec(missionCode)
	if err != nil {
		panic(err.Error())
	}

	//delete dependencies 
	stmtDependencyTo, err := tx.Prepare("DELETE FROM dependency WHERE end_mission_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtDependencyTo.Close()
	_, err = stmtDependencyTo.Exec(missionCode)
	if err != nil {
		panic(err.Error())
	}

	//delete from campaign 
	stmtCampaignNode, err := tx.Prepare("DELETE FROM graph WHERE node_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtCampaignNode.Close()
	_, err = stmtCampaignNode.Exec(missionCode)
	if err != nil {
		panic(err.Error())
	}

	//delete from campaign 
	stmtCampaignCampaignCode, err := tx.Prepare("DELETE FROM graph WHERE campaign_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtCampaignCampaignCode.Close()
	_, err = stmtCampaignCampaignCode.Exec(missionCode)
	if err != nil {
		panic(err.Error())
	}


	//delete targets
	stmtTarget, err := tx.Prepare("DELETE FROM target WHERE mission_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtTarget.Close()
	_, err = stmtTarget.Exec(missionCode)
	if err != nil {
		panic(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err.Error())
		pillarsLog.Logger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			pillarsLog.Logger.Panic(err.Error())
		}
		return false, err
	}
	return true, err
}



func QueryMissionByMissionCode(missionCode * string) (* utility.Mission, error) {
	stmt, err := DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, is_campaign, mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM mission WHERE mission_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(missionCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var mission utility.Mission
	if result.Next() {
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.IsCampaign), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.RealEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
	}
	return &mission, err
}

func QueryMissionsByProjectCode(projectCode * string) ([] utility.Mission, error) {
	stmt, err := DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, is_campaign, mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM mission where project_code = ?")
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
		&(mission.ProductType), &(mission.IsCampaign), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryCampaignsByProjectCode(projectCode * string) ([] utility.Mission, error) {
	stmt, err := DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, is_campaign, mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM mission where project_code = ? AND is_campaign=1")
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
		&(mission.ProductType), &(mission.IsCampaign), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

