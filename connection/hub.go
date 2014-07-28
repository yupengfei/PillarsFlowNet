package connection

import (
	// "github.com/gorilla/websocket"
	"PillarsFlowNet/pillarsLog"
	// "time"
	"fmt"
)

//hub handle all kind of request
//add more channel to realize more kind of request
//ugly through, modify hub and connection both to add a request
type hub struct {

	connections map[string]*connection
	chart chan []byte
	register chan * connection
	unregister chan * connection
}

var Hub = hub {
	connections: make(map[string]*connection),
	chart: make(chan [] byte),
	register: make(chan * connection),
	unregister: make(chan * connection),
}


func (h *hub) Run() {
	for {
		select {
		case c := <- h.register:
			fmt.Println(*(c.userCode))
			h.connections[*(c.userCode)] = c
		case c := <- h.unregister:
			if _, ok := h.connections[*(c.userCode)]; ok {
				close(c.send)
				delete(h.connections, *(c.userCode))			
			}
		case m := <- h.chart:
			pillarsLog.Logger.Println(m)
			//TODO
			//tansmit
			// select {
			// case c.send <- m:
			// default:
			// 	close(c.send)
			// 	delete(h.connections, c)
			// }
		}
	}
}
