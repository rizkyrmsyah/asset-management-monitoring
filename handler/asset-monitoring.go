package handler

import (
	"asset-tracker/middleware"
	"asset-tracker/model"
	"asset-tracker/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AssetMonitoringHandler(router *gin.Engine) {
	monitoring := router.Group("/monitoring").Use(middleware.AuthUser())
	monitoring.POST("/", handleSubmitMonitoring)
	monitoring.DELETE("/:id", handleDeleteMonitoring)
}

func handleSubmitMonitoring(c *gin.Context) {
	var monitoringData model.AssetMonitoringHistory

	err := c.ShouldBindJSON(&monitoringData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = usecase.AddMonitoringData(c, monitoringData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "input monitoring asset berhasil",
	})
}

func handleDeleteMonitoring(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = usecase.DeleteMonitoringDatat(c, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "hapus monitoring asset berhasil",
	})
}
