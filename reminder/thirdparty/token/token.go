package token

import "github.com/dgrijalva/jwt-go"

type JwtTokenClaim struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
	jwt.StandardClaims
}
