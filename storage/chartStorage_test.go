package storage

import (
	"testing"
	"time"
	// "fmt"
	"PillarsFlowNet/utility"
)


func TestStoreToChart(t * testing.T) {
	ConnectToMgo()
	fromUserCode := "123"
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	toUserCode := "456"
	isPicture := 0
	message := "test"

	chartCode := utility.GenerateCode(&fromUserCode)
	chart := utility.Chart {
		ChartCode: *chartCode,
		IsPicture: isPicture,
		Message: message,
		From: fromUserCode,
		SendTime: sendTime,
		To: toUserCode,
		ReceivedTime: time.Now().Format("2006-01-02 15:04:05"),
		IsReceived: 0,
		Deleted: 0,
		DeletedTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := StoreToChart(&chart)
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println(*(utility.ObjectToJsonString(result)))
	}
	CloseMgoConnection()
}

func TestMarkAsReceiveByChartCode(t * testing.T) {
	ConnectToMgo()
	chartCode := "6fce4188e44a2decab8f0bc2cfbff1fd"
	_, err := MarkAsReceiveByChartCode(&chartCode)
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println(*(utility.ObjectToJsonString(result)))
	}
	CloseMgoConnection()
}

func TestGetAllUnreceivedMessageByUserCode(t * testing.T) {
	ConnectToMgo()
	userCode := "456"
	_, err := GetAllUnreceivedMessageByUserCode(&userCode)
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println(*(utility.ObjectToJsonString(result)))
	}
	CloseMgoConnection()
}


