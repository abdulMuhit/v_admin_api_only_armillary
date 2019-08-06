package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"v_admin_api_only_armillary/model"

	"github.com/bandros/framework"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetJwt(email, password string) (string, error) {
	fmt.Println("email, pass", email, password)

	user, err := model.Login(email, password)
	fmt.Println("user, ", user)

	if err != nil {
		fmt.Println("error: ", err)
		return "", err
	}
	mybrand, err := model.GetMyBrandList(user["id"].(string))

	var list []string
	for _, u := range mybrand {
		list = append(list, u["id_brand"].(string))
		fmt.Println("u ", u["id_brand"].(string))
	}

	fmt.Println("list [] ", list)
	listbrand := []map[string]interface{}{}
	if len(list) != 0 {
		csv := strings.Join(list, ",")
		listb, err := model.GetAllBrandList(csv)
		listbrand = listb
		if err != nil {
			fmt.Println("error ", err)
			return "", err
		}
	} /*else {
		listbrand = []map[string]interface{}{
			{"role ": "SU"},
		}
	}*/

	fmt.Println("listbrand ", listbrand)
	expHour, _ := strconv.Atoi(framework.Config("jwtExp"))

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := sign.Claims.(jwt.MapClaims)
	claims["email"] = user["email"]
	claims["status"] = user["status"]
	claims["parent"] = user["parent"]
	/*	claims["no_member"] = user["no_member"]*/
	claims["brand_list"] = listbrand
	claims["role"] = user["role"]
	claims["nama"] = user["nama"]
	claims["id"] = user["id"]
	//claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(expHour)).Unix()
	token, err := sign.SignedString([]byte(framework.Config("jwtKey")))
	if err != nil {
		//framework.ErrorJson("JWT "+err.Error(),c)
		return "", err
	}
	fmt.Println("token ", token)
	return token, err
}

func LoginFrom(c *gin.Context) {
	email := c.DefaultPostForm("email", "")
	password := c.DefaultPostForm("password", "")
	fmt.Println("email ", email, " password ", password)

	token, err := GetJwt(email, password)

	if err != nil {
		fmt.Println("error1 ", err)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 300,
				"msg":  "username/password anda salah ",
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "correct login",
		"token": token,
	})
}
