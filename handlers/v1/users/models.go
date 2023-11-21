package v1Users

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
	ExpirationHours   int64
}

type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

type UserRegistrationRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegistrationResponse struct {
	Username        string  `json:"username"`
	Email           string  `json:"email"`
	CurrentBalance  float32 `json:"current_balance"`
	StartingBalance float32 `json:"starting_balance"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
