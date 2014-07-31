package authentication

import (
	"fmt"
)

//init when start
//read data from mysql and authentication.properties
var AuthMap map[* string]string

func init() {
	
}
//todo
//this function has not been complished
func GetAuthInformation(userName * string) bool {
	fmt.Println(userName)
	return true
}