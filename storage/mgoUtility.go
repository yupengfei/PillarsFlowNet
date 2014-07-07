package storage

// import "fmt"
import "labix.org/v2/mgo"
import "PillarsFlowNet/utility"
//import "labix.org/v2/mgo/bson"

var session * mgo.Session
var errMgo error

//ConnectToMgo("root:123456@172.16.253.216/PillarsFlow")

func ConnectToMgo() * mgo.Session {
	propertyMap := utility.ReadProperty("../Mgo.properties")

	var userName, password, host, database string
	userName =  propertyMap["DBUserName"]
	password = propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	//port = propertyMap["DBPort"]
	database = propertyMap["DBDatabase"]
	if session != nil {
		// fmt.Println("session already exist")
		return session
	}
	session, errMgo = mgo.Dial(userName + ":" + password + "@" + host + "/" + database)
	if errMgo != nil {
		panic(errMgo.Error())
	}
	return session
}

func CloseMgoConnection() {
	session.Close()
}


