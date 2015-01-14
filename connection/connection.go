package connection

import (
	"github.com/gorilla/websocket"
	"time"
	"net/http"
	"fmt"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/projectLogic"
	"PillarsFlowNet/missionLogic"
	"PillarsFlowNet/graphLogic"
	"PillarsFlowNet/dependencyLogic"
	"PillarsFlowNet/targetLogic"
	"PillarsFlowNet/chartLogic"
	"PillarsFlowNet/postLogic"
	"PillarsFlowNet/userLogic"
	"PillarsFlowNet/dailyLogic"
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

			if *command == "getAllProject" {
				go projectLogic.GetAllProject(c.userCode, parameter, Hub)
			} else if *command == "addProject" {
				go projectLogic.AddProject(c.userCode, parameter, Hub)
			} else if *command == "modifyProject" {
				go projectLogic.ModifyProject(c.userCode, parameter, Hub)
			} else if *command == "getProjectCampaign" {
				go missionLogic.GetProjectCampaign(c.userCode, parameter, Hub)
			} else if *command == "addMission" {
				go missionLogic.AddMission(c.userCode, parameter, Hub)
			} else if *command == "modifyMission" {
				go missionLogic.ModifyMission(c.userCode, parameter, Hub)
			} else if *command == "deleteMission" {
				go missionLogic.ModifyProject(c.userCode, parameter, Hub)
			} else if *command == "getCampaignNode" {
				go graphLogic.GetCampaignNode(c.userCode, parameter, Hub)
			} else if *command == "addNode" {
				go graphLogic.AddNode(c.userCode, parameter, Hub)
			} else if *command == "modifyNode" {
				go graphLogic.ModifyNode(c.userCode, parameter, Hub)
			} else if *command == "deleteNode" {
				go graphLogic.DeleteNode(c.userCode, parameter, Hub)
			} else if *command == "getAllDependency" {
				go dependencyLogic.GetAllDependency(c.userCode, parameter, Hub)
			} else if *command == "addDependency" {
				go dependencyLogic.AddDependency(c.userCode, parameter, Hub)
			} else if *command == "modifyDependency" {
				go dependencyLogic.ModifyDependency(c.userCode, parameter, Hub)
			} else if *command == "deleteDependency" {
				go dependencyLogic.DeleteDependency(c.userCode, parameter, Hub)
			} else if *command == "addTarget" {
				go targetLogic.AddTarget(c.userCode, parameter, Hub)
			} else if *command == "modifyTarget" {
				go targetLogic.ModifyTarget(c.userCode, parameter, Hub)
			} else if *command == "deleteTarget" {
				go targetLogic.DeleteTarget(c.userCode, parameter, Hub)
			} else if *command == "getTargetByMissionCode" {
				go targetLogic.GetTargetByMissionCode(c.userCode, parameter, Hub)
			} else if *command == "addDaily" {
				go dailytLogic.AddDaily(c.userCode, parameter, Hub)
			} else if *command == "modifyDaily" {
				go dailytLogic.ModifyDaily(c.userCode, parameter, Hub)
			} else if *command == "deleteDaily" {
				go dailytLogic.DeleteDaily(c.userCode, parameter, Hub)
			} else if *command == "getDailyByMissionCode" {
				go dailytLogic.GetDailyByMissionCode(c.userCode, parameter, Hub)
			} else if *command == "getAllUser" {
				go userLogic.GetAllUser(c.userCode, parameter, Hub)
			} else if * command == "addChart" {
				go chartLogic.AddChart(c.userCode, parameter, Hub)
			} else if * command == "receiveChart" {
				go chartLogic.ReceiveChart(c.userCode, parameter, Hub)
			} else if * command == "getAllUnreceivedChart" {
				go chartLogic.GetAllUnreceivedChart(c.userCode, parameter, Hub)
			} else if * command == "addPost" {
				go postLogic.AddPost(c.userCode, parameter, Hub)
			} else if * command == "getAllTargetPost" {
				go postLogic.GetAllTargetPost(c.userCode, parameter, Hub)
			} else if * command == "getPersonAllWaitingMission" {
				go missionLogic.GetPersonAllWaitingMission(c.userCode, parameter, Hub)
			} else if * command == "getPersonAllUndergoingMission" {
				go missionLogic.GetPersonAllUndergoingMission(c.userCode, parameter, Hub)
			} else if * command == "getPersonAllReviewingMission" {
				go missionLogic.GetPersonAllReviewingMission(c.userCode, parameter, Hub)
			} else if * command == "getPersonAllFinishedMission" {
				go missionLogic.GetPersonAllFinishedMission(c.userCode, parameter, Hub)
			} else if * command == "getAllUndesignatedMission" {
				go missionLogic.GetAllUndesignatedMission(c.userCode, parameter, Hub)
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
