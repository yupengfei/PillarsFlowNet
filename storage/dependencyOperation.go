package storage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)


//insert is a Transaction
func InsertIntoDependency(dependency * utility.Dependency) (bool, error) {

	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("INSERT INTO dependency(dependency_code, project_code, start_mission_code, end_mission_code, dependencyType) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(dependency.DependencyCode, dependency.ProjectCode, dependency.StartMissionCode, dependency.EndMissionCode, dependency.DependencyType)
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
	return dependencySlice, err

}



