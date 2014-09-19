package storage
import (
	"time"
	"fmt"
	"labix.org/v2/mgo"
)
type ChartStorage struct {
	Message * string
	From * string
	SendTime * string
	To * string
	ReceivedTime * string
	Receipt int
	IsRead int
	Deleted int
	DeletedTime * string
}
func StoreToChart(from_user_code * string, to_user_code * string, message string, receipt bool, session * mgo.Session) {
	// var charSto ChartStorage
	
	//golang's format is a mess.
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//let the ReceivedTime be now
	



}