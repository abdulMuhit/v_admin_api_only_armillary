package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//JSON lib
func JSON(c *gin.Context, code int, msg string, data interface{}) {
	var status = c.GetHeader("header_status")
	if status == "1" {
		c.JSON(code, gin.H{
			"code": code,
			"msg":  msg,
			"data": data,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
