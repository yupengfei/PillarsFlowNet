package storage
import (
	"time"
	"PillarsFlowNet/utility"
	"fmt"
)
// type Chart struct {
//     ChartCode string
//     IsPicture string
//     Message string
//     From string
//     SendTime string
//     To string
//     ReceivedTime string
//     Receipt int
//     IsRead int
//     Deleted int
//     DeletedTime string
// }
func StoreToChart(fromUserCode * string, sendTime * string, toUserCode * string, message * string, receipt bool) (bool, error){
	
    chartCode := utility.GenerateCode(fromUserCode)
	chartMgo := utility.Chart {
		ChartCode: *chartCode,
		From: * fromUserCode,
		SendTime: * sendTime,
		To: * toUserCode,
		ReceivedTime: time.Now().Format("2006-01-02 15:04:05"),
		Receipt: 0,
		IsRead: 0,
		Deleted: 0,
		DeletedTime: time.Now().Format("2006-01-02 15:04:05"),
		Message: *message,
	}
	
	// 插入记录
	fmt.Println(*(utility.ObjectToJsonString(chartMgo)))
	//chartString := utility.ObjectToJsonString(chart)
	err := ChartCollection.Insert(chartMgo)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

// func MarkAsReadByChartCode(chartCode * string) (bool, error) {
	
// }

