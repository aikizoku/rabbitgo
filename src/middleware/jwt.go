package middleware

import "net/http"

type JWTAuthenticator struct {
	Secret    []byte
	WhiteList []string
}

type JWTAuthenticatorClaims struct {
	ISS  string `json:"iss"  valid:"required"`
	SUB  string `json:"sub"  valid:"required"`
	AUD  string `json:"aud"  valid:"required"`
	JTI  string `json:"jti"  valid:"required"`
	IAT  int64  `json:"iat"  valid:"required"`
	EXP  int64  `json:"exp"  valid:"required"`
	User JWTAuthenticatorUserClaims
}

type JWTAuthenticatorUserClaims struct {
	ID        string
	CreatedAt int64
}

func Encode(jwt JWTAuthenticatorClaims) {

}

func Decode(jwt string) {

}

func (a *JWTAuthenticator) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// NewJWTAuthenticator ... JWTAuthenticatorを作成する
func NewJWTAuthenticator(jwtSec string, whiteList []string) *JWTAuthenticator {
	auth := &JWTAuthenticator{
		Secret:    []byte(jwtSec),
		WhiteList: whiteList,
	}
	return auth
}
