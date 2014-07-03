package utility

import "testing"

import "bytes"


func TestMd5sum(t *testing.T) {
    test := Md5sum("123456")
    //fmt.Printf("%x", test)
    result := []byte("e10adc3949ba59abbe56e057f20f883e")
    if len(test) != len(result) {
        t.Error("Md5sum length mismatch")
    } else {
        if ! bytes.Equal(test, result) {
            t.Error("Md5sum wrong")
        }
    }
}
