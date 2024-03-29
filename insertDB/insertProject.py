#!/usr/bin/python3

import DBoperation
import md5sum
import randomInt

def InsertProject(fileName, connection):
	with open("projects") as inFile:
		lines = inFile.readlines()
		projectCount = len(lines)/9
		print(int(projectCount))
		for i in range(2):
			project_name = lines[9 * i + 0]
			project_detail = lines[9 * i + 1]
			plan_begin_datetime = lines[9 * i + 2]
			plan_end_datetime = lines[9 * i + 3]
			real_begin_datetime = lines[9 * i + 4]
			real_end_datetime = lines[9 * i + 5]
			person_in_charge = lines[9 * i + 6]
			status = lines[9 * i + 7]
			project_code = md5sum.CalculateMd5(project_name 
				+ str(randomInt.Random()))
			sql = "INSERT INTO project " \
				+ "(project_code, project_name, project_detail, `plan_begin_datetime`, \
					plan_end_datetime, real_begin_datetime, real_end_datetime, person_in_charge, status)" \
				+ " VALUES ('" \
				+ project_code + "','" \
				+ project_name + "','" +  project_detail \
				+ "','" + plan_begin_datetime + "','" + plan_end_datetime + "','" + real_begin_datetime \
				+ "','" + real_end_datetime + "','" + person_in_charge + "','" + status \
				+ "')"
			print(sql)
			DBoperation.RunInsertCommand(connection, sql)
		
			


if __name__ == '__main__':


	connection = DBoperation.ConnectDB("172.16.246.253", "3306", "root", 
		"123456", "PillarsFlow")
	InsertProject("projects", connection)
