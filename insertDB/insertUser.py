#!/usr/bin/python3

import DBoperation
import md5sum
import randomInt

def InsertUser(fileName, connection):
	# `user_code` char(32) not null unique,#计算生成的唯一识别符
	# `user_name` char(20) NOT NULL unique,
	# `password` char(32) not null,
	# `group` varchar(20) not null,
	# `display_name` char(20) not null,
	with open("users") as inFile:
		for line in inFile.readlines():
			words = line.split()
			# print(words)
			user_name = words[0]
			password = words[1]
			group = words[2]
			display_name = words[3]
			user_code = md5sum.CalculateMd5(user_name 
				+ str(randomInt.Random()))
			sql = "INSERT INTO user " \
				+ "(user_code, user_name, password, `group`, \
					display_name)" \
				+ " VALUES ('" \
				+ user_code + "','" \
				+ user_name + "','" +  md5sum.CalculateMd5(password) \
				+ "','" + group + "','" + display_name + "')"
			print(sql)
			DBoperation.RunInsertCommand(connection, sql)
			


if __name__ == '__main__':
	connection = DBoperation.ConnectDB("172.16.246.253", "3306", "root", 
		"123456", "PillarsFlow")
	InsertUser("users", connection)
