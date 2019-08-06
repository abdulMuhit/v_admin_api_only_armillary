package middleware

import (
	"fmt"
	"net/http"
	"v_admin_api_only_armillary/lib"

	"github.com/bandros/framework"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func notfoundToken(c *gin.Context) {
	c.Abort()
	c.JSON(http.StatusExpectationFailed, gin.H{
		"code": http.StatusForbidden,
		"data": "token not found",
	})
	return
}

func AuthApi(c *gin.Context) {
	tokenString := c.GetHeader("token")
	fmt.Println("tokenstring ", tokenString)
	if tokenString == "" {
		fmt.Println("token not found")
		c.Abort()
		return
	}

	/*
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != token.Method {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(framework.Config("jwtApiKey")), nil
		})
	*/

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(framework.Config("jwtKey")), nil
	})

	if err != nil {
		fmt.Println("token not found " + err.Error())
		c.Abort()
		return
	}

	fmt.Println("token ", token)

}

func Auth(c *gin.Context) {

	var tokenString string
	var client bool
	session := sessions.Default(c)
	if v := session.Get(framework.Config("jwtName")); v == nil {
		tokenString = c.GetHeader("token")
		client = true
		fmt.Println("token in header ", tokenString)
		if tokenString == "" {
			fmt.Println("token not found")
			notfoundToken(c)
		}
	} else {
		client = false
		tokenString = v.(string)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(framework.Config("jwtKey")), nil
	})
	if err != nil {
		fmt.Println("token is " + err.Error())
		if client {
			notfoundToken(c)
		}
		// notfound(session, c)

		lib.JSON(c, http.StatusNotFound, "Not found token", gin.H{})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
		var data_jwt map[string]interface{}
		data_jwt = map[string]interface{}{}
		data_jwt["email"] = claims["email"]
		data_jwt["status"] = claims["status"]
		data_jwt["parent"] = claims["parent"]
		data_jwt["brand_list"] = claims["brand_list"]
		data_jwt["nama"] = claims["nama"]
		data_jwt["role"] = claims["role"]
		data_jwt["id"] = claims["id"]
		c.Set("jwt", data_jwt)
	} else {
		fmt.Println("not found claims")
		lib.JSON(c, http.StatusNotFound, "Not found token", gin.H{})
		// notfound(session, c)
		c.Abort()
		return
	}

}

func Cors() gin.HandlerFunc {
	fmt.Println("cors middleware ")
	return func(c *gin.Context) {

		/*
			c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Add("Access-Control-Allow-Credentials", "true")
			c.Writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
			c.Writer.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		*/

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
