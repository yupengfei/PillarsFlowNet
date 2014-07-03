package utility

import "os"
import "fmt"
import "bufio"
import "strings"
import "io"

func ReadProperty(fileName string) map[string]string {
    file, err := os.Open(fileName)
    defer file.Close()
    if err != nil {
        fmt.Println(fileName, err)
        return nil
    }
    buff := bufio.NewReader(file) //读入缓存
    propertyMap := make(map[string]string)
    for {
        line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
        if err != nil || io.EOF == err {
            break
        }
        //fmt.Print(line)  //可以对一行进行处理
        propertyPair := strings.Split(line, "=")
        //fmt.Println(propertyPair[0])
        propertyMap[propertyPair[0]] = propertyPair[1]
    }
    return propertyMap
}
