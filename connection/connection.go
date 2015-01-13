package connection

import (
	"github.com/gorilla/websocket"
	"time"
	"net/http"
	"fmt"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/user"
	"PillarsFlowNet/pillarsLog"
)

const (
	// Time allowed to write a message to the client.
	writeWait = 100 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 600 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from client.
	maxMessageSize = 512*512000000
)

//use package websocket provied by gorilla, 
//it will upgrade http connection to web socket connection
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// connection is an middleman between the websocket connection and the hub.
//It consists of a connection and a pointer to userCode
//and a []byte channel, which is used for hub.
type connection struct {
	ws * websocket.Conn
	send chan []byte
	userCode * string
}

// write writes a message with the given message type and message to client.
func (c *connection) write(mt int, message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, message)
}

//writePump pumps messages from the hub to the websocket connection
func (c *connection) writePump() {
	//ping the client, ifthere is no message in pingPeriod
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		fmt.Println("writePump closing")
		ticker.Stop()
		c.ws.Close()
		fmt.Println("writePump closed")
	}()
	for {
		select {
		case message, ok := <- c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <- ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}
//readPump pumps messages from the websocket connection to the hub
func (c * connection) readPump() {
	defer func() {
		fmt.Println("readPump closing")
		Hub.unregister <- c
		close(c.send)
		c.ws.Close()
		fmt.Println("readPump closed")

	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		command, parameter, error := utility.ParseInMessage(message)
		if error != nil {
			pillarsLog.PillarsLogger.Println("parse in message error")
			pillarsLog.PillarsLogger.Println(error.Error())
		}
		//if userCode is nil, login first
		//fmt.Println(*command)
		//fmt.Println(*parameter)
		if c.userCode == nil {
			if *command != "login" {
				continue
			} else {
				result, userCode, err := user.ValidateUser(parameter)
				fmt.Println(*userCode)
				fmt.Println(*result)

				if err == nil {
					fmt.Println("*result")
					c.send <- []byte(*result)
					if *userCode != "" {
						c.userCode = userCode
						Hub.register <- c
					}
				}
			}

		} else {//else do some other command
			userCodeAndParameter := *(c.userCode) + "@" + *parameter;
			if *command == "getAllProject" {
				Hub.getAllProject <- &userCodeAndParameter
			} else if *command == "addProject" {
				Hub.addProject <- &userCodeAndParameter
			} else if *command == "modifyProject" {
				Hub.modifyProject <- &userCodeAndParameter
			} else if *command == "getAllCampaign" {
				Hub.getAllCampaign <- &userCodeAndParameter
			} else if *command == "addMission" {
				Hub.addMission <- &userCodeAndParameter
			} else if *command == "modifyMission" {
				Hub.modifyMission <- &userCodeAndParameter
			} else if *command == "deleteMission" {
				Hub.deleteMission <- &userCodeAndParameter
			} else if *command == "getAllNode" {
				Hub.getAllNode <- &userCodeAndParameter
			} else if *command == "addNode" {
				Hub.addNode <- &userCodeAndParameter
			} else if *command == "modifyNode" {
				Hub.modifyNode <- &userCodeAndParameter
			} else if *command == "deleteNode" {
				Hub.deleteNode <- &userCodeAndParameter
			} else if *command == "getAllDependency" {
				Hub.getAllDependency <- &userCodeAndParameter
			} else if *command == "addDependency" {
				Hub.addDependency <- &userCodeAndParameter
			} else if *command == "modifyDependency" {
				Hub.modifyDependency <- &userCodeAndParameter
			} else if *command == "deleteDependency" {
				Hub.deleteDependency <- &userCodeAndParameter
			} else if *command == "addTarget" {
				Hub.addTarget <- &userCodeAndParameter
			} else if *command == "modifyTarget" {
				Hub.modifyTarget <- &userCodeAndParameter
			} else if *command == "deleteTarget" {
				Hub.deleteTarget <- &userCodeAndParameter
			} else if *command == "queryTargetByMissionCode" {
				Hub.queryTargetByMissionCode <- &userCodeAndParameter
			} else if *command == "addDaily" {
				Hub.addTarget <- &userCodeAndParameter
			} else if *command == "modifyDaily" {
				Hub.modifyTarget <- &userCodeAndParameter
			} else if *command == "deleteDaily" {
				Hub.deleteTarget <- &userCodeAndParameter
			} else if *command == "queryDailyByMissionCode" {
				Hub.queryTargetByMissionCode <- &userCodeAndParameter
			} else if *command == "getAllUser" {
				Hub.getAllUser <- &userCodeAndParameter
			} else if * command == "addChart" {
				Hub.addChart <- &userCodeAndParameter
			} else if * command == "receiveChart" {
				Hub.receiveChart <- &userCodeAndParameter
			} else if * command == "getAllUnreceivedChart" {
				Hub.getAllUnreceivedChart <- &userCodeAndParameter
			} else if * command == "addPost" {
				Hub.addPost <- &userCodeAndParameter
			} else if * command == "getAllTargetPost" {
				Hub.getAllTargetPost <- &userCodeAndParameter
			} else if * command == "getPersonAllWaitingMission" {
				Hub.getPersonAllWaitingMission <- &userCodeAndParameter
			} else if * command == "getPersonAllUndergoingMission" {
				Hub.getPersonAllUndergoingMission <- &userCodeAndParameter
			} else if * command == "getPersonAllReviewingMission" {
				Hub.getPersonAllReviewingMission <- &userCodeAndParameter
			} else if * command == "getPersonAllFinishedMission" {
				Hub.getPersonAllFinishedMission <- &userCodeAndParameter
			} else if * command == "getAllUndesignatedMission" {
				Hub.getAllUndesignatedMission <- &userCodeAndParameter
			}
		}
	}
}



// serverWs handles webocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err.Error())
		return
	}
	c := &connection{send: make(chan []byte), ws: ws, userCode: nil}
	go c.writePump()
	c.readPump()
}
