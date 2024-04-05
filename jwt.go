package kadi

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

// It contains commonly used fields
// if you wish, you can add your own struct to the custom field
type JwtModel struct {
	Custom any
}

var _tokenAuth *jwtauth.JWTAuth

func InitTokenAuth() {
	_tokenAuth = jwtauth.New("HS256", []byte(kadigo.appID), nil)
}

func GenerateToken(model JwtModel) (string, error) {
	modelAsMap, err := StructToMap(model)
	_, token, _ := _tokenAuth.Encode(modelAsMap)
	return token, err
}

func TokenAuth() *jwtauth.JWTAuth {
	return _tokenAuth
}

func GetUser(r *http.Request) JwtModel {
	_, claims, _ := jwtauth.FromContext(r.Context())

	var user JwtModel
	MapToStruct(claims, &user)

	return user
}
