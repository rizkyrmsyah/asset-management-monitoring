package usecase

import (
	"asset-tracker/database"
	"asset-tracker/model"
	"asset-tracker/repository"
	"errors"

	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context, location model.Location) (err error) {
	err = repository.CreateLocation(database.DbConnection, location)
	if err != nil {
		return errors.New(err.Error())
	}

	return
}

func GetAllLocation(c *gin.Context) (assets []model.Location, err error) {
	assets, err = repository.GetAllLocation(database.DbConnection)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return
}

func UpdateLocation(c *gin.Context, location model.LocationUpdateRequest) (err error) {
	err = repository.UpdateLocation(database.DbConnection, location)
	if err != nil {
		return errors.New(err.Error())
	}

	return
}

func DeleteLocation(c *gin.Context, locationId int) (err error) {
	err = repository.DeleteLocation(database.DbConnection, locationId)
	if err != nil {
		return errors.New(err.Error())
	}

	return
}
