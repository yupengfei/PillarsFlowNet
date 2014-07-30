package pillarsLog

import "log"
import "os"

var logFileName = "./PillarsFlowNet.log"

var Logger * log.Logger

var outFile * os.File

func init() {
	outFile, err := os.OpenFile(logFileName, os.O_RDWR, 0777)

	if err != nil {
		panic(err.Error())
	}
    // Ldate         = 1 << iota     // the date: 2009/01/23
    // Ltime                         // the time: 01:23:23
    // Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
    // Llongfile                     // full file name and line number: /a/b/c/d.go:23
    // Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
    // LstdFlags     = Ldate | Ltime // initial values for the standard logger

	Logger = log.New(outFile, "", log.Ldate|log.Ltime|log.Llongfile)
}

func Destory() {
    outFile.Close()
}