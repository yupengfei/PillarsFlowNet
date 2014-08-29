package connection

import (
	"github.com/gorilla/websocket"
	"time"
	"net/http"
	"fmt"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/pillarsLog"
)

const (
	// Time allowed to write a message to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from client.
	maxMessageSize = 512*512
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
	userName * string
	// group * string
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
		Hub.unregister <- c
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
			pillarsLog.Logger.Println("parse in message error")
			pillarsLog.Logger.Println(error.Error())
		}
		//if userCode is nil, login first
		fmt.Println(*command)
		fmt.Println(*parameter)
		if c.userName == nil {
			if *command != "login" {
				return
			} else {
				user, error := utility.ParseLoginInMessage(parameter)
				if error != nil {
					pillarsLog.Logger.Println("parse login message error")
				}
				validLogin := storage.CheckUserNameAndPassword(&((*user).UserName), &((*user).Password))
				var sysError = utility.Error {
						ErrorCode: 0,
						ErrorMessage: "",
				}
				var loginMessage utility.LoginInMessage
				if validLogin {
					c.userName = &((*user).UserName)
					Hub.register <- c
					
					loginMessage = utility.LoginInMessage {
						Auth: "success",
						AuthMessage : "",
					}
				} else {
					loginMessage = utility.LoginInMessage {
						Auth: "failed",
						AuthMessage : "userName or Password wrong",
					}
				}
				loginStr := string(utility.LoginMessageToJson(loginMessage))					
				var out = utility.OutMessage {
						Error: sysError,
						Command: "login",
						Result: loginStr,
					}
				c.send <- utility.LoginMessageToJson(out)

			}

		} else {//else do some other command
			if *command == "chart" {
				Hub.chart <- []byte(*parameter)
			} /*else if *command == "getAllProject" {
				go 
			}*/
			

		}

		//h.broadcast <- message
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
	c := &connection{send: make(chan []byte), ws: ws, userName: nil}
	go c.writePump()
	c.readPump()
}
