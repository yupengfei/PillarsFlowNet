package log

import "log"
import "os"

var logFileName = "../PillarsFlowNet.log"

var logger * log.logger

func init() {
	outFile, err := os.Open("logFileName")

	if err != nil {
		panic(err.Error())
	}
	defer outFile.Close()
	//unfinished
	logger = log.New()
	
}