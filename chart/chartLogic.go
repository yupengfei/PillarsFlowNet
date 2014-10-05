package chart

import (
	"time"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)
// ChartCode: *chartCode,
// 		IsPicture: isPicture,
// 		Message: message,
// 		From: fromUserCode,
// 		SendTime: sendTime,
// 		To: toUserCode,
// 		ReceivedTime: time.Now().Format("2006-01-02 15:04:05"),
// 		Receipt: receipt,
// 		IsRead: 0,
// 		Deleted: 0,
// 		DeletedTime: time.Now().Format("2006-01-02 15:04:05"),
func Chart(userCodeAndParameter * string) ([] byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var chart * utility.Chart
	var toUserCode * string
	if (errorCode == 0) {
		chart, _ := utility.ParseChartMessage(&(inputParameters[1]))
		toUserCode = &(chart.To)

		chart.ChartCode = *(utility.GenerateCode(&inputParameters[0]))
		chart.From = inputParameters[0]
		chart.ReceivedTime = time.Now().Format("2006-01-02 15:04:05")
		chart.IsRead = 0
		chart.Deleted = 0
		chart.DeletedTime = time.Now().Format("2006-01-02 15:04:05")
		storage.StoreToChart(chart)
	}
	return utility.ObjectToJsonByte(chart), toUserCode
}