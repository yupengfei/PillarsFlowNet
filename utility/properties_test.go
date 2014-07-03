package utility

import "testing"

func TestReadProperty(t * testing.T) {
	ReadProperty("../DB.properties")
	//fmt.Println(propertyMap["DBIP"])
}
