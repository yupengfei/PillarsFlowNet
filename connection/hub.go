package connection

import (
	"PillarsFlowNet/utility"
	// "github.com/gorilla/websocket"
	"PillarsFlowNet/project"
	"PillarsFlowNet/mission"
	// "PillarsFlowNet/campaign"
	// "PillarsFlowNet/dependency"
	// "PillarsFlowNet/target"
	// "time"
	"fmt"

)

//hub handle all kind of request
//add more channel to realize more kind of request
//ugly through, modify hub and connection both to add a request
type hub struct {

	connections map[string]*connection
	chart chan []byte
	post chan []byte
	register chan * connection
	unregister chan * connection
	getAllProject chan * connection
	addProject chan * string
	modifyProject chan * string
	getAllCompaign chan * string
	addMission chan * string
	modifyMission chan * string
	deleteMission chan * string
	getAllNode chan * string
	addNode chan * string
	modifyNode chan * string
	deleteNode chan * string
	getAllDependency chan * string
	addDependency chan * string
	deleteDependency chan * string
	addTarget chan * string
	modifyTarget chan * string
	deleteTarget chan * string
	searchTargetByMissionCode chan * string
}

var Hub = hub {
	connections: make(map[string]*connection),
	chart: make(chan [] byte),
	post: make(chan [] byte),
	register: make(chan * connection),
	unregister: make(chan * connection),
	getAllProject: make(chan * connection),
	addProject: make(chan * string),
	modifyProject: make(chan * string),
	getAllCompaign: make(chan * string),
	addMission: make(chan * string),
	modifyMission: make(chan * string),
	deleteMission: make(chan * string),
	getAllNode: make(chan * string),
	addNode: make(chan * string),
	modifyNode: make(chan * string),
	deleteNode: make(chan * string),
	getAllDependency: make(chan * string),
	addDependency: make(chan * string),
	deleteDependency: make(chan * string),
	addTarget: make(chan * string),
	modifyTarget: make(chan * string),
	deleteTarget: make(chan * string),
	searchTargetByMissionCode: make(chan * string),
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
			//should be moved to  chart logic
			fmt.Println(string(m))
			chartMessage, chartToPerson, error := utility.ParseChartMessage(m)
			if (error != nil) {
				return
			}
			fmt.Println(*chartToPerson)
			connection, ok := Hub.connections[*chartToPerson];
			if ok {
				connection.send <- []byte(*chartMessage)
			}
		case c := <- h.getAllProject:
			c.send <- project.GetAllProject()	

		case m := <- h.addProject:
			result, userCode := project.AddProject(m)
			h.connections[*userCode].send <- result

		case m := <- h.modifyProject:
			result, userCode := project.ModifyProject(m)
			h.connections[*userCode].send <- result

		case m := <- h.getAllCompaign:
			result, userCode := mission.GetAllCompaign(m)
			h.connections[*userCode].send <- result

		case m := <- h.addMission:
			result, userCode := mission.AddMission(m)
			h.connections[*userCode].send <- result

		case m := <- h.modifyMission:
			result, userCode := mission.ModifyMission(m)
			h.connections[*userCode].send <- result

		case m := <- h.deleteMission:
			result, userCode := mission.DeleteMission(m)
			h.connections[*userCode].send <- result
			
		// case m := <- h.getAllNode:
		// 	c.send <- campaign.GetAllNode(m)

		// case m := <- h.addNode:
		// 	c.send <- campaign.AddNode(m)

		// case m := <- h.modifyNode:
		// 	c.send <- campaign.ModifyNode(m)

		// case m := <- h.deleteNode:
		// 	c.send <- campaign.DeleteNode(m)

		// case m := <- h.getAllDependency:
		// 	c.send <- dependency.GetAllDependency(m)

		// case m := <- h.addDependency:
		// 	c.send <- dependency.AddDependency(m)

		// case m := <- h.deleteDependency:
		// 	c.send <- dependency.DeleteDependency(m)

		// case m := <- h.addTarget:
		// 	c.send <- target.AddTarget(m)

		// case m := <- h.modifyTarget:
		// 	c.send <- target.ModifyTarget(m)

		// case m := <- h.deleteTarget:
		// 	c.send <- target.DeleteTarget(m)

		// case m := <- h.searchTargetByMissionCode:
		// 	c.send <- target.SearchTargetByMissionCode(m)

		}

	}
}
