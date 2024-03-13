package usecase

import (
	"errors"
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

func RegisterUser(c *gin.Context, user model.User) (err error) {
	hashedPassword, err := helper.Bcrypt(user.Password)
	if err != nil {
		return errors.New(err.Error())
	}

	user.Password = hashedPassword
	err = repository.AddUser(database.DbConnection, user)
	if err != nil {
		return errors.New("email sudah terdaftar di sistem kami")
	}

	return
}

func LoginUser(c *gin.Context, loginRequest model.LoginRequest) (loginResponse *model.LoginResponse, err error) {
	inputPassword := loginRequest.Password
	user, err := repository.FindUserByEmail(database.DbConnection, loginRequest.Email)
	if err != nil {
		return nil, errors.New("email atau password yang kamu masukkan kurang tepat")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	if err != nil {
		return nil, errors.New("email atau password yang kamu masukkan kurang tepat")
	}

	accessToken, exp, err := createAccessToken(user)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	res := &model.LoginResponse{
		AccessToken: accessToken,
		Exp:         exp,
	}

	return res, nil
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

func UpdateProfile(c *gin.Context, user model.UpdateProfileRequest) (err error) {
	if user.Password != nil {
		hashedPassword, err := helper.Bcrypt(*user.Password)
		if err != nil {
			return errors.New(err.Error())
		}
		user.Password = &hashedPassword
	}

	sessionData := c.MustGet("session").(*model.JwtCustomClaims)
	user.ID = sessionData.ID

	err = repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		return errors.New(err.Error())
	}

	return
}

func GetUserProfile(c *gin.Context) {
	sessionData := c.MustGet("session").(*model.JwtCustomClaims)
	user, err := repository.FindUserById(database.DbConnection, sessionData.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email atau password yang kamu masukkan kurang tepat.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
	})
}
