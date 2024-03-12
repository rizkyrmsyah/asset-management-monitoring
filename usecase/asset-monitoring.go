package usecase

import (
	"asset-tracker/database"
	"asset-tracker/model"
	"asset-tracker/repository"
	"errors"

	"github.com/gin-gonic/gin"
)

func AddMonitoringData(c *gin.Context, monitoring model.AssetMonitoringHistory) (err error) {
	asset, err := repository.GetDetailAssetByAssetId(database.DbConnection, monitoring.AssetID)
	if err != nil {
		return errors.New(err.Error())
	}

	if asset == nil {
		return errors.New("asset tidak ditemukan")
	}

	sessionData := c.MustGet("session").(*model.JwtCustomClaims)
	monitoring.UserID = sessionData.ID

	err = repository.AddMonitoringData(database.DbConnection, monitoring)
	if err != nil {
		return errors.New(err.Error())
	}

	return
}

func DeleteMonitoringDatat(c *gin.Context, monitoringId int) (err error) {
	err = repository.DeleteMonitoringData(database.DbConnection, monitoringId)
	if err != nil {
		return errors.New(err.Error())
	}

	return
}
