package encoding

import (
	"fmt"
	"testing"
	"os"
	"io/ioutil"
)

func TestBase64Bytes(t * testing.T) {
	hello := "对比 world"
	debyte := Base64Encode([]byte(hello))

	enbyte, err := Base64Decode(debyte)
	if err != nil {
		panic(err.Error())
	}
	if hello != string(enbyte) {
		t.Error("hello is not equal to enbyte")
	}
	fmt.Print(string(enbyte))
}

func TestBase64Jpg(t * testing.T) {
	file, err := os.Open("encode_test.jpg")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	picture, err := ioutil.ReadAll(file)

	debyte := Base64Encode(picture)

	enbyte, err := Base64Decode(debyte)
	if err != nil {
		panic(err.Error())
	}
	if len(picture) != len(enbyte) {
		t.Error("picture encoding or decoding wrong")
	}
	// file, err := os.Open("encode_test_save.jpg")
	ioutil.WriteFile("encode_test_save.jpg", picture, 0644)


}