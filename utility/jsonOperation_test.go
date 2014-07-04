package utility

import "testing"
import "fmt"
import "os"

func TestParserInMessage() {
	file, err := os.Open("login.json")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	data, err + io
}