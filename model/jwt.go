package model

import "github.com/golang-jwt/jwt"

type JWT struct {
	AccessToken		string		`json:"access_token"`
	RefreshToken	string		`json:"refresh_token"`
}

type DecodedJWT struct {
	Email					string				`json:"email"`
	ExpiresAt			float64				`json:"exp"`
}

type JWTAccessTokenClaims struct {
	Email		string		`json:"email"`
	jwt.StandardClaims
}

type JWTRefreshTokenClaims struct {
	jwt.StandardClaims
}