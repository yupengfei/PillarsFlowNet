package storage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)


//insert is a Transaction
func InsertIntoDependency(dependency * utility.Dependency) (bool, error) {

	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("INSERT INTO dependency(dependency_code, campaign_code, project_code, start_mission_code, end_mission_code, dependency_type) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(dependency.DependencyCode, dependency.CampaignCode, dependency.ProjectCode, dependency.StartMissionCode, dependency.EndMissionCode, dependency.DependencyType)
	if err != nil {
		return false, err
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
		return false, err
	}
	return true, err
}

func ModifyDependency(dependency * utility.Dependency) (bool, error) {

	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("UPDATE dependency SET campaign_code=?, project_code=?, start_mission_code=?, end_mission_code=?, dependency_type=? WHERE dependency_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(dependency.CampaignCode, dependency.ProjectCode, dependency.StartMissionCode, dependency.EndMissionCode, dependency.DependencyType, dependency.DependencyCode)
	if err != nil {
		panic(err.Error())
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
		return false, err
	}
	return true, err
}

func DeleteDependencyByDependencyCode(projectCode * string) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("DELETE FROM dependency WHERE dependency_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(*projectCode)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		pillarsLog.Logger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			pillarsLog.Logger.Panic(err.Error())
		}
		return false, err
	}
	return true, err
}

func QueryDependenciesByProjectCode(projectCode * string) ([] utility.Dependency, error){
	

	stmt, err := DBConn.Prepare("SELECT dependency_code, campaign_code, project_code, start_mission_code, end_mission_code, dependency_type, insert_datetime, update_datetime FROM dependency WHERE project_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var dependencySlice [] utility.Dependency
	for result.Next() {
		var dependency utility.Dependency
		err = result.Scan(&(dependency.DependencyCode), &(dependency.CampaignCode), &(dependency.ProjectCode), &(dependency.StartMissionCode), &(dependency.EndMissionCode),
		&(dependency.DependencyType), &(dependency.InsertDatetime), &(dependency.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		dependencySlice = append(dependencySlice, dependency)
	}
	return dependencySlice, err

}

func QueryDependenciesByCampaignCode(campaignCode * string) ([] utility.Dependency, error){
	

	stmt, err := DBConn.Prepare("SELECT dependency_code, campaign_code, project_code, start_mission_code, end_mission_code, dependency_type, insert_datetime, update_datetime FROM dependency WHERE campaign_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(campaignCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var dependencySlice [] utility.Dependency
	for result.Next() {
		var dependency utility.Dependency
		err = result.Scan(&(dependency.DependencyCode), &(dependency.CampaignCode), &(dependency.ProjectCode), &(dependency.StartMissionCode), &(dependency.EndMissionCode),
		&(dependency.DependencyType), &(dependency.InsertDatetime), &(dependency.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		dependencySlice = append(dependencySlice, dependency)
	}
	return dependencySlice, err

}

func QueryDependencyByDependencyCode(dependencyCode * string) (* utility.Dependency, error){
	

	stmt, err := DBConn.Prepare("SELECT dependency_code, campaign_code, project_code, start_mission_code, end_mission_code, dependency_type, insert_datetime, update_datetime FROM dependency WHERE dependency_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(dependencyCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var dependency utility.Dependency
	if result.Next() {
		err = result.Scan(&(dependency.DependencyCode), &(dependency.CampaignCode), &(dependency.ProjectCode), &(dependency.StartMissionCode), &(dependency.EndMissionCode),
		&(dependency.DependencyType), &(dependency.InsertDatetime), &(dependency.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
	}
	return &dependency, err

}




