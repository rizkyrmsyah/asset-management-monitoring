package handler

import (
	"asset-tracker/middleware"
	"asset-tracker/model"
	"asset-tracker/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserHandler(router *gin.Engine) {
	user := router.Group("/user")
	user.POST("/register", handleRegister)
	user.POST("/login", handleLogin)
	user.GET("/profile", middleware.AuthUser(), usecase.GetUserProfile)
	user.PUT("/profile", middleware.AuthUser(), handleUpdateProfile)
}

func handleRegister(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = usecase.RegisterUser(c, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "registrasi berhasil",
	})
}

func handleLogin(c *gin.Context) {
	var loginRequest model.LoginRequest
	var loginResponse *model.LoginResponse
	var err error

	err = c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	loginResponse, err = usecase.LoginUser(c, loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    loginResponse,
	})
}

func handleUpdateProfile(c *gin.Context) {
	var user model.UpdateProfileRequest
	var err error

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = usecase.UpdateProfile(c, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update profile berhasil",
	})
}
