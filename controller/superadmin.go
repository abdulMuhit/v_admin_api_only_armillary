package controller

import (
	"fmt"
	"net/http"
	"strings"
	"v_admin_api_only_armillary/model"

	"github.com/bandros/framework"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetAllBrand(c *gin.Context) {

	id := c.PostForm("id")
	parent := c.DefaultPostForm("parent", "")

	fmt.Println("id, parent ", id, " parent ", parent)

	fmt.Println("GET ALL BRAND ")

	var result interface{}
	if parent == "" {
		res, err := model.GetAllBrandList("")
		if err != nil {
			fmt.Println("error ", err)
			//framework.ErrorJson(err.Error(), c)

			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"res":  res,
			})
			return
		}
		fmt.Println("res ", res)
		result = res
	} else {
		res2, err2 := model.GetMyBrandList(parent)
		if err2 != nil {
			fmt.Println("error ", err2)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"res":  res2,
			})
			return
		}

		fmt.Println("res2 ", res2)

		var str []string

		for _, u := range res2 {
			fmt.Println("res found u! brand ", u["id_brand"].(string))
			str = append(str, u["id_brand"].(string))
		}

		csv := strings.Join(str, ",")
		fmt.Println("csv ", csv)
		res3, err := model.GetAllBrandList(csv)
		if err != nil {
			fmt.Println("error ", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"res":  res3,
			})
			return
		}

		result = res3
		fmt.Println("result by vendor", result)
	}

	res2, err := model.GetMyBrandList(id)
	if err != nil {
		fmt.Println("error ", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"res":  res2,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":          http.StatusOK,
		"all_brandlist": result,
		"my_brandlist":  res2,
	})

}

func SUGetAllPengguna(c *gin.Context) {

	session := sessions.Default(c)
	v := session.Get(framework.Config("jwtName"))
	fmt.Println("session: ", session)
	fmt.Println("v: ", v)
	parent := c.DefaultPostForm("parent", "")

	data, err := model.GetListPenggunaByParent(parent)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
			"result": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": data,
	})
}

func GetPengguna(c *gin.Context) {
	fmt.Println("get pengguna controller")

	var datatables model.DatatablesSource
	err := c.Bind(&datatables)
	fmt.Println("get pengguna datatables ", datatables)

	if err != nil {
		//framework.ErrorJson(err.Error(), c)
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		return
	}

	extra := map[string]string{
		"status": c.DefaultPostForm("status", ""),
		"parent": c.DefaultPostForm("parent", ""),
	}

	fields := []string{
		"id",
		"email",
		"status",
		"parent",
		"role",
		"nama",
	}

	fmt.Println("fields, datatables, extra ", fields, datatables, extra)

	//data, err := model.GetListPenggunaByParent(parent)
	result, err := model.FilterUserByParent(fields, datatables, extra)
	fmt.Println("data ", result)

	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotFound,
			"result": err,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})

		//framework.ErrorJson(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, &result)
}

func DeletePengguna(c *gin.Context) {
	id := c.PostForm("id")
	//parent := c.PostForm("parent")

	/*if parent != "" {
		_, err := model.DeletePenggunaChild(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":   http.StatusNotModified,
				"result": "Error " + err.Error(),
			})
			return
		}
	}*/

	res, err := model.DeletePenggunaParent(id)
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

func TambahPengguna(c *gin.Context) {
	data := model.PenggunaStruct{}
	c.Bind(&data)
	fmt.Println("tambah pengguna data: ", data)
	fmt.Println("status pengguna ", data.Status)
	res, err := model.CekMemberPengguna(data)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotAcceptable,
			"result": "terdapat error 1" + err.Error(),
		})
		return
	}

	fmt.Println("res cont: ", res)
	if res {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotAcceptable,
			"result": "email telah terdaftar",
		})
		return

	}

	res3, err := model.TambahPengguna(data)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusNotAcceptable,
			"result": "terdapat error 3" + err.Error(),
		})
		return
	}
	fmt.Println("res3", res3)

	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": res3,
	})

}

func SignPengguna(c *gin.Context) {

	data := model.PenggunaStruct{}
	c.Bind(&data)
	fmt.Println("edit pengguna data: ", data)

	//result, err := model.EditPengguna(data)
	res, err := model.SignPenggunaBrand(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
		//framework.ErrorJson(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   http.StatusOK,
		"result": res,
	})

}

func GantiPassword(c *gin.Context) {
	token := c.GetHeader("token")
	fmt.Println("token ", token)

	email := c.PostForm("email")
	oldpass := c.PostForm("oldpass")
	newpass := c.PostForm("newpass")

	fmt.Println("oldpass, new pass ", oldpass, newpass)

	log, err := model.Login(email, oldpass)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusForbidden,
			"result": "password anda salah",
		})
		return
	}

	pass := framework.Password(newpass)
	fmt.Println("logid ", log["id"])
	fmt.Println("log ", log)

	r, err := model.ChangePasswordUser(log["id"].(string), pass)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   http.StatusForbidden,
			"result": "Ganti password gagal",
		})
		return
	}
	fmt.Println("log ", log)
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"newpass": r,
		"result":  "Ganti Password berhasil",
	})
}
