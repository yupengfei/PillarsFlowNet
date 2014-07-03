package utility

import "testing"
// import "fmt"

func TestDBConnection(t * testing.T) {
	//fmt.Println("testing")
	// ConnectToDB("root", "123456", "172.16.253.216", "3306", "PillarsFlow")
	db := ConnectToDB("root", "123456", "172.16.253.216", "3306", "PillarsFlow")
	//stmtOut, err := db.Prepare()
	// if err := nil {
	// 	panic(err.Error())
	// }
	
	rows, err := db.Query("SELECT 1")
	for rows.Next() {
		var result int
		err = rows.Scan(&result)
		//fmt.Println(result)
		if err != nil || result != 1 {
			t.Error("error")
		}
	}
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(result)
}
func TestDBConnection2(t * testing.T) {
	ConnectToDB("root", "123456", "172.16.253.216", "3306", "PillarsFlow")
}

func TestDBClose(t * testing.T) {
	CloseDBConnection()
}