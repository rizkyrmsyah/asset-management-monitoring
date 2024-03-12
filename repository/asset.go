package repository

import (
	"asset-tracker/model"
	"database/sql"
)

func AddAsset(db *sql.DB, asset model.Asset) (err error) {
	sql := "INSERT INTO assets (name, code, in_date, source, location_id, created_at, updated_at) VALUES ($1 ,$2, $3, $4, $5, NOW(), NOW())"
	errs := db.QueryRow(sql, asset.Name, asset.Code, asset.InDate, asset.Source, asset.LocationID)

	return errs.Err()
}

func GetAllAsset(db *sql.DB) (assets []model.Asset, err error) {
	sql := "SELECT * FROM assets WHERE deleted_at IS NULL ORDER BY name ASC"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var assetData = model.Asset{}
		err = rows.Scan(&assetData.ID, &assetData.Name, &assetData.Code, &assetData.InDate, &assetData.Source, &assetData.CreatedAt, &assetData.UpdatedAt, &assetData.DeletedAt)
		if err != nil {
			return
		}

		assets = append(assets, assetData)
	}

	return
}

func GetDetailAssetByAssetId(db *sql.DB, id int) (asset *model.AssetDetail, err error) {
	var assetDetail model.AssetDetail
	var AssetMonitoringHistory model.AssetMonitoringHistory
	var AssetMonitoringHistoryData []model.AssetMonitoringHistory

	sql := "SELECT * FROM assets WHERE id = $1 AND deleted_at IS NULL"
	err = db.QueryRow(sql, id).Scan(&assetDetail.ID, &assetDetail.Name, &assetDetail.Code, &assetDetail.InDate, &assetDetail.Source, &assetDetail.CreatedAt, &assetDetail.UpdatedAt, &assetDetail.DeletedAt)
	if err != nil {
		return nil, err
	}

	sql2 := "SELECT * FROM asset_monitoring_histories WHERE asset_id = $1 ORDER BY id DESC"
	rows, err := db.Query(sql2, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&AssetMonitoringHistory.ID,
			&AssetMonitoringHistory.AssetID,
			&AssetMonitoringHistory.UserID,
			&AssetMonitoringHistory.Status,
			&AssetMonitoringHistory.Notes,
			&AssetMonitoringHistory.CreatedAt,
			&AssetMonitoringHistory.UpdatedAt,
		)
		if err != nil {
			return
		}
		AssetMonitoringHistoryData = append(AssetMonitoringHistoryData, AssetMonitoringHistory)
	}

	assetDetail.ControlHistories = AssetMonitoringHistoryData

	return &assetDetail, nil
}

func UpdateAsset(db *sql.DB, asset model.Asset) (err error) {
	sql := "UPDATE assets SET name = $2, code = $3, in_date = $4, source = $5, location_id= $6, updated_at = NOW() WHERE id = $1"
	errs := db.QueryRow(sql, asset.ID, asset.Name, asset.Code, asset.InDate, asset.Source, asset.LocationID)

	return errs.Err()
}

func DeleteAsset(db *sql.DB, assetId int) (err error) {
	sql := "UPDATE assets SET deleted_at = NOW() WHERE id = $1"
	errs := db.QueryRow(sql, assetId)

	return errs.Err()
}
