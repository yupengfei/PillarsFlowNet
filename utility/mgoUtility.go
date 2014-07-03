package utility

// import "fmt"
import "labix.org/v2/mgo"
//import "labix.org/v2/mgo/bson"

var session * mgo.Session
var errMgo error

func ConnectToMgo(host string) * mgo.Session {
	if session != nil {
		return session
	}
	session, errMgo = mgo.Dial(host)
	if err != nil {
		panic(err.Error())
	}
	return session
}

func CloseMgoConnection() {
	session.Close()
}


