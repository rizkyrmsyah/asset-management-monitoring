package repository

import (
	"asset-tracker/model"
	"database/sql"
)

func AddAsset(db *sql.DB, asset model.Asset) (err error) {
	sql := "INSERT INTO assets (name, code, in_date, source, created_at, updated_at) VALUES ($1 ,$2, $3, $4, NOW(), NOW())"
	errs := db.QueryRow(sql, asset.Name, asset.Code, asset.InDate, asset.Source)

	return errs.Err()
}

func GetAllAsset(db *sql.DB) (assets []model.Asset, err error) {
	sql := "SELECT * FROM assets ORDER BY name ASC"
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
