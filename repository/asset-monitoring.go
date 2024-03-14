package repository

import (
	"asset-tracker/model"
	"database/sql"
)

func AddMonitoringData(db *sql.DB, monitoring model.AssetMonitoringHistory) (err error) {
	sql := "INSERT INTO asset_monitoring_histories (asset_id, user_id, status, notes, created_at, updated_at) VALUES ($1 ,$2, $3, $4, NOW(), NOW())"
	errs := db.QueryRow(sql, monitoring.AssetID, monitoring.UserID, monitoring.Status, monitoring.Notes)

	return errs.Err()
}

func DeleteMonitoringData(db *sql.DB, assetId int) (err error) {
	sql := "DELETE FROM asset_monitoring_histories WHERE id = $1"
	errs := db.QueryRow(sql, assetId)

	return errs.Err()
}
