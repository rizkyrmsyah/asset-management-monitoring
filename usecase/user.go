package usecase

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"asset-tracker/database"
	"asset-tracker/helper"
	"asset-tracker/model"
	"asset-tracker/repository"
)

func RegisterUser(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	hashedPassword, err := helper.Bcrypt(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Password = hashedPassword
	err = repository.AddUser(database.DbConnection, user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email sudah terdaftar di sistem kami",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "registrasi berhasil",
	})
}

func LoginUser(c *gin.Context) {
	var loginRequest model.LoginRequest
	var loginResponse model.LoginResponse
	var err error

	err = c.ShouldBindJSON(&loginRequest)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	inputPassword := loginRequest.Password
	user, err := repository.FindUserByEmail(database.DbConnection, loginRequest.Email)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email atau password yang kamu masukkan kurang tepat.",
		})
		return
	}

	accessToken, exp, err := createAccessToken(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	loginResponse.AccessToken = accessToken
	loginResponse.Exp = exp

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    loginResponse,
	})
}

func createAccessToken(user *model.User) (accessToken string, exp int64, err error) {
	var jwtSecret = os.Getenv("JWT_SECRET")
	jwtTTL, err := strconv.Atoi(os.Getenv("JWT_TTL"))
	if err != nil {
		return
	}

	exp = time.Now().Add(time.Second * time.Duration(jwtTTL)).Unix()
	claims := &model.JwtCustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(jwtSecret))

	return
}
