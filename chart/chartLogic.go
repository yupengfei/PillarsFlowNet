package chart

import (
	"time"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)

func AddChart(userCodeAndParameter * string) ([] byte, *string, *string) {//result, fromUserCode, ToUserCode
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var chartOut * utility.Chart
	var err error
	var toUserCode * string
	if (errorCode == 0) {
		chart, _ := utility.ParseChartMessage(&(inputParameters[1]))
		toUserCode = &(chart.To)

		chart.ChartCode = *(utility.GenerateCode(&inputParameters[0]))
		chart.From = inputParameters[0]
		chart.ReceivedTime = time.Now().Format("2006-01-02 15:04:05")
		chart.IsReceived = 0
		chart.Deleted = 0
		chart.DeletedTime = time.Now().Format("2006-01-02 15:04:05")
		chartOut, err = storage.StoreToChart(chart)
		if err != nil {
			errorCode = 1
		}
	}
	var command = "addChart"
	result := utility.BoolResultToOutMessage(&command, chartOut, errorCode, &inputParameters[0])
	return result, &(inputParameters[0]), toUserCode
}

func ReceiveChart(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	if (errorCode == 0) {
		chartCode, _ := utility.ParseChartCodeMessage(&(inputParameters[1]))
		_, err :=storage.MarkAsReceiveByChartCode(&(chartCode.ChartCode))
		if err != nil {
			errorCode = 1
		}
	}

	var command = "receiveChart"
	result := utility.BoolResultToOutMessage(&command, &inputParameters[1], errorCode, &inputParameters[0])
	return result, &(inputParameters[0])
}

func GetAllUnreceivedChart(userCodeAndParameter * string) ([] byte, *string) {
	//userCode, parameter 
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var opResult []utility.Chart
	var err error
	if (errorCode == 0) {
		opResult, err = storage.GetAllUnreceivedMessageByUserCode(&(inputParameters[0]))
		if err != nil {
			errorCode = 1
		}
	}

	command := "getAllUnreceivedChart"
	result := utility.SliceResultToOutMessage(&command, opResult, errorCode, &(inputParameters[0]))
	return result, &(inputParameters[0])
}