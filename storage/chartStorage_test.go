package storage

import (
	"testing"
	"time"
)


func TestStoreToChart(t * testing.T) {
	ConnectToMgo()
	fromUserCode := "123"
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	toUserCode := "456"
	message := "test"
	receipt := false

	StoreToChart(&fromUserCode, &sendTime, &toUserCode, &message, receipt)
	CloseMgoConnection()
}


