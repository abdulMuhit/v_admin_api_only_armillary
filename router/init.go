package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
	"v_admin_api_only_armillary/controller"

	"github.com/bandros/framework"
	"github.com/cxww107/asyncwork/worker"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.NoRoute(error404)
	//r.NoRoute(noRoute)
	r.NoMethod(error404)

	RouterApi(r)

	// r.POST("/daftar", controller.AddVendor)     //api
	r.POST("/login", controller.LoginFrom) //api
	r.GET("/testAsync", func(c *gin.Context) {
		result := doSomeWork()
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"result": result,
		})
	})
}

/*
func noRoute(c *gin.Context)  {
	c.Redirect(http.StatusNotFound, "index")
}
*/

func error404(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	login := false
	if v != nil {
		login = true
	}

	js := []string{
		"/asset/js/popper.min",
		"/asset/js/errorpagenotfound",
	}

	c.HTML(http.StatusNotFound, "error/404", gin.H{
		"title": "Error 404",
		"login": login,
		"js":    js,
	})
}

func doSomeWork() string {

	// Define three slow functions which will be runnning concurrently
	slowFunction := func() interface{} {
		time.Sleep(time.Second * 10)
		fmt.Println("slow function")
		return 2
	}

	verySlowFunction := func() interface{} {
		time.Sleep(time.Second * 5)
		fmt.Println("very slow function")
		return "I'm ready"
	}

	// One function returns an error
	errorFunction := func() interface{} {
		time.Sleep(time.Second * 16)
		fmt.Println("function with an error")
		return errors.New("Error in function")
	}

	tasks := []worker.TaskFunction{slowFunction, verySlowFunction, errorFunction}

	// Use context to cancel goroutines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resultChannel := worker.PerformTasks(ctx, tasks)
	// Print value from first goroutine and cancel others
	var s string
	for result := range resultChannel {
		switch result.(type) {
		case error:
			fmt.Println("Received error")
			cancel()
			s = "Received error"
			return s
		case string:
			fmt.Println("Here is a string:", result.(string))
			s = "Here is a string: " + result.(string)
			return s
		case int:
			fmt.Println("Here is an integer:", result.(int))
			s = "Here is an integer:" + result.(string)
			return s
		default:
			fmt.Println("Some unknown type ")
			s = "Some unknown type"
			return s
		}
	}

	return s

}
