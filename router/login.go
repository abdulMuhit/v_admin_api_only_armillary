package router

import (
	"v_admin_api_only_armillary/controller"

	"github.com/gin-gonic/gin"
)

func RouterHome(r *gin.Engine) {
	//rLogin.GET("/",controller.Login)
	r.POST("/process", controller.LoginFrom) //api
}
