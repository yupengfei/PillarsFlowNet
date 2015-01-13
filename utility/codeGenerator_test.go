package utility

import (
	"testing"
	"fmt"
	"strconv"
	"math/rand"
)

func TestRandomInt(t * testing.T) {
	fmt.Println("tesing random int generetor " +  strconv.Itoa(RandomInt()))
	fmt.Println("tesing random int generetor " +  strconv.Itoa(RandomInt()))
}

func TestRandomInt2(t * testing.T) {
	fmt.Println("tesing random int generetor " +  strconv.Itoa(RandomInt()))
}

func TestGenerateCode(t * testing.T) {
	str := "testing code generetor"
	fmt.Println("tesing random int generetor " +  *(GenerateCode(&str)))
}

//use same seed and magic number will generate the same number
func TestGenerateCode2(t * testing.T) {
	var seed int64 = 123
	magicNumber := 1000
	rand.Seed(seed)
	fmt.Println("tesing same seed and magic number 1 " + strconv.Itoa(rand.Intn(magicNumber)))
}

func TestGenerateCode3(t * testing.T) {
	var seed int64 = 123
	magicNumber := 1000
	rand.Seed(seed)
	fmt.Println("tesing same seed and magic number 2 " + strconv.Itoa(rand.Intn(magicNumber)))
}