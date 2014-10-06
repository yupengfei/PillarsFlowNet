package chart

import (
	"time"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)
// addChart
// {
// 	“command”:”addChart”,
// 	“parameter”:”{
// 		ChartCode 任意string,可以没有 
//     		IsPicture int #0不是，1是
//     		Message string
//     		From 任意string,可以没有
//     		SendTime 任意string,可以没有
//     		To string
//     		ReceivedTime 任意string,可以没有
//     		IsReceived 任意int,可以没有
//     		Deleted int 任意int,可以没有
//     		DeletedTime 任意string,可以没有
// 	}”
// }
// 返回值分别给发出人和受到人
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “addChart”,
// 	“result”:”{
// 		ChartCode string 
//     		IsPicture int #0不是，1是
//     		Message string
//     		From string
//     		SendTime string
//     		To string
//     		ReceivedTime 任意string,可以没有
//     		IsRecieved 任意int,可以没有
//     		Deleted int 任意int,可以没有
//     		DeletedTime 任意string,可以没有
// 	}”
// }
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
	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * utility.OutMessage
	if errorCode != 0 {
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addChart",
						Result: "{}",
					}
		out = & tempout
	} else {
		tempout :=utility.OutMessage {
						Error: sysError,
						Command: "addChart",
						Result: *utility.ObjectToJsonString(chartOut) ,
					}
		out = & tempout
	}

	var result = utility.ObjectToJsonByte(*out)
	return result, &(inputParameters[0]), toUserCode
}

// receiveChart由收到人返回给服务器
// {
// 	“command”:”receiveChart”,
// 	“parameter”:”{
// 		ChartCode string
// 	}”
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “receiveChart”,
// 	“result”:”{
// 		ChartCode string
// 	}”
// }

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

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out *  utility.OutMessage
	if errorCode != 0 {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "receiveChart",
						Result: "{}",
					}
		out = & tempout
	} else {
		var tempout = utility.OutMessage {
						Error: sysError,
						Command: "receiveChart",
						Result:inputParameters[1],
					}
		out = & tempout
	}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}

// getAllUnreceivedChart获取所有的未读信息
// {
// 	“command”:”getAllUnreceivedChart”,
// 	”parameter“：”{
// 	}“
// }
// 返回值
// {
// 	"error": {
// 		"errorCode" : 0,
// 		"errorMessage": ""
// 	},
// 	“commnd”: “getAllUnreceivedChart”,
// 	“result”:”[{
// 		ChartCode string 
//     		IsPicture int #0不是，1是
//     		Message string
//     		From string
//     		SendTime string
//     		To string
//     		ReceivedTime 任意string,可以没有
//     		IsRecieved 任意int,可以没有
//     		Deleted int 任意int,可以没有
//     		DeletedTime 任意string,可以没有
// 	}]”
// }

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

	var sysError = utility.Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out = utility.OutMessage {
						Error: sysError,
						Command: "getAllUnreceivedChart",
						Result:*utility.ObjectToJsonString(opResult),
					}
	var result = utility.ObjectToJsonByte(out)

	return result, &(inputParameters[0])
}