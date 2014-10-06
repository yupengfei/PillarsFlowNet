package connection

import (
	// "PillarsFlowNet/utility"
	// "github.com/gorilla/websocket"
	"PillarsFlowNet/project"
	"PillarsFlowNet/mission"
	"PillarsFlowNet/graph"
	"PillarsFlowNet/dependency"
	"PillarsFlowNet/target"
	"PillarsFlowNet/chart"
	"PillarsFlowNet/post"
	"PillarsFlowNet/user"
	"sync"
	"fmt"

)

//hub handle all kind of request
//add more channel to realize more kind of request
//ugly through, modify hub and connection both to add a request
type hub struct {

	connections map[string]*connection
	
	register chan * connection
	unregister chan * connection
	getAllProject chan * connection
	addProject chan * string
	modifyProject chan * string
	getAllCampaign chan * string
	addMission chan * string
	modifyMission chan * string
	deleteMission chan * string
	queryMissionByMissionCode chan * string
	getAllNode chan * string
	addNode chan * string
	modifyNode chan * string
	deleteNode chan * string
	getAllDependency chan * string
	addDependency chan * string
	deleteDependency chan * string
	modifyDependency chan * string
	addTarget chan * string
	modifyTarget chan * string
	deleteTarget chan * string
	queryTargetByMissionCode chan * string
	getAllUser chan * string

	addChart chan * string
	receiveChart chan * string
	getAllUnreceivedChart chan * string

	addPost chan * string
	getAllTargetPost chan * string


	productionMutex * sync.Mutex
	//postMutex * sync.Mutex
}

var Hub = hub {
	connections: make(map[string]*connection),
	
	register: make(chan * connection),
	unregister: make(chan * connection),
	getAllProject: make(chan * connection),
	addProject: make(chan * string),
	modifyProject: make(chan * string),
	getAllCampaign: make(chan * string),
	addMission: make(chan * string),
	modifyMission: make(chan * string),
	deleteMission: make(chan * string),
	queryMissionByMissionCode: make(chan * string),
	getAllNode: make(chan * string),
	addNode: make(chan * string),
	modifyNode: make(chan * string),
	deleteNode: make(chan * string),
	getAllDependency: make(chan * string),
	addDependency: make(chan * string),
	modifyDependency: make(chan * string),
	deleteDependency: make(chan * string),
	addTarget: make(chan * string),
	modifyTarget: make(chan * string),
	deleteTarget: make(chan * string),
	queryTargetByMissionCode: make(chan * string),

	getAllUser: make(chan * string),

	addChart: make(chan * string),
	receiveChart: make(chan * string),
	getAllUnreceivedChart: make(chan * string),

	addPost: make(chan * string),
	getAllTargetPost: make(chan * string),


	productionMutex: new(sync.Mutex),
}


func (h *hub) Run() {
	for {
		select {
		case c := <- h.register:
			fmt.Println(*(c.userCode))
			h.connections[*(c.userCode)] = c
		case c := <- h.unregister:
			close(c.send)
			//if x.userCode is not nil, then h.connections contains the conresponding connection
			if c.userCode != nil {
				if _, ok := h.connections[*(c.userCode)]; ok {				
					delete(h.connections, *(c.userCode))			
				}
			}
		case m := <- h.addChart:
			result, fromUserCode, toUserCode := chart.AddChart(m)
			h.SendToUserCode(result, fromUserCode)
			h.SendToUserCode(result, toUserCode)
		case m := <- h.receiveChart:
			result, userCode := chart.ReceiveChart(m)
			h.SendToUserCode(result, userCode)
		case m := <- h.getAllUnreceivedChart:
			result, userCode := chart.GetAllUnreceivedChart(m)
			h.SendToUserCode(result, userCode)
		case m := <- h.addPost:
			result, _ := post.AddPost(m)
			h.Dispatch(result)

		case m := <- h.getAllTargetPost:
			result, userCode := post.AddPost(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.getAllUser:
			result, userCode := user.GetAllUser(m)
			h.SendToUserCode(result, userCode)

		case c := <- h.getAllProject:
			c.send <- project.GetAllProject()	


		case m := <- h.addProject:
			h.productionMutex.Lock()
			result, _ := project.AddProject(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.modifyProject:
			h.productionMutex.Lock()
			result, _ := project.ModifyProject(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.getAllCampaign:
			result, userCode := mission.GetAllCampaign(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.addMission:
			h.productionMutex.Lock()
			result, _ := mission.AddMission(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.modifyMission:
			h.productionMutex.Lock()
			result, _ := mission.ModifyMission(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.deleteMission:
			h.productionMutex.Lock()
			result, _ := mission.DeleteMission(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.queryMissionByMissionCode:
			result, userCode := mission.QueryMissionByMissionCode(m)
			h.connections[*userCode].send <- result

		case m := <- h.getAllNode:
			result, userCode := graph.GetAllNode(m)
			h.connections[*userCode].send <- result

		case m := <- h.addNode:
			h.productionMutex.Lock()
			result, _ := graph.AddNode(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.modifyNode:
			h.productionMutex.Lock()
			result, _ := graph.ModifyNode(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.deleteNode:
			h.productionMutex.Lock()
			result, _ := graph.DeleteNode(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.getAllDependency:
			result, userCode := dependency.GetAllDependency(m)
			h.connections[*userCode].send <- result

		case m := <- h.addDependency:
			h.productionMutex.Lock()
			result, _ := dependency.AddDependency(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.deleteDependency:
			h.productionMutex.Lock()
			result, _ := dependency.DeleteDependency(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.modifyDependency:
			h.productionMutex.Lock()
			result, _ := dependency.ModifyDependency(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.addTarget:
			h.productionMutex.Lock()
			result, _ := target.AddTarget(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.modifyTarget:
			h.productionMutex.Lock()
			result, _ := target.ModifyTarget(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.deleteTarget:
			h.productionMutex.Lock()
			result, _ := target.DeleteTarget(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.queryTargetByMissionCode:
			result, userCode := target.QueryTargetByMissionCode(m)
			h.connections[*userCode].send <- result

		}

	}
}

func (h *hub) Dispatch(result []byte) {
	fmt.Println(string(result))
	for userCode := range h.connections {
		select {
		case h.connections[userCode].send <- result:
		default:
			close(h.connections[userCode].send)
			delete(h.connections, userCode)
		}
	}
}

func (h *hub) SendToUserCode(result []byte, userCode * string) {
	select {
	case h.connections[*userCode].send <- result:
	default:
		close(h.connections[*userCode].send)
		delete(h.connections, *userCode)
	}
}
