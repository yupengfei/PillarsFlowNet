package main

import (
    "log"
    "net/http"
    "PillarsFlowNet/connection"
)

func main() {
    // http.HandleFunc("/", websocketServer)
    go connection.Hub.Run()
    http.HandleFunc("/", connection.ServeWs)
    if err := http.ListenAndServe(":1234", nil); err!= nil {
        log.Fatal("ListenAndServe", err)
    }
}
