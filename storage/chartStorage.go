package storage
import (

	"PillarsFlowNet/utility"
	"labix.org/v2/mgo/bson"
	"time"
)
// type Chart struct {
//     ChartCode string
//     IsPicture string
//     Message string
//     From string
//     SendTime string
//     To string
//     ReceivedTime string
//     IsReceived int
//     Deleted int
//     DeletedTime string
// }
func StoreToChart(chart * utility.Chart) (* utility.Chart, error){
	
	err := ChartCollection.Insert(chart)
	if err != nil {
		return chart, err
	}
	ChartCollection.Find(bson.M{"chartcode":chart.ChartCode}).One(chart)
	return chart, err
}

func MarkAsReceiveByChartCode(chartCode * string) (* string, error) {
	err := ChartCollection.Update(bson.M{"chartcode": *chartCode}, bson.M{"$set": bson.M{"isreceived": 1, 
		"receivedtime": time.Now().Format("2006-01-02 15:04:05")}})
	if err != nil {
		return chartCode, err
	}
	return chartCode, err
}

func GetAllUnreceivedMessageByUserCode(userCode * string) ([] utility.Chart, error) {
	var chartSlice [] utility.Chart
	iter := ChartCollection.Find(bson.M{"to":*userCode, "isreceived":0}).Iter()
	err := iter.All(&chartSlice)
	if err != nil {
		return chartSlice, err
	}
	return chartSlice, err
}

