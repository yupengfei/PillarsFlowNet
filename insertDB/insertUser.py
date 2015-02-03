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
			email = words[0]
			password = words[1]
			group = words[2]
			display_name = words[3]
			picture = words[4]
			user_code = md5sum.CalculateMd5(email 
				+ str(randomInt.Random()))
			sql = "INSERT INTO user " \
				+ "(user_code, email, password, `group`, \
					display_name, picture)" \
				+ " VALUES ('" \
				+ user_code + "','" \
				+ email + "','" +  md5sum.CalculateMd5(password) \
				+ "','" + group + "','" + display_name + "','" + picture + "')"
			print(sql)
			DBoperation.RunInsertCommand(connection, sql)
			


if __name__ == '__main__':
	connection = DBoperation.ConnectDB("localhost", "3306", "root", 
		"123456", "PillarsFlow")
	InsertUser("users", connection)
