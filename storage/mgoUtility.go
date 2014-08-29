package storage


import "labix.org/v2/mgo"
import "PillarsFlowNet/utility"
import "PillarsFlowNet/pillarsLog"
var Session * mgo.Session

//ConnectToMgo("root:123456@172.16.253.216/PillarsFlow")
func init() {
	Session = ConnectToMgo()
}

func ConnectToMgo() * mgo.Session {
	if Session != nil {
		// fmt.Println("session already exist")
		return Session
	}

	propertyMap := utility.ReadProperty("../Mgo.properties")

	var  host, database string
	// userName =  propertyMap["DBUserName"]
	// password = propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	//port = propertyMap["DBPort"]
	database = propertyMap["DBDatabase"]
	//userName + ":" + password + "@" + 
	Session, errMgo := mgo.Dial(host + "/" + database)
	if errMgo != nil {
		pillarsLog.Logger.Panic("can not connect to mongo server")
	}
	return Session
}

func CloseMgoConnection() {
	Session.Close()
	Session = nil
}


