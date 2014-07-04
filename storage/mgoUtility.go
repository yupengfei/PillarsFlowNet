package storage

// import "fmt"
import "labix.org/v2/mgo"
//import "labix.org/v2/mgo/bson"

var session * mgo.Session
var errMgo error

//ConnectToMgo("root:123456@172.16.253.216/PillarsFlow")

func ConnectToMgo(userName string, password string, host string, port string, database string) * mgo.Session {
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


