package main

import (
    "github.com/gorilla/websocket"
    "fmt"
    "log"
    "net/http"
)
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func websocketServer(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    for {
        messageType, p, err := conn.ReadMessage()
        fmt.Println(messageType)
        fmt.Println(p)
        if err != nil {
            return
        }
        result := "success"
        var test []byte= []byte(result)
        if err = conn.WriteMessage(messageType, test); err != nil {
            return
        }
    }
}

func main() {
    //http.Handle("/", websocket.Hander(websocketServer))
    http.HandleFunc("/", websocketServer)
    if err := http.ListenAndServe(":1234", nil); err!= nil {
        log.Fatal("ListenAndServe", err)
    }
}
