package helper

import (
	"os"
	"strconv"
	"time"

	"github.com/bandros/framework"
	"github.com/dgrijalva/jwt-go"
)

func CreateTokenApiJwt(from, to string) (string, error) {
	expMinute, _ := strconv.Atoi(framework.Config("jwtApiExp"))

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["from"] = from
	claims["to"] = to
	//claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expMinute)).Unix()
	token, err := sign.SignedString([]byte("bandrostokenkey"))
	if err != nil {
		return "", err
	}
	return token, err
}

func SetEnv() {
	forTarget := os.Getenv("TARGET")
	mode := os.Getenv("MODE")
	switch forTarget {
	case "gcloud":
		switch mode {
		case "A1":
			os.Setenv("portHost", "8080")
			mysqlHostLocalGcloud := os.Getenv("mysqlHostLocalGcloud")
			os.Setenv("mysqlGoogle", "0")
			os.Setenv("mysqlHost", mysqlHostLocalGcloud)
		}
	default:
		panic("unrecognized escape character")
	case "linode":
		switch mode {
		case "B1":
			os.Setenv("portHost", "8080")
			mysqlHostLocal := os.Getenv("mysqlHostLocal")
			os.Setenv("mysqlGoogle", "0")
			os.Setenv("mysqlHost", mysqlHostLocal)

			os.Setenv("storageLinkDefault", "http://localhost:8080")
			os.Setenv("storageBucketVendorDefault", "public")
			os.Setenv("storageBucketFolderDefalut", "img")
		default:
			panic("unrecognized escape character")
		}
	}

}
