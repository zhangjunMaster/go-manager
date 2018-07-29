package auth

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

const (
	SecretKey = "dan-Gan*pO-jIe+wo#De!keY"
)

//func New(method SigningMethod) *Token
//type Token struct {
//		Raw       string                 // The raw token.  Populated when you Parse a token
//		Method    SigningMethod          // The signing method used or to be used
//		Header    map[string]interface{} // The first segment of the token
//		Claims    Claims                 // The second segment of the token
//		Signature string                 // The third segment of the token.  Populated when you Parse a token
//		Valid     bool                   // Is the token valid?  Populated when you Parse/Verify a token
//}
//type MapClaims map[string]interface{}
//Claims: make(jwt.MapClaims) 生成的Claims 一般将数据挂载到Claims上

// 生成token
func CreateToken(admin string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(30)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["admin"] = admin
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

// parse token
func ParseToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SecretKey), nil
		})

	if err == nil {
		//token.Claims转成jwt.MapClaims类型
		//exp是interface{}，类型断言成int64才能比较
		exp := token.Claims.(jwt.MapClaims)["exp"]
		v := exp.(int64)
		if token.Valid && v > time.Now().Unix() {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}
