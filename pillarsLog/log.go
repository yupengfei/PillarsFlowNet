package pillarsLog

import "log"
import "os"
import "CGWorldlineWeb/utility"
import "fmt"

//Usage
//PillarsLogger.Print
var PillarsLogger *log.Logger

var outFile *os.File

func init() {

    if PillarsLogger == nil {
        propertyMap := utility.ReadProperty("../log.properties")
        logFileName := propertyMap["LogFile"]
        fmt.Println(logFileName)
        var err error
        outFile, err = os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)

        if err != nil {
            panic(err.Error())
        }

        PillarsLogger = log.New(outFile, "", log.Ldate|log.Ltime|log.Llongfile)
    }
}

func CloseLogFile() {
    outFile.Close()
}