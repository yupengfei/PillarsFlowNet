package storage

import (
	"testing"
	"time"
	// "fmt"
	"PillarsFlowNet/utility"
)


func TestStoreToChart(t * testing.T) {
	fromUserCode := "123"
	sendTime := time.Now().Format("2006-01-02 15:04:05")
	toUserCode := "456"
	isPicture := 0
	message := "test"

	chartCode := utility.GenerateCode(&fromUserCode)
	chart := utility.Chart {
		Id: *chartCode,
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
}

func TestMarkAsReceiveByChartCode(t * testing.T) {
	chartCode := "75551d43829706c067a5f8053eccc60c"
	_, err := MarkAsReceiveByChartCode(&chartCode)
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println(*(utility.ObjectToJsonString(result)))
	}
}

func TestGetAllUnreceivedMessageByUserCode(t * testing.T) {
	userCode := "456"
	_, err := GetAllUnreceivedMessageByUserCode(&userCode)
	if err != nil {
		panic(err.Error())
	} else {
		//fmt.Println(*(utility.ObjectToJsonString(result)))
	}
}


