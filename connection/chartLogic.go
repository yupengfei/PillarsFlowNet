package connection

import (
	"PillarsFlowNet/authentication"
	"PillarsFlowNet/chartStorage"
	"PillarsFlowNet/utility"
	"time"
	// "fmt"
)

//向chart表添加一条记录
//其中inputParameters[0]是执行该操作的用户的code
//inputParameters[1]包含了更多的聊天信息
func AddChart(userCode *string, parameter *string) { //result, fromUserCode, ToUserCode
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var chartOut *utility.Chart
	var err error
	var toUserCode *string
	if errorCode == 0 {
		chart, _ := utility.ParseChartMessage(parameter)
		toUserCode = &(chart.To)

		chart.Id = *(utility.GenerateCode(userCode))
		chart.From = *userCode
		chart.ReceivedTime = time.Now().Format("2006-01-02 15:04:05")
		chart.IsReceived = 0
		chart.Deleted = 0
		chart.DeletedTime = time.Now().Format("2006-01-02 15:04:05")
		chartOut, err = chartStorage.StoreToChart(chart)
		if err != nil {
			errorCode = 1
		}
	}
	var command = "addChart"
	result := utility.BoolResultToOutMessage(&command, chartOut, errorCode, userCode)
	Hub.SendToUserCode(result, userCode) ////这个有必要在发给自己吗
	Hub.SendToUserCode(result, toUserCode)
}

//用户在阅读完某条信息后会将该条信息标记为已读
func ReceiveChart(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	if errorCode == 0 {
		chartCode, _ := utility.ParseChartCodeMessage(parameter)
		_, err := chartStorage.MarkAsReceiveByChartCode(&(chartCode.ChartCode))
		if err != nil {
			errorCode = 1
		}
	}

	var command = "receiveChart"
	result := utility.BoolResultToOutMessage(&command, parameter, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}

//用户登陆后，会向服务器请求所有发给自己的未读信息
func GetAllUnreceivedChart(userCode *string, parameter *string) {
	auth := authentication.GetAuthInformation(userCode)
	var errorCode int
	if auth == false {
		errorCode = 3
	}
	var opResult []utility.Chart
	var err error
	if errorCode == 0 {
		opResult, err = chartStorage.GetAllUnreceivedMessageByUserCode(parameter)
		if err != nil {
			errorCode = 1
		}
	}

	command := "getAllUnreceivedChart"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, userCode)
	Hub.SendToUserCode(result, userCode)
}
