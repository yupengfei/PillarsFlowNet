package storage

import "testing"

func TestConnectToMgo(t * testing.T) {
	//ConnectToMgo("root", "123456", "172.16.253.216", "27017", "PillarsFlow")
	ConnectToMgo()
} 

func TestConnectToMgo2(t * testing.T) {
	ConnectToMgo()
} 

func TestMgoClose(t * testing.T) {
	CloseMgoConnection()
}
