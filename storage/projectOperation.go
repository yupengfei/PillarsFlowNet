package storage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)
//we should use persistence layer instead, but the logic id not so confusing
//we are in hurry 


//insert is a Transaction
func InsertIntoProject(project * utility.Project) bool {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare("INSERT INTO project(project_code, project_name, project_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		// fmt.Print(err.Error())
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.ProjectCode, project.ProjectName, project.ProjectDetail,
		project.PlanBeginDatetime, project.PlanEndDatetime, project.RealBeginDatetime, 
		project.RealEndDatetime, project.PersonInCharge,
		project.Status, project.Picture)
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

func QueryProjectByProjectCode(projectCode * string) * utility.Project {
	stmt, err := DBConn.Prepare("SELECT project_code, project_name, project_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM project WHERE project_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var project utility.Project
	if result.Next() {
		err = result.Scan(&(project.ProjectCode), &(project.ProjectName), &(project.ProjectDetail),
		&(project.PlanBeginDatetime), &(project.PlanEndDatetime), &(project.RealBeginDatetime), &(project.PlanEndDatetime), &(project.PersonInCharge),
		&(project.Status), &(project.Picture), &(project.InsertDatetime), &(project.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
	}
	return &project
}

func QueryAllProject() [] utility.Project {
	stmt, err := DBConn.Prepare("SELECT project_code, project_name, project_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM project")
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
	var projectSlice [] utility.Project
	for result.Next() {
		var project utility.Project
		err = result.Scan(&(project.ProjectCode), &(project.ProjectName), &(project.ProjectDetail),
		&(project.PlanBeginDatetime), &(project.PlanEndDatetime), &(project.RealBeginDatetime), &(project.PlanEndDatetime), &(project.PersonInCharge),
		&(project.Status), &(project.Picture), &(project.InsertDatetime), &(project.UpdateDatetime))
		if err != nil {
			pillarsLog.Logger.Print(err.Error())
		}
		projectSlice := append(projectSlice, project)
	}
	return projectSlice
}
