package mongoUtility

import (
	"testing"

)

func TestConnectToMgo(t * testing.T) {
	session := ConnectToMgo()
	if session == nil {
		t.Error("connect to mongo 1 failed")
	}
}

func TestConnectToMgo2(t * testing.T) {
	session := ConnectToMgo()
	if session == nil {
		t.Error("connect to mongo 1 failed")
	}
}

func TestCloseMgoConnection(t * testing.T) {
	CloseMgoConnection()
}