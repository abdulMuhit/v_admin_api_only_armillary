package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bandros/framework"
	"github.com/gin-gonic/gin"
)

type VendorData struct {
	Id           string `form:"id"`
	Nama_lengkap string `form:"nama_lengkap"`
	No_hp        string `form:"no_hp"`
	Email        string `form:"email"`
	Alamat       string `form:"alamat"`
	Nama_brand   string `form:"nama_brand"`
	Kota         string `form:"kota"`
	Provinsi     string `form:"provinsi"`
	Kodepos      string `form:"kodepos"`
	Website      string `form:"website"`
	Facebook     string `form:"facebook"`
	Instagram    string `form:"instagram"`
	Marketplace  string `form:"marketplace"`
	Kategori     string `form:"kategori"`
	Status       string `form:"status"`
	Foto         []Pic
}

type Pic struct {
	Id       string
	Owner    string
	Tipe     string
	Path_url string
}

func DataVendorOne(id string) (VendorData, error) {
	var results VendorData

	db := framework.Database{}
	db.Select("id, nama, no_hp, email, alamat, nama_brand, kota, provinsi, kodepos, website, facebook, instagram, marketplace, kategori, status")
	db.Where("id", id)
	db.From("vendor")
	data, err := db.Row()
	if err != nil {
		return results, err
	}
	results.Id = data["id"].(string)
	results.Nama_lengkap = data["nama"].(string)
	results.No_hp = data["no_hp"].(string)
	results.Email = data["email"].(string)
	results.Alamat = data["alamat"].(string)
	results.Nama_brand = data["nama_brand"].(string)
	results.Kota = data["kota"].(string)
	results.Provinsi = data["provinsi"].(string)
	results.Kodepos = data["kodepos"].(string)
	results.Website = data["website"].(string)
	results.Facebook = data["facebook"].(string)
	results.Instagram = data["instagram"].(string)
	results.Marketplace = data["marketplace"].(string)
	results.Kategori = data["kategori"].(string)
	results.Status = data["status"].(string)

	db.Clear()
	db.Select("id, owner, tipe, path_url")
	db.Where("owner", id)
	db.Where("tipe", "vendorpic")
	db.From("picture")
	pic, err := db.Result()
	if err != nil {
		return results, err
	}

	var p Pic
	var ps []Pic

	for _, c := range pic {
		p.Id = c["id"].(string)
		p.Owner = c["owner"].(string)
		p.Tipe = c["tipe"].(string)
		p.Path_url = c["path_url"].(string)
		ps = append(ps, p)
	}

	if len(pic) < 3 {
		for i := len(pic); i < 3; i++ {
			p.Id = ""
			p.Owner = id
			p.Tipe = "vendorpic"
			p.Path_url = ""
			ps = append(ps, p)
		}
	}

	results.Foto = ps
	return results, nil
}

func ChangePasswordUser(id string, newpass string) (map[string]interface{}, error) {

	db := framework.Database{}
	defer db.Close()
	stmt := map[string]interface{}{
		"password": newpass,
	}

	db.From("pengguna")
	db.Where("id", id)

	err := db.Update(stmt)
	if err != nil {
		fmt.Println("error ", err)
		return nil, err
	}
	return stmt, err
}

func ChangeStatusVendor(c *gin.Context) (map[string]interface{}, error) {
	data := VendorData{}
	c.Bind(&data)

	db := framework.Database{}
	defer db.Close()
	stmt := map[string]interface{}{
		"status": data.Status,
	}

	db.From("vendor")
	db.Where("id", data.Id)

	err := db.Update(stmt)
	if err != nil {
		//framework.ErrorHtml(err.Error(),c)
		return nil, err
	}

	return stmt, err
}

func EditVendorM(c *gin.Context) (map[string]interface{}, error) {

	data := VendorData{}
	c.Bind(&data)

	db := framework.Database{}
	defer db.Close()

	stmt := map[string]interface{}{
		"nama":        data.Nama_lengkap,
		"no_hp":       data.No_hp,
		"email":       data.Email,
		"alamat":      data.Alamat,
		"nama_brand":  data.Nama_brand,
		"kota":        data.Kota,
		"provinsi":    data.Provinsi,
		"kodepos":     data.Kodepos,
		"website":     data.Website,
		"facebook":    data.Facebook,
		"instagram":   data.Instagram,
		"marketplace": data.Marketplace,
		"kategori":    data.Kategori,
		"status":      data.Status,
	}

	db.From("vendor")
	db.Where("id", data.Id)

	err := db.Update(stmt)
	if err != nil {
		//		framework.ErrorHtml(err.Error(),c)
		return nil, err
	}

	return stmt, err

}

func GetVendorPic(id string, tipe string) ([]map[string]interface{}, error) {
	db := framework.Database{}
	defer db.Close()

	db.Select("id, owner, tipe, path_url")
	db.Where("owner", id)
	db.Where("tipe", tipe)
	db.From("picture")
	pic, err := db.Result()
	if err != nil {
		return nil, err
	}

	return pic, err
}

func AddOneVendorPic(c *gin.Context) (interface{}, error) {

	id := c.PostForm("myid")
	tipe := c.PostForm("tipe")

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("error get file", err)
		return nil, err
	}
	ext := filepath.Ext(file.Filename)
	name := strings.TrimSuffix(file.Filename, ext)
	namefile := strings.Replace(name, " ", "_", -1)
	// todo make this flexible
	urlAddress := os.Getenv("imgPath") + tipe + id + namefile + ext
	toUrl := "." + urlAddress

	if err := c.SaveUploadedFile(file, toUrl); err != nil {
		println("error save: ", err)
		return nil, err
	}

	db := framework.Database{}
	defer db.Close()

	db.From("picture")
	insert := map[string]interface{}{
		"owner":    id,
		"tipe":     tipe,
		"path_url": urlAddress,
	}

	res, err := db.Insert(insert)
	return res, err
}

func DeleteOneVendorPic(id string, url string) (string, error) {

	pathname := "." + strings.TrimPrefix(url, os.Getenv("baseUrl")+":"+os.Getenv("portHost"))
	err := framework.RemoveFile(pathname)
	if err != nil {
		fmt.Println("error ", err)
		return "error", err
	}
	db := framework.Database{}
	defer db.Close()

	db.From("picture")
	db.Where("id", id)

	return "success", db.Delete()

}

func DeleteVendor(myid string) error {
	db := framework.Database{}
	defer db.Close()

	db.From("vendor")
	db.Where("id", myid)

	return db.Delete()
}

func DeleteAllVendorPic(owner string, tipe string) error {
	db := framework.Database{}
	defer db.Close()

	db.Select("owner, path_url")
	db.From("picture")
	db.Where("owner", owner)
	db.Where("tipe", tipe)
	res, err := db.Result()
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	for _, pre := range res {

		pathname := "." + strings.TrimPrefix(pre["path_url"].(string), os.Getenv("baseUrl")+":"+os.Getenv("portHost"))
		err := framework.RemoveFile(pathname)
		if err != nil {
			fmt.Println("error ", err)
			return err
		}
	}
	//
	db.Clear()
	db.From("picture")
	db.Where("owner", owner)
	db.Where("tipe", tipe)
	return db.Delete()
}

func FilterVendor(fields []string, data DatatablesSource, extra map[string]string) (*DatatablesResult, error) {

	result := DatatablesResult{}
	result.Draw = data.Draw
	db := framework.Database{}
	defer db.Close()
	db.Select("a.id")
	db.From("vendor a")
	//db.Join("karyawan k","k.id_absen=a.id_peg","inner")

	if extra["start"] == "1970-01-01" || extra["end"] == "1970-01-01" {

	} else if extra["start"] == "" || extra["end"] == "" {

	} else {
		db.Where("DATE(tanggal) >=", extra["start"])
		db.Where("DATE(tanggal) <=", extra["end"])
	}

	//db.Where("role !=" , 7)
	if extra["status"] != "all" {
		db.Where("a.status", extra["status"])
	}

	num, err := db.Result()
	if err != nil {
		return nil, err
	}
	result.RecordsTotal = len(num)

	if data.Keyword != "" {
		//TODO uncomeent this if you want to filter by date
		db.StartGroup("AND")
		for _, v := range fields {
			db.WhereOr(v+" like", "%"+data.Keyword+"%")
		}
		db.EndGroup()
	}

	num, err = db.Result()
	if err != nil {
		return nil, err
	}

	result.RecordsFiltered = len(num)

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

func FilterVendorExcel(fields []string, data DatatablesSource, extra map[string]string) ([]map[string]interface{}, error) {

	db := framework.Database{}
	defer db.Close()
	db.Select("a.id")
	db.From("vendor a")

	if extra["start"] == "1970-01-01" || extra["end"] == "1970-01-01" {

	} else if extra["start"] == "" || extra["end"] == "" {

	} else {
		db.Where("DATE(tanggal) >=", extra["start"])
		db.Where("DATE(tanggal) <=", extra["end"])
	}

	//db.Where("role !=" , 7)
	if extra["status"] != "all" {
		db.Where("a.status", extra["status"])
	}

	if data.Keyword != "" {
		//TODO uncoment this if you want to filter by date
		db.StartGroup("AND")
		for _, v := range fields {
			db.WhereOr(v+" like", "%"+data.Keyword+"%")
		}
		db.EndGroup()
	}

	db.Select(strings.Join(fields, ","))
	//db.Limit(data.Length,data.Start)
	db.OrderBy(fields[data.OrdCol], data.OrdDir)
	dt, err := db.Result()
	if err != nil {
		fmt.Println("error ", err)
		return nil, err
	}

	var rdata = []map[string]interface{}{}

	if dt != nil {
		rdata = dt
	}

	return rdata, nil
}
