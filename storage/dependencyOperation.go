package storage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)
//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry 


//insert is a Transaction
func InsertIntoDependency(dependency * utility.Dependency) bool {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("INSERT INTO dependency(project_code, start_mission_code, end_mission_code, dependencyType) VALUES(?, ?, ?, ?)")
	if err != nil {
		// fmt.Print(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(dependency.ProjectCode, dependency.StartMissionCode, dependency.EndMissionCode, dependency.DependencyType)
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

func QueryDependenciesByProjectCode(projectCode * string) [] utility.Dependency{
	

	stmt, err := DBConn.Prepare("SELECT project_code, start_mission_code, end_mission_code, dependencyType, insert_datetime, update_datetime FROM dependency WHERE project_code = ?")
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
	if result.Next() {
		var dependency utility.Dependency
		err = result.Scan(&(dependency.ProjectCode), &(dependency.StartMissionCode), &(dependency.EndMissionCode),
		&(dependency.DependencyType), &(dependency.InsertDatetime), &(dependency.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		dependencySlice = append(dependencySlice, dependency)
	}
	return dependencySlice

}

