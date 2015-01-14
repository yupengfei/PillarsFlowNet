package missionStorage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/mysqlUtility"
	"PillarsFlowNet/pillarsLog"
	"fmt"
)
//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry 


//insert is a Transaction
func InsertIntoMission(mission * utility.Mission) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO mission(mission_code, mission_name, 
		project_code, product_type, is_campaign, is_assert,
		mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, 
		real_begin_datetime, real_end_datetime, person_in_charge, status, 
		picture) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(mission.MissionCode, mission.MissionName, mission.ProjectCode,
		mission.ProductType, mission.IsCampaign, mission.IsAssert, mission.MissionType, mission.MissionDetail,
		mission.PlanBeginDatetime, mission.PlanEndDatetime, mission.RealBeginDatetime, 
		mission.RealEndDatetime, mission.PersonIncharge,
		mission.Status, mission.Picture)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

//is_campaign is not checked
func ModifyMission(mission * utility.Mission) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE mission SET mission_name=?, project_code=?, 
		product_type=?, is_campaign=?, is_assert=?,
		mission_type=?, mission_detail=?, plan_begin_datetime=?, plan_end_datetime=?, 
		real_begin_datetime=?, real_end_datetime=?, person_in_charge=?, status=?, 
		picture=? WHERE mission_code=?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(mission.MissionName, mission.ProjectCode,
		mission.ProductType, mission.IsCampaign, mission.IsAssert, mission.MissionType, mission.MissionDetail,
		mission.PlanBeginDatetime, mission.PlanEndDatetime, mission.RealBeginDatetime, 
		mission.RealEndDatetime, mission.PersonIncharge,
		mission.Status, mission.Picture, mission.MissionCode)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

//delete dependencies and targets related to mission
func DeleteMissionByMissionCode(missionCode * string) (bool, error) {
	tx, err := mysqlUtility.DBConn.Begin()
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
		pillarsLog.PillarsLogger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			pillarsLog.PillarsLogger.Panic(err.Error())
		}
		return false, err
	}
	return true, err
}



func QueryMissionByMissionCode(missionCode * string) (* utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, 
		is_campaign, is_assert, mission_type, mission_detail, plan_begin_datetime, 
		plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, 
		status, picture, insert_datetime, update_datetime FROM mission WHERE mission_code=?")
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
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.RealEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &mission, err
}

func QueryMissionsByProjectCode(projectCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, 
		is_campaign, is_assert,
		mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM mission where project_code = ?")
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
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryCampaignsByProjectCode(projectCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, project_code, product_type, 
		is_campaign, is_assert, 
		mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, 
		real_begin_datetime, real_end_datetime, person_in_charge, status, picture, 
		insert_datetime, update_datetime FROM mission where project_code = ? AND is_campaign=1")
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
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryAssertCampaignsByProjectCode(projectCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, 
		project_code, product_type, 
		is_campaign, is_assert, 
		mission_type, mission_detail, plan_begin_datetime, plan_end_datetime, 
		real_begin_datetime, real_end_datetime, person_in_charge, status, picture, 
		insert_datetime, update_datetime FROM mission where project_code = ? AND is_campaign=1 and is_assert=1")
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
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryUnassertCampaignsByProjectCode(projectCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, project_code, 
		product_type, is_campaign, is_assert, mission_type, mission_detail, 
		plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, 
		person_in_charge, status, picture, insert_datetime, update_datetime FROM mission 
		where project_code = ? AND is_campaign=1 and is_assert=0")
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
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryWaitingMissionByUserCode(userCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, project_code, 
		product_type, is_campaign, is_assert, mission_type, mission_detail, 
		plan_begin_datetime, plan_end_datetime, real_begin_datetime, 
		real_end_datetime, person_in_charge, status, picture, insert_datetime, 
		update_datetime FROM mission where person_in_charge = ? AND status=0")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(userCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var missionSlice [] utility.Mission
	for result.Next() {
		var mission utility.Mission
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryUndergoingMissionByUserCode(userCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, 
		project_code, product_type, is_campaign, mission_type, mission_detail, 
		plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, 
		person_in_charge, status, picture, insert_datetime, update_datetime 
		FROM mission where person_in_charge = ? AND status=3")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(userCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var missionSlice [] utility.Mission
	for result.Next() {
		var mission utility.Mission
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryReviewingMissionByUserCode(userCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, 
		project_code, product_type, is_campaign, is_assert, mission_type, mission_detail, 
		plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, 
		person_in_charge, status, picture, insert_datetime, update_datetime FROM mission 
		where person_in_charge = ? AND status=1")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(userCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var missionSlice [] utility.Mission
	for result.Next() {
		var mission utility.Mission
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}

func QueryFinishedMissionByUserCode(userCode * string) ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, 
		project_code, product_type, is_campaign, is_assert, mission_type, mission_detail, 
		plan_begin_datetime, plan_end_datetime, real_begin_datetime, 
		real_end_datetime, person_in_charge, status, picture, insert_datetime, 
		update_datetime FROM mission where person_in_charge = ? AND status=2")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(userCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var missionSlice [] utility.Mission
	for result.Next() {
		var mission utility.Mission
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}


func QueryAllUndesignatedMission() ([] utility.Mission, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT mission_code, mission_name, project_code, 
		product_type, is_campaign, is_assert, mission_type, mission_detail, plan_begin_datetime, 
		plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, 
		status, picture, insert_datetime, update_datetime FROM mission where status=4")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var missionSlice [] utility.Mission
	for result.Next() {
		var mission utility.Mission
		err = result.Scan(&(mission.MissionCode), &(mission.MissionName), &(mission.ProjectCode),
		&(mission.ProductType), &(mission.IsCampaign), &(mission.IsAssert), &(mission.MissionType), &(mission.MissionDetail),
		&(mission.PlanBeginDatetime), &(mission.PlanEndDatetime), &(mission.RealBeginDatetime), &(mission.PlanEndDatetime), 
		&(mission.PersonIncharge), &(mission.Status), &(mission.Picture), 
		&(mission.InsertDatetime), 
		&(mission.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		missionSlice = append(missionSlice, mission)
	}
	return missionSlice, err
}