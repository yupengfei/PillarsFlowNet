package connection

import (
	"PillarsFlowNet/pillarsLog"
	"PillarsFlowNet/utility"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
	// "PillarsFlowNet/chartLogic"
	// "PillarsFlowNet/dailyLogic"
	// "PillarsFlowNet/dependencyLogic"
	// "PillarsFlowNet/graphLogic"
	// "PillarsFlowNet/missionLogic"
	// "PillarsFlowNet/postLogic"
	// "PillarsFlowNet/projectLogic"
	// "PillarsFlowNet/targetLogic"
	// "PillarsFlowNet/userLogic"
)

const (
	// Time allowed to write a message to the client.
	writeWait = 100 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 600 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from client.
	maxMessageSize = 512 * 512000000
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
	ws       *websocket.Conn
	send     chan []byte
	userCode *string
}

// write writes a message with the given message type and message to client.
func (c *connection) write(mt int, message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, message)
}

//writePump pumps messages from the hub to the websocket connection
func (c *connection) writePump() {
	//ping the client, ifthere is no message in pingPeriod
	ticker := time.NewTicker(pingPeriod) //计时器
	defer func() {
		fmt.Println("writePump closing")
		ticker.Stop()
		c.ws.Close()
		fmt.Println("writePump closed")
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

//readPump pumps messages from the websocket connection to the hub
func (c *connection) readPump() {
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
			} else { //认为是执行登陆处理
				result, userCode, err := ValidateUser(parameter)
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

		} else { //else do some other command
			fmt.Println(*command)
			if *command == "getAllProject" {
				go GetAllProject(c.userCode, parameter)
			} else if *command == "addProject" {
				go AddProject(c.userCode, parameter)
			} else if *command == "modifyProject" {
				go ModifyProject(c.userCode, parameter)
			} else if *command == "getProjectAssertCampaign" {
				go GetProjectAssertCampaign(c.userCode, parameter) //
			} else if *command == "getProjectShotCampaign" {
				go GetProjectShotCampaign(c.userCode, parameter) ///////meiyou
			} else if *command == "addMission" {
				go AddMission(c.userCode, parameter)
			} else if *command == "getMissionByMissionCode" {
				go GetMissionByMissionCode(c.userCode, parameter)
			} else if *command == "modifyMission" {
				go ModifyMission(c.userCode, parameter)
			} else if *command == "deleteMission" {
				go DeleteMission(c.userCode, parameter) //ModifyProject()
			} else if *command == "getCampaignNode" {
				go GetCampaignNode(c.userCode, parameter)
			} else if *command == "addNode" {
				go AddNode(c.userCode, parameter)
			} else if *command == "modifyNode" {
				go ModifyNode(c.userCode, parameter)
			} else if *command == "deleteNode" {
				go DeleteNode(c.userCode, parameter)
			} else if *command == "getCampaignDependency" {
				go GetCampaignDependency(c.userCode, parameter)
			} else if *command == "addDependency" {
				go AddDependency(c.userCode, parameter)
			} else if *command == "modifyDependency" {
				go ModifyDependency(c.userCode, parameter)
			} else if *command == "deleteDependency" {
				go DeleteDependency(c.userCode, parameter)
			} else if *command == "addTarget" {
				go AddTarget(c.userCode, parameter)
			} else if *command == "modifyTarget" {
				go ModifyTarget(c.userCode, parameter)
			} else if *command == "deleteTarget" {
				go DeleteTarget(c.userCode, parameter)
			} else if *command == "getTargetByMissionCode" {
				go GetTargetByMissionCode(c.userCode, parameter)
			} else if *command == "addDaily" {
				go AddDaily(c.userCode, parameter)
			} else if *command == "modifyDaily" {
				go ModifyDaily(c.userCode, parameter)
			} else if *command == "deleteDaily" {
				go DeleteDaily(c.userCode, parameter)
			} else if *command == "getCompanyDaily" {
				go GetCompanyDaily(c.userCode, parameter)
			} else if *command == "getDailyByMissionCode" {
				go GetDailyByMissionCode(c.userCode, parameter)
			} else if *command == "getAllUser" {
				go GetAllUser(c.userCode, parameter)
			} else if *command == "addChart" {
				go AddChart(c.userCode, parameter)
			} else if *command == "receiveChart" {
				go ReceiveChart(c.userCode, parameter)
			} else if *command == "getAllUnreceivedChart" {
				go GetAllUnreceivedChart(c.userCode, parameter)
			} else if *command == "addPost" {
				go AddPost(c.userCode, parameter)
			} else if *command == "getAllTargetPost" {
				go GetAllTargetPost(c.userCode, parameter)
			} else if *command == "getPersonAllWaitingMission" {
				go GetPersonAllWaitingMission(c.userCode, parameter)
			} else if *command == "getPersonAllUndergoingMission" {
				go GetPersonAllUndergoingMission(c.userCode, parameter)
			} else if *command == "getPersonAllReviewingMission" {
				go GetPersonAllReviewingMission(c.userCode, parameter)
			} else if *command == "getPersonAllFinishedMission" {
				go GetPersonAllFinishedMission(c.userCode, parameter)
			} else if *command == "getAllUndesignatedMission" {
				go GetAllUndesignatedMission(c.userCode, parameter)
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
