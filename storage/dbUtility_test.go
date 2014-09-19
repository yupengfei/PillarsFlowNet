package storage

import "testing"
// import "fmt"

func TestDBConnection(t * testing.T) {
	ConnectToDB()
	rows, err := DBConn.Query("SELECT 1")
	// if err != nil {
	// 	panic(err.Error())
	// }
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
	//ConnectToDB("root", "123456", "172.16.253.216", "3306", "PillarsFlow")
	ConnectToDB()
}

func TestDBClose(t * testing.T) {
	CloseDBConnection()
}