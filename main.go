package main

import (
	"v_admin_api_only_armillary/helper"
	"v_admin_api_only_armillary/router"

	"github.com/bandros/framework"
)

func main() {
	fw := framework.Init{}
	fw.Get()
	helper.SetEnv()
	r := fw.Begin
	//r.Use(middleware.Cors())
	router.Router(r)
	fw.Run()
}
