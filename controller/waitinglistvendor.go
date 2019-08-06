package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"v_admin_api_only_armillary/model"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/bandros/framework"
	"github.com/gin-gonic/gin"
)

/*
todo: add vendor?
func AddVendor(c *gin.Context) {

	result, err := model.AddVendor(c)
	if err != nil {
		//framework.ErrorJson(err.Error(),c)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Register vendor gagal " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Berhasil input data vendor " + result["new_id"].(string),
	})
} */

func GetVendorDataFiltered(c *gin.Context) {
	var data model.DatatablesSource
	err := c.Bind(&data)
	fmt.Println("vendor data filtered ", data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(),c)
		return
	}

	extra := map[string]string{
		"start":  c.DefaultPostForm("startDate", ""),
		"end":    c.DefaultPostForm("endDate", ""),
		"status": c.DefaultPostForm("status", "all"),
	}

	fmt.Println("str date, ", c.DefaultPostForm("startDate", ""))

	fields := []string{
		"a.id",
		"a.nama",
		"a.no_hp",
		"a.email",
		"a.alamat",
		"a.nama_brand",
		"a.kota",
		"a.provinsi",
		"a.kodepos",
		"a.website",
		"a.facebook",
		"a.instagram",
		"a.marketplace",
		"a.kategori",
		"a.tanggal",
		"a.status",
	}

	result, err := model.FilterVendor(fields, data, extra)
	if err != nil {
		fmt.Println("error a ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(),c)
		return
	}

	c.JSON(http.StatusOK, &result)

}
func ExportExcelWaitingList(c *gin.Context) {
	var data model.DatatablesSource
	err := c.Bind(&data)
	fmt.Println("vendor data filtered ", data)

	if err != nil {
		framework.ErrorJson(err.Error(), c)
		/*	c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})*/
		return
	}

	extra := map[string]string{
		"start":  c.DefaultPostForm("startDate", ""),
		"end":    c.DefaultPostForm("endDate", ""),
		"status": c.DefaultPostForm("status", "all"),
	}

	fmt.Println("str date, ", c.DefaultPostForm("startDate", ""))

	fields := []string{
		"a.id",
		"a.nama",
		"a.no_hp",
		"a.email",
		"a.alamat",
		"a.nama_brand",
		"a.kota",
		"a.provinsi",
		"a.kodepos",
		"a.website",
		"a.facebook",
		"a.instagram",
		"a.marketplace",
		"a.kategori",
		"a.tanggal",
		"a.status",
	}

	result, err := model.FilterVendorExcel(fields, data, extra)
	if err != nil {
		fmt.Println("error a ", err)
		/*	c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})*/
		framework.ErrorJson(err.Error(), c)
		return
	}

	fmt.Println("result ", result)

	categories := map[string]string{
		"A7": "No", "B7": "ID", "C7": "Nama", "D7": "Nama Brand", "E7": "No hp", "F7": "Email", "G7": "Alamat",
		"H7": "Kota", "I7": "Provinsi", "J7": "Kodepos", "K7": "Website", "L7": "Facebook", "M7": "Instagram",
		"N7": "Marketplace", "O7": "Kategori", "P7": "Tanggal", "Q7": "Status",
	}

	valu := map[string]string{}

	fmt.Println("resutl len ", len(result))

	for i := 0; i < len(result); i++ {
		k := 8 + i
		j := 1 + i
		//fmt.Println("dfdsfd ", result[i]["id"])

		valu["A"+strconv.Itoa(k)] = strconv.Itoa(j)
		valu["B"+strconv.Itoa(k)] = result[i]["id"].(string)
		valu["C"+strconv.Itoa(k)] = result[i]["nama"].(string)
		valu["D"+strconv.Itoa(k)] = result[i]["nama_brand"].(string)
		valu["E"+strconv.Itoa(k)] = result[i]["no_hp"].(string)
		valu["F"+strconv.Itoa(k)] = result[i]["email"].(string)
		valu["G"+strconv.Itoa(k)] = result[i]["alamat"].(string)
		valu["H"+strconv.Itoa(k)] = result[i]["kota"].(string)
		valu["I"+strconv.Itoa(k)] = result[i]["provinsi"].(string)
		valu["J"+strconv.Itoa(k)] = result[i]["kodepos"].(string)
		valu["K"+strconv.Itoa(k)] = result[i]["website"].(string)
		valu["L"+strconv.Itoa(k)] = result[i]["facebook"].(string)
		valu["M"+strconv.Itoa(k)] = result[i]["instagram"].(string)
		valu["N"+strconv.Itoa(k)] = result[i]["marketplace"].(string)
		valu["O"+strconv.Itoa(k)] = result[i]["kategori"].(string)
		valu["P"+strconv.Itoa(k)] = result[i]["tanggal"].(string)
		valu["Q"+strconv.Itoa(k)] = result[i]["status"].(string)
	}
	fmt.Println("value w e ", valu)

	current_time := time.Now().Local()
	fmt.Println("The Current time is ", current_time.Format("2006-01-02"))
	var time = current_time.Format("2006-01-02")

	xlsx := excelize.NewFile()

	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center","ident":1,"justify_last_line":true,"reading_order":0,"relative_indent":1,"shrink_to_fit":true,"vertical":"top","wrap_text":true}}`)
	if err != nil {
		fmt.Println("error ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(),c)
		return
	}

	xlsx.SetCellStyle("Sheet1", "a1", "q1", style)

	style, err = xlsx.NewStyle(`{"border":[{"type":"left","color":"000000","style":-1},{"type":"top","color":"000000","style":5},{"type":"bottom","color":"000000","style":5},{"type":"right","color":"000000","style":6}]}`)
	if err != nil {
		fmt.Println("error ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(),c)
		return
	}

	vcel := "Q" + strconv.Itoa(len(result)+8)

	xlsx.SetCellStyle("Sheet1", "A7", vcel, style)

	xlsx.SetCellValue("Sheet1", "A1", "Waiting List Vendor")
	xlsx.MergeCell("Sheet1", "A1", "Q1")

	xlsx.SetCellValue("Sheet1", "A2", "Keterangan: ")
	xlsx.SetCellValue("Sheet1", "A3", "Status : ")
	xlsx.SetCellValue("Sheet1", "B3", extra["status"])

	xlsx.SetCellValue("Sheet1", "A4", "Dari Tanggal ")
	xlsx.SetCellValue("Sheet1", "B4", extra["start"])
	xlsx.SetCellValue("Sheet1", "C4", "Sampai Tanggal ")
	xlsx.SetCellValue("Sheet1", "D4", extra["end"])
	xlsx.SetCellValue("Sheet1", "A5", "Filter ")
	xlsx.SetCellValue("Sheet1", "B5", data.Keyword)

	xlsx.SetCellValue("Sheet1", "O2", "Tanggal ambil data ")
	xlsx.SetCellValue("Sheet1", "P2", time)
	for k, v := range categories {
		xlsx.SetCellValue("Sheet1", k, v)
	}

	for k, v := range valu {
		xlsx.SetCellValue("Sheet1", k, v)
	}

	c.Writer.Header().Set("Content-Type", "application/octet-stream")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Writer.Header().Set("Content-Transfer-Encoding", "binary")
	c.Writer.Header().Set("Expires", "0")
	xlsx.Write(c.Writer)

}

func GetVendorData(c *gin.Context) {
	id := c.Query("id")
	data, err := model.DataVendorOne(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(),c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func GetVendorPic(c *gin.Context) {
	id := c.Query("myid")
	tipe := c.Query("tipe")

	fmt.Println("getvendorpic ", id, " tipe ", tipe)

	data, err := model.GetVendorPic(id, tipe)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(),c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "get picture ok",
		"data": data,
	})

}

func ChangeStatusVendor(c *gin.Context) {
	result, err := model.ChangeStatusVendor(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": result,
	})
}

// todo ganti flexible, for all edit
func EditVendor(c *gin.Context) {
	result, err := model.EditVendorM(c)
	if err != nil {
		//framework.ErrorJson(err.Error(), c)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": result,
	})
}

// todo ganti flexible
func DeleteVendor(c *gin.Context) {

	myid := c.PostForm("myid")
	tipe := c.PostForm("tipe")

	res := model.DeleteAllVendorPic(myid, tipe)

	if res != nil {
		fmt.Println("error: ", res)
		c.JSON(http.StatusOK, gin.H{
			"code": http.Error,
			"msg":  fmt.Sprintf("Error deleted user: %s", res),
		})
		return
	}

	err := model.DeleteVendor(myid)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  fmt.Sprintf("Successfully deleted user: %s", err),
	})

}

// todo ganti flexible
func DeleteOneVendorPic(c *gin.Context) {
	myid := c.PostForm("id")
	url := c.PostForm("url")
	res, err := model.DeleteOneVendorPic(myid, url)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotModified,
			"result": "Error " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": res,
	})

}

// todo ganti flexible
func AddOneVendorPic(c *gin.Context) {

	//fmt.Println(reflect.TypeOf(file))

	res, err := model.AddOneVendorPic(c)
	if err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusOK, gin.H{
			"code":   http.Error,
			"result": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": res,
	})

}

func DeleteAllVendorPic(c *gin.Context) {
	owner := c.PostForm("owner")
	tipe := c.PostForm("tipe")

	err := model.DeleteAllVendorPic(owner, tipe)
	if err != nil {
		fmt.Println("error: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": err,
	})

}
