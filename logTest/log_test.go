package logTest

import (
	"PillarsFlowNet/pillarsLog"
	"testing"
)

func TestLogFile(t * testing.T) {
	pillarsLog.Logger.Println("test")
}

func TestLogFile2(t * testing.T) {
	pillarsLog.Logger.Println("test2")
}

func TestLogFile3(t * testing.T) {
	pillarsLog.Destory()
}