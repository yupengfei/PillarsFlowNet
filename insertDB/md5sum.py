#!/usr/bin/python3
import hashlib

def CalculateMd5(str):
	'''
	Calculate md5 of str
	'''
	temp = hashlib.md5(str.encode())
	md5value = temp.hexdigest()
	return md5value

if __name__ == '__main__':
	print(CalculateMd5("123456"))

