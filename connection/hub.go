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
	"PillarsFlowNet/daily"
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
	getAllProject chan * string
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
	addDaily chan * string
	modifyDaily chan * string
	deleteDaily chan * string
	queryDailyByMissionCode chan * string
	getAllUser chan * string

	addChart chan * string
	receiveChart chan * string
	getAllUnreceivedChart chan * string

	addPost chan * string
	getAllTargetPost chan * string

	getPersonAllWaitingMission chan * string
	getPersonAllUndergoingMission chan * string
	getPersonAllReviewingMission chan * string
	getPersonAllFinishedMission chan * string


	productionMutex * sync.Mutex
	//postMutex * sync.Mutex
}

var Hub = hub {
	connections: make(map[string]*connection),
	
	register: make(chan * connection),
	unregister: make(chan * connection),
	getAllProject: make(chan * string),
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

	addDaily: make(chan * string),
	modifyDaily: make(chan * string),
	deleteDaily: make(chan * string),
	queryDailyByMissionCode: make(chan * string),
	getAllUser: make(chan * string),

	addChart: make(chan * string),
	receiveChart: make(chan * string),
	getAllUnreceivedChart: make(chan * string),

	addPost: make(chan * string),
	getAllTargetPost: make(chan * string),

	getPersonAllWaitingMission: make(chan * string),
	getPersonAllUndergoingMission: make(chan * string),
	getPersonAllReviewingMission: make(chan * string),
	getPersonAllFinishedMission: make(chan * string),

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
			result, userCode := post.GetAllTargetPost(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.getAllUser:
			result, userCode := user.GetAllUser(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.getAllProject:
			result, userCode := project.GetAllProject(m)	
			h.SendToUserCode(result, userCode)


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
			h.SendToUserCode(result, userCode)

		case m := <- h.getAllNode:
			result, userCode := graph.GetAllNode(m)
			h.SendToUserCode(result, userCode)

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
			h.SendToUserCode(result, userCode)

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
			h.SendToUserCode(result, userCode)

		case m := <- h.addDaily:
			h.productionMutex.Lock()
			result, _ := daily.AddDaily(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.modifyDaily:
			h.productionMutex.Lock()
			result, _ := daily.ModifyDaily(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.deleteDaily:
			h.productionMutex.Lock()
			result, _ := daily.DeleteDaily(m)
			h.Dispatch(result)
			h.productionMutex.Unlock()

		case m := <- h.queryDailyByMissionCode:
			result, userCode := daily.QueryDailyByMissionCode(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.getPersonAllWaitingMission:
			result, userCode := mission.GetPersonAllWaitingMission(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.getPersonAllUndergoingMission:
			result, userCode := mission.GetPersonAllUndergoingMission(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.getPersonAllReviewingMission:
			result, userCode := mission.GetPersonAllReviewingMission(m)
			h.SendToUserCode(result, userCode)

		case m := <- h.getPersonAllFinishedMission:
			result, userCode := mission.GetPersonAllFinishedMission(m)
			h.SendToUserCode(result, userCode)
		}

	}
}

func (h *hub) Dispatch(result []byte) {
	fmt.Println(string(result[:]))
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
	fmt.Println(string(result[:]))
	select {
	case h.connections[*userCode].send <- result:
	default:
		close(h.connections[*userCode].send)
		delete(h.connections, *userCode)
	}
}
