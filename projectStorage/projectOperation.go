package projectStorage

import (
	"PillarsFlowNet/utility"
	"PillarsFlowNet/mysqlUtility"
	"PillarsFlowNet/pillarsLog"
	// "fmt"
)

func InsertIntoProject(project * utility.Project) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("INSERT INTO project(project_code, project_name, project_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.ProjectCode, project.ProjectName, project.ProjectDetail,
		project.PlanBeginDatetime, project.PlanEndDatetime, project.RealBeginDatetime, 
		project.RealEndDatetime, project.PersonInCharge,
		project.Status, project.Picture)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

func ModifyProject(project * utility.Project) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("UPDATE project SET project_name=?, project_detail=?, plan_begin_datetime=?, plan_end_datetime=?, real_begin_datetime=?, real_end_datetime=?, person_in_charge=?, status=?, picture=? WHERE project_code=?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(project.ProjectName, project.ProjectDetail,
		project.PlanBeginDatetime, project.PlanEndDatetime, project.RealBeginDatetime, 
		project.RealEndDatetime, project.PersonInCharge,
		project.Status, project.Picture, project.ProjectCode)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

func DeleteProjectByProjectCode(projectCode * string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("DELETE FROM project WHERE project_code = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(projectCode)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

func QueryProjectByProjectCode(projectCode * string) (* utility.Project, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT project_code, project_name, project_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM project WHERE project_code=?")
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
			pillarsLog.PillarsLogger.Print(err.Error())
		}
	}
	return &project, err
}

func QueryAllProject() ([] utility.Project, error) {
	stmt, err := mysqlUtility.DBConn.Prepare("SELECT project_code, project_name, project_detail, plan_begin_datetime, plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status, picture, insert_datetime, update_datetime FROM project")
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
	var projectSlice [] utility.Project
	for result.Next() {
		var project utility.Project
		err = result.Scan(&(project.ProjectCode), &(project.ProjectName), &(project.ProjectDetail),
		&(project.PlanBeginDatetime), &(project.PlanEndDatetime), &(project.RealBeginDatetime), &(project.RealEndDatetime), &(project.PersonInCharge),
		&(project.Status), &(project.Picture), &(project.InsertDatetime), &(project.UpdateDatetime))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		projectSlice = append(projectSlice, project)
	}
	return projectSlice, err
}

