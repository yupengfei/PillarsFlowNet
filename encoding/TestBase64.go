package picture

import (
	"fmt"
	"testing"
)

func TestBase64Bytes(t * testing.T) {
	hello := "hello world"
	debyte := Base64Encode([]byte(hello))

	enbyte, err := Base64Decode(debyte)
	if err != nil {
		panic(err.Error())
	}
	if hello != string(enbyte) {
		t.Error("hello is not equal to enbyte")
	}
}

// func TestBase64Decode(t * testing.T) {


// }