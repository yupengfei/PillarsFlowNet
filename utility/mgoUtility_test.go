package utility

import "testing"

func TestConnectToMgo(t * testing.T) {
	ConnectToMgo("root:123456@172.16.253.216/PillarsFlow")
} 

func TestConnectToMgo2(t * testing.T) {
	ConnectToMgo("172.16.253.216")
} 

func TestMgoClose(t * testing.T) {
	CloseMgoConnection()
}
