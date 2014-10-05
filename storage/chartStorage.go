package storage
import (

	"PillarsFlowNet/utility"
	// "fmt"
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
func StoreToChart(chart * utility.Chart) (bool, error){
	
	err := ChartCollection.Insert(chart)
	if err != nil {
		return false, err
	}
	return true, err
}

// func MarkAsReadByChartCode(chartCode * string) (bool, error) {
	
// }

