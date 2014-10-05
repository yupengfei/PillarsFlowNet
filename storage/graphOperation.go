package storage

import (
	"PillarsFlowNet/utility"
)


func InsertIntoGraph(graph * utility.Graph) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`INSERT INTO graph(graph_code, campaign_code, project_code, node_code, 
		width, height, x_coordinate, y_coordinate) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(graph.GraphCode, graph.CampaignCode, graph.ProjectCode, graph.NodeCode,
		graph.Width, graph.Height, graph.XCoordinate, graph.YCoordinate)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
		err = tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		return false, err
	}
	return true, err
}

func ModifyGraph(graph * utility.Graph) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`UPDATE graph SET campaign_code=?, project_code=?, node_code=?, 
		width=?, height=?, x_coordinate=?, y_coordinate=? WHERE graph_code=?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(graph.CampaignCode, graph.ProjectCode, graph.NodeCode,
		graph.Width, graph.Height, graph.XCoordinate, graph.YCoordinate, graph.GraphCode)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		panic(err.Error())
		err = tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		return false, err
	}
	return true, err
}


func DeleteGraphByGraphCode(graphCode * string) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`DELETE FROM graph WHERE graph_code = ?`)
	defer stmt.Close()
	_, err = stmt.Exec(graphCode)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		// pillarsLog.Logger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		return false, err
	}
	return true, err
}
func DeleteNodeByNodeCode(nodeCode * string) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`DELETE FROM graph WHERE node_code = ?`)
	defer stmt.Close()
	_, err = stmt.Exec(nodeCode)
	if err != nil {
		panic(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		// pillarsLog.Logger.Print(err.Error())
		err = tx.Rollback()
		if err != nil {
			panic(err.Error())
		}
		return false, err
	}
	return true, err
}

func QueryGraphNodesByCampaignCode(campaignCode * string) ([] utility.Graph, error) {
	stmt, err := DBConn.Prepare(`SELECT graph_code, campaign_code, project_code, node_code, width, height, x_coordinate, y_coordinate, 
		insert_datetime, update_datetime 
		FROM graph WHERE campaign_code = ?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(campaignCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var graphSlice [] utility.Graph
	for result.Next() {
		var graph utility.Graph
		err = result.Scan(&(graph.GraphCode), &(graph.CampaignCode), &(graph.ProjectCode), &(graph.NodeCode), &(graph.Width),
			&(graph.Height), &(graph.XCoordinate), &(graph.YCoordinate))
		if err != nil {
			panic(err.Error())
		}
		graphSlice = append(graphSlice, graph)
	}
	return graphSlice, err
}

func QueryGraphNodeByGraphCode(graphCode * string) (* utility.Graph, error) {
	stmt, err := DBConn.Prepare(`SELECT graph_code, campaign_code, project_code, node_code, width, height, x_coordinate, y_coordinate, 
		insert_datetime, update_datetime 
		FROM graph WHERE graph_code = ?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(graphCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var graph * utility.Graph
	if result.Next() {
		err = result.Scan(&(graph.GraphCode), &(graph.CampaignCode), &(graph.ProjectCode), &(graph.NodeCode), &(graph.Width),
			&(graph.Height), &(graph.XCoordinate), &(graph.YCoordinate))
		if err != nil {
			panic(err.Error())
		}
	}
	return graph, err
}