package graphStorage

import (
	"CGWorldlineWeb/pillarsLog"
	"PillarsFlowNet/mysqlUtility"
	"PillarsFlowNet/utility"
)

func InsertIntoGraph(graph *utility.Graph) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`INSERT INTO graph(graph_code, campaign_code, project_code, node_code, product_type ,
		width, height, x_coordinate, y_coordinate) VALUES(?, ?, ?, ?, ?, ?, ?, ?,?)`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(graph.GraphCode, graph.CampaignCode, graph.ProjectCode, graph.NodeCode, graph.ProductType,
		graph.Width, graph.Height, graph.XCoordinate, graph.YCoordinate)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

func ModifyGraph(graph *utility.Graph) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`UPDATE graph SET campaign_code=?, project_code=?, node_code=?, product_type=?,
		width=?, height=?, x_coordinate=?, y_coordinate=? WHERE graph_code=?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(graph.CampaignCode, graph.ProjectCode, graph.NodeCode, graph.ProductType,
		graph.Width, graph.Height, graph.XCoordinate, graph.YCoordinate, graph.GraphCode)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

func DeleteGraphByGraphCode(graphCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`DELETE FROM graph WHERE graph_code = ?`)
	defer stmt.Close()
	_, err = stmt.Exec(graphCode)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}
func DeleteNodeByNodeCode(nodeCode *string) (bool, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`DELETE FROM graph WHERE node_code = ?`)
	defer stmt.Close()
	_, err = stmt.Exec(nodeCode)
	if err != nil {
		panic(err.Error())
	}
	return true, err
}

func QueryAssertNodesByCampaignCode(projectCode *string) ([]utility.Graph, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT graph_code, campaign_code,
		project_code,node_code,product_type,x_coordinate,y_coordinate,width,height  FROM graph where campaign_code = ? AND is_campaign=1 AND product_type=1`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var graphSlice []utility.Graph
	for result.Next() {
		var graph utility.Graph
		err = result.Scan(&(graph.GraphCode), &(graph.CampaignCode), &(graph.ProjectCode),
			&(graph.NodeCode), &(graph.ProductType),
			&(graph.XCoordinate), &(graph.YCoordinate), &(graph.Width), &(graph.Height))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		graphSlice = append(graphSlice, graph)
	}
	return graphSlice, err
}
func QueryShotNodesByCampaignCode(projectCode *string) ([]utility.Graph, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT graph_code, campaign_code,
		project_code,node_code,product_type,x_coordinate,y_coordinate,width,height  FROM graph where campaign_code = ? AND is_campaign=1 AND product_type=0`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(projectCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	//this is easy to imply but may not very fast
	var graphSlice []utility.Graph
	for result.Next() {
		var graph utility.Graph
		err = result.Scan(&(graph.GraphCode), &(graph.CampaignCode), &(graph.ProjectCode),
			&(graph.NodeCode), &(graph.ProductType),
			&(graph.XCoordinate), &(graph.YCoordinate), &(graph.Width), &(graph.Height))
		if err != nil {
			pillarsLog.PillarsLogger.Print(err.Error())
		}
		graphSlice = append(graphSlice, graph)
	}
	return graphSlice, err
}
func QueryGraphNodesByCampaignCode(campaignCode *string, types int) ([]utility.Graph, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT graph_code, campaign_code, project_code, node_code, product_type,width, height, x_coordinate, y_coordinate, 
		insert_datetime, update_datetime 
		FROM graph WHERE campaign_code = ? AND product_type=?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(campaignCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var graphSlice []utility.Graph
	for result.Next() {
		var graph utility.Graph
		err = result.Scan(&(graph.GraphCode), &(graph.CampaignCode), &(graph.ProjectCode), &(graph.NodeCode), &(graph.ProductType), &(graph.Width),
			&(graph.Height), &(graph.XCoordinate), &(graph.YCoordinate), &(graph.InsertDatetime), &(graph.UpdateDatetime), types)
		if err != nil {
			panic(err.Error())
		}
		graphSlice = append(graphSlice, graph)
	}
	return graphSlice, err
}

func QueryGraphNodeByGraphCode(graphCode *string) (*utility.Graph, error) {
	stmt, err := mysqlUtility.DBConn.Prepare(`SELECT graph_code, campaign_code, project_code, node_code,product_type, width, height, x_coordinate, y_coordinate, 
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
	var graph utility.Graph
	if result.Next() {
		err = result.Scan(&(graph.GraphCode), &(graph.CampaignCode), &(graph.ProjectCode), &(graph.NodeCode), &(graph.ProductType), &(graph.Width),
			&(graph.Height), &(graph.XCoordinate), &(graph.YCoordinate), &(graph.InsertDatetime), &(graph.UpdateDatetime))
		if err != nil {
			panic(err.Error())
		}
	}
	return &graph, err
}
