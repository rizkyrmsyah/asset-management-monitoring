package handler

import (
	"asset-tracker/middleware"
	"asset-tracker/model"
	"asset-tracker/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LocationHandler(router *gin.Engine) {
	location := router.Group("/location").Use(middleware.AuthUser())
	location.POST("/", handleCreateLocation)
	location.GET("/", handleGetAllLocation)
	location.PUT("/:id", handleUpdateLocation)
	location.DELETE("/:id", handleDeleteLocation)
}

func handleCreateLocation(c *gin.Context) {
	var location model.Location

	err := c.ShouldBindJSON(&location)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = usecase.CreateLocation(c, location)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tambah lokasi berhasil",
	})
}

func handleGetAllLocation(c *gin.Context) {
	res, err := usecase.GetAllLocation(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    res,
	})
}

func handleUpdateLocation(c *gin.Context) {
	var locationRequest model.LocationUpdateRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&locationRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	locationRequest.ID = id
	err = usecase.UpdateLocation(c, locationRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ubah lokasi berhasil",
	})
}

func handleDeleteLocation(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = usecase.DeleteLocation(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "hapus lokasi berhasil",
	})
}
