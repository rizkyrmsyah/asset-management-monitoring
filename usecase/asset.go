package usecase

import (
	"asset-tracker/database"
	"asset-tracker/model"
	"asset-tracker/repository"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAsset(c *gin.Context, asset model.Asset) (err error) {
	time, err := time.Parse("2006-01-02", asset.InDate)
	if err != nil {
		return errors.New(err.Error())
	}
	asset.InDate = time.Format("2006-01-02")
	err = repository.AddAsset(database.DbConnection, asset)
	if err != nil {
		return errors.New("kode asset sudah terdaftar")
	}

	return
}

func GetAllAsset(c *gin.Context) (assets []model.Asset, err error) {
	assets, err = repository.GetAllAsset(database.DbConnection)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}
