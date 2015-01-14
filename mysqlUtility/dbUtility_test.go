package storage

import "testing"
// import "fmt"

func TestDBConnection(t * testing.T) {
	ConnectToDB()

}
func TestDBConnection2(t * testing.T) {
	//ConnectToDB("root", "123456", "172.16.253.216", "3306", "PillarsFlow")
	ConnectToDB()
}

func TestDBClose(t * testing.T) {
	CloseDBConnection()
}