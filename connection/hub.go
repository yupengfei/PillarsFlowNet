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
	// "time"
	"fmt"

)

//hub handle all kind of request
//add more channel to realize more kind of request
//ugly through, modify hub and connection both to add a request
type hub struct {

	connections map[string]*connection
	chart chan * string
	post chan * string
	register chan * connection
	unregister chan * connection
	getAllProject chan * connection
	addProject chan * string
	modifyProject chan * string
	getAllCompaign chan * string
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
}

var Hub = hub {
	connections: make(map[string]*connection),
	chart: make(chan * string),
	post: make(chan * string),
	register: make(chan * connection),
	unregister: make(chan * connection),
	getAllProject: make(chan * connection),
	addProject: make(chan * string),
	modifyProject: make(chan * string),
	getAllCompaign: make(chan * string),
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
		case m := <- h.chart:
			result, userCode := chart.Chart(m)
			h.connections[*userCode].send <- result	
		case m := <- h.post:
			post.Post(m)
			//h.connections[*userCode].send <- result	
		case c := <- h.getAllProject:
			c.send <- project.GetAllProject()	


		case m := <- h.addProject:
			result, _ := project.AddProject(m)
			h.Dispatch(result)

		case m := <- h.modifyProject:
			result, _ := project.ModifyProject(m)
			h.Dispatch(result)

		case m := <- h.getAllCompaign:
			result, userCode := mission.GetAllCompaign(m)
			h.connections[*userCode].send <- result

		case m := <- h.addMission:
			result, _ := mission.AddMission(m)
			h.Dispatch(result)

		case m := <- h.modifyMission:
			result, _ := mission.ModifyMission(m)
			h.Dispatch(result)

		case m := <- h.deleteMission:
			result, _ := mission.DeleteMission(m)
			h.Dispatch(result)

		case m := <- h.queryMissionByMissionCode:
			result, userCode := mission.QueryMissionByMissionCode(m)
			h.connections[*userCode].send <- result

		case m := <- h.getAllNode:
			result, userCode := graph.GetAllNode(m)
			h.connections[*userCode].send <- result

		case m := <- h.addNode:
			result, _ := graph.AddNode(m)
			h.Dispatch(result)

		case m := <- h.modifyNode:
			result, _ := graph.ModifyNode(m)
			h.Dispatch(result)

		case m := <- h.deleteNode:
			result, _ := graph.DeleteNode(m)
			h.Dispatch(result)

		case m := <- h.getAllDependency:
			result, userCode := dependency.GetAllDependency(m)
			h.connections[*userCode].send <- result

		case m := <- h.addDependency:
			result, _ := dependency.AddDependency(m)
			h.Dispatch(result)

		case m := <- h.deleteDependency:
			result, _ := dependency.DeleteDependency(m)
			h.Dispatch(result)

		case m := <- h.modifyDependency:
			result, _ := dependency.ModifyDependency(m)
			h.Dispatch(result)

		case m := <- h.addTarget:
			result, _ := target.AddTarget(m)
			h.Dispatch(result)

		case m := <- h.modifyTarget:
			result, _ := target.ModifyTarget(m)
			h.Dispatch(result)

		case m := <- h.deleteTarget:
			result, _ := target.DeleteTarget(m)
			h.Dispatch(result)

		case m := <- h.queryTargetByMissionCode:
			result, userCode := target.QueryTargetByMissionCode(m)
			h.connections[*userCode].send <- result

		}

	}
}

func (h *hub) Dispatch(result []byte) {
	for userCode := range h.connections {
		select {
		case h.connections[userCode].send <- result:
		default:
			close(h.connections[userCode].send)
			delete(h.connections, userCode)
		}
	}
}
