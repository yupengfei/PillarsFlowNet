#! /usr/bin/python3

import random

def Random():
	magicNumber = 100
	return random.randint(0, magicNumber)

if __name__ == "__main__":
	print(Random())