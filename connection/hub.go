package connection

import (
	"fmt"
)

//hub handle all kind of request
//add more channel to realize more kind of request
//ugly through, modify hub and connection both to add a request
type HubStruct struct {

	connections map[string]*connection
	
	register chan * connection
	unregister chan * connection
}

var Hub = hub {
	connections: make(map[string]*connection),
	
	register: make(chan * connection),
	unregister: make(chan * connection),
}


func (h *HubStruct) Run() {
	for {
		select {
		case c := <- h.register:
			fmt.Println(*(c.userCode))
			h.connections[*(c.userCode)] = c
		case c := <- h.unregister:
			//close(c.send)
			//if x.userCode is not nil, then h.connections contains the conresponding connection
			if c.userCode != nil {
				if _, ok := h.connections[*(c.userCode)]; ok {				
					delete(h.connections, *(c.userCode))			
				}
			}
		}

	}
}

func (h *HubStruct) Dispatch(result []byte) {
	//fmt.Println(string(result[:]))
	fmt.Println("Dispatch")
	for userCode := range h.connections {
		h.connections[userCode].send <- result
	}
}

func (h *HubStruct) SendToUserCode(result []byte, userCode * string) {
	//fmt.Println(string(result[:]))
	fmt.Println("send to user", *userCode)
	if _, ok := h.connections[*userCode]; ok {				
		h.connections[*userCode].send <- result			
	} else {
		fmt.Println("can not find user")
	}
	
}
