package token

import "github.com/go-chi/jwtauth/v5"

var Token *jwtauth.JWTAuth

func init() {
	Token = jwtauth.New("HS256", []byte("secret"), nil)
}
