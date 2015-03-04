package connection

import (
	"PillarsFlowNet/userStorage"
	"fmt"
)

//hub handle all kind of request
type HubStruct struct {
	connections map[string]*connection

	register   chan *connection
	unregister chan *connection
}

var Hub = HubStruct{
	connections: make(map[string]*connection),

	register:   make(chan *connection),
	unregister: make(chan *connection),
}

func (h *HubStruct) Run() {
	for {
		select {
		case c := <-h.register:
			fmt.Println(*(c.userCode))
			h.connections[*(c.userCode)] = c
		case c := <-h.unregister:
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

func (h *HubStruct) Dispatch(result []byte, code string) {
	err, users := userStorage.QueryCompanyUser(&code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Dispatch Company Numbers: ", len(users))
	for _, value := range users {
		if h.connections[*value] != nil {
			h.connections[*value].send <- result
		}
	}

	//for userCode := range h.connections {
	//	h.connections[userCode].send <- result
	//}
}

func (h *HubStruct) SendToUserCode(result []byte, userCode *string) {
	//fmt.Println(string(result[:]))
	fmt.Println("send to user", *userCode)
	if _, ok := h.connections[*userCode]; ok {
		h.connections[*userCode].send <- result
	} else {
		fmt.Println("can not find user")
	}
}
