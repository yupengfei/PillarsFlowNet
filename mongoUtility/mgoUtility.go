package mongoUtility

import "labix.org/v2/mgo"
import "PillarsFlowNet/utility"

var Session *mgo.Session
var ChartCollection *mgo.Collection
var PostCollection *mgo.Collection

//ConnectToMgo("root:123456@172.16.253.216/PillarsFlow")
func init() {
	Session = ConnectToMgo()
}

func ConnectToMgo() *mgo.Session {
	if Session != nil {
		// fmt.Println("session already exist")
		return Session
	}

	propertyMap := utility.ReadProperty("./Mgo.properties")

	var host, database string
	userName := propertyMap["DBUserName"]
	password := propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	//port = propertyMap["DBPort"]
	database = propertyMap["DBDatabase"]
	//userName + ":" + password + "@" +
	Session, errMgo := mgo.Dial(userName + ":" + password + "@" + host + "/" + database)
	if errMgo != nil {
		//panic("can not connect to mongo server")
		panic(errMgo.Error())
	}
	ChartCollection = Session.DB("PillarsFlow").C("Chart")
	PostCollection = Session.DB("PillarsFlow").C("Post")
	return Session
}

func CloseMgoConnection() {
	if Session != nil {
		Session.Close()
		Session = nil
	}
}
