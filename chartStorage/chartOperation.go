package storage
import (
	"PillarsFlowNet/mongoUtility"
	"PillarsFlowNet/utility"
	"labix.org/v2/mgo/bson"
	"time"
)

func StoreToChart(chart * utility.Chart) (* utility.Chart, error){
	
	err := mongoUtility.ChartCollection.Insert(chart)
	if err != nil {
		return chart, err
	}
	mongoUtility.ChartCollection.Find(bson.M{"_id":chart.Id}).One(chart)
	return chart, err
}

func MarkAsReceiveByChartCode(chartCode * string) (* string, error) {
	err := mongoUtility.ChartCollection.Update(bson.M{"_id": *chartCode}, bson.M{"$set": bson.M{"isreceived": 1, 
		"receivedtime": time.Now().Format("2006-01-02 15:04:05")}})
	if err != nil {
		return chartCode, err
	}
	return chartCode, err
}

func GetAllUnreceivedMessageByUserCode(userCode * string) ([] utility.Chart, error) {
	var chartSlice [] utility.Chart
	iter := mongoUtility.ChartCollection.Find(bson.M{"to":*userCode, "isreceived":0}).Iter()
	err := iter.All(&chartSlice)
	if err != nil {
		return chartSlice, err
	}
	return chartSlice, err
}

