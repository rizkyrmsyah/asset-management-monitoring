package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Exp         int64  `json:"exp"`
}

type UpdateProfileRequest struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Password *string `json:"password"`
}

type JwtCustomClaims struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
