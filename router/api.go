package router

import (
	"v_admin_api_only_armillary/controller"
	"v_admin_api_only_armillary/middleware"

	"github.com/gin-gonic/gin"
)

func RouterApi(r *gin.Engine) {
	rApi := r.Group("/api")
	rApi.Use(middleware.AuthApi)
	// todo why?
	// rApi.POST("/list_files", controller.ExportImage) //api

	//rDashboard.POST("/get_list_pengguna", controller.SUGetAllPengguna) //api todo not sure this is needed
	rApi.POST("/get_list_pengguna_datatables", controller.GetPengguna) //api
	rApi.POST("/get_list_brand", controller.GetAllBrand)               //api

	rApi.POST("/delete_pengguna", controller.DeletePengguna) //api
	rApi.POST("/tambah_pengguna", controller.TambahPengguna) //api
	rApi.POST("/sign_pengguna", controller.SignPengguna)     //api

	rApi.POST("/gantipassword", controller.GantiPassword) //api

	// waiting list vendor
	rVendor := rApi.Group("/vendor")

	//rVendor.GET("/getAllData", controller.GetAllVendorData) //api todo not sure this is needed
	rVendor.POST("/getAllDataFilter", controller.GetVendorDataFiltered) //api
	rVendor.GET("/getData", controller.GetVendorData)                   // api

	rVendor.GET("/edit_vendor_get_pic", controller.GetVendorPic) //api
	// todo not yet implemented for gcloud
	//rVendor.POST("/edit_vendor", controller.EditVendor) //api

	rVendor.POST("/change_status_vendor", controller.ChangeStatusVendor)       //api
	rVendor.POST("/exportExcelWaitingList", controller.ExportExcelWaitingList) // api

}
