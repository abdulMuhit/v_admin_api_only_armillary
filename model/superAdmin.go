package model

import (
	"fmt"
	"strings"

	"github.com/bandros/framework"
)

type Brand struct {
	Brand     string
	Id        string
	Id_member string
	Logo      string
}

type PenggunaStruct struct {
	Id          string `form:"id"`
	Email       string `form:"email"`
	Status      string `form:"status"`
	Password    string `form:"password"`
	Brand_list  string `form:"brand_list"`
	Parent      string `form:"parent"`
	Role        string `form:"role"`
	Nama        string `form:"nama"`
	Vendor_logo string
}

type Pengguna struct {
	Id     string `json:"id"`
	Email  string `json:"email"`
	Parent string `json:"parent"`
	Status string `json:"status"`
	Role   string `json:"role"`
	Nama   string `json:"nama"`
}

func FilterUserByParent(fields []string, data DatatablesSource, extra map[string]string) (*DatatablesResult, error) {

	result := DatatablesResult{}
	result.Draw = data.Draw
	db := framework.Database{}
	defer db.Close()

	db.Select("a.email")
	db.From("pengguna a")
	if extra["parent"] != "" {
		db.Where("a.parent", extra["parent"])
	}

	num, err := db.Result()
	if err != nil {
		return nil, err
	}

	result.RecordsTotal = len(num)
	fmt.Println("result record total ", len(num))

	if data.Keyword != "" {
		db.StartGroup("AND")
		for _, v := range fields {
			db.WhereOr(v+" like", "%"+data.Keyword+"%")

		}
		db.EndGroup()
	}

	num, err = db.Result()
	fmt.Println("SQL:", db.QueryView())
	if err != nil {
		return nil, err
	}

	result.RecordsFiltered = len(num)
	fmt.Println("result record filtered ", len(num))

	db.Select(strings.Join(fields, ","))
	db.Limit(data.Length, data.Start)
	db.OrderBy(fields[data.OrdCol], data.OrdDir)
	dt, err := db.Result()

	if dt == nil {
		result.Data = []map[string]interface{}{}
	} else {
		result.Data = dt
	}

	return &result, nil
}

func GetListPenggunaByParent(parent string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	db.Select("id, email, status, parent, role, nama")
	db.From("pengguna")
	fmt.Println("parent: ", parent)
	if parent != "" {
		db.Where("parent", parent)
	}

	data, err := db.Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func DeletePenggunaParent(parent string) (interface{}, error) {
	db := framework.Database{}
	defer db.Close()
	insert := []string{
		parent,
	}
	fmt.Println("data insert ", insert)

	db.Call("delete_by_vendor", insert)

	return db.Result()
}

func CekMemberPengguna(data PenggunaStruct) (bool, error) {
	db := framework.Database{}
	defer db.Close()
	db.From("pengguna")
	//db.Where("no_member", data.No_member)
	//db.WhereOr("no_member", data.No_member)
	db.Where("email", data.Email)
	fmt.Println("datano member ", data)
	res, err := db.Result()
	if err != nil {
		return false, err
	}

	fmt.Println("res: ", res)
	if len(res) == 0 {
		return false, err
	}

	return true, err
}

func GetAllBrandList(listbrand string) ([]map[string]interface{}, error) {

	// just some mockup
	test := []map[string]interface{}{
		{
			"id":        "109",
			"brand":     "somebrand1",
			"id_member": "",
			"logo":      "https://images.fastcompany.net/image/upload/w_596,c_limit,q_auto:best,f_auto/fc/3034007-inline-i-applelogo.jpg",
			"slug":      "some slug here",
		},
		{
			"id":        "119",
			"brand":     "whateverbrand",
			"id_member": "2",
			"logo":      "https://about.canva.com/wp-content/uploads/sites/3/2016/08/Band-Logo.png",
			"slug":      "some slug here 2",
		},
	}

	/* 	type FruitBasket struct {
	   		Name    string
	   		Fruit   []string
	   		Id      int64  `json:"ref"`
	   		private string // An unexported field is not encoded.
	   		Created time.Time
	   	}

	   	basket := FruitBasket{
	   		Name:    "Standard",
	   		Fruit:   []string{"Apple", "Banana", "Orange"},
	   		Id:      999,
	   		private: "Second-rate",
	   		Created: time.Now(),
	   	}

	   		var jsonData []byte
	   	jsonData, err1 := json.Marshal(basket)
	   	if err1 != nil {
	   		log.Println(err1)
	   	}
	   	fmt.Println("JSON DATA : ", string(jsonData))

	*/

	return test, nil

}

func GetMyBrandList(user string) ([]map[string]interface{}, error) {
	fmt.Println("user str", user)
	db := framework.Database{}
	defer db.Close()
	db.From("brand_user")
	db.Where("id_user", user)
	res, err := db.Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("mybrandlist ", res)

	return res, err
}

func TambahPengguna(data PenggunaStruct) (interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	pass := framework.Password(data.Password)
	fmt.Println("data pengguna, ", data)

	ins := []string{
		data.Email,
		data.Status, //
		data.Parent, //
		pass,
		data.Brand_list, //
		data.Role,
		data.Nama,
	}

	fmt.Println("insert ", ins)
	db.Call("add_child_user", ins)
	return db.Result()
}

func SignPenggunaBrand(data PenggunaStruct) (interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	fmt.Println("data sign ", data)
	fmt.Println("data id ", data.Id)
	fmt.Println("data status ", data.Status)

	insert := []string{
		data.Email,
		data.Status,
		data.Brand_list,
		data.Nama,
		data.Id,
	}

	fmt.Println("data insert ", insert)

	db.Call("assign_brand", insert)

	return db.Result()

}
