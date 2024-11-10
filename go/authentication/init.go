package authentication

import "github.com/go-chi/jwtauth/v5"

var Token *jwtauth.JWTAuth

func init() {
	// TODO: replace with secret key
	Token = jwtauth.New("HS256", []byte("secret"), nil)
}
