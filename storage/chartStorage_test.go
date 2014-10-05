package storage

import (
	"testing"
	"time"
	"PillarsFlowNet/utility"
)


func TestStoreToChart(t * testing.T) {
	ConnectToMgo()
	fromUserCode := "123"
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	toUserCode := "456"
	isPicture := 0
	message := "test"
	receipt := 0

	chartCode := utility.GenerateCode(&fromUserCode)
	chart := utility.Chart {
		ChartCode: *chartCode,
		IsPicture: isPicture,
		Message: message,
		From: fromUserCode,
		SendTime: sendTime,
		To: toUserCode,
		ReceivedTime: time.Now().Format("2006-01-02 15:04:05"),
		Receipt: receipt,
		IsRead: 0,
		Deleted: 0,
		DeletedTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	StoreToChart(&chart)
	CloseMgoConnection()
}


