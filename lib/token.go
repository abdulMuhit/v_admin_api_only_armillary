package lib

import (
	"github.com/bandros/framework"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func GenerateTokenMember(user map[string]interface{}) (string, error) {
	expHour, _ := strconv.Atoi(framework.Config("jwtExp"))
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["jti"] = framework.Password(user["id"].(string) + string(time.Now().Unix()))
	claims["id"] = user["id"]
	claims["nama"] = user["nama"]
	claims["email"] = user["email"]
	claims["no_hp"] = user["no_hp"]
	claims["status"] = user["status"]
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expHour)).Unix()
	return sign.SignedString([]byte(framework.Config("jwtKeyApi")))
}

func GenerateTokenAdmin(admin map[string]interface{}) (string, error) {
	expHour, _ := strconv.Atoi(framework.Config("jwtExp"))
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["jti"] = framework.Password(admin["id"].(string) + string(time.Now().Unix()))
	claims["id"] = admin["id"]
	claims["nama"] = admin["nama"]
	claims["email"] = admin["email"]
	claims["username"] = admin["username"]
	claims["level"] = admin["level"]
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expHour)).Unix()
	return sign.SignedString([]byte(framework.Config("jwtKeyApi")))
}
