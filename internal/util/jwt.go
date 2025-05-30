package util

import (
	"errors"
	"go-board-api/config"
	"go-board-api/model"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(email string) (string, error) {
	hour, _ := strconv.Atoi(config.EnvVar.JWTAccessTokenExpires)

	exp := time.Now().Add(time.Hour * time.Duration(hour)).Unix()
	claims := &model.JWTAccessTokenClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	access_token, err := t.SignedString([]byte(config.EnvVar.JWTSecret))
	if err != nil {
		return "", err
	}

	return access_token, nil
}

func CreateRefreshToken() (string, error) {
	hour, _ := strconv.Atoi(config.EnvVar.JWTRefreshTokenExpires)

	exp := time.Now().Add(time.Hour * time.Duration(hour)).Unix()
	claims := &model.JWTRefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refresh_token, err := t.SignedString([]byte(config.EnvVar.JWTSecret))
	if err != nil {
		return "", err
	}

	return refresh_token, nil
}

func ValidateToken(token string) (bool, error) {
	p, err := jwt.Parse(token, func (token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvVar.JWTSecret), nil
	})

	if err != nil {
		return false, err
	}

	if !p.Valid {
		return false, errors.New("invalid jwt token")
	}

	return true, nil
}

func DecodeToken(token string) (*model.DecodedJWT, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvVar.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	decoded := &model.DecodedJWT{}

	decoded.Email = claims["email"].(string)
	decoded.ExpiresAt = claims["exp"].(float64)

	return decoded, nil
}