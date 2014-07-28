package main

import (
    "log"
    "net/http"
    "PillarsFlowNet/connection"
)


// func websocketServer(w http.ResponseWriter, r *http.Request) {
//     conn, err := upgrader.Upgrade(w, r, nil)
//     if err != nil {
//         fmt.Println(err)
//         return
//     }
//     for {
//         //
//         messageType, p, err := conn.ReadMessage()
//         fmt.Println(messageType)
//         fmt.Println(string(p))
//         if err != nil {
//             return
//         }
//         result := "success"
//         var test []byte= []byte(result)
//         if err = conn.WriteMessage(messageType, test); err != nil {
//             return
//         }
//     }

// }

func main() {
    // http.HandleFunc("/", websocketServer)
    go connection.Hub.Run()
    http.HandleFunc("/", connection.ServeWs)
    if err := http.ListenAndServe(":1234", nil); err!= nil {
        log.Fatal("ListenAndServe", err)
    }
}
