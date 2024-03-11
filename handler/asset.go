package handler

import (
	"asset-tracker/middleware"
	"asset-tracker/model"
	"asset-tracker/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AssetHandler(router *gin.Engine) {
	asset := router.Group("/asset").Use(middleware.AuthUser())
	asset.POST("/", handleCreateAsset)
	// asset.GET("/", usecase.GetUserProfile)
	// asset.PUT("/:id", usecase.UpdateProfile)
	// asset.DELETE("/:id", usecase.UpdateProfile)
}

func handleCreateAsset(c *gin.Context) {
	var asset model.Asset

	err := c.ShouldBindJSON(&asset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = usecase.CreateAsset(c, asset)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tambah asset berhasil",
	})
}
