package picture

import (
	"fmt"
	"testing"
)

func TestBase64Bytes(t * testing.T) {
	hello := "hello world"
	debyte := Base64Encode([]byte(hello))

	enbyte := Base64Decode(debyte)

}

// func TestBase64Decode(t * testing.T) {
	

// }