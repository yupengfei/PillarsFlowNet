package main

import (
	"PillarsFlowNet/connection"
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc("/", websocketServer)
	go connection.Hub.Run()
	http.HandleFunc("/", connection.ServeWs)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
