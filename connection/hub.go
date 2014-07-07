package connection

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type hub struct {

	connections map[string]*connection
	transmit chan []byte
	login chan * connection
	logout chan * connection
}

var h = hub {
	connections: make(map[string]*connection)
	transmit: make(chan [] byte)
	login: make(chan * connection)
	logout: make(chan * connection)
}

func (h *hub) run() {
	for {
		select {
		case c := <- h.login:
			h.connections[&c.User.UserName] = c
		case c := <- h.logout:
			if _, ok := h.connections[&c.User.UserName]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := h.transmit:
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
