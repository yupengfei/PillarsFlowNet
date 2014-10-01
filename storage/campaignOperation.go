package storage

import (
	"PillarsFlowNet/utility"
)


func InsertIntoCampaign(campaign * utility.Campaign) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`INSERT INTO campaign(campaign_code, project_code, node_code, 
		width, height, x_coordinate, y_coordinate) VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(campaign.CampaignCode, campaign.ProjectCode, campaign.NodeCode,
		campaign.Width, campaign.Height, campaign.XCoordinate, campaign.YCoordinate)
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


func DeleteCampaignByCampaignCode(campaignCode * string) (bool, error) {
	tx, err := DBConn.Begin()
	stmt, err := tx.Prepare(`DELETE FROM campaign WHERE campaign_code = ?`)
	defer stmt.Close()
	_, err = stmt.Exec(campaignCode)
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
	stmt, err := tx.Prepare(`DELETE FROM campaign WHERE node_code = ?`)
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

func QueryCampaignNodesByCampaignCode(campaignCode * string) ([] utility.Campaign, error) {
	stmt, err := DBConn.Prepare(`SELECT campaign_code, project_code, node_code, width, height, x_coordinate, y_coordinate, 
		insert_datetime, update_datetime 
		FROM campaign WHERE campaign_code = ?`)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	result, err := stmt.Query(campaignCode)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var campaignSlice [] utility.Campaign
	for result.Next() {
		var campaign utility.Campaign
		err = result.Scan(&(campaign.CampaignCode), &(campaign.ProjectCode), &(campaign.NodeCode), &(campaign.Width),
			&(campaign.Height), &(campaign.XCoordinate), &(campaign.YCoordinate))
		if err != nil {
			panic(err.Error())
		}
		campaignSlice = append(campaignSlice, campaign)
	}
	return campaignSlice, err
}