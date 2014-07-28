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
			fmt.Println(*(c.userName))
			h.connections[*(c.userName)] = c
		case c := <- h.unregister:
			close(c.send)
			//if x.userName is not nil, then h.connections contains the conresponding connection
			if c.userName != nil {
				if _, ok := h.connections[*(c.userName)]; ok {				
					delete(h.connections, *(c.userName))			
				}
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