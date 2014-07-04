package storage

import "testing"

func TestConnectToMgo(t * testing.T) {
	ConnectToMgo("root", "123456", "172.16.253.216", "27017", "PillarsFlow")
} 

func TestConnectToMgo2(t * testing.T) {
	ConnectToMgo("root", "123456", "172.16.253.216", "27017", "PillarsFlow")
} 

func TestMgoClose(t * testing.T) {
	CloseMgoConnection()
}
