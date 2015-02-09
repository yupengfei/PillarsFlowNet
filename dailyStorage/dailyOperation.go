package dailyStorage

import (
	"PillarsFlowNet/mysqlUtility"
	"PillarsFlowNet/utility"
	//"PillarsFlowNet/pillarsLog"
	// "fmt"
)

//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry

//insert is a Transaction
func InsertIntoDaily(daily *utility.Daily) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO daily(daily_code,company_code, mission_code, project_code, version_tag, storage_position, picture) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(daily.DailyCode, daily.CompanyCode, daily.MissionCode, daily.ProjectCode, daily.VersionTag, daily.StoragePosition, daily.Picture)
	if err != nil {
		//pillarsLog.Logger.Print(err.Error())
		panic(err.Error())
	}
	return true, err
}

func ModifyDaily(daily *utility.Daily) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE daily SET  mission_code=?, project_code=?, version_tag=?, storage_position=?, picture=? WHERE daily_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(daily.MissionCode, daily.ProjectCode, daily.VersionTag, daily.StoragePosition, daily.Picture, daily.DailyCode)
	if err != nil {
		//pillarsLog.Logger.Print(err.Error())
		panic(err.Error())
	}
	return true, err
}

func DeleteDailyByDailyCode(dailyCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("DELETE FROM daily WHERE daily_code = ?")
	defer stmt.Close()
	_, err = stmt.Exec(dailyCode)
	if err != nil {
		//pillarsLog.Logger.Print(err.Error())
		panic(err.Error())
	}
	return true, err
}
func QueryDailysByCompanyCode(companyCode *string) ([]utility.Daily, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT daily_code,company_code, company_code,mission_code, project_code, version_tag, storage_position, picture, insert_datetime, update_datetime FROM daily WHERE company_code = ? AND Date(insert_datetime)=Date(Now())")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(companyCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var targetSlice []utility.Daily
	for result.Next() {
		var daily utility.Daily
		err = result.Scan(&(daily.DailyCode), &(daily.CompanyCode), &(daily.MissionCode), &(daily.ProjectCode), &(daily.VersionTag), &(daily.StoragePosition),
			&(daily.Picture), &(daily.InsertDatetime), &(daily.UpdateDatetime))
		if err != nil {
			//pillarsLog.Logger.Print(err.Error())
		}
		targetSlice = append(targetSlice, daily)
	}
	return targetSlice, err
}
func QueryDailysByMissionCode(missionCode *string) ([]utility.Daily, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT daily_code,company_code, mission_code, project_code, version_tag, storage_position, picture, insert_datetime, update_datetime FROM daily WHERE mission_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(missionCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var targetSlice []utility.Daily
	for result.Next() {
		var daily utility.Daily
		err = result.Scan(&(daily.DailyCode), &(daily.CompanyCode), &(daily.MissionCode), &(daily.ProjectCode), &(daily.VersionTag), &(daily.StoragePosition),
			&(daily.Picture), &(daily.InsertDatetime), &(daily.UpdateDatetime))
		if err != nil {
			//pillarsLog.Logger.Print(err.Error())
		}
		targetSlice = append(targetSlice, daily)
	}
	return targetSlice, err
}

func QueryDailyByDailyCode(dailyCode *string) (*utility.Daily, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT daily_code,company_code, mission_code, project_code, version_tag, storage_position, picture, insert_datetime, update_datetime FROM daily WHERE daily_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(dailyCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var daily utility.Daily
	if result.Next() {
		err = result.Scan(&(daily.DailyCode), &(daily.CompanyCode), &(daily.MissionCode), &(daily.ProjectCode), &(daily.VersionTag), &(daily.StoragePosition),
			&(daily.Picture), &(daily.InsertDatetime), &(daily.UpdateDatetime))
		if err != nil {
			//pillarsLog.Logger.Print(err.Error())
		}
	}
	return &daily, err
}
